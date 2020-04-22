package authorityservice

import (
	`go-authmanager/dbModule/dbauthority`
	`go-authmanager/dbModule/dbrole`
	"go-authmanager/module"
	`go-authmanager/response`
	"github.com/gin-gonic/gin"
	uuid `github.com/satori/go.uuid`
	`log`
	`time`
)

func CreateAuthority(c *gin.Context) {
	var authParam module.AuthAuthorityParamModule
	errA := c.ShouldBind(&authParam)
	if errA == nil {
		/*
			//检查当前的权限名称key是否存在
			        AuthAuthorityDataModule query = new AuthAuthorityDataModule();
			        query.setAuthorityName(authorityEntity.getAuthorityName());
			        AuthAuthorityDataModule authorityReturn = authAuthorityMapper.selectByAuthorityName(query);
			        if (authorityReturn != null) {
			            responseVo.setMessage("该权限名[" + authorityEntity.getAuthorityName() + "]已存在！");
			            return responseVo;
			        }
			        //检查是否存在父级权限
			        if (StringUtils.isNotBlank(authorityEntity.getAuthParent())) {
			            query.setAuthorityName(authorityEntity.getAuthParent());
			            AuthAuthorityDataModule pAuth = authAuthorityMapper.selectByAuthorityName(query);
			            if (pAuth == null) {
			                responseVo.setMessage("输入的父级名[" + authorityEntity.getAuthParent() + "]在系统中不存在！");
			                return responseVo;
			            } else {
			                authorityEntity.setParentId(pAuth.getGuid());
			            }
			        }
			        authorityEntity.setGuid(UuidCreate.createUuuid());
			        authorityEntity.setCreateTime(new Date());
			        authorityEntity.setModifyTime(new Date());
			        authAuthorityMapper.insertSelective(authorityEntity);
			        responseVo.setCode(ExpresswayResponCode.SUCCESS.getCode());
			        responseVo.setMessage("执行成功");
			        return responseVo;
		*/
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
func UpdateAuthority(c *gin.Context) {
	var authParam module.AuthAuthorityParamModule
	errA := c.ShouldBind(&authParam)
	if errA == nil {
		/*
		   		//检查是否存在父级权限
		              if (authorityEntity.getAuthParent() != null) {
		                  AuthAuthorityDataModule query = new AuthAuthorityDataModule();
		                  query.setAuthorityName(authorityEntity.getAuthParent());
		                  AuthAuthorityDataModule pAuth = authAuthorityMapper.selectByAuthorityName(query);
		                  if (pAuth == null) {
		                      responseVo.setMessage("输入的父级名[" + authorityEntity.getAuthParent() + "]在系统中不存在！");
		                      return responseVo;
		                  } else {
		                      authorityEntity.setParentId(pAuth.getGuid());
		                  }
		              }
		              authorityEntity.setModifyTime(new Date());
		              authAuthorityMapper.updateByPrimaryKeySelective(authorityEntity);
		*/
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

func DeleteAuthority(c *gin.Context) {
	var authParam module.AuthAuthorityParamModule
	errA := c.ShouldBind(&authParam)
	if errA == nil {
		/*
		  //删除权限关联的角色数据
		        AuthRoleAuthorityDataModule record = new AuthRoleAuthorityDataModule();
		        record.setAuthorityId(authorityEntity.getGuid());
		        authRoleAuthorityMapper.deleteById(record);
		        //删除具体权限数据，由于权限是树形结构，当前暂要求删除当前的该权限，对于当前权限的子集暂时不做删除处理
		        AuthAuthorityDataModule delete = new AuthAuthorityDataModule();
		        delete.setGuid(authorityEntity.getGuid());
		        authAuthorityMapper.deleteByPrimaryKey(authorityEntity.getGuid());
		*/
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

func QueryPageListAuthority(c *gin.Context) {
	var authParam module.AuthAuthorityParamModule
	errA := c.ShouldBind(&authParam)
	if errA == nil {
		/*
		 PageHelper.startPage(authorityEntity.getPage(), authorityEntity.getPageSize());
		        List<AuthAuthorityDataModule> list = authAuthorityMapper.listAllByCondition(authorityEntity);
		        PageInfo<AuthAuthorityDataModule> pageInfo = new PageInfo<AuthAuthorityDataModule>(list);
		        //查询分页数据
		        PageModule<AuthorityEntityResponsePo> pageModule = new PageModule<>();
		        List<AuthorityEntityResponsePo> voList = new ArrayList<>();
		        for (AuthAuthorityDataModule tr : pageInfo.getList()) {
		            AuthorityEntityResponsePo responseVo1 = DozerBeanMapperUtil.map(tr, AuthorityEntityResponsePo.class);
		            voList.add(responseVo1);
		        }
		        pageModule.setPage(authorityEntity.getPage());
		        pageModule.setPageSize(pageInfo.getPageSize());
		        pageModule.setTotal(pageInfo.getTotal());
		        pageModule.setRows(voList);
		*/
		queryAuth:=dbauthority.AuthAuthority{}
		pageList :=queryAuth.ListAllByCondition(&authParam)
		response.ShowData(c, pageList)
	} else {
		response.ShowError(c, "执行失败")
	}

}
