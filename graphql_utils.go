package medutils

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

func GinContextToContextMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := context.WithValue(c.Request.Context(), "GinContextKey", c)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}

func GinContextFromContext(ctx context.Context) (*gin.Context, error) {
	ginContext := ctx.Value("GinContextKey")
	if ginContext == nil {
		err := errors.New("could not retrieve gin.Context")
		return nil, err
	}

	gc, ok := ginContext.(*gin.Context)
	if !ok {
		err := errors.New("gin.Context has wrong type")
		return nil, err
	}
	return gc, nil
}

func ConvertToString(s *string) string {
	if s != nil {
		return *s
	}

	return ""
}

func ConvertToInt(i *int) int {
	if i != nil {
		return *i
	}

	return 0
}
