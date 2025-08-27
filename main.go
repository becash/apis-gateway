package main

import (
	"apis_gateway/suppliers/get_your_guide"
	"embed"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	//"github.com/oapi-codegen/oapi-codegen/v2/pkg/securityprovider"
)

//go:embed suppliers/get_your_guide/api.yaml
var openapiFS embed.FS

func main() {
	//
	//	// Serve the OpenAPI spec
	//	http.Handle("/suppliers/get_your_guide/api.yaml", http.FileServer(http.FS(openapiFS)))
	//
	//	// Minimal Swagger UI page referencing the spec and CDN assets
	//	http.HandleFunc("/docs", func(w http.ResponseWriter, r *http.Request) {
	//		w.Header().Set("Content-Type", "text/html; charset=utf-8")
	//		w.Write([]byte(`<!doctype html>
	//<html>
	//  <head>
	//    <title>API Docs</title>
	//    <link rel="stylesheet" href="https://unpkg.com/swagger-ui-dist@5/swagger-ui.css">
	//  </head>
	//  <body>
	//    <div id="swagger-ui"></div>
	//    <script src="https://unpkg.com/swagger-ui-dist@5/swagger-ui-bundle.js"></script>
	//    <script>
	//      window.ui = SwaggerUIBundle({
	//        url: '/suppliers/get_your_guide/api.yaml',
	//        dom_id: '#swagger-ui',
	//        presets: [SwaggerUIBundle.presets.apis, SwaggerUIBundle.SwaggerUIStandalonePreset]
	//      });
	//    </script>
	//  </body>
	//</html>`))
	//	})

	// start server

	server := get_your_guide.NewServer()

	r := gin.Default()

	get_your_guide.RegisterHandlers(r, server)

	// And we serve HTTP until the world ends.

	s := &http.Server{
		Handler: r,
		Addr:    "0.0.0.0:8080",
	}

	// And we serve HTTP until the world ends.
	log.Fatal(s.ListenAndServe())

}
