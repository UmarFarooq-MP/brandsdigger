package auth

type Login struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}

type SignUp struct {
	Email     string `json:"email"`
	Password  string `json:"password"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Age       int    `json:"age"`
}

type Auth interface {
	Login(input Login) (string, error)
	Signup(input SignUp) (string, error)
}
