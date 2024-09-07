package entities

type User struct {
	Id        int64
	Role      string `validate:"required"`
	Nama      string `validate:"required"`
	Email     string `validate:"required" label:"E-Mail"`
	Password  string `validate:"required"`
	Cpassword string `validate:"required"`
}

type UserUpdate struct {
	Id    int64
	Role  string `validate:"required"`
	Nama  string `validate:"required"`
	Email string `validate:"required" label:"E-Mail"`
}
