package auth

type paramReset struct {
	Token   string `query:"token" json:"token" validate:"required" example:"XXXXXXXXX"`
	NewPass string `query:"new_password" json:"new_password" validate:"required" example:"PASSNEW"`
	RetPass string `query:"retype_password" json:"retype_password" validate:"required" example:"PASSNEW"`
}

type respon200Reset struct {
	Response string `json:"response" example:"Success, jawaban dari sistem"`
	Status   string `json:"status" example:"200, kode nilai response dari sistem"`
}

type respon422Reset struct {
	Response Respon422DReset `json:"error"`
	Status   string          `json:"status" example:"422, kode nilai response dari sistem"`
}

type Respon422DReset struct {
	Invtoken string `json:"Invalid_token" example:"Invalid link. Try requesting again, jawaban dari sistem"`
	Emptpass string `json:"Empty_passwords" example:"Please ensure both field are entered, jawaban dari sistem"`
}

// @Tags Auth
// @Summary Reset password
// @Description Reset password, param wajib token, new_password,retype_password
// @Accept  json
// @Produce  json
// @Param payload body auth.paramReset true "payload"
// @Success 200 {object} auth.respon200Reset
// @Success 422 {object} auth.respon422Reset
// @Router /password/reset [post]
func dummyReset() {}
