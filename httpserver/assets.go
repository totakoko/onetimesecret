package httpserver

import (
	"errors"
	"fmt"
	"io/ioutil"
	"mime"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func (s *HTTPServer) loadStaticAssets() error {
	assetPaths, err := filepath.Glob(publicPath + "/**")
	if err != nil {
		return err
	}
	if len(assetPaths) == 0 {
		return errors.New("no assets were loaded")
	}

	for _, assetPath := range assetPaths {
		log.Debug().Msgf("Reading asset %s", assetPath)
		assetURL := strings.TrimPrefix(assetPath, publicPath)

		assetContent, err := ioutil.ReadFile(assetPath)
		if err != nil {
			return fmt.Errorf("could not read asset %s: %v", assetURL, err)
		}
		s.assetsCache[assetURL] = assetContent

		s.router.HEAD(assetURL, s.serveAsset)
		s.router.GET(assetURL, s.serveAsset)
		log.Debug().Msgf("Loaded asset %s from %s", assetURL, assetPath)
	}
	log.Debug().Msgf("Loaded %d assets", len(assetPaths))
	return nil
}

func (s *HTTPServer) serveAsset(c *gin.Context) {
	c.Status(http.StatusOK)
	assetPath := c.Request.URL.Path
	assetContent := s.assetsCache[assetPath]
	c.Writer.Header().Set("Content-Type", mime.TypeByExtension(filepath.Ext(assetPath)))
	c.Writer.Write(assetContent)
}
