package roleservice

import (
	`go-authmanager/dbModule/dbrole`
	`go-authmanager/dbModule/dbuser`
	"go-authmanager/module"
	"go-authmanager/response"
	"github.com/gin-gonic/gin"
	uuid `github.com/satori/go.uuid`
	`log`
	`time`
)

type RoleEntityResponsePo struct {
	RoleId       string `json:"roleId" `
	RoleName     string `json:"roleName" `
	RoleDescribe string `json:"roleDescribe" `
}

func CreateRole(c *gin.Context) {
	var roleParam module.AuthRoleParamModule
	errA := c.ShouldBind(&roleParam)
	if errA == nil {
		insertRole := dbrole.AuthRoleModule{}
		// 判断是否存在该用角色
		roleReturn, err := insertRole.SelectByRoleName(&roleParam)
		if err == nil {
			response.ShowError(c, "该角色名["+roleReturn.RoleName+"]已存在！")
			return
		}
		insertRole.Guid = uuid.NewV4().String()
		insertRole.RoleDescribe = roleParam.RoleDescribe
		insertRole.RoleName=roleParam.RoleName
		insertRole.CreateTime = time.Now()
		insertRole.ModifyTime = time.Now()
		insertRole.InsertSelective()
		response.ShowSuccess(c, "执行成功")
	} else {
		response.ShowError(c, "执行失败")
	}
	/*jsonstr, _ := ioutil.ReadAll(c.Request.Body)
	var data map[string]interface{}
	err := json.Unmarshal(jsonstr, &data)
	if err != nil {
		response.ShowError(c, "fail")
		return
	}
	if _, ok := data["id"]; !ok {
		response.ShowError(c, "fail")
		return
	}
	id := int(data["id"].(float64))
	fmt.Println(id)
	datas := map[string]string{"status": "success"}
	response.ShowData(c, datas)*/
}

func UpdateRole(c *gin.Context) {
	var roleParam module.AuthRoleParamModule
	errA := c.ShouldBind(&roleParam)
	if errA == nil {
		updateRole := dbrole.AuthRoleModule{}
		// 判断是否存在该用角色
		roleParam.Notself = roleParam.RoleId
		roleReturn, err := updateRole.SelectByRoleName(&roleParam)
		if err == nil {
			response.ShowError(c, "该角色名["+roleReturn.RoleName+"]已存在！")
			return
		}
		updateRole.Guid = roleParam.RoleId
		updateRole.RoleName = roleParam.RoleName
		updateRole.RoleDescribe = roleParam.RoleDescribe
		updateRole.RoleName=roleParam.RoleName
		updateRole.UpdateByPrimaryKeySelective()
		response.ShowSuccess(c, "执行成功")
	} else {
		response.ShowError(c, "执行失败")
	}

}

func DeleteRole(c *gin.Context) {
	var roleParam module.AuthRoleParamModule
	errA := c.ShouldBind(&roleParam)
	if errA == nil {

		/*
		  //查询角色
		            AuthRoleDataModule role = authRoleMapper.selectByPrimaryKey(roleEntity.getGuid());
		            if (role != null) {
		                //删除角色捆绑的权限
		                AuthRoleAuthorityDataModule delete = new AuthRoleAuthorityDataModule();
		                delete.setRoleId(roleEntity.getGuid());
		                authRoleAuthorityMapper.deleteById(delete);
		                //删除用户捆绑的角色
		                AuthUsersRoleDataModule record = new AuthUsersRoleDataModule();
		                record.setRoleId(roleEntity.getGuid());
		                authUsersRoleMapper.deleteById(record);
		            //删除角色
		            authRoleMapper.deleteByPrimaryKey(role.getGuid());
		        } else {
		            log.error("根据角色id :{}找不到角色数据", roleEntity.getGuid());
		        }
		*/
		// //删除角色捆绑的权限
		deleteRole := dbrole.AuthRoleModule{}
		deleteRole.Guid = roleParam.RoleId
		roleReturn, err := deleteRole.SelectByPrimaryKey()
		if err == nil {
			deleteRoleAuth := dbrole.AuthRoleAuthority{}
			deleteRoleAuth.RoleId = roleReturn.Guid
			deleteRoleAuth.DeleteById()
			//删除用户捆绑的角色
			deleteUserROle := dbuser.AuthUsersRoleModule{}
			deleteUserROle.RoleId = roleReturn.Guid
			deleteUserROle.DeleteById()
			//删除角色
			deleteRole.DeleteByPrimaryKey()
		} else {
			log.Println("根据角色id :{}找不到角色数据", roleParam.RoleId)
		}
		response.ShowSuccess(c, "执行成功")
	} else {
		response.ShowError(c, "执行失败")
	}
}

func QueryListAllRole(c *gin.Context) {
	var roleParam module.AuthRoleParamModule
	errA := c.ShouldBind(&roleParam)
	if errA == nil {
		/*
			 BaseResponseModule responseVo = BaseResponseModule.createSuccess();
			        AuthRoleDataModule query = new AuthRoleDataModule();
			query.setUiview("c6afd0ef42eb455f98aeb2c3e334562a");
			List<AuthRoleDataModule> list = authRoleMapper.selectAll(query);
			List<RoleEntityResponsePo> voList = new ArrayList<>();
			for (AuthRoleDataModule tr : list) {
				RoleEntityResponsePo roleEntityResponseVo = new RoleEntityResponsePo();
				roleEntityResponseVo.setRoleId(tr.getGuid());
				roleEntityResponseVo.setRoleName(tr.getRoleName());
				roleEntityResponseVo.setRoleDescribe(tr.getRoleDescribe());
				voList.add(roleEntityResponseVo);
			}
			responseVo.setData(voList);
			return responseVo;
		*/
		queryRole := dbrole.AuthRoleModule{}
		list := queryRole.SelectAll()
		var voList []RoleEntityResponsePo
		if list != nil {
			for _, role := range list {
				roleResponse := RoleEntityResponsePo{}
				roleResponse.RoleId = role.Guid
				roleResponse.RoleName = role.RoleName
				roleResponse.RoleDescribe = role.RoleDescribe
				voList = append(voList, roleResponse)
			}
		}
		response.ShowData(c, voList)
		response.ShowSuccess(c, "执行成功")
	} else {
		response.ShowError(c, "执行失败")
	}

}

func UpdateRoleAuth(c *gin.Context) {
	var roleParam module.AuthRoleParamModule
	errA := c.ShouldBind(&roleParam)
	if errA == nil {
		/*
		   		 BaseResponseModule responseVo = BaseResponseModule.createSuccess();
		              //删除角色捆绑的权限
		              AuthRoleAuthorityDataModule delete = new AuthRoleAuthorityDataModule();
		              delete.setRoleId(roleEntity.getGuid());
		              authRoleAuthorityMapper.deleteById(delete);
		              //重新捆绑角色的权限
		              if (roleEntity.getRoleAuthIds() != null) {
		                  List<AuthRoleAuthorityDataModule> list = new ArrayList<AuthRoleAuthorityDataModule>();
		                  for (String authid : roleEntity.getRoleAuthIds()) {
		                      AuthRoleAuthorityDataModule roa = new AuthRoleAuthorityDataModule();
		                      roa.setGuid(UuidCreate.createUuuid());
		                      roa.setRoleId(roleEntity.getGuid());
		                      roa.setAuthorityId(authid);
		                      list.add(roa);
		                  }
		                  //批量插入
		                  authRoleAuthorityMapper.insertList(list);
		              }

		*/
		deleteRoleAuth := dbrole.AuthRoleAuthority{}
		deleteRoleAuth.RoleId = roleParam.RoleId
		deleteRoleAuth.DeleteById()
		if roleParam.RoleAuthIds != nil {
			roleAuth := dbrole.AuthRoleAuthority{}
			for _, authid := range roleParam.RoleAuthIds {
				roleAuth.Guid = uuid.NewV4().String()
				roleAuth.RoleId = roleParam.RoleId
				roleAuth.AuthorityId = authid
				roleAuth.CreateTime = time.Now()
				roleAuth.ModifyTime = time.Now()
				roleAuth.InsertList()
			}
		}
		response.ShowSuccess(c, "执行成功")
	} else {
		response.ShowError(c, "执行失败")
	}

}

func QueryListRoleAuth(c *gin.Context) {
	var roleParam module.AuthRoleParamModule
	errA := c.ShouldBind(&roleParam)
	if errA == nil {
		/*
		   		 List<String> list = new ArrayList<String>();
		              AuthRoleAuthorityDataModule query = new AuthRoleAuthorityDataModule();
		              query.setRoleId(roleEntity.getGuid());
		              query.setAuthorityState("all");
		              List<AuthRoleAuthorityDataModule> listData = authRoleAuthorityMapper.getForList(query);
		              if (listData != null) {
		                  for (AuthRoleAuthorityDataModule r : listData) {
		                      list.add(r.getAuthorityId());
		                  }
		              }
		              responseVo.setData(list);
		*/
		var list []string
		queryRoleAuth := dbrole.AuthRoleAuthority{}
		paramAuth :=module.AuthRoleParamModule{}
		paramAuth.RoleId=roleParam.RoleId
		paramAuth.AuthorityState="all"
		listData := queryRoleAuth.GetForList(&paramAuth)
		if listData != nil {
			for _, roleAuthority := range listData {
				list = append(list, roleAuthority.AuthorityId)
			}
		}
		response.ShowData(c,list)
	} else {
		response.ShowError(c, "执行失败")
	}
}
