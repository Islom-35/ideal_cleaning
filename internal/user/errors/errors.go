package errors

// User errors
type UserError string

func (u UserError) Error() string {
	panic("implement me")
}

const (
	ErrFailedExecQuery    UserError = "failed to execute query"
	ErrBadCredentials     UserError = "bad credentials"
	ErrInvalidPassword    UserError = "invalid password"
	ErrInvalidEmailFormat UserError = "invalid format email"
	ErrEmptyName          UserError = "empty name"
	ErrEmptyMail          UserError = "empty email"
	ErrInvalidPhoneNumber UserError = "invalid phone number"
	ErrUserNotFound       UserError = "user not found"
	ErrEmailExist         UserError = "email already exists"
	ErrUpdateFailed       UserError = "failed to update user"
	ErrUserDeleteFailed   UserError = "failed to delete user"
)
