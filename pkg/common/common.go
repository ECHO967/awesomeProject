package common

type UserInfo struct {
	Username string `json:"username"`
	Nickname string `json:"nickname"`
	Profile  string `json:"profile""`
}

type UpdateInfo struct {
	Username string `json:"username"`
	Nickname string `json:"nickname"`
	Profile  string `json:"profile"`
	Token    string `json:"token"`
}
