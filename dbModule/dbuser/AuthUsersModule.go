package dbuser

import (
	"go-authmanager/database"
	"go-authmanager/module"
	"fmt"
	"time"
)

type AuthUsersModule struct {
	Guid       string    `json:"guid" gorm:"primary_key;column:guid"`
	LoginName  string    `json:"loginName" gorm:"column:login_name"`
	LoginPwd   string    `json:"loginPwd" gorm:"column:login_pwd"`
	RealName   string    `json:"realName" gorm:"column:real_name"`
	Phone      string    `json:"phone" gorm:"column:phone;default:null"`
	HeadIcon   string    `json:"headIcon" gorm:"column:head_icon"`
	CreateTime time.Time `json:"createTime" gorm:"column:create_time"`
	ModifyTime time.Time `json:"modifyTime" gorm:"column:modify_time"`
}

func (a *AuthUsersModule) InsertSelective() int64 {
	var result = database.Db.Table("auth_users").Create(a)
	if result.Error == nil {
		return result.RowsAffected
	}
	return 0
}

func (a *AuthUsersModule) DeleteByPrimaryKey() int64 {
	var result = database.Db.Table("auth_users").Delete(a)
	if result.Error == nil {
		return result.RowsAffected
	}
	return 0
}

func (a *AuthUsersModule) SelectByPrimaryKey(userId string) AuthUsersModule {

	return AuthUsersModule{}
}

func (a *AuthUsersModule) UpdateByPrimaryKeySelective() int64 {
	var updateSet = make(map[string]interface{})
	updateSet["real_name"] = a.RealName
	updateSet["head_icon"] = a.HeadIcon
	updateSet["phone"] = a.Phone
	updateSet["login_pwd"] = a.LoginPwd
	updateSet["modify_time"] = time.Now()
	result := database.Db.Table("auth_users").Where("guid = ?", a.Guid).Updates(updateSet)
	if result.Error == nil {
		return result.RowsAffected
	}
	return 0
}

func (a *AuthUsersModule) ListPageByCondition(user *module.AuthUsersParamModule) []AuthUsersModule {
	var userList []AuthUsersModule
	Db := database.Db.Table("auth_users")
	//result := database.Db.Table("auth_users").Find(&userList)
	if len(user.LoginName) > 0 {
		Db = Db.Where("login_name like ?", "%"+user.LoginName+"%")

	}
	if len(user.RealName) > 0 {
		Db = Db.Where("real_name like ?", "%"+user.RealName+"%")

	}
	if len(user.RoleId) > 0 {
		Db = Db.Where(" guid  in (select user_id from auth_users_role where role_id = ?)", user.RoleId)
	}
	if user.Page > 0 && user.PageSize > 0 {
		Db = Db.Limit(user.PageSize).Offset((user.Page - 1) * user.PageSize)
	}
	if err := Db.Find(&userList).Error; err != nil {
		fmt.Println(err.Error())
		return []AuthUsersModule{}
	}
	return userList
}

func (a *AuthUsersModule) SelectByLoginName() (AuthUsersModule,error) {
	var user AuthUsersModule
	err:= database.Db.Table("auth_users").Where("login_name = ?", a.LoginName).Find(&user).Error
	return user,err
}
