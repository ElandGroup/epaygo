package main

import (
	"epaygo"
	. "epaygo/core/common"
	"epaygo/core/helper"
	"net/http"

	"github.com/labstack/echo"
)

func DirectPayAL(c echo.Context) error {
	directPayDto := new(AlDirectPayDto)
	if err := c.Bind(directPayDto); err != nil {
		//return c.JSON(http.StatusBadRequest, APIResult{Success: false, Error: APIError{Code: 10012, Message: BadRequestMessage(directPayDto)}})
		return c.JSON(http.StatusBadRequest, helper.CheckRequestFormat(helper.MessageString(20004, "Object")))
	}
	directPayDto.OutTradeNo = helper.UuIdForPay(UuIdAlOutTradeNo)

	//payService := new(epaygo.AlPayService)
	payService, _ := epaygo.CreatePayment("AL")

	directPayDtoP := structToMap(directPayDto)

	if result, err := payService.DirectPay(directPayDtoP); err != nil {
		return c.JSON(http.StatusOK, APIResult{Success: false, Error: *err})
	} else {
		//c.JSON(http.StatusOK, APIResult{Success: true, Result: result})
		return c.JSON(http.StatusOK, APIResult{Success: true, Result: result})
		//c.String(http.StatusOK, result)
	}

}

func OrderQueryAL(c echo.Context) error {
	dto := new(AlOrderQueryDto)
	if err := c.Bind(dto); err != nil {
		//return c.JSON(http.StatusBadRequest, APIResult{Success: false, Error: APIError{Code: 10012, Message: BadRequestMessage(dto)}})
		return c.JSON(http.StatusBadRequest, helper.CheckRequestFormat(helper.MessageString(20004, "Object")))
	}

	payService, _ := epaygo.CreatePayment("AL")

	dtoP := structToMap(dto)
	if result, err := payService.OrderQuery(dtoP); err != nil {
		return c.JSON(http.StatusOK, APIResult{Success: false, Error: *err})
	} else {
		//c.JSON(http.StatusOK, APIResult{Success: true, Result: result})
		return c.JSON(http.StatusOK, APIResult{Success: true, Result: result})
		//c.String(http.StatusOK, result)
	}

}

func RefundAL(c echo.Context) error {
	dto := new(AlRefundDto)
	if err := c.Bind(dto); err != nil {
		// return c.JSON(http.StatusBadRequest, APIResult{Success: false, Error: APIError{Code: 10012, Message: BadRequestMessage(dto)}})
		return c.JSON(http.StatusBadRequest, helper.CheckRequestFormat(helper.MessageString(20004, "Object")))

	}

	dto.OutRequestNo = helper.UuIdForPay(UuIdAlRefundNo)

	payService, _ := epaygo.CreatePayment("AL")
	dtoP := structToMap(dto)
	//1.query transNo by OutTradeNo
	// q, _ := payService.OrderQueryAl(&AlOrderQueryDto{AlPayBaseDto: dto.AlPayBaseDto})
	// js, _ := simplejson.NewJson([]byte(q))
	// tradeNo, _ := js.Get(alConst.TradeNo).String()
	// dto.TradeNo = tradeNo
	if result, err := payService.Refund(dtoP); err != nil {
		return c.JSON(http.StatusOK, APIResult{Success: false, Error: *err})
	} else {
		//c.JSON(http.StatusOK, APIResult{Success: true, Result: result})
		return c.JSON(http.StatusOK, APIResult{Success: true, Result: result})
		//c.String(http.StatusOK, result)
	}

}

func ReverseAL(c echo.Context) error {
	dto := new(AlReverseDto)
	if err := c.Bind(dto); err != nil {
		// return c.JSON(http.StatusBadRequest, APIResult{Success: false, Error: APIError{Code: 10012, Message: BadRequestMessage(dto)}})
		return c.JSON(http.StatusBadRequest, helper.CheckRequestFormat(helper.MessageString(20004, "Object")))

	}

	payService, _ := epaygo.CreatePayment("AL")
	dtoP := structToMap(dto)
	if result, err := payService.OrderReverse(dtoP, 10); err != nil {
		return c.JSON(http.StatusOK, APIResult{Success: false, Error: *err})
	} else {
		//c.JSON(http.StatusOK, APIResult{Success: true, Result: result})
		return c.JSON(http.StatusOK, APIResult{Success: true, Result: result})
		//c.String(http.StatusOK, result)
	}

}
