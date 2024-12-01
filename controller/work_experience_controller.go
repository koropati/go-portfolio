package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/koropati/go-portfolio/bootstrap"
	"github.com/koropati/go-portfolio/domain"
	"github.com/koropati/go-portfolio/internal/converter"
	"github.com/koropati/go-portfolio/internal/cryptos"
	"github.com/koropati/go-portfolio/internal/uploader"
	"github.com/koropati/go-portfolio/internal/validator"
)

type WorkExperienceController struct {
	WorkExperienceUsecase domain.WorkExperienceUsecase
	AccessTokenUsecase    domain.AccessTokenUsecase
	RefreshTokenUsecase   domain.RefreshTokenUsecase
	Config                *bootstrap.Config
	Cryptos               cryptos.Cryptos
	Validator             *validator.Validator
}

func (ctr *WorkExperienceController) Index(c *gin.Context) {
	c.HTML(http.StatusOK, "work_experience.tmpl", nil)
}

func (ctr *WorkExperienceController) Create(c *gin.Context) {

	var request domain.WorkExperience

	err := c.ShouldBind(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.JsonResponse{Message: err.Error(), Success: false})
		return
	}

	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.JsonResponse{Message: err.Error(), Success: false})
		return
	}

	filePath, err := uploader.UploadFile(c, file)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.JsonResponse{Message: err.Error(), Success: false})
		return
	} else {
		request.SetFileURL(filePath)
	}

	request.SetActive()

	err = request.GenerateID()
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.JsonResponse{Message: err.Error(), Success: false})
		return
	}

	err = ctr.Validator.Validate(request)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.JsonResponse{Message: err.Error(), Success: false})
		return
	}

	err = ctr.WorkExperienceUsecase.Create(c, request)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.JsonResponse{Message: err.Error(), Success: false})
		return
	}

	c.JSON(http.StatusOK, domain.JsonResponse{
		Message: "Success Create Data",
		Success: true,
	})
}

func (ctr *WorkExperienceController) Retrieve(c *gin.Context) {
	var filter domain.Filter
	err := c.BindQuery(&filter)

	if err != nil {
		c.JSON(http.StatusBadRequest, domain.JsonResponse{Message: err.Error(), Success: false})
		return
	}

	data, meta, err := ctr.WorkExperienceUsecase.Retrieve(c, filter)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.JsonResponse{Message: err.Error(), Success: false})
		return
	}

	c.JSON(http.StatusOK, domain.JsonResponse{
		Message: "Success Retrieve Data",
		Data:    data,
		Meta:    meta,
		Success: true,
	})
}

func (ctr *WorkExperienceController) Update(c *gin.Context) {

	var request domain.WorkExperience

	err := c.ShouldBind(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.JsonResponse{Message: err.Error(), Success: false})
		return
	}

	uuidString := c.Query("id")
	request.ID, err = converter.StringToUUID(uuidString)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.JsonResponse{Message: err.Error(), Success: false})
		return
	}

	err = ctr.Validator.Validate(request)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.JsonResponse{Message: err.Error(), Success: false})
		return
	}

	data, err := ctr.WorkExperienceUsecase.Update(c, request.ID, request)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.JsonResponse{Message: err.Error(), Success: false})
		return
	}

	c.JSON(http.StatusOK, domain.JsonResponse{
		Message: "Success Update Data",
		Data:    data,
		Success: true,
	})
}

func (ctr *WorkExperienceController) Delete(c *gin.Context) {
	uuidString := c.Query("id")
	id, err := converter.StringToUUID(uuidString)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.JsonResponse{Message: err.Error(), Success: false})
		return
	}

	err = ctr.WorkExperienceUsecase.Delete(c, id)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.JsonResponse{Message: err.Error(), Success: false})
		return
	}

	c.JSON(http.StatusOK, domain.JsonResponse{
		Message: "Success Delete Data",
		Success: true,
	})
}

func (ctr *WorkExperienceController) Get(c *gin.Context) {
	uuidString := c.Query("id")
	id, err := converter.StringToUUID(uuidString)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.JsonResponse{Message: err.Error(), Success: false})
		return
	}

	data, err := ctr.WorkExperienceUsecase.GetById(c, id)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.JsonResponse{Message: err.Error(), Success: false})
		return
	}

	c.JSON(http.StatusOK, domain.JsonResponse{
		Message: "Success Get Data",
		Data:    data,
		Success: true,
	})
}
