package user

type paramCreateUser struct {
	Username string `query:"username" json:"username" validate:"required" example:"ryanthe"`
	Email    string `query:"email" json:"email" validate:"required" example:"ryanthe@gmail.com"`
	Password string `query:"password" json:"password" validate:"required" example:"paijo321"`
}

type respon201CreateUser struct {
	Response ResponseCreateUser `json:"response"`
	Status   string             `json:"status" example:"201, kode nilai response dari sistem"`
}

type ResponseCreateUser struct {
	AvatarPath string `json:"avatar_path" example:"lokasi path avatar"`
	Email      string `json:"email" example:"email dari account"`
	Id         string `json:"id" example:"id yang tersimpan didalam sistem"`
	Token      string `json:"token" example:"nilai balik token dari sistem"`
	Username   string `json:"username" example:"username yang tersimpan"`
	Creat      string `json:"created_at" example:"tanggal dan jam data dibuat"`
	Updat      string `json:"updated_at" example:"tanggal dan jam data diupdate"`
}

type respon422CreateUser struct {
	Response Respon422DCreateUser `json:"error"`
	Status   string               `json:"status" example:"422, kode nilai response dari sistem"`
}

type Respon422DCreateUser struct {
	Reqemail string `json:"required_email" example:"Required Email"`
	Reqpass  string `json:"required_password" example:"Required Password"`
	Requname string `json:"Required_username" example:"Required Username"`
	Usertak  string `json:"Taken_username" example:"Username Already Taken"`
	Emailtak string `json:"Taken_email" example:"Username Already Taken"`
}

type respon500CreateUser struct {
	Response Respon500DCreateUser `json:"error"`
	Status   string               `json:"status" example:"500, kode nilai response dari sistem"`
}

type Respon500DCreateUser struct {
	Usertak  string `json:"Taken_username" example:"Username Already Taken"`
	Emailtak string `json:"Taken_email" example:"Username Already Taken"`
}

// @Tags User
// @Summary Create pengguna sistem
// @Description Create pengguna sistem, param wajib username, email dan password
// @Accept  json
// @Produce  json
// @Param payload body user.paramCreateUser true "payload"
// @Success 201 {object} user.respon201CreateUser
// @Success 422 {object} user.respon422CreateUser
// @Success 500 {object} user.respon500CreateUser
// @Router /users [post]
func dummyCreateUser()
