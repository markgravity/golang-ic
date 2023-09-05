package forms

type AccessTokenForm struct {
	AccessToken string `header:"Authorization" binding:"required"`
}
