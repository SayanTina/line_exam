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

	c.Request.ParseMultipartForm(10 << 20)

	file, _, err := c.Request.FormFile("file")
	if err != nil {
		fmt.Println("Error Retrieving the File")
		fmt.Println(err)
		return
	}

	csvLines, err := csv.NewReader(file).ReadAll()
	if err != nil {
		fmt.Println(err)
	}

	up_status_amount := 0
	down_status_amount := 0

	ch := make(chan int)

	for _, line := range csvLines {
		go services.HealthCheckService(line[0], ch)
	}

	defer file.Close()

	for range csvLines {
		result := <-ch
		if result == 0 {
			down_status_amount++
		} else {
			up_status_amount++
		}
	}

	elapsed := time.Since(start)
	log.Printf("healthcheck2  took %s", elapsed)

	response := models.HealthCheckResp{
		ServiceAmount:     len(csvLines),
		UpServiceAmount:   up_status_amount,
		DownServiceAmount: down_status_amount,
		ExecutionTime:     elapsed,
	}

	c.JSON(http.StatusOK, response)
}
