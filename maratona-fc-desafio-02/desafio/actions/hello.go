package actions

import (
    "net/http"
    
	"github.com/gobuffalo/buffalo"
)

// HelloIndex default implementation.
func HelloIndex(c buffalo.Context) error {
	return c.Render(http.StatusOK, r.HTML("hello/index.html"))
}

