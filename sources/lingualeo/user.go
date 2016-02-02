package lingualeo

type User struct {
	Username     string `json:"nickname"`
	Id           int    `json:"user_id"`
	AutologinKey string `json:"autologin_key"`
}
