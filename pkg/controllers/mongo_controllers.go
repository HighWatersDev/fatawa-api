package controllers

import (
	"net/http"
	"strconv"
	"strings"

	"fatawa-api/pkg/models"
	"fatawa-api/pkg/services"
	"github.com/gin-gonic/gin"
)

type FatwaController struct {
	fatwaService services.FatwaService
}

func NewFatwaController(fatwaService services.FatwaService) FatwaController {
	return FatwaController{fatwaService}
}

func (pc *FatwaController) CreateFatwa(ctx *gin.Context) {
	var fatwa *models.Fatwa

	if err := ctx.ShouldBindJSON(&fatwa); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	newFatwa, err := pc.fatwaService.CreateFatwa(fatwa)

	if err != nil {
		if strings.Contains(err.Error(), "title already exists") {
			ctx.JSON(http.StatusConflict, gin.H{"status": "fail", "message": err.Error()})
			return
		}

		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"status": "success", "data": newFatwa})
}

func (pc *FatwaController) UpdateFatwa(ctx *gin.Context) {
	fatwaId := ctx.Param("fatwaId")

	var fatwa *models.FatwaDb
	if err := ctx.ShouldBindJSON(&fatwa); err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	updatedFatwa, err := pc.fatwaService.UpdateFatwa(fatwaId, fatwa)
	if err != nil {
		if strings.Contains(err.Error(), "Id exists") {
			ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": err.Error()})
			return
		}
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": updatedFatwa})
}

func (pc *FatwaController) FindFatwaById(ctx *gin.Context) {
	fatwaId := ctx.Param("fatwaId")

	fatwa, err := pc.fatwaService.FindFatwaById(fatwaId)

	if err != nil {
		if strings.Contains(err.Error(), "Id exists") {
			ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": err.Error()})
			return
		}
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": fatwa})
}

func (pc *FatwaController) FindFatawa(ctx *gin.Context) {
	var page = ctx.DefaultQuery("page", "1")
	var limit = ctx.DefaultQuery("limit", "10")

	intPage, err := strconv.Atoi(page)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	intLimit, err := strconv.Atoi(limit)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	fatawa, err := pc.fatwaService.FindFatawa(intPage, intLimit)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "results": len(fatawa), "data": fatawa})
}

func (pc *FatwaController) DeleteFatwa(ctx *gin.Context) {
	fatwaId := ctx.Param("fatwaId")

	err := pc.fatwaService.DeleteFatwa(fatwaId)

	if err != nil {
		if strings.Contains(err.Error(), "Id exists") {
			ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": err.Error()})
			return
		}
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}
