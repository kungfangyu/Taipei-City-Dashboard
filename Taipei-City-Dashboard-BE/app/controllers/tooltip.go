package controllers

import (
	"TaipeiCityDashboardBE/app/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetTooltipData(c *gin.Context) {
    district := c.Param("district")

    data, err := models.FetchTooltipDataByDistrict(district)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "error": "資料庫查詢失敗",
        })
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "district": district,
        "data":     data,
    })
}