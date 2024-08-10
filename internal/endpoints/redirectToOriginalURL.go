package endpoints

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

func (c *LinkEndpoints) RedirectToOriginalURL(ctx *gin.Context) {
    shortURL := ctx.Param("shortURL")

    linkDto, err := c.usecase.GetByShortLink(shortURL)
    if err != nil || linkDto == nil {
        ctx.JSON(http.StatusNotFound, gin.H{"error": "Link not found"})
        return
    }

    ctx.HTML(http.StatusOK, "redirect.html", gin.H{"OriginalURL": linkDto.OriginalURL})
}
