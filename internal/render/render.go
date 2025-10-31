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
	pages      map[string]*template.Template
	components *template.Template
	routes     map[string]*template.Template
}

func LoadTemplates() (*Template, error) {
	// load layouts and components into template tree
	layouts := template.Must(template.ParseGlob("web/layouts/*.html"))
	components := template.Must(template.ParseGlob("web/components/*.html"))

	// init pages [route]/template map
	pages := make(map[string]*template.Template)
	routes := make(map[string]*template.Template)

	homeFilePath, _ := filepath.Glob("web/routes/index.html")
	routeFilePaths, _ := filepath.Glob("web/routes/**/*.html")

	routeFiles := append(homeFilePath, routeFilePaths...)

	// build page and route templates
	for _, filePath := range routeFiles {

		// clone base layout templates for page templates
		pageTmpl, err := layouts.Clone()
		if err != nil {
			return nil, err
		}

		// create new Templates for route templates
		routeTmpl := template.New(filepath.Base(filePath))

		// merge component template tree to page/ route tree
		for _, ct := range components.Templates() {
			// AddParseTree will override exsiting templates with the same name
			if _, err := pageTmpl.AddParseTree(ct.Name(), ct.Tree); err != nil {
				return nil, err
			}

			// should be more refined, to only add parse tree on components which actually require it
			if _, err := routeTmpl.AddParseTree(ct.Name(), ct.Tree); err != nil {
				return nil, err
			}
		}

		// parse route specific file into merged layout/component tree
		if _, err := pageTmpl.ParseFiles(filePath); err != nil {
			return nil, err
		}

		// build and parse route templates
		routeTmpl, err = routeTmpl.ParseFiles(filePath)
		if err != nil {
			return nil, err
		}

		// store route template clone in map under route path name (blog/blog)
		routePath := strings.TrimPrefix(filePath, "web/routes/")
		routePath = routePath[:len(routePath)-len(filepath.Ext(routePath))]

		pages[routePath] = pageTmpl
		routes[routePath] = routeTmpl
	}

	return &Template{
		pages:      pages,
		components: components,
		routes:     routes,
	}, nil
}

func (t *Template) Render(w io.Writer, name string, data any, c echo.Context) error {
	if t == nil {
		return fmt.Errorf("render: templates not initialized")
	}

	if c.Request().Header.Get("HX-Request") == "true" {
		return t.renderRouteFragment(w, name, data)
	}

	return t.renderPage(w, name, data)
}

func (t *Template) renderRouteFragment(w io.Writer, name string, data any) error {
	if t == nil {
		return fmt.Errorf("render: templates not initialized")
	}

	routeKey := name

	tmpl, ok := t.routes[routeKey]
	if !ok || tmpl == nil {
		return fmt.Errorf("render: route template %q not found (available: %v)", routeKey, t.ListRouteRoutes())
	}

	return tmpl.ExecuteTemplate(w, name, data)
}

func (t *Template) renderPage(w io.Writer, name string, data any) error {
	if t == nil {
		return fmt.Errorf("render: templates not initialized")
	}

	routeKey := name

	tmpl, ok := t.pages[routeKey]
	if !ok || tmpl == nil {
		return fmt.Errorf("render: route template %q not found (available: %v)", routeKey, t.ListPageRoutes())
	}

	if tmpl.Lookup("base") == nil {
		return fmt.Errorf("render: base template not found for route %q", routeKey)
	}

	return tmpl.ExecuteTemplate(w, "base", data)
}

func (t *Template) ListRouteRoutes() []string {
	keys := make([]string, 0, len(t.routes))
	for k := range t.routes {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}

func (t *Template) ListPageRoutes() []string {
	keys := make([]string, 0, len(t.routes))
	for k := range t.pages {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}
