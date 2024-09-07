package req

type NewUser struct {
	Nama     string `validate:"required"`
	Email    string `validate:"required"`
	Password string `validate:"required"`
}

type CreateUser struct {
	Role     int
	Nama     string `validate:"required"`
	Email    string `validate:"required"`
	Password string `validate:"required"`
}
