package routes

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Popoola-Opeyemi/sprintBackend/models"
	"github.com/Popoola-Opeyemi/sprintBackend/utils"
	"github.com/jinzhu/copier"
	"github.com/labstack/echo"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/gommon/log"
)

// Students ...
type Students struct {
	Env    *utils.Enviroment
	Prefix string
}

// Init ...
func (s *Students) Init(prefix string, env *utils.Enviroment) error {
	s.Env = env
	s.Prefix = prefix

	g := env.Echo.Group(prefix)
	// rt := g.Group("/auth")

	rt := g.Group("",
		GatePass(CreatePass(
			models.RoleAdmin,
			http.MethodPost,
			http.MethodPut,
			http.MethodGet,
			http.MethodDelete,
		), s.Env.Log),
	)

	rt.GET("/students", s.GetAllStudents)
	rt.POST("/students", s.CreateStudents)
	rt.GET("/students/:id", s.GetStudents)
	rt.DELETE("/students/:id", s.DelStudents)
	rt.POST("/students/:id", s.UpdateStudent)
	return nil
}

// CreateStudents ...
func (s *Students) CreateStudents(c echo.Context) error {
	res := &utils.Response{}
	count := &models.Settings{}

	err := count.Get(s.Env.DB)

	if err != nil {
		log.Info("an error occured", err)
	}

	// Generating the registration iD and updating the count in settings table
	registrationID := utils.GenerateID(count)
	count.Update(s.Env.DB, count.RegistrationCount)

	student := &models.Students{}

	if err := c.Bind(&student); err != nil {
		res.SetError(err)
		return c.JSON(http.StatusBadRequest, res)
	}
	student.Firstname = utils.Capitalize(&student.Firstname)
	student.Lastname = utils.Capitalize(&student.Lastname)
	student.Middlename = utils.Capitalize(&student.Middlename)
	student.RegistrationID = registrationID
	img, err := utils.ImageHandler(student.Image, s.Env)

	if err != nil {
		student.Image, _ = json.Marshal(struct{}{})
	}
	student.Image, _ = json.Marshal(img)

	if err := c.Validate(student); err != nil {
		res.SetError(err)
		return c.JSON(http.StatusOK, res)
	}

	err = student.Create(s.Env.DB)

	if err != nil {
		res.SetError(err)
		return c.JSON(http.StatusOK, res)
	}

	res.SetData(struct{ Message string }{"done"})
	return c.JSON(http.StatusOK, res)

}

// UpdateStudent ...
func (s *Students) UpdateStudent(c echo.Context) error {
	res := &utils.Response{}
	id := c.Param("id")
	student := &models.Students{}

	if err := c.Bind(student); err != nil {
		res.SetError(err)
		return c.JSON(http.StatusOK, res)
	}

	student.ID = id
	// log.Info(student.DateRegistered)

	img, err := utils.ImageHandler(student.Image, s.Env)

	if err != nil {
		student.Image, _ = json.Marshal(struct{}{})
	}
	student.Image, _ = json.Marshal(img)

	if err := c.Validate(student); err != nil {
		res.SetError(err)
		return c.JSON(http.StatusOK, res)
	}

	if err := student.UPDATE(s.Env.DB); err != nil {
		res.SetError(err)
		return c.JSON(http.StatusOK, res)
	}

	sess, _ := session.Get(models.AuthSessionName, c)
	adminID := sess.Values["id"]
	NewID := adminID.(int)
	StudentID, _ := strconv.Atoi(student.ID)
	Jdata, _ := json.Marshal(struct{}{})

	logger := models.SaveLog(
		NewID,
		StudentID,
		c.Request().Header.Get("user-agent"),
		utils.RequestIP(c.Request()),
		"update",
		Jdata,
	)
	// logger.AdminID = NewID
	// logger.StudentID, _ = strconv.Atoi(student.ID)
	// logger.UserAgent = c.Request().Header.Get("user-agent")
	// logger.IPAddress = utils.RequestIP(c.Request())
	// logger.Action = "Update"
	// logger.Data, _ = json.Marshal(struct{}{})

	err = logger.Save(s.Env.DB)

	if err != nil {
		res.SetError(err)
		return c.JSON(http.StatusOK, res)
	}

	// log.Info("id", id)

	res.SetData(struct{ Message string }{"done"})
	return c.JSON(http.StatusOK, res)

}

// GetStudents a single student from database and returns array containing student ...
func (s *Students) GetStudents(c echo.Context) error {
	res := &utils.Response{}
	id := c.Param("id")
	student := &models.Students{
		ID: id,
	}
	err := student.GET(s.Env.DB, id)
	if err != nil {
		res.SetError(err)
		return c.JSON(http.StatusOK, res)
	}

	res.SetData(student)
	return c.JSON(http.StatusOK, res)

}

// GetAllStudents all students from the database and returns the list in an array ...
func (s *Students) GetAllStudents(c echo.Context) error {

	res := &utils.Response{}

	student := &models.Students{}

	var err error

	page, err := utils.ToNum(c.QueryParam("page"))

	perPage, err := utils.ToNum(c.QueryParam("perPage"))

	limit := (page - 1) * perPage

	if err != nil {
		res.SetError(err)
		return c.JSON(http.StatusOK, res)
	}

	list, count, err := student.All(s.Env.DB, perPage, limit)

	resp := models.StudentListArray{Items: list, Total: count}

	if err != nil {
		res.SetError(err)
		return c.JSON(http.StatusOK, res)
	}

	res.Set("student", resp)
	return c.JSON(http.StatusOK, res)

}

// DelStudents ...
func (s *Students) DelStudents(c echo.Context) error {
	res := utils.Response{}
	LogData := models.LogData{}

	id := c.Param("id")

	stdnt := &models.Students{}

	err := stdnt.GET(s.Env.DB, id)
	if err != nil {
		res.SetError(err)
		return c.JSON(http.StatusOK, res)
	}

	student := &models.Students{ID: id}

	if err := student.DELETE(s.Env.DB); err != nil {
		res.SetError(err)
		return c.JSON(http.StatusOK, res)
	}

	sess, _ := session.Get(models.AuthSessionName, c)
	adminID := sess.Values["id"]
	NewID := adminID.(int)
	StudentID, _ := strconv.Atoi(stdnt.ID)
	copier.Copy(&LogData, &stdnt)
	Jdata, _ := json.Marshal(LogData)

	logger := models.SaveLog(
		NewID,
		StudentID,
		c.Request().Header.Get("user-agent"),
		utils.RequestIP(c.Request()),
		"delete",
		Jdata,
	)

	err = logger.Save(s.Env.DB)

	if err != nil {
		res.SetError(err)
		return c.JSON(http.StatusOK, res)
	}

	res.Set("deleted", true)
	return c.JSON(http.StatusOK, res)
}
