package medutils

import "github.com/gin-gonic/gin"

func GetTokensFromCookie(context *gin.Context) (string, string, string, error) {
	authToken, err := context.Cookie("auth-token")
	if err != nil {
		return "", "", "", err
	}
	refreshToken, err := context.Cookie("refresh-token")
	if err != nil {
		return "", "", "", err
	}
	csrfKey, err := context.Cookie("csrf-key")
	if err != nil {
		return "", "", "", err
	}

	if authToken == "" || refreshToken == "" || csrfKey == "" {
		return "", "", "", ErrEmptyToken
	}

	return authToken, refreshToken, csrfKey, err
}
