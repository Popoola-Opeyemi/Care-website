package routes

import (
	"api/models"
	"api/utils"
	"errors"
	"net/http"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/gommon/log"
)

var (
	errInvalidCredentials = errors.New("Email or Password is incorrect")
)

// Users ...
type Users struct {
	Env    *utils.Enviroment
	Prefix string
}

// Init ...
func (s *Users) Init(prefix string, env *utils.Enviroment) error {
	s.Env = env
	s.Prefix = prefix

	g := env.Echo.Group(prefix)

	auth := g.Group("/auth")
	auth.POST("/login", s.Login)
	auth.GET("/logout", s.Logout)
	auth.GET("/status", s.Status)
	auth.GET("/users", s.GetUsers)
	return nil
}

//  Status ...
func (s *Users) Status(c echo.Context) error {
	response := new(utils.Response)
	user := &models.Users{}

	sess, err := session.Get(models.AuthSessionName, c)

	if err != nil {
		response.SetError(errors.New("session has expired"))
		return c.JSON(http.StatusUnauthorized, response)
	}

	sessVal, ok := sess.Values["id"]
	if !ok {
		response.SetError(errors.New("session has invalid"))
		return c.JSON(http.StatusOK, response)
	}

	ID, ok := sessVal.(int)
	if !ok {
		response.SetError(errors.New("session is invalidd"))
		return c.JSON(http.StatusOK, response)
	}

	user.ID = ID
	if err := user.GET(s.Env.DB, user.ID); err != nil {
		response.SetError(err)
		return c.JSON(http.StatusOK, response)
	}

	user.Password = "***" // conceal password after use
	response.Set("admin", user)
	return c.JSON(http.StatusOK, response)

}

// GetUsers ...
func (s *Users) GetUsers(c echo.Context) error {
	res := &utils.Response{}
	admin := models.Users{}

	list, err := admin.All(s.Env.DB)

	if err != nil {
		res.SetError(err)
		return c.JSON(http.StatusOK, res)
	}

	for _, usr := range list {
		usr.Password = "***"
	}

	log.Info(list)

	res.Set("users", list)
	return c.JSON(http.StatusOK, res)

}

// login routes ...
func (s *Users) Login(c echo.Context) error {
	res := &utils.Response{}
	adminUser := &models.Users{}
	form := &models.LoginForm{}

	// Bind request body to form
	if err := c.Bind(form); err != nil {
		res.SetError(err)
		return c.JSON(http.StatusBadRequest, res)
	}
	// Validate form
	if err := c.Validate(form); err != nil {
		res.SetError(err)
		return c.JSON(http.StatusOK, res)
	}
	// Get the users email
	err := adminUser.GetByEmail(s.Env.DB, form.Email)

	log.Info(err)
	if err != nil {
		res.SetError(errInvalidCredentials)
		return c.JSON(http.StatusOK, res)
	}

	if !utils.CheckPasswordHash(form.Password, adminUser.Password) {
		res.SetError(errInvalidCredentials)
		return c.JSON(http.StatusOK, res)
	}

	adminUser.Password = "***" // conceal password after use
	sess, _ := session.Get(models.AuthSessionName, c)
	sess.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   86400, // setting cookie expires in a day
		HttpOnly: false,
	}
	sess.Values["id"] = adminUser.ID
	sess.Values["email"] = utils.Base64Encode(adminUser.Email)
	sess.Values["isLoggedIn"] = true
	sess.Values["role"] = adminUser.RoleID
	sess.Save(c.Request(), c.Response())

	// LogData := models.LogData{}

	res.SetData(adminUser)
	return c.JSON(http.StatusOK, res)

}

// Logout ...
func (s *Users) Logout(c echo.Context) error {

	res := &utils.Response{}
	sess, err := session.Get(models.AuthSessionName, c)
	if err != nil {
		res.SetError(err)
		return c.JSON(http.StatusBadRequest, res)
	}

	sess.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   -1, // delete cookie for the session
		HttpOnly: false,
	}

	sess.Save(c.Request(), c.Response())
	res.Set("logout", true)
	return c.JSON(http.StatusOK, res)

}
