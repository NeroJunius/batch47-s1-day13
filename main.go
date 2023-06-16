package main

import (
	"batch47/connection"
	"context"
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo"
)

type Projects struct {
	ID           int
	Title        string
	Author       string
	StartDate    time.Time
	EndDate      time.Time
	Duration     string
	DescProjects string
	NodeJS       bool
	ReactJS      bool
	NextJS       bool
	TypeScript   bool
	Image        string
	StartFormat  string
	EndFormat    string
}

func main() {
	connection.DatabaseConnect()
	e := echo.New()
	e.Static("/assets", "assets")

	e.GET("/", Home)
	e.GET("/contactMe", contactMe)
	e.GET("/project", createProject)
	e.GET("/projectDetail/:id", projectDetail)
	e.GET("/update-project/:id", updateProject)

	e.POST("/add-project", addProject)
	e.POST("/update-project/:id", sendUpdatedProject)
	e.POST("/delete-project/:id", deleteProject)

	e.Logger.Fatal(e.Start("localhost:5000"))
}

// home
func Home(c echo.Context) error {
	data, _ := connection.Conn.Query(context.Background(), "SELECT id, image, title, detail_date, duration, author_id, description, node_js, react_js, next_js, type_script  FROM tb_projects")
	var result []Projects
	for data.Next() {
		var each = Projects{}

		err := data.Scan(&each.ID, &each.Title, &each.Author, &each.StartDate, &each.EndDate, &each.Duration, &each.DescProjects, &each.NodeJS, &each.ReactJS, &each.NextJS, &each.TypeScript, &each.Image)
		if err != nil {
			fmt.Println(err.Error())
			return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
		}

		each.Author = "Nero Achmad"

		result = append(result, each)
	}

	post := map[string]interface{}{
		"Project": result,
	}
	var tmpl, err = template.ParseFiles("views/index.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return tmpl.Execute(c.Response(), post)
}

// contact
func contactMe(c echo.Context) error {
	var tmpl, err = template.ParseFiles("views/contact-me.html")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	return tmpl.Execute(c.Response(), nil)
}

// project page
func createProject(c echo.Context) error {

	var tmpl, err = template.ParseFiles("views/project.html")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return tmpl.Execute(c.Response(), nil)
}

// detail time
func countingDuration(startDate, endDate string) string {
	startTime, _ := time.Parse("2006-01-02", startDate)
	endTime, _ := time.Parse("2006-01-02", endDate)

	durationTime := int(endTime.Sub(startTime).Hours())
	durationDays := durationTime / 24
	durationWeeks := durationDays / 7
	durationMonths := durationWeeks / 4
	durationYears := durationMonths / 12
	var duration string
	if durationYears > 1 {
		duration = strconv.Itoa(durationYears) + " years"
	} else if durationYears > 0 {
		duration = strconv.Itoa(durationYears) + " year"
	} else {
		if durationMonths > 1 {
			duration = strconv.Itoa(durationMonths) + " months"
		} else if durationMonths > 0 {
			duration = strconv.Itoa(durationMonths) + " month"
		} else {
			if durationWeeks > 1 {
				duration = strconv.Itoa(durationWeeks) + " weeks"
			} else if durationWeeks > 0 {
				duration = strconv.Itoa(durationWeeks) + " week"
			} else {
				if durationDays > 1 {
					duration = strconv.Itoa(durationDays) + " days"
				} else {
					duration = strconv.Itoa(durationDays) + " day"
				}
			}
		}
	}
	return duration
}

// add project
func addProject(c echo.Context) error {
	c.Request().ParseForm()

	title := c.FormValue("inputTitle")
	startDate := c.FormValue("startDate")
	endDate := c.FormValue("endDate")
	DescProjects := c.FormValue("inputDescription")

	var nodeJS bool
	if c.FormValue("nodeJS") == "yes" {
		nodeJS = true
	}
	var nextJS bool
	if c.FormValue("nextJS") == "yes" {
		nextJS = true
	}
	var reactJS bool
	if c.FormValue("reactJS") == "yes" {
		reactJS = true
	}
	var typeScript bool
	if c.FormValue("typeScript") == "yes" {
		typeScript = true
	}

	image := c.FormValue("inputImage")

	_, err := connection.Conn.Exec(context.Background(),
		"INSERT INTO tb_projects (image, title, start_date, end_date, duration, description, node_js, react_js, next_js, type_script) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)",
		image, title, startDate, endDate, countingDuration(startDate, endDate), DescProjects, nodeJS, reactJS, nextJS, typeScript)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.Redirect(http.StatusMovedPermanently, "/")
}

// project detail
func projectDetail(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	var listProjects = Projects{}

	err := connection.Conn.QueryRow(context.Background(),
		" SELECT id, image, title, start_date, end_date, duration, author_id, description, node_js, react_js, next_js, type_script  FROM tb_projects WHERE id=$1", id).Scan(
		&listProjects.ID, &listProjects.Image, &listProjects.Title, &listProjects.StartDate, &listProjects.EndDate, &listProjects.Duration,
		&listProjects.NodeJS, &listProjects.ReactJS, &listProjects.NextJS, &listProjects.TypeScript)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	data := map[string]interface{}{
		"Project":   listProjects,
		"StartDate": listProjects.StartDate.Format("28 Juni 1999"),
		"EndDate":   listProjects.EndDate.Format("28 Juni 1999"),
	}

	var tmpl, errTemp = template.ParseFiles("views/project-detail.html")

	if errTemp != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": errTemp.Error()})
	}

	return tmpl.Execute(c.Response(), data)
}

// edit project
func updateProject(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	var listProjects = Projects{}
	err := connection.Conn.QueryRow(context.Background(),
		" SELECT id, image, title, start-date, end-date, duration, description, node_js, react_js, next_js, type_script, auhtor_id  FROM tb_projects WHERE id=$1", id).Scan(
		&listProjects.ID, &listProjects.Image, &listProjects.Title, &listProjects.StartDate, &listProjects.EndDate, &listProjects.Duration,
		&listProjects.NodeJS, &listProjects.ReactJS, &listProjects.NextJS, &listProjects.TypeScript)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	data := map[string]interface{}{
		"Project":   listProjects,
		"StartDate": listProjects.StartDate.Format("28 Juni 1999"),
		"EndDate":   listProjects.EndDate.Format("28 Juni 1999"),
	}

	var tmpl, errTemp = template.ParseFiles("views/project-detail.html")

	if errTemp != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": errTemp.Error()})
	}

	return tmpl.Execute(c.Response(), data)
}

// post edited project
func sendUpdatedProject(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	title := c.FormValue("InputTitle")
	startDate := c.FormValue("InputStartDate")
	endDate := c.FormValue("InputEndDate")
	DescProjects := c.FormValue("inputDescriptions")

	var nodeJS bool
	if c.FormValue("nodeJS") == "yes" {
		nodeJS = true
	}
	var nextJS bool
	if c.FormValue("nextJS") == "yes" {
		nextJS = true
	}
	var reactJS bool
	if c.FormValue("reactJS") == "yes" {
		reactJS = true
	}
	var typeScript bool
	if c.FormValue("typeScript") == "yes" {
		typeScript = true
	}

	image := c.FormValue("InputAnImage")

	_, err := connection.Conn.Exec(context.Background(),
		"UPDATE tb_projects SET image=$1, title=$2, start_date=$3, end_date=$4, duration=$5, description=$6, node_js=$7, react_js=$8, next_js=$9, type_script=$10) WHERE id=$1",
		image, title, startDate, endDate, countingDuration(startDate, endDate), DescProjects, nodeJS, reactJS, nextJS, typeScript, id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.Redirect(http.StatusMovedPermanently, "/")
}

// delete project
func deleteProject(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	_, err := connection.Conn.Exec(context.Background(), "DELETE FROM tb_projects WHERE id=$1", id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.Redirect(http.StatusMovedPermanently, "/")
}
