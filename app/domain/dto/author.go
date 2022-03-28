package dto

type AuthorizeDTO struct {
	CasbinUser string `json:"casbin_user"`
	RequestURI string `json:"request_uri"`
	Method     string `json:"method"`
}
