package userservice

import (
	"go-authmanager/dbModule/dbuser"
	"go-authmanager/response"
	"github.com/gin-gonic/gin"

	"github.com/satori/go.uuid"
	"time"
)
import "go-authmanager/module"

func CreateUser(c *gin.Context) {
	var user module.AuthUsersParamModule
	errA := c.ShouldBind(&user)
	if errA == nil {
		//输出json结果给调用方
		insert := dbuser.AuthUsersModule{}
		insert.Guid = uuid.NewV4().String()
		insert.LoginName = user.LoginName
		insert.HeadIcon = user.HeadIcon
		insert.LoginPwd = user.LoginPwd
		insert.RealName = user.RealName
		insert.Phone = user.Phone
		insert.CreateTime = time.Now()
		insert.ModifyTime = time.Now()
		insert.InsertSelective()
		//添加用户的角色
		if len(user.RoleIds) > 0 {
			for _, roleId := range user.RoleIds {
				userRole := dbuser.AuthUsersRoleModule{}
				userRole.Guid = uuid.NewV4().String()
				userRole.UserId = roleId
				userRole.CreateTime = time.Now()
				userRole.ModifyTime = time.Now()
				userRole.InsertSelective()
			}
		}
		response.ShowSuccess(c, "执行成功")
	} else {
		response.ShowError(c, "执行失败")
	}
}
func UpdateUser(c *gin.Context) {
	var user module.AuthUsersParamModule
	errA := c.ShouldBind(&user)
	if errA == nil {
		//输出json结果给调用方
		update := dbuser.AuthUsersModule{}
		update.Guid = user.UserId
		update.LoginName = user.LoginName
		update.HeadIcon = user.HeadIcon
		update.LoginPwd = user.LoginPwd
		update.RealName = user.RealName
		update.Phone = user.Phone
		update.CreateTime = time.Now()
		update.ModifyTime = time.Now()
		update.UpdateByPrimaryKeySelective()
		response.ShowSuccess(c, "执行成功")
	} else {
		response.ShowError(c, "执行失败")
	}
}

func DeleteUser(c *gin.Context) {
	var user module.AuthUsersParamModule
	errA := c.ShouldBind(&user)
	if errA == nil {
		//输出json结果给调用方
		delete := dbuser.AuthUsersModule{}
		delete.Guid = user.UserId
		delete.DeleteByPrimaryKey()
		response.ShowSuccess(c, "执行成功")
	} else {
		response.ShowError(c, "执行失败")
	}
}

func QueryPageListUser(c *gin.Context) {
	var user module.AuthUsersParamModule
	errA := c.ShouldBind(&user)
	if errA == nil {
		//输出json结果给调用方
		query := dbuser.AuthUsersModule{}
		//query.Guid = user.UserId
		userList := query.ListPageByCondition(&user)
		response.ShowData(c, userList)
	} else {
		response.ShowError(c, "执行失败")
	}
}
func QueryListUserRoles(c *gin.Context) {
	var user module.AuthUsersParamModule
	errA := c.ShouldBind(&user)
	if errA == nil {
		//查询用户已授权的角色
		query := dbuser.AuthUsersRoleModule{}
		user.RoleState = "all"
		roleList := query.GetForList(&user)
		roleListId := make([]string, 0)
		if len(roleList) > 0 {
			for _, value := range roleList {
				roleListId = append(roleListId, value.RoleId)
			}
		}
		response.ShowData(c, roleListId)
	} else {
		response.ShowError(c, "执行失败")
	}
}
