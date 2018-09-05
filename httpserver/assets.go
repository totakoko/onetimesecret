package httpserver

import (
	"fmt"
	"io/ioutil"
	"mime"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"gitlab.com/totakoko/onetimesecret/helpers"
)

func (s *HTTPServer) loadStaticAssets() error {
	return filepath.Walk(publicPath, func(assetPath string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		assetURL := strings.TrimPrefix(assetPath, publicPath)

		assetContent, err := ioutil.ReadFile(assetPath)
		if err != nil {
			return fmt.Errorf("could not read asset %s: %v", assetURL, err)
		}
		s.assetsCache[assetURL] = assetContent

		s.router.HEAD(assetURL, s.serveAsset)
		s.router.GET(assetURL, s.serveAsset)
		log.Debug().Msgf("Loaded asset %s from %s", assetURL, assetPath)
		return nil
	})
}

func (s *HTTPServer) serveAsset(c *gin.Context) {
	c.Status(http.StatusOK)
	assetPath := c.Request.URL.Path
	assetContent := s.assetsCache[assetPath]
	c.Writer.Header().Set("Content-Type", mime.TypeByExtension(filepath.Ext(assetPath)))
	_, err := c.Writer.Write(assetContent)
	helpers.LogOnError(err)
}
