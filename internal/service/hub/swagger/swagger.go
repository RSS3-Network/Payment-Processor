// Modified from https://github.com/shanbay/gobay/blob/master/echo/swagger/swagger.go

package swagger

import (
	"bytes"
	"fmt"
	"html/template"
	"net/http"
	"path"
	"strings"

	"github.com/labstack/echo/v4"
)

type Opts func(*config)

// func SetSwaggerHost(host string) Opts {
//	return func(o *config) {
//		o.SwaggerHost = host
//	}
//}
//
// func SetSwaggerAuthorizer(f func(*http.Request) bool) Opts {
//	return func(o *config) {
//		o.Authorizer = f
//	}
//}
//
// func SetSwaggerIsHTTPS(b bool) Opts {
//	return func(o *config) {
//		o.IsHTTPS = b
//	}
//}

// SwaggerOpts configures the Doc middlewares
type config struct {
	// SpecURL the url to find the spec for
	SpecURL string
	// SwaggerHost for the js that generates the swagger ui site, defaults to: http://petstore.swagger.io/
	SwaggerHost string
	// When this return value is false, 403 will be responsed.
	Authorizer func(*http.Request) bool

	IsHTTPS bool
}

func prepare(basePath string, cfg *config, swaggerJSON []byte) (string, string, []byte) {
	docPath := path.Join(basePath, "apidocs")

	// swagger html
	tmpl := template.Must(template.New("swaggerdoc").Parse(swaggerTemplateV2))
	buf := bytes.NewBuffer(nil)
	_ = tmpl.Execute(buf, cfg)
	uiHTML := buf.String()

	// swagger json
	responseSwaggerJSON := swaggerJSON
	if cfg.IsHTTPS {
		responseSwaggerJSON = []byte(strings.Replace(
			string(swaggerJSON),
			`"schemes": [
    "http"
  ],`,
			`"schemes": [
    "https"
  ],`,
			1))
	}

	return docPath, uiHTML, responseSwaggerJSON
}

func strIn(target string, source ...string) bool {
	for _, s := range source {
		if target == s {
			return true
		}
	}

	return false
}

// Doc creates a middleware to serve a documentation site for a swagger spec.
// This allows for altering the spec before starting the http listener.
func Doc(basePath string, swaggerJSON []byte, opts ...Opts) echo.MiddlewareFunc {
	cfg := &config{
		SpecURL:     path.Join(basePath, "swagger.json"),
		SwaggerHost: "https://petstore.swagger.io",
	}
	for _, opt := range opts {
		opt(cfg)
	}

	docPath, uiHTML, responseSwaggerJSON := prepare(basePath, cfg, swaggerJSON)

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			reqPath := c.Request().URL.Path
			if strIn(reqPath, basePath, docPath, cfg.SpecURL) {
				if cfg.Authorizer != nil && !cfg.Authorizer(c.Request()) {
					return c.String(403, "Forbidden")
				}

				switch reqPath {
				case docPath:
					return c.HTML(http.StatusOK, uiHTML)
				case cfg.SpecURL:
					return c.JSONBlob(http.StatusOK, responseSwaggerJSON)
				case basePath:
					return c.Redirect(http.StatusFound, docPath)
				}
			}

			if next == nil {
				return c.String(http.StatusNotFound, fmt.Sprintf("%q not found", reqPath))
			}

			return next(c)
		}
	}
}

const swaggerTemplateV2 = `
	<!-- HTML for static distribution bundle build -->
<!DOCTYPE html>
<html lang="en">
  <head>
    <meta charset="UTF-8">
    <title>API documentation</title>
    <link rel="stylesheet" type="text/css" href="{{ .SwaggerHost }}/swagger-ui.css" >
    <link rel="icon" type="image/png" href="{{ .SwaggerHost }}/favicon-32x32.png" sizes="32x32" />
    <link rel="icon" type="image/png" href="{{ .SwaggerHost }}/favicon-16x16.png" sizes="16x16" />
    <style>
      html
      {
        box-sizing: border-box;
        overflow: -moz-scrollbars-vertical;
        overflow-y: scroll;
      }

      *,
      *:before,
      *:after
      {
        box-sizing: inherit;
      }

      body
      {
        margin:0;
        background: #fafafa;
      }
    </style>
  </head>

  <body>
    <div id="swagger-ui"></div>

    <script src="{{ .SwaggerHost }}/swagger-ui-bundle.js"> </script>
    <script src="{{ .SwaggerHost }}/swagger-ui-standalone-preset.js"> </script>
    <script>
    window.onload = function() {
      
      // Begin Swagger UI call region
      const ui = SwaggerUIBundle({
        "dom_id": "#swagger-ui",
        deepLinking: true,
        presets: [
          SwaggerUIBundle.presets.apis,
          SwaggerUIStandalonePreset
        ],
        plugins: [
          SwaggerUIBundle.plugins.DownloadUrl
        ],
        layout: "StandaloneLayout",
        validatorUrl: "https://validator.swagger.io/validator",
        url: "{{ .SpecURL }}",
      })

      // End Swagger UI call region
      window.ui = ui
    }
  </script>
  </body>
</html>`
