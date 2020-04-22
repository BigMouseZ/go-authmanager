package dbauthority

import (
	"go-authmanager/database"
	`go-authmanager/module`
	`fmt`
	"time"
)

type AuthAuthority struct {
	Guid          string    `json:"guid" gorm:"primary_key;column:guid"`
	AuthorityName string    `json:"authorityName" gorm:"column:authority_name"`
	AuthorityType string    `json:"authorityType" gorm:"column:authority_type"`
	StartPath     string    `json:"startPath" gorm:"column:start_path"`
	DisplayName   string    `json:"displayName" gorm:"column:display_name"`
	Sort          int64     `json:"sort" gorm:"column:sort"`
	Icon          string    `json:"icon" gorm:"column:icon"`
	Command       string    `json:"command" gorm:"column:command"`
	ParentId      string    `json:"parentId" gorm:"column:parent_id"`
	CreateTime    time.Time `json:"createTime" gorm:"column:create_time"`
	ModifyTime    time.Time `json:"modifyTime" gorm:"column:modify_time"`

}

type AuthAuthorityForList struct {
	Guid           string `json:"guid" gorm:"column:guid"`
	AuthorityName  string `json:"authorityName" gorm:"column:authority_name"`
	AuthorityType  string `json:"authorityType" gorm:"column:authority_type"`
	StartPath      string `json:"startPath" gorm:"column:start_path"`
	DisplayName    string `json:"displayName" gorm:"column:display_name"`
	Sort           int64 `json:"sort" gorm:"column:sort"`
	Icon           string `json:"icon" gorm:"column:icon"`
	Command        string `json:"command" gorm:"column:command"`
	ParentId       string `json:"parentId" gorm:"column:parent_id"`
	AuthorityState string `json:"parentId" gorm:"column:authority_state"`
	ParentName    string    `json:"parentName" gorm:"column:parent_name"`
}

func (a *AuthAuthority) InsertSelective() int64 {
	var result = database.Db.Table("auth_authority").Create(a)
	if result.Error == nil {
		return result.RowsAffected
	}
	return 0
}

func (a *AuthAuthority) DeleteByPrimaryKey() int64 {
	var result = database.Db.Table("auth_authority").Delete(a)
	if result.Error == nil {
		return result.RowsAffected
	}
	return 0
}

func (a *AuthAuthority) SelectByCommand(command string) AuthAuthority {

	return AuthAuthority{}
}

func (a *AuthAuthority) UpdateByPrimaryKeySelective() int64 {
	var updateSet = make(map[string]interface{})
	/*updateSet["real_name"] = a.RealName
	updateSet["head_icon"] = a.HeadIcon
	updateSet["phone"] = a.Phone
	updateSet["login_pwd"] = a.LoginPwd*/
	updateSet["modify_time"] = time.Now()
	result := database.Db.Table("auth_role").Where("guid = ?", a.Guid).Updates(updateSet)
	if result.Error == nil {
		return result.RowsAffected
	}
	return 0
}

func (a *AuthAuthority) GetForList(param *module.AuthAuthorityParamModule) []AuthAuthorityForList {
	var authList []AuthAuthorityForList
	/*
		直接Raw("xxxxx")或者Table("xxx").Select("xxxxx").Joins("xxxxxx").Where("xxxxxx") 这样的
	*/

	Db := database.Db
	Db = Db.Table("auth_authority tba").Select("tba.guid ,tba.authority_name,tba.authority_type,tba.start_path,"+
		"tba.display_name,tba.icon,tba.command,	tbb.authority_id as authority_state,? as user_id,tba.parent_id,tba.sort ", param.UserId).Joins(
		" left join (select distinct authority_id from auth_role_authority"+
			" where role_id in(select role_id from auth_users_role where user_id =?)) tbb on tba.guid = tbb.authority_id ", param.UserId).Where("1=1")

	/*
		rawSQL := "select tba.guid ,tba.authority_name,tba.authority_type,tba.start_path," +
			" tba.display_name,tba.icon,tba.command,	tbb.authority_id as authority_state,? as user_id," +
			" tba.parent_id,	tba.sort from auth_authority tba LEFT JOIN 	(select distinct authority_id from auth_role_authority" +
			" where role_id in(select role_id from auth_users_role where user_id =?)) tbb on tba.guid = tbb.authority_id where 1=1 "
		//sqlParam:=[]interface{param.UserId, param.UserId}
		var sqlParam = make([]interface{}, 5)
		sqlParam = append(append(sqlParam, param.UserId), param.UserId)
		if len(param.AuthorityState) > 0 {
			rawSQL += " and tbb.authority_id is not null "
		}
		if len(param.ParentId) > 0 && param.ParentId == "top" {
			rawSQL += " and tba.parent_id is null "
		} else if len(param.ParentId) > 0 {
			rawSQL += " and tba.parent_id = ？" + param.ParentId
			sqlParam = append(sqlParam, param.ParentId)
		}
		Db = Db.Raw(rawSQL, param.UserId, param.UserId)*/
	if len(param.AuthorityState) > 0 {
		Db = Db.Where("  tbb.authority_id is not null ")
	}
	if len(param.ParentId) > 0 && param.ParentId == "top" {
		Db = Db.Where("  tba.parent_id is null ")
	} else if len(param.ParentId) > 0 {
		Db = Db.Where("   tba.parent_id = ?", param.ParentId)
	}
	if err := Db.Scan(&authList).Error; err != nil {
		fmt.Println(err.Error())
	}
	return authList
}
func (a *AuthAuthority) ListAllByCondition(param *module.AuthAuthorityParamModule) []AuthAuthorityForList {
	var authList []AuthAuthorityForList
	sqlString := "  select  tbc.* from	(select tba.*,tbb.authority_name as parent_name from auth_authority tba" +
		" LEFT JOIN auth_authority tbb on tba.parent_id = tbb.guid) tbc where 1=1"
	Db := database.Db
	Db = Db.Table("auth_authority tba").Select("tba.*,tbb.authority_name as parent_name").Joins(" LEFT JOIN auth_authority tbb on tba.parent_id = tbb.guid").Where("1=1")

	if len(param.DisplayName) > 0 {
		Db = Db.Where("  tba.display_name like ?", "%"+param.DisplayName+"%")
	}
	if len(param.AuthorityName) > 0 {
		Db = Db.Where(" tba.authority_name like ?", "%"+param.AuthorityName+"%")
	}
	if len(param.OnlyUserAuto) > 0 {
		sqlString += ""
		Db = Db.Where("   tbc.guid in(select distinct authority_id from auth_role_authority where role_id in("+
			"select role_id from auth_users_role where user_id = '? )	)", param.OnlyUserAuto)
	}
	if param.Page > 0 && param.PageSize > 0 {
		Db = Db.Limit(param.PageSize).Offset((param.Page - 1) * param.PageSize)
	}
	result := Db.Scan(&authList)
	if result.Error == nil {
		return authList
	}
	return nil

}

func (a *AuthAuthority) SelectByAuthorityName() (AuthAuthority, error) {
	var authority AuthAuthority
	err := database.Db.Table("auth_authority").Where("authority_name = ?", a.AuthorityName).Find(&authority).Error
	return authority, err
}
