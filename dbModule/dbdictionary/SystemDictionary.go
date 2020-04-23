package dbconfig

import (
	`fmt`
	`go-authmanager/database`
	`go-authmanager/module`
	`time`
)

type ConfigDictionaries struct {
	Guid                 string    `json:"guid" gorm:"primary_key;column:guid"`
	DictionariesType     string    `json:"dictionariesType" gorm:"column:dictionaries_type"`
	ParentCode           string    `json:"parentCode" gorm:"column:parent_code"`
	DictionariesCode     string    `json:"dictionariesCode" gorm:"column:dictionaries_code"`
	DictionariesName     string    `json:"dictionariesName" gorm:"column:dictionaries_name"`
	Sort                 int64     `json:"sort" gorm:"column:sort"`
	DictionariesValue    string    `json:"dictionariesValue" gorm:"column:dictionaries_value"`
	DictionariesDescribe string    `json:"dictionariesDescribe" gorm:"column:dictionaries_describe"`
	CreateTime           time.Time `json:"createTime" gorm:"column:create_time"`
	ModifyTime           time.Time `json:"modifyTime" gorm:"column:modify_time"`
}

func (a *ConfigDictionaries) SelectByDictionariesCode() (ConfigDictionaries, error) {
	var dict ConfigDictionaries
	Db := database.Db.Table("config_dictionaries")
	if len(a.DictionariesCode) > 0 {
		Db = Db.Where("dictionaries_code =?", a.DictionariesCode)
	}
	err := Db.Find(&dict).Error
	return dict, err
}

func (a *ConfigDictionaries) SelectByPrimaryKey() (ConfigDictionaries, error) {
	var dict ConfigDictionaries
	Db := database.Db.Table("config_dictionaries")
	if len(a.Guid) > 0 {
		Db = Db.Where("guid =?", a.Guid)
	}
	err := Db.Find(&dict).Error
	return dict, err
}
func (a *ConfigDictionaries) GetCountByExample() int64 {
	count := 0
	Db := database.Db.Table("config_dictionaries")
	if len(a.Guid) > 0 {
		Db = Db.Where("guid =?", a.Guid)
	}
	err := Db.Count(&count).Error
	if err == nil {
		return int64(count)
	}
	return 0

}

func (a *ConfigDictionaries) InsertSelective() int64 {
	var result = database.Db.Table("config_dictionaries").Create(a)
	if result.Error == nil {
		return result.RowsAffected
	}
	return 0
}

func (a *ConfigDictionaries) DeleteByPrimaryKey() int64 {
	var result = database.Db.Table("config_dictionaries").Delete(a)
	if result.Error == nil {
		return result.RowsAffected
	}
	return 0
}

func (a *ConfigDictionaries) UpdateByPrimaryKeySelective() int64 {
	var updateSet = make(map[string]interface{})
	/*updateSet["real_name"] = a.RealName
	updateSet["head_icon"] = a.HeadIcon
	updateSet["phone"] = a.Phone
	updateSet["login_pwd"] = a.LoginPwd*/
	updateSet["modify_time"] = time.Now()
	result := database.Db.Table("config_dictionaries").Where("guid = ?", a.Guid).Updates(updateSet)
	if result.Error == nil {
		return result.RowsAffected
	}
	return 0
}

func (a *ConfigDictionaries) listAllByCondition(user *module.AuthUsersParamModule) []ConfigDictionaries {
	var dictList []ConfigDictionaries
	Db := database.Db.Table("config_dictionaries")
	/*if len(user.LoginName) > 0 {
		Db = Db.Where("login_name like ?", "%"+user.LoginName+"%")

	}
	if len(user.RealName) > 0 {
		Db = Db.Where("real_name like ?", "%"+user.RealName+"%")

	}
	if len(user.RoleId) > 0 {
		Db = Db.Where(" guid  in (select user_id from auth_users_role where role_id = ?)", user.RoleId)
	}*/
	if err := Db.Find(&dictList).Error; err != nil {
		fmt.Println(err.Error())
		return []ConfigDictionaries{}
	}
	return dictList
}
