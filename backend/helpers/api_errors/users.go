package apierrors

import (
	"strings"

	"github.com/gin-gonic/gin"
)

func UserRegistrationError(err error) gin.H {
	if strings.Contains(err.Error(), "users_email_key") {
		return gin.H{"error": "email already exists"}
	} else if strings.Contains(err.Error(), "users_phone_key") {
		return gin.H{"error": "phone already exists"}
	} else if strings.Contains(err.Error(), "users_line_id_key") {
		return gin.H{"error": "line id already exists"}
	}
	return gin.H{"error": err.Error()}
}
