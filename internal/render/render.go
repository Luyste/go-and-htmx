package render

import (
	"fmt"
	"html/template"
	"io"
	"path/filepath"
	"sort"
	"strings"

	"github.com/labstack/echo/v4"
)

type Template struct {
	templates map[string]*template.Template
}

func LoadTemplates() (*Template, error) {
	// load layouts and components into template tree
	rootLayout := template.Must(template.ParseFiles("web/routes/layout.html"))
	// rootPage := template.Must(template.ParseGlob("web/routes/index.html"))
	// routePages := template.Must(template.ParseGlob("web/routes/*/index.html"))

	// write recursive tree walker
	// routeFragments := template.Must(template.ParseGlob("web/routes/*/!{index, layout}.html"))

	components := template.Must(template.ParseGlob("web/components/*.html"))

	homeFilePath, _ := filepath.Glob("web/routes/*.html")
	routeFilePaths, _ := filepath.Glob("web/routes/**/*.html")

	routeFiles := append(homeFilePath, routeFilePaths...)

	// build templates

	// we are going to build two sets of templates which we will store in one big map.
	// 1. Full Page templates: page, blog/page, form/page
	// 2. Page content templates: index, blog/index, form/index
	// 3. Route Fragment templates: {fragmentName}, blog/{fragmentName}, form/{fragmentName}
	// 4. Route handlers now return two kinds, 1. full page, 2. page content on hx-swap request.
	pageMap := make(map[string](*template.Template))
	indexMap := make(map[string](*template.Template))
	// fragmentsMap := make(map[string](*template.Template))

	for _, filePath := range routeFiles {

		// clone base layout templates for page templates
		pageTmpl, err := rootLayout.Clone()
		if err != nil {
			return nil, err
		}

		// create new Templates for each route: [/blog], [/form], [/]
		// indexTmpl := template.New(filepath.Base(filePath))

		// merge component template tree to page/ route tree
		for _, ct := range components.Templates() {
			// AddParseTree will override exsiting templates with the same name
			if _, err := pageTmpl.AddParseTree(ct.Name(), ct.Tree); err != nil {
				return nil, err
			}

			// should be more refined, to only add parse tree on components which actually require it
			// if _, err := indexTmpl.AddParseTree(ct.Name(), ct.Tree); err != nil {
			// 	return nil, err
			// }
		}

		// parse route specific file into merged layout/component tree
		if _, err := pageTmpl.ParseFiles(filePath); err != nil {
			return nil, err
		}

		// build and parse route templates
		// indexTmpl, err = indexTmpl.ParseFiles(filePath)
		// if err != nil {
		// 	return nil, err
		// }

		// store route template clone in map under route path name (blog/blog)
		routePath := strings.TrimPrefix(filePath, "web/routes/")
		routePath = routePath[:len(routePath)-(len(filepath.Base(routePath)))]

		// store templates under dynamic keys built from routePath
		pageMap[routePath+"page"] = pageTmpl
		// indexMap[routePath+"/index"] = indexTmpl
	}

	templates := make(map[string]*template.Template, len(pageMap)+len(indexMap))
	for k, v := range pageMap {
		templates[k] = v
	}
	for k, v := range indexMap {
		templates[k] = v
	}

	return &Template{
		templates: templates,
	}, nil
}

func (t *Template) Render(w io.Writer, name string, data any, c echo.Context) error {
	logger := c.Logger()

	if t == nil {
		return fmt.Errorf("render: templates not initialized\n")
	}

	logger.Debugf("available templates: %v\n", t.ListTemplates())

	tmpl, ok := t.templates[name]
	logger.Debugf("selected template: %v\n", tmpl.Name())

	for _, t := range tmpl.Templates() {
		logger.Debugf("template templates: %v\n", t.Name())
	}

	if !ok || tmpl == nil {
		return fmt.Errorf("render: route template %q not found (available: %v)\n", name, t.ListTemplates())
	}

	return tmpl.ExecuteTemplate(w, name, data)
}

func (t *Template) ListTemplates() []string {
	keys := make([]string, 0, len(t.templates))
	for k := range t.templates {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}
