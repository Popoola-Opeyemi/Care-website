package routes

import (
	"api/models"
	"api/utils"
	"encoding/json"
	"net/http"

	"github.com/jinzhu/copier"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"github.com/rs/xid"

	"github.com/labstack/gommon/log"
)

type Settings struct {
	Env    *utils.Enviroment
	Prefix string
}

// Init ...
func (pg *Settings) Init(prefix string, env *utils.Enviroment) error {
	pg.Env = env
	pg.Prefix = prefix

	g := env.Echo.Group(prefix)

	// All users have access to route
	g.GET("/settings/:id", pg.Get)
	g.GET("/settings", pg.Getall)

	rt := g.Group("",
		GatePass(CreatePass(
			models.RoleAdmin,
			http.MethodPost,
			http.MethodPut,
			http.MethodGet,
			http.MethodDelete,
		), pg.Env.Log),
	)

	// Admin routes only.
	rt.POST("/settings", pg.Create)
	rt.POST("/settings/:id", pg.Update)
	return nil
}

// Getall ...
func (pg *Settings) Getall(c echo.Context) error {
	response := new(utils.Response)
	page := new(models.Settings)

	list, err := page.All(pg.Env.DB)
	if err != nil {
		response.SetError(err)
		return c.JSON(http.StatusOK, response)
	}

	response.Set("data", list)
	return c.JSON(http.StatusOK, response)
}

// Create ...
func (pg *Settings) Create(c echo.Context) error {
	response := new(utils.Response)
	page := new(models.Settings)
	form := struct {
		Attributes json.RawMessage `json:"attributes" validate:"required"`
		UUID       string          `json:"uuid"`
	}{}

	err := c.Bind(&form)
	if err != nil {
		response.SetError(err)
		return c.JSON(http.StatusOK, response)
	}
	if err = c.Validate(&form); err != nil {
		response.SetError(err)
		return c.JSON(http.StatusOK, response)
	}
	form.UUID = xid.New().String()

	copier.Copy(page, form)

	if err != nil {
		response.SetError(err)
		return c.JSON(http.StatusOK, response)
	}

	err = utils.Transact(pg.Env.DB, func(db *gorm.DB) error {
		if err := page.Create(db); err != nil {
			return err
		}
		return nil
	})

	response.Set("Settings", form)
	return c.JSON(http.StatusOK, response)

}

// Get ...
func (pg *Settings) Get(c echo.Context) error {
	res := &utils.Response{}

	id, _ := utils.ToNum(c.Param("id"))

	page := &models.Settings{
		ID: id,
	}

	log.Info("id", id)

	err := page.GET(pg.Env.DB, id)
	if err != nil {
		res.SetError(err)
		return c.JSON(http.StatusOK, res)
	}

	res.SetData(page)
	return c.JSON(http.StatusOK, res)
}

//Delete ...
func (pg *Settings) Delete(c echo.Context) error {
	response := new(utils.Response)
	id, _ := utils.ToNum(c.Param("id"))
	page := models.Settings{ID: id}

	err := utils.Transact(pg.Env.DB, func(db *gorm.DB) error {
		if err := page.DELETE(db); err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		response.SetError(err)
		return c.JSON(http.StatusInternalServerError, response)
	}

	response.SetData(struct{ message string }{"done"})
	return c.JSON(http.StatusOK, response)

}

// Update ...
func (pg *Settings) Update(c echo.Context) error {
	response := new(utils.Response)
	page := new(models.Settings)
	id, _ := utils.ToNum(c.Param("id"))

	form := struct {
		Attributes json.RawMessage `json:"attributes" validate:"required"`
	}{}

	if err := page.GET(pg.Env.DB, id); err != nil {
		response.SetError(err)
		return c.JSON(http.StatusOK, response)
	}

	if err := c.Bind(&form); err != nil {
		response.SetError(err)
		return c.JSON(http.StatusOK, response)
	}

	page.ID = id

	if err := c.Validate(form); err != nil {
		response.SetError(err)
		return c.JSON(http.StatusOK, response)
	}

	copier.Copy(page, form)

	if err := page.UPDATE(pg.Env.DB); err != nil {
		response.SetError(err)
		return c.JSON(http.StatusOK, response)
	}

	response.SetData(struct{ Message string }{"done"})
	return c.JSON(http.StatusOK, response)

}
