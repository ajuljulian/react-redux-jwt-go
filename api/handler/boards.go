package handler

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *Handler) PublicBoard(c echo.Context) (err error) {
	log.Print("Getting public content")
	return c.HTML(http.StatusOK, "public board")
}

func (h *Handler) UserBoard(c echo.Context) (err error) {
	return c.HTML(http.StatusOK, "user board")
}

func (h *Handler) ModeratorBoard(c echo.Context) (err error) {
	return c.HTML(http.StatusOK, "moderator board")
}

func (h *Handler) AdminBoard(c echo.Context) (err error) {
	return c.HTML(http.StatusOK, "admin board")
}
