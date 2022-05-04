package main

import (
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"strconv"
)

func main() {
	e := echo.New()
	e.GET("/", saludar)
	e.GET("/dividir", dividir)

	persons := e.Group("/people")
	persons.Use(middlewareLogPeople)
	persons.POST("", crear)
	persons.GET("/:id", consultar)
	persons.PUT("/:id", actualizar)
	persons.DELETE("/:id", borrar)

	e.Logger.Fatal(e.Start(":8080"))
}

func saludar(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{"saludo": "Hola mundo"})
}

func dividir(c echo.Context) error {
	d := c.QueryParam("id")
	f, _ := strconv.Atoi(d)
	if f == 0 {
		return c.String(http.StatusBadRequest, "El valor no puede ser cero")
	}

	r := 3000 / f
	return c.String(http.StatusOK, strconv.Itoa(r))
}

func crear(c echo.Context) error {
	return c.String(http.StatusOK, "crear")
}

func actualizar(c echo.Context) error {
	id := c.Param("id")
	return c.String(http.StatusOK, "actualizar "+id)
}

func borrar(c echo.Context) error {
	id := c.Param("id")
	return c.String(http.StatusOK, "borrar "+id)
}

func consultar(c echo.Context) error {
	id := c.Param("id")
	return c.String(http.StatusOK, "consultar "+id)
}

func middlewareLogPeople(f echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		log.Println("Request done for people")
		return f(c)
	}
}
