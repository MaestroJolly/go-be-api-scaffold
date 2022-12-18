package helpers

var ErrorMessageNormalizer map[string]string = map[string]string{
	"crypto/bcrypt: hashedPassword is not the hash of the given password":                           "Invalid username or password.",
	"ERROR: duplicate key value violates unique constraint \"users_username_key\" (SQLSTATE 23505)": "username already exists",
	"ERROR: duplicate key value violates unique constraint \"users_email_key\" (SQLSTATE 23505)":    "email already exists",
}
