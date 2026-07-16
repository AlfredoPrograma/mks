package users

import (
	"net/http"
	"strconv"

	"github.com/alfredoprograma/mks/internal/queries"
	"github.com/labstack/echo/v5"
)

type Controller interface {
	RegisterRoutes(root *echo.Group)

	CreateUser(c *echo.Context) error
	GetUser(c *echo.Context) error
}
type controller struct {
	service Service
}

func (h *controller) CreateUser(c *echo.Context) error {
	type CreateUserBody struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	var body CreateUserBody
	if err := c.Bind(&body); err != nil {
		return err
	}

	_, err := h.service.CreateOne(c.Request().Context(), queries.CreateUserParams(body))

	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, nil)
}

func (h *controller) GetUser(c *echo.Context) error {
	raw := c.Param("id")
	id, err := strconv.Atoi(raw)
	if err != nil {
		return err
	}

	user, err := h.service.GetByID(c.Request().Context(), int32(id))
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, user)
}

func (h *controller) RegisterRoutes(r *echo.Group) {
	r.POST("/", h.CreateUser)
	r.GET("/:id", h.GetUser)
}

func NewController(service Service) Controller {
	return &controller{service}
}
