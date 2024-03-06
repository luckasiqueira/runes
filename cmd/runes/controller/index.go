package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

/*
Index set our context.HTML to:
StatusOK sets a 200 status
Set index.html as page rendered
gin.H defines "variables" to use on rendered page
*/
func Index(context *gin.Context) {
	context.HTML(http.StatusOK, "index.html", gin.H{
		"Title": "Runes",
	})
}
