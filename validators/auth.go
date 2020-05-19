package validators

type AuthUser struct {
	Account  string `form:"phone" valid:"required,stringlength(3|50)"`
	Password string `form:"password" valid:"required,stringlength(3|50)"`
}
