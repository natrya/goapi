package auth

type paramForget struct {
	Email string `query:"paijan" json:"email" validate:"required" example:"ryanthe@gmail.com" default:"ryanthe@gmail.com"`
}

type respon200Forget struct {
	Response string `json:"response" example:"jawaban dari sistem"`
	Status   string `json:"status" example:"422, kode nilai response dari sistem"`
}

type respon422Forget struct {
	Response Respon422DForget `json:"error"`
	Status   string           `json:"status" example:"422, kode nilai response dari sistem"`
}

type Respon422DForget struct {
	Reqemail string `json:"Required_email" example:"Required Email"`
	Noemail  string `json:"No_email" example:"Sorry, we do not recognize this email"`
}

// @Tags Auth
// @Summary Forget password
// @Description Lupa password, param wajib email
// @Accept  json
// @Produce  json
// @Param payload body auth.paramForget true "payload"
// @Success 200 {object} auth.respon200Forget
// @Success 422 {object} auth.respon422Forget
// @Router /password/forgot [post]
func dummyForget() {}
