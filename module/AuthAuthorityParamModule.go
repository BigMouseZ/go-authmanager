package module

type AuthAuthorityParamModule struct {
	AuthId        string `json:"authId"`
	AuthorityName string `json:"authorityName"`
	AuthorityType string `json:"authorityType"`
	StartPath     string `json:"startPath"`
	Command       string `json:"command"`
	AuthParent    string `json:"authParent"`
	Sort          int64 `json:"sort"`
	Icon          string `json:"icon"`
	Page          int64  `json:"page"`
	PageSize      int64  `json:"pageSize"`
	DisplayName   string `json:"displayName"`

	AuthorityState string `json:"authorityState"`
	UserId         string `json:"userId"`
	OnlyUserAuto  string `json:"onlyUserAuto"`
	ParentId  string `json:"parentId"`

}
