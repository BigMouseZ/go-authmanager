package aop

import (
	"github.com/gin-gonic/gin"
	`go-authmanager/cache/rediskey`
	`go-authmanager/cache/redisutil/redisitem`
	`go-authmanager/response`
	"log"
	"net/url"
	"strconv"
	`strings`
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		u, err := url.Parse(c.Request.RequestURI)
		if err != nil {
			panic(err)
		}
		//response.ShowSuccess(c,"aop校验")
		log.Println("请求路径：" + u.Path)
		// access the status we are sending
		status := c.Writer.Status()
		log.Println("当前状态：" + strconv.Itoa(status)) //状态 200
		sessionId := c.GetHeader("sessionId")
		baseurl := u.Path

		if !strings.Contains(baseurl, "/common/login.htm") {
			if len(sessionId) == 0 {
				log.Println("登录状态验证-参数无效没有sessionid")
				response.ShowError(c, "登录状态验证-参数无效没有sessionid")
				//中断后续执行
				c.Abort()
			}
			login, _ := redisitem.Exists(rediskey.USER_LOGIN_INFO_TABLE + ":" + sessionId)
			if !login {
				//没有登录
				log.Println("没有通过登录效验，或已过期")
				response.ShowError(c, "没有通过登录效验，或已过期")
				c.Abort()
				return
			} else {
				//登录效验通过
				log.Println("通过登录效验")
				//刷新时效
				//Long seconds = cmsApi.getParamItemLongValue(ParameterRedisKey.param_userAuthority_option.param_userAuthority_invalidSeconds);
				secondsLo := 7200
				redisitem.Expire(rediskey.USER_LOGIN_INFO_TABLE+":"+sessionId, secondsLo)
			}
		}

		/*
		  //权限校验
		        //进行全新校验
		        BaseResponseModule checkResult = checkAuth(sessionId, baseurl, responseVo);
		        if (checkResult != null) {
		            return checkResult;
		        }
		       String loginName = "null";
		        if (!baseurl.contains("/common/login.htm")) {
		            CurrentUserInfo currentUserInfo = cmsApi.getCurrentLoginUserInfo(sessionId);
		            loginName = currentUserInfo.getLoginName();
		        }
		        //后续程序
		        Object proceed = joinPoint.proceed();
		        //日志操作
		        handleLog(sessionId, baseurl, joinPoint, proceed, request,loginName);
		        return proceed;

		*/
	}
}
