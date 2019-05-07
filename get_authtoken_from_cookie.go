package medutils

import "github.com/gin-gonic/gin"

func GetAuthTokenFromCookie(context *gin.Context) (authtoken string, err error) {
	authToken, err := context.Cookie("token")

	if err != nil {
		return "", err
	}

	if authToken == "" {
		return "", ErrEmptyAuthToken
	}

	return authToken, nil
}
