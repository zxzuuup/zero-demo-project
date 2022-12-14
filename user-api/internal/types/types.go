// Code generated by goctl. DO NOT EDIT.
package types

type UserInfoReq struct {
	UserId int64 `json:"userid"`
}

type UserInfoResp struct {
	UserId   int64  `json:"userid"`
	NickName string `json:"nickname"`
	Mobile   string `json:"mobile"`
}

type UserCreateReq struct {
	Mobile   string `json:"mobile"`
	NickName string `json:"nickname"`
	Data     string `json:"data"`
}

type UserCreateResp struct {
	Flag bool `json:"flag"`
}
