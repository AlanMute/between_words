package main

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.Static("/static", "./static")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})

	r.POST("/combine", func(c *gin.Context) {
		a := c.PostForm("input1")
		b := c.PostForm("input2")
		result := between(a, b)
		c.JSON(http.StatusOK, gin.H{"result": result})
	})

	r.Run(":8080")

}

func between(a string, b string) string {
	aw := strings.Fields(a)
	bw := strings.Fields(b)

	al := len(aw)
	bl := len(bw)
	m := al

	if m < bl {
		m = bl
	}

	var ans []string

	for i := 0; i < m; i++ {
		if al > i {
			ans = append(ans, aw[i])
		}
		if bl > i {
			ans = append(ans, bw[i])
		}
	}
	out := ""

	for _, w := range ans {
		out += w + " "
	}

	out = strings.TrimSpace(out)

	return out
}
