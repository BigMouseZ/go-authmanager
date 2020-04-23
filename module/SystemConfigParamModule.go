package module

import (
	`time`
)

type ConfigParameterParamModule struct {
	Guid              string    `json:"guid" `
	ParameterType     string    `json:"parameterType" `
	ParameterKey      string    `json:"parameterKey" `
	ParameterName     string    `json:"parameterName" `
	ParameterValue    string    `json:"parameterValue" `
	Sort              int64     `json:"sort" `
	ParameterDescribe string    `json:"parameterDescribe"`
	ParentId          string    `json:"parentId" `
	CreateTime        time.Time `json:"createTime" `
	ModifyTime        time.Time `json:"modifyTime" `
}
