package response

type AuthorityEntityResponsePo struct {
	/*
	 @Mapping(value = "guid")
	    private String authId;//	是	String	权限id
	    private String authorityName;//	是	String	权限名称
	    private String authorityType	;//是	String	类型
	    private String startPath;//	否	String	前端路由
	    private String command;//	否	String	后端命令字
	    private String displayName	;//否	String	显示名
	    @Mapping(value = "parentName")
	    private String  authParent;//	否	String	父级
	    private Integer sort;//	否	String	排序
	    private String icon;//	否	String	图标

	*/
	AuthId        string `json:"authId"`
	AuthorityName string `json:"authorityName"`
	AuthorityType string `json:"authorityType"`
	StartPath     string `json:"startPath"`
	Command       string `json:"command"`
	AuthParent    string `json:"authParent"`
	Sort          int64 `json:"sort"`
	Icon          string `json:"icon"`
	DisplayName   string `json:"displayName"`




}