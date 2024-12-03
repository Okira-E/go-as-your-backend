package utils

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/org/example/opt"
)

type Where struct {
	And map[string]any `json:"and"`
	Or  map[string]any `json:"or"`
}

type Filter struct {
	Where   Where    `json:"where"`
	GroupBy []string `json:"group_by"`
	OrderBy []string `json:"order_by"`
}

type Response struct {
	Success bool        `json:"success"`
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func Err(c *fiber.Ctx, code int, message string, data any) error {
	response := Response{
		Success: false,
		Code:    code,
		Message: message,
		Data:    data,
	}

	return c.Status(code).JSON(response)
}

func Ok(c *fiber.Ctx, code int, message string, data any) error {
	response := Response{
		Success: true,
		Code:    code,
		Message: message,
		Data:    data,
	}

	return c.Status(code).JSON(response)
}

// ParseLimitAndOffset extracts both limit and offset params while providing validation
// and a default value for both if they were not provided.
func ParseLimitAndOffset(params url.Values) (int, int, error) {
	var limit int = 100
	var offset int = 0

	// Parse the limit.
	if value := params.Get("limit"); value != "" {
		valueInt, err := strconv.Atoi(value)
		if err != nil {
			return 0, 0, fmt.Errorf("Failed to convert limit to an int. %s", err)
		}

		if valueInt <= 100 {
			limit = valueInt
		}
	}

	// Parse the offset.
	if value := params.Get("offset"); value != "" {
		valueInt, err := strconv.Atoi(value)
		if err != nil {
			return 0, 0, fmt.Errorf("Failed to convert offset to an int. %s", err)
		}

		offset = valueInt
	}

	return limit, offset, nil
}

// ParseFilterClause parses the filter param in an endpoint.
//
// Example:
//
// "http://localhost:3200?filter={ \"where\": {\"id\": \"123-456-789\"}, \"order_by\": [\"created_by\"] }&limit=50&offset=10"
//
// Returns:
//
// Filter: {map[id:123-456-789] [] [created_by]}
func ParseFilterClause(params url.Values) (opt.Option[Filter], error) {
	filterStr := params.Get("filter")
	if filterStr == "" {
		return opt.None[Filter](), nil
	}

	var filterObj Filter
	err := json.Unmarshal([]byte(filterStr), &filterObj)
	if err != nil {
		return opt.None[Filter](), fmt.Errorf("Failed to unmarshal the where object. %s", err)
	}

	return opt.Some(filterObj), nil
}

