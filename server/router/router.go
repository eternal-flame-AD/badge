package router

import (
	"strconv"

	"github.com/eternal-flame-AD/badge/handler"
	"github.com/eternal-flame-AD/badge/internal/shields"
	"github.com/gin-gonic/gin"
)

var Router = gin.Default()

func badgeHandler(b *shields.B, err error) func(c *gin.Context) {
	var badge shields.B
	if err != nil {
		badge = shields.B{
			Subject: "error",
			Status:  err.Error(),
			Color:   "red",
		}
	} else {
		badge = *b
	}
	return func(c *gin.Context) {
		c.Header("cache-control", "max-age=30")
		if label := c.Query("label"); label != "" {
			badge.Subject = label
		}
		if style := c.Query("style"); style != "" {
			badge.Style = style
		}
		if logo := c.Query("logo"); logo != "" {
			badge.Logo = &shields.Logo{
				URLOrName: logo,
			}
			if logoColor := c.Query("logoColor"); logoColor != "" {
				badge.Logo.Color = logoColor
			}
			if logoWidth := c.Query("logoWidth"); logoWidth != "" {
				badge.Logo.Width, _ = strconv.Atoi(logoWidth)
			}
		}
		if links, ok := c.GetQueryArray("link"); ok {
			if len(links) == 1 {
				badge.Link = &shields.Link{
					L: links[0],
				}
			} else {
				badge.Link = &shields.Link{
					L: links[0],
					R: links[1],
				}
			}
		}
		if colorA := c.Query("colorA"); colorA != "" {
			badge.ColorLOverride = colorA
		}
		if colorB := c.Query("colorB"); colorB != "" {
			badge.Color = colorB
		}
		if maxAge := c.Query("maxAge"); maxAge != "" {
			maxAgeInt, _ := strconv.Atoi(maxAge)
			if maxAgeInt >= 30 {
				c.Header("cache-control", "max-age="+maxAge)
			}
		}

		c.Redirect(302, shields.URL(badge))
	}
}

func init() {
	img := Router.Group("/img")
	{
		img.GET("/css", func(c *gin.Context) {
			selector := c.Query("selector")
			url := c.Query("url")
			badgeHandler(handler.CSSSelector(selector, url))(c)
		})

		greasyFork := img.Group("/greasyfork")
		{
			greasyFork.GET("/install/:scriptname", func(c *gin.Context) {
				scriptname := c.Param("scriptname")
				badgeHandler(handler.GreasyForkTotalInstall(scriptname))(c)
			})

			greasyFork.GET("/version/:scriptname", func(c *gin.Context) {
				scriptname := c.Param("scriptname")
				badgeHandler(handler.GreasyForkVersion(scriptname))(c)
			})
		}
	}
}
