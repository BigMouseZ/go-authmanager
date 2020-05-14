package main

import (
	"github.com/gin-gonic/gin"
	"go-authmanager/aop"
	"go-authmanager/database"
	"go-authmanager/service/authorityservice"
	`go-authmanager/service/commonservice`
	`go-authmanager/service/configservice`
	`go-authmanager/service/dictionaryservice`
	"go-authmanager/service/roleservice"
	"go-authmanager/service/userservice"
)

func main() {

/*	url := "http://172.16.2.20:8100/expressway_track/shortpathtest/ShortPathTest.htm"
	abody:=make(map[string]string,0)
	abody["startGantryCode"] = "G008496949511004236"
	abody["endGantryCode"] = "G008162237386915590"
	resp,err:=goutils.PostUtil(url,abody,nil,nil )
	if err != nil {
		fmt.Println(err.Error())
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(string(body))*/
	database.Dbinit()
	//Default返回一个默认的路由引擎
	r := gin.Default()
	r.Use(aop.Auth())
	user := r.Group("/user")
	{
		user.POST("/createUser.htm", userservice.CreateUser)
		user.POST("/updateUser.htm", userservice.UpdateUser)
		user.POST("/deleteUser.htm", userservice.DeleteUser)
		user.POST("/queryPageListUser.htm", userservice.QueryPageListUser)
		user.POST("/queryListUserRoles.htm", userservice.QueryListUserRoles)

	}
	role := r.Group("/role")
	{
		role.POST("/createRole.htm", roleservice.CreateRole)
		role.POST("/updateRole.htm", roleservice.UpdateRole)
		role.POST("/deleteRole.htm", roleservice.DeleteRole)
		role.POST("/queryListAllRole.htm", roleservice.QueryListAllRole)
		role.POST("/updateRoleAuth.htm", roleservice.UpdateRoleAuth)
		role.POST("/queryListRoleAuth.htm", roleservice.QueryListRoleAuth)
	}
	authority := r.Group("/authority")
	{
		authority.POST("/createAuthority.htm", authorityservice.CreateAuthority)
		authority.POST("/updateAuthority.htm", authorityservice.UpdateAuthority)
		authority.POST("/deleteAuthority.htm", authorityservice.DeleteAuthority)
		authority.POST("/queryPageListAuthority.htm", authorityservice.QueryPageListAuthority)
	}

	// TODO 系统配置
	config := r.Group("/config")
	{
		config.POST("/createConfigGroup.htm", configservice.CreateConfigGroup)
		config.POST("/updateConfigGroup.htm", configservice.UpdateConfigGroup)
		config.POST("/deleteConfigGroup.htm", configservice.DeleteConfigGroup)
		config.POST("/queryListConfigGroup.htm", configservice.QueryListConfigGroup)
		config.POST("/createConfig.htm", configservice.CreateConfig)
		config.POST("/updateConfig.htm", configservice.UpdateConfig)
		config.POST("/deleteConfig.htm", configservice.DeleteConfig)
		config.POST("/queryListConfig.htm", configservice.QueryListConfig)
	}
	//TODO 字典
	dictionary:= r.Group("/dictionary")
	{
		dictionary.POST("/createDictionaryGroup.htm", dictionaryservice.CreateDictionaryGroup)
		dictionary.POST("/updateDictionaryGroup.htm", dictionaryservice.UpdateDictionaryGroup)
		dictionary.POST("/deleteDictionaryGroup.htm", dictionaryservice.DeleteDictionaryGroup)
		dictionary.POST("/queryListDictionaryGroup.htm", dictionaryservice.QueryListDictionaryGroup)
		dictionary.POST("/createDictionary.htm", dictionaryservice.CreateDictionary)
		dictionary.POST("/updateDictionary.htm", dictionaryservice.UpdateDictionary)
		dictionary.POST("/deleteDictionary.htm", dictionaryservice.DeleteDictionary)
		dictionary.POST("/queryListDictionary.htm", dictionaryservice.QueryListDictionary)
	}
	common := r.Group("/common")
	{
		common.POST("/login.htm", commonservice.Login)
		common.POST("/logout.htm", commonservice.Logout)
		common.POST("/queryRoleListAllAuthority.htm", commonservice.QueryRoleListAllAuthority)
		common.POST("/queryListUserAuth.htm", commonservice.QueryListUserAuth)
		common.POST("/upload/files.htm", commonservice.UploadFiles)
	}
	r.Run(":8080") // listen and serve on 0.0.0.0:8080
}
