package aop

import (
	"github.com/gin-gonic/gin"
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
		log.Println("请求路径："+u.Path)
		// access the status we are sending
		status := c.Writer.Status()
		log.Println("当前状态："+ strconv.Itoa(status)) //状态 200
		sessionId := c.GetHeader("sessionId")
		baseurl:=u.Path

		if strings.Contains(baseurl,"/common/login.htm"){
			if len(sessionId)==0{

			}
		}
		/*
		 if (!baseurl.contains("/common/login.htm")) {
		            if (sessionId == null) {
		                log.error("登录状态验证-参数无效没有sessionid");
		                responseVo.setCode(ExpresswayResponCode.NOLOGIN.getCode());
		                responseVo.setMessage("登录状态验证-参数无效没有sessionid");
		                return responseVo;
		            }
		            boolean login = redisItemCacheService.exists(CmsRedisKey.USER_LOGIN_INFO_TABLE + ":" + sessionId);
		            if (!login) {
		                //没有登录
		                log.info("没有通过登录效验，或已过期");
		                responseVo.setCode(ExpresswayResponCode.NOLOGIN.getCode());
		                responseVo.setMessage("没有通过登录效验，或已过期");
		                return responseVo;
		            } else {
		                //登录效验通过
		                log.info("通过登录效验");
		                //刷新时效
		                Long seconds = cmsApi.getParamItemLongValue(ParameterRedisKey.param_userAuthority_option.param_userAuthority_invalidSeconds);
		                long secondsLo = seconds == null ? 7200 : seconds.longValue();
		                redisItemCacheService.expire(CmsRedisKey.USER_LOGIN_INFO_TABLE + ":" + sessionId, secondsLo);
		            }
		        }
		*/
	}
}
