package response

type UserAuthStateResponsePo struct {
	AuthId        string                    `json:"authId"`
	AuthorityName string                    `json:"authorityName"`
	AuthorityType string                    `json:"authorityType"`
	StartPath     string                    `json:"startPath"`
	Command       string                    `json:"command"`
	AuthHas       bool                      `json:"authHas"`
	Sort          int64                     `json:"sort"`
	Icon          string                    `json:"icon"`
	DisplayName   string                    `json:"displayName"`
	Children      []*UserAuthStateResponsePo `json:"children"`
}
