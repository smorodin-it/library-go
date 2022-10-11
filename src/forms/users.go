package forms

type UserSignUpForm struct {
	Email           string `json:"email" validate:"required,email,lte=100"`
	Password        string `json:"password" validate:"required,gte=8,lte=20"`
	PasswordConfirm string `json:"passwordConfirm" validate:"eqfield=Password"`
}

type UserSignInForm struct {
	Email    string `json:"email" validate:"required,email,lte=100"`
	Password string `json:"password" validate:"required,gte=8,lte=20"`
}

type ProfileAddEditForm struct {
	FirstName  string  `json:"firstName" validate:"required,gte=2,lte=255"`
	LastName   string  `json:"lastName" validate:"required,gte=2,lte=255"`
	Patronymic *string `json:"patronymic" validate:"omitempty,gte=2,lte=255"`
	Phone      string  `json:"phone" validate:"required,len=11"`
	Address    string  `json:"address" validate:"required"`
	UserId     string  `json:"userId" validate:"required,uuid4"`
}
