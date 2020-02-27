package main

import (
	"reflect"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"gopkg.in/go-playground/validator.v9"
)

type Booking struct {
	CheckIn  time.Time `form:"check_in"  binding:"required,bookabledate" time_format:"2006-01-01"`
	CheckOut time.Time `form:"check_out"  binding:"required,gtfield=checkIn" time_format:"2006-01-01"`
}

func customFunc(v *validator.Validate, topStruct reflect.Value, currentStructOrField reflect.Value, field reflect.Value, fieldType reflect.Type, fieldKind reflect.Kind, param string) bool {
	if date, ok := field.Interface().(time.Time); ok {
		today := time.Now()
		if date.Unix() > today.Unix() {
			return true
		}
	}
	return false
}

func main() {
	r := gin.Default()
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("bookabledate", customFunc)
	}
	r.GET("/bookable", testHandler)
	r.Run()
}

func testHandler(c *gin.Context) {
	var b Booking
	if err := c.ShouldBind(&b); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "ok!", "booking": b})
}
