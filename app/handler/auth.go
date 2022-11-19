package handler

import (
	"net/http"

	usecaseauth "github.com/absormu/dans_test/app/usecase/auth"
	"github.com/labstack/echo/v4"

	"github.com/absormu/dans_test/app/entity"
	md "github.com/absormu/dans_test/app/middleware"
	lg "github.com/absormu/dans_test/pkg/response"
	resp "github.com/absormu/dans_test/pkg/response"
	sdk "github.com/absormu/dans_test/pkg/sdk"
)

func LoginHandler(c echo.Context) (e error) {
	logger := md.GetLogger(c)
	logger.Info("handler: LoginHandler")

	req := entity.Auth{}
	if e = c.Bind(&req); e != nil {
		logger.WithField("error", e.Error()).Error("Catch error bind request")
		e = resp.CustomError(c, http.StatusBadRequest, sdk.ERR_PARAM_ILLEGAL,
			lg.Language{Bahasa: nil, English: e.Error()}, nil, nil)
		return
	}

	e = usecaseauth.Login(c, req)

	return
}
