package services

type ILoginService interface {
	AdminLogin(email, password string) string
}
