package routes

import (
	"net/http"

	"github.com/Popoola-Opeyemi/sprintBackend/models"
	"github.com/Popoola-Opeyemi/sprintBackend/utils"
	"github.com/labstack/echo"
)

// Dashboard ...
type Dashboard struct {
	Env    *utils.Enviroment
	Prefix string
}

// Init ...
func (d *Dashboard) Init(prefix string, env *utils.Enviroment) error {
	d.Env = env
	d.Prefix = prefix

	g := env.Echo.Group(prefix)
	// rt := g.Group("/auth")

	rt := g.Group("",
		GatePass(CreatePass(
			models.RoleAdmin,
			http.MethodPost,
			http.MethodPut,
			http.MethodGet,
			http.MethodDelete,
		), d.Env.Log),
	)

	rt.GET("/dashboard", d.GetDashboard)
	return nil
}

// Dashboard ...
func (d *Dashboard) GetDashboard(c echo.Context) error {
	res := &utils.Response{}
	ages := []models.AgeCount{}
	
	var err error

	count := make(map[string]int)

	participate := make(map[string][]models.Participate)

	count["students"], _ = models.GetTotalStudents(d.Env.DB)
	count["male"], _ = models.GetGenderCount(d.Env.DB, 1)
	count["female"], _ = models.GetGenderCount(d.Env.DB, 2)

	participate["participated2018"], _ = models.ParticipateCount(d.Env.DB, "participated2018")
	participate["participate2020"], _ = models.ParticipateCount(d.Env.DB, "participate2020")

	ages, _ = models.GetAgeCount(d.Env.DB)

	if err != nil {
		res.SetError(err)
		return c.JSON(http.StatusOK, res)
	}

	dashboard := &models.Dashboard{Count: count, AgeCount: ages, Participate: participate}

	res.SetData(dashboard)
	return c.JSON(http.StatusOK, res)

	// res.SetData(struct{ Message string }{"done"})
	// return c.JSON(http.StatusOK, res)
}
