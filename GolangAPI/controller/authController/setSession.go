package authController

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
)

var store = sessions.NewCookieStore([]byte("123"))

func SetSession(c *gin.Context, key string, value interface{}) {

	session, _ := store.Get(c.Request, "session")
	session.Values[key] = value
	session.Save(c.Request, c.Writer)
}

func GetSession(c *gin.Context, key string) interface{} {
	session, _ := store.Get(c.Request, "session")
	return session.Values[key]
}

// ClearSession clears the session
func ClearSession(c *gin.Context) {
	session, _ := store.Get(c.Request, "session")
	session.Options.MaxAge = -1
	session.Save(c.Request, c.Writer)
}
