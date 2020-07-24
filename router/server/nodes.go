package server

import (
	"github.com/labstack/echo/v4"
	"github.com/tbuchaillot/dkvs/router/errors"
	"net/http"
)

// GetKey handles the GET Key requests to the nodes.
func (s *Server) RegisterNode(c echo.Context) error {
	e := errors.Error{}
	u := &struct{
		Host string `json:"host"`
		Port int `json:"port"`
	}{}
	if err := c.Bind(u); err != nil {
		e.StatusCode = http.StatusBadRequest
		e.Message = "Invalid Body"
		return c.JSON(e.StatusCode,e)
	}

	if err := s.nodes.RegisterNode(u.Host,u.Port);err != nil{
		e.StatusCode = http.StatusInternalServerError
		e.Message =  err.Error()
		return c.JSON(e.StatusCode,e)
	}
	response := echo.Map{}
	response["result"] = "Node registered"

	return c.JSON(http.StatusOK, response)
}

