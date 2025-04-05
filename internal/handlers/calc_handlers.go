package calc_handlers

import (
	"go-api-calc/api"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CalcServer struct{}

func NewCalcServer() CalcServer {
	return CalcServer{}
}

func (CalcServer) AddNumbers(c *gin.Context) {
	var addRequest api.CalculationRequestNumbers
	if err := c.ShouldBindBodyWithJSON(&addRequest); err != nil {
		c.JSON(http.StatusBadRequest, api.Error{Message: "400 - Bad Request"})
		return
	}

	c.JSON(http.StatusOK, api.Result{Result: addRequest.Number1 + addRequest.Number2})

}

func (CalcServer) DivideNumbers(c *gin.Context) {
	var divideRequest api.DivisionRequest
	if err := c.ShouldBindBodyWithJSON(&divideRequest); err != nil {
		c.JSON(http.StatusBadRequest, api.Error{Message: "400 - Bad Request"})
		return
	}

	if divideRequest.Divisor == 0 {
		c.JSON(http.StatusBadRequest, api.Error{Message: "Division by zero is not allowed."})
		return
	}

	c.JSON(http.StatusOK, api.Result{Result: divideRequest.Dividend / divideRequest.Divisor})
}

func (CalcServer) MultiplyNumbers(c *gin.Context) {
	var multiplyRequest api.CalculationRequestNumbers
	if err := c.ShouldBindBodyWithJSON(&multiplyRequest); err != nil {
		c.JSON(http.StatusBadRequest, api.Error{Message: "400 - Bad Request"})
		return
	}

	c.JSON(http.StatusOK, api.Result{Result: multiplyRequest.Number1 * multiplyRequest.Number2})
}

func (CalcServer) SubtractNumbers(c *gin.Context) {
	var subtractRequest api.CalculationRequestNumbers
	if err := c.ShouldBindBodyWithJSON(&subtractRequest); err != nil {
		c.JSON(http.StatusBadRequest, api.Error{Message: "400 - Bad Request"})
		return
	}

	c.JSON(http.StatusOK, api.Result{Result: subtractRequest.Number1 - subtractRequest.Number2})
}

func (CalcServer) SumNumbers(c *gin.Context) {
	var sumNumbers api.SumRequest
	if err := c.ShouldBindBodyWithJSON(&sumNumbers); err != nil {
		c.JSON(http.StatusBadRequest, api.Error{Message: "400 - Bad Request"})
		return
	}

	result := api.Result{Result: 0}

	for _, v := range sumNumbers {
		result.Result += v
	}

	c.JSON(http.StatusOK, result)
}
