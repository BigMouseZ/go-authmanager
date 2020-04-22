package commonservice

import (
	"fmt"
	"github.com/gin-gonic/gin"
	`github.com/gomodule/redigo/redis`
	uuid `github.com/satori/go.uuid`
	`go-authmanager/cache/rediskey`
	`go-authmanager/cache/redisutil/redisitem`
	`go-authmanager/cache/redisutil/redismap`
	`go-authmanager/dbModule/dbauthority`
	`go-authmanager/dbModule/dbuser`
	`go-authmanager/gofastdfs`
	"go-authmanager/module"
	"go-authmanager/response"
	`log`
	`strings`
)

func Login(c *gin.Context) {
	var user module.AuthUsersParamModule
	errA := c.ShouldBind(&user)
	if errA == nil {
		query := dbuser.AuthUsersModule{}
		query.LoginName = user.LoginName
		userReturn, err := query.SelectByLoginName()
		if err != nil {
			response.ShowError(c, "该用户名["+user.LoginName+"]不存在！")
			return
		}
		if user.LoginPwd != userReturn.LoginPwd {
			response.ShowError(c, "该用户名["+user.LoginName+"]的登录密码不正确！")
			return
		}
		//登录密码正确，开始获取用户权限
		last_session_id, err := redismap.GetByKey(rediskey.USER_SESSIONID_TABLE, userReturn.LoginName)

		if last_session_id != nil {
			lastsess, _ := redis.String(last_session_id, nil)
			//删除上次的登录缓存
			redismap.DeleteTable(lastsess)
			redisitem.DeleteByKey(rediskey.USER_LOGIN_INFO_TABLE + ":" + lastsess)
		}
		current_session_id := uuid.NewV4().String()
		redismap.SetByKey(rediskey.USER_SESSIONID_TABLE, userReturn.LoginName, current_session_id)
		//查询并装载用户权限到缓存
		authority := dbauthority.AuthAuthority{}
		authorityParam := module.AuthAuthorityParamModule{}
		authorityParam.UserId = userReturn.Guid
		authorityParam.AuthorityState = "all"
		user_auth_list := authority.GetForList(&authorityParam)
		for _, value := range user_auth_list {
			commands := value.Command
			if len(commands) > 0 && len(value.AuthorityState) > 0 {
				command_array := strings.Split(commands, ",")
				for _, command := range command_array {
					redismap.SetByKey(current_session_id, command, value)
				}
			}
		}
		/*
		   //组织数据返回
		         Map<String, Object> response_map = new HashMap<>();
		         response_map.put("loginName", userReturn.getLoginName());
		         response_map.put("realName", userReturn.getRealName());
		         response_map.put("headIcon", userReturn.getHeadIcon());
		         response_map.put("sessionId", current_session_id);
		         response_map.put("userId", userReturn.getGuid());
		         //缓存用户信息和sessionid的捆绑
		         Long seconds = cmsApi.getParamItemLongValue(ParameterRedisKey.param_userAuthority_option.param_userAuthority_invalidSeconds);
		         long secondsLo = seconds == null ? 7200 : seconds.longValue();
		         redisItemCacheService.setByKey(CmsRedisKey.USER_LOGIN_INFO_TABLE + ":" + current_session_id, userReturn, secondsLo);
		         responseVo.setCode(ExpresswayResponCode.SUCCESS.getCode());
		         responseVo.setMessage("登录成功");
		         responseVo.setData(response_map);

		*/
		response_map := make(map[string]interface{})
		response_map["loginName"] = userReturn.LoginName
		response_map["realName"] = userReturn.RealName
		response_map["headIcon"] = userReturn.HeadIcon
		response_map["sessionId"] = current_session_id
		response_map["userId"] = userReturn.Guid
		//缓存用户信息和sessionid的捆绑
		redisitem.SetByKey(rediskey.USER_LOGIN_INFO_TABLE+":"+current_session_id, userReturn)
		response.ShowData(c, response_map)
	} else {
		log.Println(errA.Error())
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
	fmt.Println(id)*/
}

func Logout(c *gin.Context) {
	var b module.AuthUsersParamModule
	errA := c.ShouldBind(&b)
	sessionId := c.GetHeader("sessionId")
	if errA == nil {
		/*
		   		 //清除登录缓存
		              redisMapCacheService.deleteTable(sessoinid);
		              AuthUsersDataModule userInfo = redisItemCacheService.getByKey(CmsRedisKey.USER_LOGIN_INFO_TABLE + ":" + sessoinid);
		              if (userInfo != null && userInfo.getLoginName() != null) {
		                  redisMapCacheService.deleteByKey(CmsRedisKey.USER_SESSIONID_TABLE, userInfo.getLoginName());
		              }
		              redisItemCacheService.deleteByKey(CmsRedisKey.USER_LOGIN_INFO_TABLE + ":" + sessoinid);
		              redisMapCacheService.deleteTable(sessoinid);
		*/

		redismap.DeleteTable(sessionId)
		user, err := redisitem.GetByKey(rediskey.USER_LOGIN_INFO_TABLE + ":" + sessionId)
		if err == nil {
			p, ok := user.(dbuser.AuthUsersModule)
			if ok {
				redismap.DeleteByKey(rediskey.USER_SESSIONID_TABLE, p.LoginName)
			} else {
				fmt.Println("e is not an People")
			}

		}
		redisitem.DeleteByKey(rediskey.USER_LOGIN_INFO_TABLE + ":" + sessionId)

		redismap.DeleteTable(sessionId)
	} else {
		log.Println(errA.Error())
		response.ShowError(c, "执行失败")
	}

}

func QueryRoleListAllAuthority(c *gin.Context) {
	var b module.AuthUsersParamModule
	errA := c.ShouldBind(&b)
	//sessionId:=c.GetHeader("sessionId")
	if errA == nil {
		/*
		   		 AuthAuthorityDataModule queryUserPage = new AuthAuthorityDataModule();
		              AuthUsersDataModule usersEntity = redisItemCacheService.getByKey(CmsRedisKey.USER_LOGIN_INFO_TABLE + ":" + sessionId);
		   		//        queryUserPage.setOnlyUserAuto(usersEntity.getGuid());
		   		List<AuthAuthorityDataModule> list = authAuthorityMapper.listAllByCondition(queryUserPage);
		   		List<AuthorityEntityResponsePo> voList = new ArrayList<>();
		   		for (AuthAuthorityDataModule tr : list) {
		   			AuthorityEntityResponsePo responseVo1 = DozerBeanMapperUtil.map(tr, AuthorityEntityResponsePo.class);
		   			voList.add(responseVo1);
		   		}
		*/
		authQuer := dbauthority.AuthAuthority{}
		list := authQuer.ListAllByCondition(&module.AuthAuthorityParamModule{})
		voList := make([]response.AuthorityEntityResponsePo, 0)
		for _, auth := range list {
			authRep := response.AuthorityEntityResponsePo{}
			authRep.AuthId = auth.Guid
			authRep.AuthorityName = auth.AuthorityName
			authRep.AuthorityType = auth.AuthorityType
			authRep.StartPath = auth.StartPath
			authRep.Command = auth.Command
			authRep.AuthParent = auth.ParentName
			authRep.Sort = auth.Sort
			authRep.Icon = auth.Icon
			authRep.DisplayName = auth.DisplayName
			voList = append(voList, authRep)
		}
		response.ShowData(c, voList)
	} else {
		log.Println(errA.Error())
		response.ShowError(c, "执行失败")
	}
}

func QueryListUserAuth(c *gin.Context) {
	var userParam module.AuthUsersParamModule
	errA := c.ShouldBind(&userParam)
	if errA == nil {
		list := make([]response.UserAuthStateResponsePo, 0)
		authority := dbauthority.AuthAuthority{}
		param := module.AuthAuthorityParamModule{}
		param.UserId = userParam.UserId
		param.ParentId = "top"
		authList := authority.GetForList(&param)
		temp_list := make([]*response.UserAuthStateResponsePo, 0)
		for _, ua := range authList {
			uab := response.UserAuthStateResponsePo{}
			uab.AuthId = ua.Guid
			uab.AuthorityName = ua.AuthorityName
			uab.AuthorityType = ua.AuthorityType
			uab.Command = ua.Command
			uab.DisplayName = ua.DisplayName
			uab.StartPath = ua.StartPath
			uab.Icon = ua.Icon
			uab.Sort = ua.Sort
			if len(ua.AuthorityState) == 0 {
				uab.AuthHas = false
			} else {
				uab.AuthHas = true
			}
			//查询子节点
			param.ParentId = ua.Guid
			child_list := authority.GetForList(&param)
			c_list := make([]*response.UserAuthStateResponsePo, 0)
			for _, u := range child_list {
				ucb := response.UserAuthStateResponsePo{}
				ucb.AuthId = u.Guid
				ucb.AuthorityName = u.AuthorityName
				ucb.AuthorityType = u.AuthorityType
				ucb.Command = u.Command
				ucb.DisplayName = u.DisplayName
				ucb.StartPath = u.StartPath
				ucb.Icon = u.Icon
				ucb.Sort = u.Sort
				if len(u.AuthorityState) == 0 {
					ucb.AuthHas = false
				} else {
					ucb.AuthHas = true
				}
				c_list = append(c_list, &ucb)
				temp_list = append(temp_list, &ucb)
			}
			uab.Children = c_list
			list = append(list, uab)
		}
		for len(temp_list) > 0 {
			t_list := make([]*response.UserAuthStateResponsePo, 0)
			for _, ua2 := range temp_list {
				param.ParentId = ua2.AuthId
				child_list2 := authority.GetForList(&param)
				ttlist := make([]*response.UserAuthStateResponsePo, 0)
				for _, uu := range child_list2 {
					ucb2 := response.UserAuthStateResponsePo{}
					ucb2.AuthId = uu.Guid
					ucb2.AuthorityName = uu.AuthorityName
					ucb2.AuthorityType = uu.AuthorityType
					ucb2.Command = uu.Command
					ucb2.DisplayName = uu.DisplayName
					ucb2.StartPath = uu.StartPath
					ucb2.Icon = uu.Icon
					ucb2.Sort = uu.Sort
					if len(uu.AuthorityState) == 0 {
						ucb2.AuthHas = false
					} else {
						ucb2.AuthHas = true
					}
					ttlist = append(ttlist, &ucb2)
					t_list = append(t_list, &ucb2)
				}
				ua2.Children = ttlist
			}
			temp_list = temp_list[0:0]
			temp_list = t_list[0:]
		}
		response.ShowData(c,list)
	} else {
		log.Println(errA.Error())
		response.ShowError(c, "执行失败")
	}
}
func UploadFiles(c *gin.Context) {

	form, _ := c.MultipartForm()
	//这里是多文件上传 在之前单文件upload上传的基础上加 [] 变成upload[] 类似文件数组的意思
	files := form.File["files"]
	//循环存文件到服务器本地
	for _, file := range files {
		c.SaveUploadedFile(file, file.Filename)
		gofastdfs.UploadFiles(file)
	}
	response.ShowSuccess(c,"执行成功")
}