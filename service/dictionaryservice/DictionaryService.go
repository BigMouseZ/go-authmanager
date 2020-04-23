package dictionaryservice

import (
	"github.com/gin-gonic/gin"
	uuid `github.com/satori/go.uuid`
	`go-authmanager/dbModule/dbauthority`
	`go-authmanager/dbModule/dbrole`
	"go-authmanager/module"
	`go-authmanager/response`
	`log`
	`time`
)

func CreateDictionaryGroup(c *gin.Context) {
	var authParam module.AuthAuthorityParamModule
	errA := c.ShouldBind(&authParam)
	if errA == nil {
		//检查当前的权限名称key是否存在

		insert := dbauthority.AuthAuthority{}
		insert.AuthorityName = authParam.AuthorityName
		_, err := insert.SelectByAuthorityName()
		if err == nil {
			response.ShowError(c, "该权限名["+authParam.AuthorityName+"]已存在！")
			return
		}
		//检查是否存在父级权限
		if len(authParam.AuthParent) > 0 {
			insert.AuthorityName = authParam.AuthParent
			pAuth, err := insert.SelectByAuthorityName()
			if err != nil {
				response.ShowError(c, "输入的父级名["+authParam.AuthParent+"]在系统中不存在！")
				return
			} else {
				insert.ParentId = pAuth.Guid
			}
		}
		insert.Guid = uuid.NewV4().String()
		insert.AuthorityName = authParam.AuthorityName
		insert.Command = authParam.Command
		insert.AuthorityType = authParam.AuthorityType
		insert.DisplayName = authParam.DisplayName
		insert.Icon = authParam.Icon
		insert.Sort = authParam.Sort
		insert.StartPath = authParam.StartPath
		insert.CreateTime = time.Now()
		insert.ModifyTime = time.Now()
		insert.InsertSelective()
		response.ShowSuccess(c, "执行成功")
	} else {
		log.Println(errA.Error())
		response.ShowError(c, "执行失败")
	}

}
func UpdateDictionaryGroup(c *gin.Context) {
	var authParam module.AuthAuthorityParamModule
	errA := c.ShouldBind(&authParam)
	if errA == nil {
		update := dbauthority.AuthAuthority{}
		//检查是否存在父级权限
		if len(authParam.AuthParent) > 0 {
			update.AuthorityName = authParam.AuthParent
			pAuth, err := update.SelectByAuthorityName()
			if err != nil {
				response.ShowError(c, "输入的父级名["+authParam.AuthParent+"]在系统中不存在！")
				return
			} else {
				update.ParentId = pAuth.Guid
			}
		}
		update.Guid = authParam.AuthId
		update.AuthorityName = authParam.AuthorityName
		update.Command = authParam.Command
		update.AuthorityType = authParam.AuthorityType
		update.DisplayName = authParam.DisplayName
		update.Icon = authParam.Icon
		update.Sort = authParam.Sort
		update.StartPath = authParam.StartPath
		update.ModifyTime = time.Now()
		response.ShowSuccess(c, "执行成功")
	} else {
		response.ShowError(c, "执行失败")
	}

}

func DeleteDictionaryGroup(c *gin.Context) {
	var authParam module.AuthAuthorityParamModule
	errA := c.ShouldBind(&authParam)
	if errA == nil {
		//删除权限关联的角色数据
		deleteRoleAuth  := dbrole.AuthRoleAuthority{}
		deleteRoleAuth.AuthorityId = authParam.AuthId
		deleteRoleAuth.DeleteById()
		//删除具体权限数据，由于权限是树形结构，当前暂要求删除当前的该权限，对于当前权限的子集暂时不做删除处理
		deleteAuth:=dbauthority.AuthAuthority{}
		deleteAuth.Guid = authParam.AuthId
		deleteAuth.DeleteByPrimaryKey()
		response.ShowSuccess(c, "执行成功")
	} else {
		response.ShowError(c, "执行失败")
	}

}

func QueryListDictionaryGroup(c *gin.Context) {
	var authParam module.AuthAuthorityParamModule
	errA := c.ShouldBind(&authParam)
	if errA == nil {
		queryAuth:=dbauthority.AuthAuthority{}
		pageList :=queryAuth.ListAllByCondition(&authParam)
		response.ShowData(c, pageList)
	} else {
		response.ShowError(c, "执行失败")
	}

}


func CreateDictionary(c *gin.Context) {
	var authParam module.AuthAuthorityParamModule
	errA := c.ShouldBind(&authParam)
	if errA == nil {
		//检查当前的权限名称key是否存在

		insert := dbauthority.AuthAuthority{}
		insert.AuthorityName = authParam.AuthorityName
		_, err := insert.SelectByAuthorityName()
		if err == nil {
			response.ShowError(c, "该权限名["+authParam.AuthorityName+"]已存在！")
			return
		}
		//检查是否存在父级权限
		if len(authParam.AuthParent) > 0 {
			insert.AuthorityName = authParam.AuthParent
			pAuth, err := insert.SelectByAuthorityName()
			if err != nil {
				response.ShowError(c, "输入的父级名["+authParam.AuthParent+"]在系统中不存在！")
				return
			} else {
				insert.ParentId = pAuth.Guid
			}
		}
		insert.Guid = uuid.NewV4().String()
		insert.AuthorityName = authParam.AuthorityName
		insert.Command = authParam.Command
		insert.AuthorityType = authParam.AuthorityType
		insert.DisplayName = authParam.DisplayName
		insert.Icon = authParam.Icon
		insert.Sort = authParam.Sort
		insert.StartPath = authParam.StartPath
		insert.CreateTime = time.Now()
		insert.ModifyTime = time.Now()
		insert.InsertSelective()
		response.ShowSuccess(c, "执行成功")
	} else {
		log.Println(errA.Error())
		response.ShowError(c, "执行失败")
	}

}
func UpdateDictionary(c *gin.Context) {
	var authParam module.AuthAuthorityParamModule
	errA := c.ShouldBind(&authParam)
	if errA == nil {
		update := dbauthority.AuthAuthority{}
		//检查是否存在父级权限
		if len(authParam.AuthParent) > 0 {
			update.AuthorityName = authParam.AuthParent
			pAuth, err := update.SelectByAuthorityName()
			if err != nil {
				response.ShowError(c, "输入的父级名["+authParam.AuthParent+"]在系统中不存在！")
				return
			} else {
				update.ParentId = pAuth.Guid
			}
		}
		update.Guid = authParam.AuthId
		update.AuthorityName = authParam.AuthorityName
		update.Command = authParam.Command
		update.AuthorityType = authParam.AuthorityType
		update.DisplayName = authParam.DisplayName
		update.Icon = authParam.Icon
		update.Sort = authParam.Sort
		update.StartPath = authParam.StartPath
		update.ModifyTime = time.Now()
		response.ShowSuccess(c, "执行成功")
	} else {
		response.ShowError(c, "执行失败")
	}

}

func DeleteDictionary(c *gin.Context) {
	var authParam module.AuthAuthorityParamModule
	errA := c.ShouldBind(&authParam)
	if errA == nil {
		//删除权限关联的角色数据
		deleteRoleAuth  := dbrole.AuthRoleAuthority{}
		deleteRoleAuth.AuthorityId = authParam.AuthId
		deleteRoleAuth.DeleteById()
		//删除具体权限数据，由于权限是树形结构，当前暂要求删除当前的该权限，对于当前权限的子集暂时不做删除处理
		deleteAuth:=dbauthority.AuthAuthority{}
		deleteAuth.Guid = authParam.AuthId
		deleteAuth.DeleteByPrimaryKey()
		response.ShowSuccess(c, "执行成功")
	} else {
		response.ShowError(c, "执行失败")
	}

}

func QueryListDictionary(c *gin.Context) {
	var authParam module.AuthAuthorityParamModule
	errA := c.ShouldBind(&authParam)
	if errA == nil {
		queryAuth:=dbauthority.AuthAuthority{}
		pageList :=queryAuth.ListAllByCondition(&authParam)
		response.ShowData(c, pageList)
	} else {
		response.ShowError(c, "执行失败")
	}

}
