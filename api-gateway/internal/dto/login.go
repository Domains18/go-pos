package dto


type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}


type LoginResult struct {
	Token       string `json:"token"`
	User        string `json:"user"`
	Expire      int    `json:"expire"`
	DisplayName string `json:"display_name"`
}

