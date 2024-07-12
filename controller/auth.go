package controller

import (
	"fmt"
	"net/http"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

type formData struct {
	Username string `form:"username"`
	Password string `form:"password"`
}

func CreateSession(c echo.Context) error {
	fd := formData{}
	err := c.Bind(&fd)
	if err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}

	sess, err := session.Get("session", c)
	if err != nil {
		fmt.Println(err)
		// TODO: handle error
		return err
	}

	sess.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   86400 * 7,
		HttpOnly: true,
	}

	if fd.Username != "user" || fd.Password != "pass" {
		return c.NoContent(http.StatusUnauthorized)
	}

	sess.Values["username"] = fd.Username
	sess.Values["authenticated"] = true
	if err := sess.Save(c.Request(), c.Response()); err != nil {
		return err
	}

	return c.Redirect(http.StatusFound, "/admin")
}

func Logout(c echo.Context) error {
	sess, _ := session.Get("session", c)

	sess.Values["authenticated"] = false
	sess.Save(c.Request(), c.Response())

	return c.Redirect(http.StatusFound, "/")
}
