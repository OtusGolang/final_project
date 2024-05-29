package models

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type ListRequest struct {
	SUBNET string `json:"subnet"`
	//Mask string `json:"mask"`
}

type ClearBucket struct {
	Login string `json:"login"`
	Ip    string `json:"ip"`
}
