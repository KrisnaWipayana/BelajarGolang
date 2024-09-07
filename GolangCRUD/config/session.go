package config

import "github.com/gorilla/sessions"

const SESSION_ID = "go_auth"

var Store = sessions.NewCookieStore([]byte("111111111111111"))
