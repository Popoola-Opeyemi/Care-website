package routes

import (
	"api/models"
	"api/utils"
	"net/http"

	"github.com/jinzhu/copier"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"

	"github.com/labstack/gommon/log"
)

// Testimonials ...
type Testimonials struct {
	Env    *utils.Enviroment
	Prefix string
}

// Init ...
func (pg *Testimonials) Init(prefix string, env *utils.Enviroment) error {
	pg.Env = env
	pg.Prefix = prefix

	g := env.Echo.Group(prefix)

	// All users have access to route
	g.GET("/testimonials", pg.GetAll)
	g.GET("/testimonials/:id", pg.Get)

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
	rt.POST("/testimonials", pg.Create)
	rt.DELETE("/testimonials/:id", pg.Delete)
	rt.POST("/testimonials/:id", pg.Update)
	return nil
}

// GetAll ...
func (pg *Testimonials) GetAll(c echo.Context) error {
	res := new(utils.Response)
	faq := new(models.Testimonials)

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

	list, count, err := faq.All(pg.Env.DB, perPage, limit)

	resp := models.TestimonialsListArray{Items: list, Total: count}

	if err != nil {
		res.SetError(err)
		return c.JSON(http.StatusOK, res)
	}

	res.Set("data", resp)
	return c.JSON(http.StatusOK, res)
}

// Create ...
func (pg *Testimonials) Create(c echo.Context) error {
	response := new(utils.Response)
	page := new(models.Testimonials)
	form := struct {
		Author  string `json:"author" validate:"required"`
		Content string `json:"content" validate:"required"`
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

	response.Set("Testimonials", form)
	return c.JSON(http.StatusOK, response)

}

// Get ...
func (pg *Testimonials) Get(c echo.Context) error {
	res := &utils.Response{}

	id, err := utils.ToNum(c.Param("id"))
	if err != nil {
	}

	page := &models.Testimonials{
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
func (pg *Testimonials) Delete(c echo.Context) error {
	response := new(utils.Response)
	id, _ := utils.ToNum(c.Param("id"))
	page := models.Testimonials{ID: id}

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
func (pg *Testimonials) Update(c echo.Context) error {
	response := new(utils.Response)
	page := new(models.Testimonials)
	id, _ := utils.ToNum(c.Param("id"))
	form := struct {
		Author  string `json:"author" validate:"required"`
		Content string `json:"content" validate:"required"`
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
