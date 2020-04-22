package module

type AuthUsersParamModule struct {
	UserId    string   `json:"userId"`
	LoginName string   `json:"loginName"`
	LoginPwd  string   `json:"loginPwd"`
	RealName  string   `json:"realName"`
	Phone     string   `json:"phone"`
	HeadIcon  string   `json:"headIcon"`
	RoleIds   []string `json:"roleIds"`
	Page      int64    `json:"page"`
	PageSize  int64    `json:"pageSize"`
	RoleId    string   `json:"roleId"`
	/**
	 * null表示用户为授权该角色，不是null则表示用户已授权该角色
	 */
	RoleState    string   `json:"roleState"`
}
