package job

import (
	"net/http"

	"github.com/absormu/dans_test/app/entity"
	md "github.com/absormu/dans_test/app/middleware"
	repojob "github.com/absormu/dans_test/app/repository/job"
	pg "github.com/absormu/dans_test/pkg/pagination"
	lg "github.com/absormu/dans_test/pkg/response"
	resp "github.com/absormu/dans_test/pkg/response"
	sdk "github.com/absormu/dans_test/pkg/sdk"
	"github.com/labstack/echo/v4"
)

func GetJobList(c echo.Context, extractToken entity.ExtractToken) (e error) {
	logger := md.GetLogger(c)
	logger.WithField("request", extractToken).Info("usecase: GetJobList")

	description := c.QueryParam("description")
	location := c.QueryParam("location")
	fullTime := c.QueryParam("full_time")

	meta, e := pg.Pagination(c)
	if e != nil {
		logger.WithField("error", e.Error()).Error("Catch error Pagination")
		e = resp.CustomError(c, http.StatusBadRequest, sdk.ERR_PARAM_ILLEGAL,
			lg.Language{Bahasa: nil, English: "Bad Request"}, nil, nil)
		return
	}

	var jobLists []entity.JobList
	jobLists, e = repojob.RequestJobList(c, description, location, fullTime)
	if e != nil {
		logger.WithField("error", e.Error()).Error("Catch error RequestJobList request")
		return
	}

	total := int64(len(jobLists))
	var responseData []entity.JobList
	var data entity.JobList
	for _, jobList := range jobLists {
		data.ID = jobList.ID
		data.Type = jobList.Type
		data.Url = jobList.Url
		data.CreatedAt = jobList.CreatedAt
		data.Company = jobList.Company
		data.CompanyUrl = jobList.CompanyUrl
		data.Location = jobList.Location
		data.Title = jobList.Title
		data.Description = jobList.Description
		data.HowToApply = jobList.HowToApply
		data.CompanyLogo = jobList.CompanyLogo
		responseData = append(responseData, data)
	}

	metaPagination := pg.GenerateMeta(c, total, meta.Limit, meta.Page, meta.Offset, meta.Pagination, nil)

	e = resp.CustomError(c, http.StatusOK, sdk.ERR_SUCCESS,
		lg.Language{Bahasa: "Sukses", English: "Success"}, metaPagination, responseData)
	return
}
