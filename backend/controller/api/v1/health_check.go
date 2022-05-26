package v1

import (
	"encoding/csv"
	"fmt"

	"log"

	"net/http"
	"time"

	"github.com/chaiyo/line-exam/models"
	"github.com/chaiyo/line-exam/services"

	"github.com/gin-gonic/gin"
)

func HealthCheck(c *gin.Context) {

	start := time.Now()

	//10 << 20 magaic number for support file size (10MB)
	c.Request.ParseMultipartForm(10 << 20)

	file, _, err := c.Request.FormFile("file")
	if err != nil {
		fmt.Println("Error Retrieving the File")
		fmt.Println(err)
		c.Error(err).SetType(gin.ErrorTypeAny)
		c.JSON(400, map[string]interface{}{"message": err.Error()})
		return
	}

	csvLines, err := csv.NewReader(file).ReadAll()
	if err != nil {
		fmt.Println("Error Reading the File")
		fmt.Println(err)
		c.JSON(400, map[string]interface{}{"message": err.Error()})
		return
	}

	up_status_amount := 0
	down_status_amount := 0

	ch := make(chan int)

	for _, line := range csvLines {
		go services.HealthCheckService(line[0], ch)
	}

	defer file.Close()

	for range csvLines {
		switch result := <- ch; result {
		case 0:
			down_status_amount++
		case 1:
			up_status_amount++
		default:
			//have some cell contain space or empty cell
		}
	}

	elapsed := time.Since(start)
	log.Printf("healthcheck2  took %s", elapsed)

	response := models.HealthCheckResp{
		ServiceAmount:     up_status_amount + down_status_amount,
		UpServiceAmount:   up_status_amount,
		DownServiceAmount: down_status_amount,
		ExecutionTime:     elapsed,
	}

	c.JSON(http.StatusOK, response)
}
