package dbrole

import (
	"go-authmanager/database"
	`go-authmanager/module`
	"time"
)

type AuthRoleModule struct {
	Guid         string    `json:"guid" gorm:"primary_key;column:guid"`
	RoleName     string    `json:"roleName" gorm:"column:role_name"`
	RoleDescribe string    `json:"roleDescribe" gorm:"column:role_describe"`
	CreateTime   time.Time `json:"createTime" gorm:"column:create_time"`
	ModifyTime   time.Time `json:"modifyTime" gorm:"column:modify_time"`
}

func (a *AuthRoleModule) InsertSelective() int64 {
	var result = database.Db.Table("auth_role").Create(a)
	if result.Error == nil {
		return result.RowsAffected
	}
	return 0
}

func (a *AuthRoleModule) DeleteByPrimaryKey() int64 {
	var result = database.Db.Table("auth_role").Delete(a)
	if result.Error == nil {
		return result.RowsAffected
	}
	return 0
}

func (a *AuthRoleModule) SelectByPrimaryKey()(AuthRoleModule,error)  {

	var user AuthRoleModule
	err:= database.Db.Table("auth_role").Where("guid = ?", a.Guid).Find(&user).Error
	if err!=nil{
		return AuthRoleModule{},err
	}
	return user,err
}

func (a *AuthRoleModule) UpdateByPrimaryKeySelective() int64 {
	var updateSet = make(map[string]interface{})
	updateSet["role_name"] = a.RoleName
	updateSet["role_describe"] = a.RoleDescribe
	updateSet["modify_time"] = time.Now()
	result := database.Db.Table("auth_role").Where("guid = ?", a.Guid).Updates(updateSet)
	if result.Error == nil {
		return result.RowsAffected
	}
	return 0
}

func (a *AuthRoleModule) SelectAll() []AuthRoleModule {
	var userList []AuthRoleModule
	result := database.Db.Table("auth_role").Find(&userList)
	if result.Error == nil {
		return userList
	}
	return nil
}
func (a *AuthRoleModule) SelectByRoleName(roleParam *module.AuthRoleParamModule) (AuthRoleModule, error) {
	var role AuthRoleModule
	Db:= database.Db.Table("auth_role")
	if len(roleParam.RoleName)>0{
		Db = Db.Where("role_name = ?", roleParam.RoleName)
	}
	if len(roleParam.Notself)>0{
		Db = Db.Where("guid ! = ?",roleParam.Notself)
	}
	err :=Db.Find(&role).Error
	return role, err
}
