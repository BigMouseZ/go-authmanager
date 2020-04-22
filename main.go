package main

import (
	"go-authmanager/aop"
	"go-authmanager/database"
	"go-authmanager/service/authorityservice"
	`go-authmanager/service/commonservice`
	"go-authmanager/service/roleservice"
	"go-authmanager/service/userservice"
	"github.com/gin-gonic/gin"
)

func main() {


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
