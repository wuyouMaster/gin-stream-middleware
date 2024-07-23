package main

import (
	"fmt"
	"time"
	"io"
	"github.com/gin-gonic/gin"
	gsm "github.com/wuyouMaster/gin-stream-middleware"
)

type MyCallbackInstance struct {
	gsm.CallbackInstance
}

func (c MyCallbackInstance) Call(b []byte) ([]byte, error) {
	fmt.Printf("call2 >>>>>, message: %s", string(b))
	return b, nil
}


func main() {
	var middlewareList []gsm.Callback
	instance := gsm.CallbackInstance{Name: "hello world"}

	middlewareList = append(middlewareList, instance)
	r := gin.Default()

	r.Use(gsm.Register(middlewareList))
	r.GET("/stream", func(c *gin.Context) {
		w := c.Writer
		c.Stream(func(wc io.Writer) bool {
			for i := 0; i <= 10; i++ {
				fmt.Fprintf(w, "message %d\n", i)
				time.Sleep(time.Second)
			}
			return false
		})
	})

	r.Run(":8090")
}