package user

type respon200Query struct {
	Response ResponseQuery `json:"response"`
	Status   string        `json:"status" example:"200, kode nilai response dari sistem"`
}

type ResponseQuery struct {
	AvatarPath string `json:"avatar_path" example:"lokasi path avatar"`
	Email      string `json:"email" example:"email dari account"`
	Id         string `json:"id" example:"id yang tersimpan didalam sistem"`
	Token      string `json:"token" example:"nilai balik token dari sistem"`
	Username   string `json:"username" example:"username yang tersimpan disistem"`
	Creat      string `json:"created_at" example:"tanggal dan jam saat dibuat"`
	Updat      string `json:"updated_at" example:"tanggal dan jam saat diupdate"`
}

type respon404Query struct {
	Response Respon404DQuery `json:"error"`
	Status   string          `json:"status" example:"404, kode nilai response dari sistem"`
}

type Respon404DQuery struct {
	Nouser string `json:"No_user" example:"No User Found, jawaban dari sistem"`
}

// @Tags User
// @Summary Query id user
// @Description Query id user
// @Accept  json
// @Produce  json
// @Param   some_id     path    int     true        "ID User"
// @Success 200 {object} user.respon200Query
// @Success 404 {object} user.respon404Query
// @Router /users/{some_id} [get]
func dummyQueryUser() {}
