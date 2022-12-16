package helpers

var ErrorMessageNormalizer map[string]string = map[string]string{
	"crypto/bcrypt: hashedPassword is not the hash of the given password": "Invalid email or password.",
}
