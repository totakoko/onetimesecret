package httpserver

import (
	"html/template"
	"io/ioutil"
	"path"
	"path/filepath"

	"github.com/rs/zerolog/log"
)

var (
	templatesCache = map[string]*template.Template{}
)

// list and parse the files in the templates directory
func init() {
	templatePaths, err := filepath.Glob(".build/templates/*.html")
	if err != nil {
		log.Fatal().Msg(err.Error())
	}

	for _, templatePath := range templatePaths {
		templateNameWithExt := path.Base(templatePath)
		templateName := templateNameWithExt[:len(templateNameWithExt)-len(".html")] // removes .html

		templateContent, err := ioutil.ReadFile(templatePath)
		if err != nil {
			log.Fatal().Msgf("could not parse template %s: %v", templateName, err)
		}

		htmlTemplate, err := template.New("").Parse(string(templateContent))
		if err != nil {
			log.Fatal().Msgf("could not parse template %s: %v", templateName, err)
		}

		log.Debug().Msgf("Loaded template %s from %s", templateName, templatePath)
		templatesCache[templateName] = htmlTemplate
	}

	if len(templatePaths) == 0 {
		log.Fatal().Msg("No templates were loaded!")
	}
	log.Debug().Msgf("Loaded %d templates", len(templatePaths))
}
