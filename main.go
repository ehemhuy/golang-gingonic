package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/gocolly/colly/v2"
	"github.com/joho/godotenv"
)

type ReqBody struct {
	Url  string `json:"url"`
	Page string
}

type ReqQuery struct {
	Url  string `json:"url"`
	Page int    `json:"page" form:"page"`
	Size int    `json:"size" form:"size"`
}

type Zone struct {
	Url   string `json:"url"`
	Title string `json:"title"`
}

type Image struct {
	Uri string `json:"uri"`
}

func main() {
	var router = gin.Default()
	godotenv.Load()
	port := os.Getenv("PORT")
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, "ok")

	})

	//get khu vuc
	router.GET("/zone/list", func(c *gin.Context) {
		var zones []interface{}

		collector := colly.NewCollector()

		collector.OnHTML(".node-title a", func(element *colly.HTMLElement) {
			var zone Zone
			zone.Url = "https://checkerviet.me" + element.Attr("href")
			zone.Title = element.Text
			zones = append(zones, zone)
		})

		collector.OnScraped(func(r *colly.Response) {

			c.JSON(http.StatusOK, zones)
		})

		collector.Visit("https://checkerviet.me/forums/gai-goi-ha-noi.6/")
	})

	// get chi tiet 1 khu vuc
	router.POST("/zone/detail", func(ctx *gin.Context) {
		var zone Zone
		ctx.ShouldBindJSON(&zone)
		var zones []Zone

		c := colly.NewCollector()

		// Find and visit all links
		c.OnHTML(".block-body .structItem-title a:last-child", func(e *colly.HTMLElement) {
			zone := Zone{
				Url:   "https://checkerviet.me" + e.Attr("href"),
				Title: e.Text,
			}
			zones = append(zones, zone)
		})

		c.OnScraped(func(r *colly.Response) {
			ctx.JSON(http.StatusOK, zones)
		})

		c.Visit(zone.Url)
	})

	// get anh tu link
	router.POST("/images", func(ctx *gin.Context) {
		var body map[string]string
		var images []Image
		ctx.ShouldBindJSON(&body)
		c := colly.NewCollector()

		c.OnHTML("img.bbImage", func(h *colly.HTMLElement) {
			var image = Image{
				Uri: h.Attr("data-url"),
			}
			images = append(images, image)
		})

		c.OnScraped(func(r *colly.Response) {
			ctx.JSON(http.StatusOK, images)
		})

		c.Visit(body["url"])
	})

	router.Run(":" + port)
}
