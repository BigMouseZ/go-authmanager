package dbconfig

import (
	`fmt`
	`go-authmanager/database`
	`go-authmanager/module`
	`time`
)

type ConfigParameter struct {
	Guid              string    `json:"guid" gorm:"primary_key;column:guid"`
	ParameterType     string    `json:"parameterType" gorm:"column:parameter_type"`
	ParameterKey      string    `json:"parameterKey" gorm:"column:parameter_key"`
	ParameterName     string    `json:"parameterName" gorm:"column:parameter_name"`
	ParameterValue    string    `json:"parameterValue" gorm:"column:parameter_value"`
	Sort              int64     `json:"sort" gorm:"column:sort"`
	ParameterDescribe string    `json:"parameterDescribe" gorm:"column:parameter_describe"`
	ParentId          string    `json:"parentId" gorm:"column:parent_id"`
	CreateTime        time.Time `json:"createTime" gorm:"column:create_time"`
	ModifyTime        time.Time `json:"modifyTime" gorm:"column:modify_time"`
}
func (a *ConfigParameter) SelectByPrimaryKey() int64 {
	var config ConfigParameter
	Db := database.Db.Table("config_parameter")
	if len(a.Guid) > 0 {
		Db = Db.Where("guid =?", a.Guid)
	}
	var result = Db.Find(&config)
	if result.Error == nil {
		return result.RowsAffected
	}
	return 0
}

func (a *ConfigParameter) InsertSelective() int64 {
	var result = database.Db.Table("config_parameter").Create(a)
	if result.Error == nil {
		return result.RowsAffected
	}
	return 0
}

func (a *ConfigParameter) DeleteByPrimaryKey() int64 {
	var result = database.Db.Table("config_parameter").Delete(a)
	if result.Error == nil {
		return result.RowsAffected
	}
	return 0
}


func (a *ConfigParameter) UpdateByPrimaryKeySelective() int64 {
	var updateSet = make(map[string]interface{})
	/*updateSet["real_name"] = a.RealName
	updateSet["head_icon"] = a.HeadIcon
	updateSet["phone"] = a.Phone
	updateSet["login_pwd"] = a.LoginPwd*/
	updateSet["modify_time"] = time.Now()
	result := database.Db.Table("config_parameter").Where("guid = ?", a.Guid).Updates(updateSet)
	if result.Error == nil {
		return result.RowsAffected
	}
	return 0
}


func (a *ConfigParameter) listAllByCondition(user *module.AuthUsersParamModule) []ConfigParameter {
	var configList []ConfigParameter
	Db := database.Db.Table("config_parameter")
	/*if len(user.LoginName) > 0 {
		Db = Db.Where("login_name like ?", "%"+user.LoginName+"%")

	}
	if len(user.RealName) > 0 {
		Db = Db.Where("real_name like ?", "%"+user.RealName+"%")

	}
	if len(user.RoleId) > 0 {
		Db = Db.Where(" guid  in (select user_id from auth_users_role where role_id = ?)", user.RoleId)
	}*/
	if err := Db.Find(&configList).Error; err != nil {
		fmt.Println(err.Error())
		return []ConfigParameter{}
	}
	return configList
}