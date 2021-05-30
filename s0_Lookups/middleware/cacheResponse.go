package middleware

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/gin-contrib/cache/persistence"
	"github.com/gin-gonic/gin"
	"github.com/jaswanth-gorripati/PGK/s0_Lookups/models"
)

// bodyCacheWriter is used to cache responses in gin.
type bodyCacheWriter struct {
	gin.ResponseWriter
	cache      persistence.InMemoryStore
	requestURI string
}

// Write a JSON response to gin and cache the response.
func (w bodyCacheWriter) Write(b []byte) (int, error) {
	// Write the response to the cache only if a success code
	status := w.Status()

	if 200 <= status && status <= 299 {
		var ald models.AllLutData
		err := json.Unmarshal(b, &ald)
		if err != nil {
			log.Errorf("Failed to convert the data into ald , error %v", err)
		}
		w.cache.Set(w.requestURI, ald, 10*time.Minute)
	}

	// Then write the response to gin
	return w.ResponseWriter.Write(b)
}

// CacheCheck sees if there are any cached responses and returns
// the cached response if one is available.
func CacheCheck(cache persistence.InMemoryStore) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the ignoreCache parameter
		ignoreCache := strings.ToLower(c.Query("ignoreCache")) == "true"

		// See if we have a cached response
		log.Debugf("=================" + c.Request.RequestURI)
		response, exists := cache.Cache.Get(c.Request.RequestURI)
		//var ald models.AllLutData
		_, ok := response.(models.AllLutData)

		log.Debugf("Is interface in AllLutData format - ", ok)

		fmt.Println(!ignoreCache && exists, exists, ignoreCache)
		if !ignoreCache && exists && ok {

			log.Debug("Sending cache response")

			c.JSON(200, response.(models.AllLutData))
			c.Abort()
		} else if ignoreCache {
			c.Next()
		} else {
			// If not, pass our cache writer to the next middleware
			bcw := &bodyCacheWriter{cache: cache, requestURI: c.Request.RequestURI, ResponseWriter: c.Writer}
			c.Writer = bcw
			c.Next()
		}
	}
}
