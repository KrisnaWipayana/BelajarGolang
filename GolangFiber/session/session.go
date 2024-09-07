package session

import (
	"time"

	"github.com/gofiber/fiber/v2/middleware/session"
)

var Store *session.Store

func Session() {
	Store = session.New(session.Config{
		CookieHTTPOnly: true,
		CookieSecure:   false,
		Expiration:     24 * time.Hour,
		KeyLookup:      "cookie:session_id",
	})
}
