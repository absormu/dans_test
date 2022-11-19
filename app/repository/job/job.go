package job

import (
	"encoding/json"

	"github.com/absormu/dans_test/app/entity"
	md "github.com/absormu/dans_test/app/middleware"
	cm "github.com/absormu/dans_test/pkg/configuration"
	sdk "github.com/absormu/dans_test/pkg/sdk"
	"github.com/labstack/echo/v4"
)

func RequestJobList(c echo.Context, description, location, fullTime string) (res []entity.JobList, e error) {

	logger := md.GetLogger(c)
	logger.WithField("request", "").Info("repository: RequestJobList")

	headers := map[string]string{}
	headers["MsgID"] = "ceocececggg"

	queryParams := map[string]string{}
	if description != "" {
		queryParams["description"] = description
	}
	if location != "" {
		queryParams["location"] = location
	}
	if fullTime == "true" {
		queryParams["type"] = "Full Time"
	}

	rawResponse, e := sdk.RawGetRequest(logger, cm.Config.JobListUrl, cm.Config.Timeout, queryParams)
	if e != nil {
		return
	}

	if e = json.Unmarshal(rawResponse, &res); e != nil {
		return
	}

	return
}
