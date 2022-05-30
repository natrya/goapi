package auth

type paramLogin struct {
	Email    string `query:"email" json:"email" validate:"required" example:"ryanthe@gmail.com" default:"ryanthe@gmail.com"`
	Password string `query:"password" json:"password" validate:"required" example:"paijo321"`
}

type respon200Login struct {
	Response ResponseLogin `json:"response"`
	Status   string        `json:"status" example:"200, kode nilai response dari sistem"`
}

type ResponseLogin struct {
	AvatarPath string `json:"avatar_path" example:"lokasi path avatar"`
	Email      string `json:"email" example:"email dari account"`
	Id         string `json:"id" example:"id yang tersimpan didalam sistem"`
	Token      string `json:"token" example:"nilai balik token dari sistem"`
	Username   string `json:"username" example:"username yang tersimpan disistem"`
}

type respon422Login struct {
	Response Respon422DLogin `json:"error"`
	Status   string          `json:"status" example:"422, kode nilai response dari sistem"`
}

type Respon422DLogin struct {
	Reqemail string `json:"required_email" example:"Required Email"`
	Reqpass  string `json:"required_password" example:"Required Password"`
	Incemail string `json:"No_record" example:"No Record Found"`
	Incpass  string `json:"Incorrect_password" example:"Incorrect Password"`
}

// @Tags Auth
// @Summary Login ke sistem
// @Description Login ke sistem, param wajib email dan password
// @Accept  json
// @Produce  json
// @Param payload body auth.paramLogin true "payload"
// @Success 200 {object} auth.respon200Login
// @Success 422 {object} auth.respon422Login
// @Security BearerAuth
// @Router /login [post]
func dummyLogin() {}
