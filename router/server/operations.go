package server

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	"io/ioutil"
	"net/http"
	"github.com/tbuchaillot/dkvs/router/errors"
)

// GetKey handles the GET Key requests to the nodes.
func (s *Server) GetKey(c echo.Context) error {
	e := errors.Error{}
	key  :=  c.Param("key")
	if key  == ""{
		e.StatusCode = http.StatusBadRequest
		e.Message = "Invalid Params"
		c.JSON(e.StatusCode,e)

	}

	node,err := s.nodes.GetNode(key)
	if err != nil {
		e.StatusCode = http.StatusBadRequest
		e.Message = err.Error()
		return c.JSON(e.StatusCode,e)

	}

	value, err := s.cl.GetValue(node,key)
	if err != nil {
		e.StatusCode = 500
		e.Message = err.Error()
		return c.JSON(e.StatusCode,e)

	}
	response := echo.Map{}
	if err := json.Unmarshal(value, &response); err != nil {
		e.StatusCode = 500
		e.Message = err.Error()
		return c.JSON(e.StatusCode,e)
	}


	return c.JSON(http.StatusOK, response)
}

// GetKey handles the GET Key requests to the nodes.
func (s *Server) SetKey(c echo.Context) error {
	e := errors.Error{}
	key  :=  c.Param("key")
	if key  == ""{
		e.StatusCode = http.StatusBadRequest
		e.Message = "Invalid Params"
		return c.JSON(e.StatusCode,e)
	}

	if c.Request().Body == nil {
		e.StatusCode = http.StatusBadRequest
		e.Message = "Empty body"
		return c.JSON(e.StatusCode,e)
	}
	bodyBytes, _ := ioutil.ReadAll(c.Request().Body)

	node,err := s.nodes.GetNode(key)
	if err != nil {
		e.StatusCode = http.StatusBadRequest
		e.Message = err.Error()
		return c.JSON(e.StatusCode,e)
	}

	err = s.cl.SetValue(node,key,bodyBytes)
	if err != nil {
		e.StatusCode = 500
		e.Message = err.Error()
		return c.JSON(e.StatusCode,e)
	}

	response := echo.Map{}
	response["result"] = "Key "+key+ " saved."
	return c.JSON(http.StatusOK, response)
}
