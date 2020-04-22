package dbuser

import (
	"go-authmanager/database"
	`go-authmanager/module`
	`fmt`
	"time"
)

type AuthUsersRoleModule struct {
	Guid       string    `gorm:"primary_key;column:guid"`
	UserId     string    `gorm:"column:user_id"`
	RoleId     string    `gorm:"column:role_id"`
	CreateTime time.Time `gorm:"column:create_time"`
	ModifyTime time.Time `gorm:"column:modify_time"`
}
type AuthUsersRoleModuleForList struct {
	UserId    string `gorm:"column:user_id"`
	RoleId    string `gorm:"column:role_id"`
	RoleState string `gorm:"column:role_state"`
}

func (a *AuthUsersRoleModule) DeleteById() int64 {
	Db := database.Db.Table("auth_users_role")
	if len(a.Guid) > 0 {
		Db = Db.Where("guid =?", a.Guid)
	}
	if len(a.RoleId) > 0 {
		Db = Db.Where("role_id =?", a.RoleId)
	}
	var result = Db.Delete(a)
	if result.Error == nil {
		return result.RowsAffected
	}
	return 0
}

func (a *AuthUsersRoleModule) InsertSelective() int64 {
	var result = database.Db.Table("auth_users_role").Create(a)
	if result.Error == nil {
		return result.RowsAffected
	}
	return 0

}

func (a *AuthUsersRoleModule) GetForList(userParam *module.AuthUsersParamModule) []AuthUsersRoleModuleForList {
	var roleList []AuthUsersRoleModuleForList
	Db := database.Db
	Db = Db.Table("auth_role").Select(" auth_role.guid as role_id,auth_role.role_name,tba.role_id as role_state ,"+
		"tba.user_id").Joins("LEFT JOIN (select role_id,user_id from auth_users_role where user_id = ?) tba on tba.role_id = auth_role.guid", userParam.UserId).Where("1=1")
	if len(userParam.RoleState) > 0 {
		Db = Db.Where("tba.role_id is not null")
	}
	if err := Db.Scan(&roleList).Error; err != nil {
		fmt.Println(err.Error())
		return []AuthUsersRoleModuleForList{}
	}
	return roleList
	/*
	   select
	    auth_role.guid as role_id,
	    auth_role.role_name,
	    tba.role_id as role_state ,
	    tba.user_id
	    from auth_role LEFT JOIN (
	    select role_id,user_id from auth_users_role where user_id = '${userId}'
	    ) tba on tba.role_id = auth_role.guid
	    <where>
	      <if test="roleState != null and roleState != ''">
	        <![CDATA[ and  tba.role_id is not null ]]>
	      </if>
	    </where>
	*/
}
