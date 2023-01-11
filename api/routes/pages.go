package routes

import (
	"api/models"
	"api/utils"
	"encoding/json"
	"net/http"

	"github.com/rs/xid"

	"github.com/jinzhu/copier"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
)

type Pages struct {
	Env    *utils.Enviroment
	Prefix string
}

// Init ...
func (pg *Pages) Init(prefix string, env *utils.Enviroment) error {
	pg.Env = env
	pg.Prefix = prefix

	g := env.Echo.Group(prefix)

	//  Plain routes all users have access
	g.GET("/pages", pg.GetAll)
	g.GET("/pages/:id", pg.GetPage)

	rt := g.Group("",
		GatePass(CreatePass(
			models.RoleAdmin,
			http.MethodPost,
			http.MethodPut,
			http.MethodGet,
			http.MethodDelete,
		), pg.Env.Log),
	)

	// Admin Routes only

	rt.POST("/pages", pg.CreatePages)
	rt.DELETE("/pages/:id", pg.DelPages)
	rt.POST("/pages/:id", pg.UpdatePages)
	return nil
}

// GetAll ...
func (pg *Pages) GetAll(c echo.Context) error {
	res := new(utils.Response)
	pages := new(models.Pages)

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

	list, count, err := pages.All(pg.Env.DB, perPage, limit)

	for _, items := range list {
		items.Attributes, _ = json.Marshal(struct{}{})

	}

	resp := models.PagesListArray{Items: list, Total: count}

	if err != nil {
		res.SetError(err)
		return c.JSON(http.StatusOK, res)
	}

	res.Set("data", resp)
	return c.JSON(http.StatusOK, res)
}

// CreatePages ...
func (pg *Pages) CreatePages(c echo.Context) error {
	response := new(utils.Response)
	page := new(models.Pages)
	form := struct {
		Name       string          `json:"name"`
		UUID       string          `json:"uuid"`
		Route      string          `json:"route" validate:"required"`
		Attributes json.RawMessage `json:"attributes" validate:"required"`
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
	form.Attributes, err = json.Marshal(form.Attributes)
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

	response.Set("Pages", form)
	return c.JSON(http.StatusOK, response)

}

// GetPage ...
func (pg *Pages) GetPage(c echo.Context) error {
	res := &utils.Response{}

	id := c.Param("id")
	log.Info(id)

	page := &models.Pages{
		UUID: id,
	}

	err := page.GETBYUUID(pg.Env.DB, id)
	if err != nil {
		res.SetError(err)
		return c.JSON(http.StatusOK, res)
	}

	res.SetData(page)
	return c.JSON(http.StatusOK, res)
}

//DelPages ...
func (pg *Pages) DelPages(c echo.Context) error {
	response := new(utils.Response)
	id := c.Param("id")
	page := models.Pages{UUID: id}

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

// UpdatePages ...
func (pg *Pages) UpdatePages(c echo.Context) error {
	response := new(utils.Response)
	page := new(models.Pages)
	// Attrib := models.Attributes{}
	id := c.Param("id")
	form := struct {
		Name       string          `json:"name" validate:"required"`
		Route      string          `json:"route" validate:"required"`
		Attributes json.RawMessage `json:"attributes" validate:"required"`
	}{}

	if err := page.GETBYUUID(pg.Env.DB, id); err != nil {
		response.SetError(err)
		return c.JSON(http.StatusOK, response)
	}

	if err := c.Bind(&form); err != nil {
		response.SetError(err)
		return c.JSON(http.StatusOK, response)
	}
	page.UUID = id

	if err := c.Validate(form); err != nil {
		response.SetError(err)
		return c.JSON(http.StatusOK, response)
	}

	var Attrib interface{}

	err := json.Unmarshal(form.Attributes, &Attrib)
	if err != nil {
		log.Fatal("an error occured", err)
	}
	err = utils.SaveImageInJson(pg.Env, Attrib.(map[string]interface{}))
	if err != nil {
		log.Fatal("an error occured", err)
	}
	form.Attributes, err = json.Marshal(Attrib)
	if err != nil {
		log.Fatal("an error occured", err)
	}

	copier.Copy(page, form)

	if err := page.UPDATE(pg.Env.DB); err != nil {
		response.SetError(err)
		return c.JSON(http.StatusOK, response)
	}

	response.SetData(struct{ Message string }{"done"})
	return c.JSON(http.StatusOK, response)

}
