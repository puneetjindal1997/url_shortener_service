package router

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc func(*gin.Context)
}

type routes struct {
	router *gin.Engine
}

type Routes []Route

func (r routes) UrlShortner(rg *gin.RouterGroup) {
	orderRouteGrouping := rg.Group("/url")
	for _, route := range urlShortner {
		switch route.Method {
		case http.MethodGet:
			orderRouteGrouping.GET(route.Pattern, route.HandlerFunc)
		case http.MethodPost:
			orderRouteGrouping.POST(route.Pattern, route.HandlerFunc)
		case http.MethodOptions:
			orderRouteGrouping.OPTIONS(route.Pattern, route.HandlerFunc)
		case http.MethodPut:
			orderRouteGrouping.PUT(route.Pattern, route.HandlerFunc)
		case http.MethodDelete:
			orderRouteGrouping.DELETE(route.Pattern, route.HandlerFunc)
		default:
			orderRouteGrouping.GET(route.Pattern, func(c *gin.Context) {
				c.JSON(200, gin.H{
					"result": "Specify a valid http method with this route.",
				})
			})
		}
	}
}

func ClientRoutes() {

	r := routes{
		router: gin.Default(),
	}
	v1 := r.router.Group(os.Getenv("API_VERSION"))
	r.UrlShortner(v1)

	if err := r.router.Run(":" + os.Getenv("PORT")); err != nil {
		log.Println("Failed to run server: %v", err)
	}
}
