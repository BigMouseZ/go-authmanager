package module

type AuthRoleParamModule struct {
	RoleId       string   `json:"roleId"`
	RoleName     string   `json:"roleName"`
	RoleDescribe string   `json:"roleDescribe"`
	RoleAuthIds  []string `json:"roleAuthIds"`
	Notself      string   `json:"notself"`
	AuthorityState  string   `json:"authorityState"`


}
