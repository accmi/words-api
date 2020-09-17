package search

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

// FindParticularEntity to search words/expressions
func FindParticularEntity(c *gin.Context) {
	word := c.Query("search")
	url := fmt.Sprintf("https://wordsapiv1.p.rapidapi.com/words/%s", word)

	req, err := http.NewRequest("GET", url, nil)

	if err != nil && req != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})

		return
	}

	req.Header.Add("x-rapidapi-host", "wordsapiv1.p.rapidapi.com")
	req.Header.Add("x-rapidapi-key", "c2b2b0e05amsh98116736bec3314p145134jsnae6b4ddac8e6")

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})

		return
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})

		return
	}

	var result WholeResult

	if err := json.Unmarshal(body, &result); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}

	c.JSON(http.StatusOK, result)
}
