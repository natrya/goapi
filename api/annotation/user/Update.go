package user

type paramUpdateUser struct {
	Email    string `query:"email" json:"email" validate:"required" example:"ryanthe@gmail.com"`
	Password string `query:"password" json:"password" validate:"required" example:"paijo321"`
	Token    string `query:"token json:"token" validate:"required" example:"XXXXXXXXX"`
}

type respon200UpdateUser struct {
	Response ResponseUpdateUser `json:"response"`
	Status   string             `json:"status" example:"200, kode nilai response dari sistem"`
}

type ResponseUpdateUser struct {
	AvatarPath string `json:"avatar_path" example:"lokasi path avatar"`
	Email      string `json:"email" example:"email dari account"`
	Id         string `json:"id" example:"id yang tersimpan didalam sistem"`
	Token      string `json:"token" example:"nilai balik token dari sistem"`
	Username   string `json:"username" example:"username yang tersimpan disistem"`
}

type respon422UpdateUser struct {
	Response Respon422DUpdateUser `json:"error"`
	Status   string               `json:"status" example:"422, kode nilai response dari sistem"`
}

type Respon422DUpdateUser struct {
	Reqemail string `json:"required_email" example:"Required Email"`
	Reqpass  string `json:"required_password" example:"Required Password"`
	Incemail string `json:"No_record" example:"No Record Found"`
	Incpass  string `json:"Incorrect_password" example:"Incorrect Password"`
}

type respon401UpdateUser struct {
	Response Respon422DUpdateUser `json:"error"`
	Status   string               `json:"status" example:"422, kode nilai response dari sistem"`
}

type Respon401DUpdateUser struct {
	Reqemail string `json:"required_email" example:"Required Email"`
	Reqpass  string `json:"required_password" example:"Required Password"`
	Incemail string `json:"No_record" example:"No Record Found"`
	Incpass  string `json:"Incorrect_password" example:"Incorrect Password"`
}

// @Tags User
// @Summary Update user sistem
// @Description Update sistem, param wajib token ,email dan password
// @Accept  json
// @Produce  json
// @Param payload body user.paramUpdateUser true "payload"
// @Param   some_id     path    int     true        "ID User"
// @Success 200 {object} user.respon200UpdateUser
// @Success 401 {object} user.respon401UpdateUser
// @Success 422 {object} user.respon422UpdateUser
// @Security BearerAuth
// @Router /users/{some_id} [put]
func dummyUpdateUser() {}
