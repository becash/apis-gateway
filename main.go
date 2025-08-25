package main

import (
	"embed"
	"fmt"
	"net/http"

	"github.com/becash/apis/gen_openapi/demo"
)

//go:embed openapi.yaml
var openapiFS embed.FS

type svc struct{}

func (svc) ServiceDemoGetOne(w http.ResponseWriter, r *http.Request, id int32) {
	// do your work (lookup by id, etc.)
	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write([]byte(fmt.Sprintf(`{"id": %d, "name": "item-%d"}`, id, id)))
}

func main() {
	// simplest: use the generated mux and default error handler
	api := demo.Handler(svc{})

	// Serve the OpenAPI spec
	http.Handle("/openapi.yaml", http.FileServer(http.FS(openapiFS)))

	// Minimal Swagger UI page referencing the spec and CDN assets
	http.HandleFunc("/docs", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.Write([]byte(`<!doctype html>
<html>
  <head>
    <title>API Docs</title>
    <link rel="stylesheet" href="https://unpkg.com/swagger-ui-dist@5/swagger-ui.css">
  </head>
  <body>
    <div id="swagger-ui"></div>
    <script src="https://unpkg.com/swagger-ui-dist@5/swagger-ui-bundle.js"></script>
    <script>
      window.ui = SwaggerUIBundle({
        url: '/openapi.yaml',
        dom_id: '#swagger-ui',
        presets: [SwaggerUIBundle.presets.apis, SwaggerUIBundle.SwaggerUIStandalonePreset]
      });
    </script>
  </body>
</html>`))
	})

	// Mount your API under root (or use http.NewServeMux if you prefer)
	http.Handle("/", api)

	// start server
	http.ListenAndServe(":8080", nil)

}
