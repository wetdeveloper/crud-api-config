package crud_api

import (
	"hash/fnv"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/wetdeveloper/connection"
)

func CrudForm(c echo.Context) error {
	return c.Render(http.StatusOK, "crudPage.html", map[string]interface{}{})
}

func hash(s string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(s))
	return h.Sum32()
}

//Create,update,Delete
func Cud(c echo.Context) error {
	operation := c.FormValue("operation")
	username := c.FormValue("username")
	password := c.FormValue("password")
	mydb, _ := connection.Connect()
	if operation == "C" {
		if connection.InsertUser(mydb, username, strconv.FormatUint(uint64(hash(password)), 10)) {
			return c.String(http.StatusOK, "Created")
		}
		return c.String(http.StatusOK, "There is an error")
	} else if operation == "U" {
		if connection.UpdateUser(mydb, username, strconv.FormatUint(uint64(hash(password)), 10)) {
			return c.String(http.StatusOK, "Updated")
		}
		return c.String(http.StatusOK, "There is an error")
	} else if operation == "D" {
		if connection.DeleteUser(mydb, username) {
			return c.String(http.StatusOK, "Deleted")
		}
		return c.String(http.StatusOK, "There is an error")
	}
	return c.String(http.StatusOK, "Not valid operation")

}

//Read
func Read(c echo.Context) error {
	mydb, _ := connection.Connect()
	return c.Render(http.StatusOK, "userslist.html", connection.ListUsers(mydb))
}
