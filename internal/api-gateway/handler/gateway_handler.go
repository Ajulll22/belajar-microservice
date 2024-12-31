package handler

import (
	"fmt"
	"io"

	"github.com/Ajulll22/belajar-microservice/internal/api-gateway/config"
	"github.com/Ajulll22/belajar-microservice/internal/api-gateway/service"
	"github.com/Ajulll22/belajar-microservice/pkg/handling"
	"github.com/gin-gonic/gin"
)

func (h *gatewayHandler) UserProxy(c *gin.Context) {
	url := fmt.Sprintf("http://user-service:%s%s", h.cfg.USER_SERVICE_PORT, c.Request.URL.Path)

	res, err := h.gatewayService.ForwardRequest(c, url)
	if err != nil {
		res := handling.ResponseError(c, err)
		c.JSON(res.Code, res)
		return
	}
	defer res.Body.Close()

	// Copy response headers and status code
	for k, v := range res.Header {
		c.Writer.Header()[k] = v
	}
	c.Writer.WriteHeader(res.StatusCode)

	io.Copy(c.Writer, res.Body)
}

func (h *gatewayHandler) MediaProxy(c *gin.Context) {
	url := fmt.Sprintf("http://media-service:%s%s", h.cfg.MEDIA_SERVICE_PORT, c.Request.URL.Path)

	res, err := h.gatewayService.ForwardRequest(c, url)
	if err != nil {
		res := handling.ResponseError(c, err)
		c.JSON(res.Code, res)
		return
	}
	defer res.Body.Close()

	// Copy response headers and status code
	for k, v := range res.Header {
		c.Writer.Header()[k] = v
	}
	c.Writer.WriteHeader(res.StatusCode)

	io.Copy(c.Writer, res.Body)
}

func (h *gatewayHandler) ProductProxy(c *gin.Context) {
	url := fmt.Sprintf("http://product-service:%s%s", h.cfg.PRODUCT_SERVICE_PORT, c.Request.URL.Path)

	res, err := h.gatewayService.ForwardRequest(c, url)
	if err != nil {
		res := handling.ResponseError(c, err)
		c.JSON(res.Code, res)
		return
	}
	defer res.Body.Close()

	// Copy response headers and status code
	for k, v := range res.Header {
		c.Writer.Header()[k] = v
	}
	c.Writer.WriteHeader(res.StatusCode)

	io.Copy(c.Writer, res.Body)
}

type gatewayHandler struct {
	cfg            config.Config
	gatewayService service.GatewayService
}

func NewGatewayHandler(cfg config.Config, gatewayService service.GatewayService) gatewayHandler {
	return gatewayHandler{cfg, gatewayService}
}
