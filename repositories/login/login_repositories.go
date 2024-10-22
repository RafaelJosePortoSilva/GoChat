package login_repositories

func GetPasswordByUsername(username string) (string, error) {
	return "pass", nil
}

func CreateNewLogin(username string, hash string) error {
	if !verifyUsernameDuplicity(username) {
		return nil
	} else {
		return nil // nao lembro como retornar erro
	}
}

func verifyUsernameDuplicity(username string) bool {
	return false
}

func VerifyUsernameExists(username string) error {
	return nil
}
