package main

import (
    "net/http"
    "math/rand"
    "strconv"
    "github.com/gin-gonic/gin"
)

func main() {
    router := gin.Default()
    router.POST("/gen_snils", genSnils)
    router.GET("/gen_snils", genSnils)
    router.Run("localhost:8080")
}

func genSnils(c *gin.Context) {
    var snils string
    var thesum int = 0
    var digit int = 0
    var remainder int = 0
    for i := 9; i > 0; i-- {
        digit = rand.Intn(10)
        thesum = thesum + digit
        if (i == 6) || (i == 3) {
            snils = snils + "-" + strconv.Itoa(digit)
        } else {
            snils = snils + strconv.Itoa(digit)
        }
    }
    remainder = thesum % 101
    if (remainder == 100) || (remainder == 0) {
        snils = snils + "00"
    } else {
        if (len(strconv.Itoa(remainder))==1) {
            snils = snils + " " + "0"+ strconv.Itoa(remainder)
        } else {
            snils = snils + " " + strconv.Itoa(remainder)
        }
    }
    c.IndentedJSON(http.StatusOK, snils)
}
