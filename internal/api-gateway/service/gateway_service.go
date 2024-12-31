package service

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type GatewayService interface {
	ForwardRequest(c *gin.Context, url string) (res *http.Response, err error)
}

func NewGatewayService() GatewayService {
	return &gatewayService{}
}

type gatewayService struct {
}

func (s *gatewayService) ForwardRequest(c *gin.Context, url string) (res *http.Response, err error) {
	// Create a new HTTP request
	req, err := http.NewRequest(c.Request.Method, url, c.Request.Body)
	if err != nil {
		return res, err
	}

	// Copy headers from the original request
	for k, v := range c.Request.Header {
		req.Header[k] = v
	}

	// Perform the request
	client := &http.Client{}
	res, err = client.Do(req)

	return res, err
}
