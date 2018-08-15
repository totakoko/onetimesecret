package httpserver

import (
	"errors"
	"fmt"
	"html/template"
	"io/ioutil"
	"path"
	"path/filepath"

	"github.com/rs/zerolog/log"
)

func (s *HTTPServer) loadTemplates() error {
	templatePaths, err := filepath.Glob(templatesPath + "/*.html")
	if err != nil {
		return err
	}
	if len(templatePaths) == 0 {
		return errors.New("no templates were loaded")
	}

	for _, templatePath := range templatePaths {
		templateNameWithExt := path.Base(templatePath)
		templateName := templateNameWithExt[:len(templateNameWithExt)-len(".html")] // removes .html

		templateContent, err := ioutil.ReadFile(templatePath)
		if err != nil {
			return fmt.Errorf("could not read template %s: %v", templateName, err)
		}

		htmlTemplate, err := template.New("").Parse(string(templateContent))
		if err != nil {
			return fmt.Errorf("could not parse template %s: %v", templateName, err)
		}

		log.Debug().Msgf("Loaded template %s from %s", templateName, templatePath)
		s.templatesCache[templateName] = htmlTemplate
	}
	log.Debug().Msgf("Loaded %d templates", len(templatePaths))
	return nil
}
