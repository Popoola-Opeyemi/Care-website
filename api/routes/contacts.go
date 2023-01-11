package routes

import (
	"api/models"
	"api/utils"
	"net/http"

	"github.com/rs/xid"

	"github.com/jinzhu/copier"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"

	"github.com/labstack/gommon/log"
)

type Contacts struct {
	Env    *utils.Enviroment
	Prefix string
}

// Init ...
func (pg *Contacts) Init(prefix string, env *utils.Enviroment) error {
	pg.Env = env
	pg.Prefix = prefix

	// All persons have access to route
	g := env.Echo.Group(prefix)
	g.POST("/contacts", pg.Create)

	//  admin routes only thats needs authentication
	rt := g.Group("",
		GatePass(CreatePass(
			models.RoleAdmin,
			http.MethodPost,
			http.MethodPut,
			http.MethodGet,
			http.MethodDelete,
		), pg.Env.Log),
	)

	rt.GET("/contacts", pg.GetAll)
	rt.GET("/contacts/:id", pg.Get)
	rt.DELETE("/contacts/:id", pg.Delete)
	return nil
}

// GetAll ...
func (pg *Contacts) GetAll(c echo.Context) error {
	res := new(utils.Response)
	contacts := new(models.Contacts)

	// var err error

	page, err := utils.ToNum(c.QueryParam("page"))

	perPage, err := utils.ToNum(c.QueryParam("perPage"))

	if err != nil {
		res.SetError(err)
		return c.JSON(http.StatusOK, res)
	}

	limit := (page - 1) * perPage

	if err != nil {
		res.SetError(err)
		return c.JSON(http.StatusOK, res)
	}

	list, count, err := contacts.All(pg.Env.DB, perPage, limit)

	resp := models.ContactsListArray{Items: list, Total: count}

	if err != nil {
		res.SetError(err)
		return c.JSON(http.StatusOK, res)
	}

	res.Set("data", resp)
	return c.JSON(http.StatusOK, res)
}

// Create ...
func (pg *Contacts) Create(c echo.Context) error {
	response := new(utils.Response)
	page := new(models.Contacts)
	form := struct {
		Name         string `json:"name"`
		UUID         string `json:"uuid"`
		Email        string `json:"email" validate:"required"`
		MobileNumber string `json:"mobile_number" validate:"required"`
		Message      string `json:"message" validate:"required"`
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

	response.Set("Contacts", form)
	return c.JSON(http.StatusOK, response)

}

// Get ...
func (pg *Contacts) Get(c echo.Context) error {
	res := &utils.Response{}

	id, err := utils.ToNum(c.Param("id"))
	if err != nil {
	}

	page := &models.Contacts{
		ID: id,
	}

	log.Info("id", id)

	err = page.GET(pg.Env.DB, id)
	if err != nil {
		res.SetError(err)
		return c.JSON(http.StatusOK, res)
	}

	res.SetData(page)
	return c.JSON(http.StatusOK, res)
}

//Delete ...
func (pg *Contacts) Delete(c echo.Context) error {
	response := new(utils.Response)
	id, _ := utils.ToNum(c.Param("id"))
	page := models.Contacts{ID: id}

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
