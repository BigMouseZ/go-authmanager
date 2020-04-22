package dbrole

import (
	"go-authmanager/database"
	`go-authmanager/module`
	"time"
)

type AuthRoleAuthority struct {
	Guid        string    `json:"guid" gorm:"primary_key;column:guid"`
	RoleId      string    `json:"roleId" gorm:"column:role_id"`
	AuthorityId string    `json:"authorityId" gorm:"column:authority_id"`
	CreateTime  time.Time `json:"createTime" gorm:"column:create_time"`
	ModifyTime  time.Time `json:"modifyTime" gorm:"column:modify_time"`
}

func (a *AuthRoleAuthority) InsertList() int64 {
	var result = database.Db.Table("auth_role_authority").Create(a)
	if result.Error == nil {
		return result.RowsAffected
	}
	return 0
}

func (a *AuthRoleAuthority) DeleteById() int64 {
	Db := database.Db.Table("auth_role_authority")
	if len(a.Guid) > 0 {
		Db = Db.Where("guid =?", a.Guid)
	}
	if len(a.RoleId) > 0 {
		Db = Db.Where("role_id =?", a.RoleId)
	}
	if len(a.RoleId) > 0 {
		Db = Db.Where("authority_id =?", a.AuthorityId)
	}
	var result = Db.Delete(a)
	if result.Error == nil {
		return result.RowsAffected
	}
	return 0
}

func (a *AuthRoleAuthority) GetForList(param *module.AuthRoleParamModule) []AuthRoleAuthority {
	var roleAuthList []AuthRoleAuthority
	/*
	select tbc.* from
	    (
	    SELECT
	    tba.guid as authority_id
	    from auth_authority tba LEFT JOIN auth_authority tbb on tba.parent_id = tbb.guid
	    ) tbc LEFT JOIN (select authority_id,role_id from auth_role_authority where role_id = '${roleId}'
	    ) tbd on tbd.authority_id = tbc.authority_id
	    <where>
	      <if test="authorityState != null and authorityState != ''">
	        <![CDATA[and  tbd.authority_id  is not null ]]>
	      </if>
	    </where>

	*/
	sqlString:="	select tbc.* from (SELECT tba.guid as authority_id from auth_authority tba LEFT JOIN auth_authority tbb on tba.parent_id = tbb.guid "+
	" ) tbc LEFT JOIN (select authority_id,role_id from auth_role_authority where role_id = ?) tbd on tbd.authority_id = tbc.authority_id"
	Db := database.Db
	var sqlParam = make([]interface{}, 0)
	sqlParam = append(sqlParam, param.RoleId)
	if len(param.AuthorityState) > 0 {
		sqlString += " and  tbd.authority_id  is not null "
	}
	Db = Db.Raw(sqlString,sqlParam)
	result := Db.Scan(&roleAuthList)
	if result.Error == nil {
		return roleAuthList
	}
	return nil
}
