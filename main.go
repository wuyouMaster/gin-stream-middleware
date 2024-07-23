package ginStreamMiddleware

import (
	"bytes"
	"fmt"
	"time"
	"io"
	"net/http"
	"github.com/gin-gonic/gin"
)

type Callback interface {
	Call([]byte) ([]byte, error)
	GetName() string
}


type StreamWriter struct {
	gin.ResponseWriter
	Buffer  *bytes.Buffer
	Context *gin.Context
	Callback []Callback
	Flusher http.Flusher
}

func (r StreamWriter) Write(b []byte) (int, error) {
	resp := b
	var err error
	for _, callback := range r.Callback {
		resp, err = callback.Call(resp)
		if err != nil {
			return 0, fmt.Errorf("error occurred in '%s' stream middleware， reason: %v", callback.GetName(), err)
		}
	}
	written, err := r.ResponseWriter.Write(resp)
	r.Flush()
	return written, err
}

func (r *StreamWriter) Flush() {
	r.Flusher.Flush()
}


func Register(callbacks []Callback) func(c *gin.Context) {
	return func(c *gin.Context) {
		buffer := new(bytes.Buffer)
		streamWriter := &StreamWriter{ResponseWriter: c.Writer, Buffer: buffer, Context: c, Callback: callbacks, Flusher: c.Writer.(http.Flusher)}
		c.Writer = streamWriter
		c.Next()
	}
}

type CallbackInstance struct {
	Callback
	Name string
}

func (c CallbackInstance) Call(b []byte)([]byte, error) {
	fmt.Printf("call >>>>>, message: %s", string(b))
	return b, nil
}

func (c CallbackInstance) GetName() string {
	return c.Name
}