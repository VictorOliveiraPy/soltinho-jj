package db

func ValidatePassword(role string) bool {
	err := role == "admin" || role == "instructor"
	return err
}
