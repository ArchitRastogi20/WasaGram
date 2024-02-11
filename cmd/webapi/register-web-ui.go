// //go:build webui

// package main

// import (
// 	"fmt"
// 	"io/fs"
// 	"net/http"
// 	"strings"

// 	// "wasa/webui"
// 	"webui"
// )

// func registerWebUI(hdl http.Handler) (http.Handler, error) {
// 	distDirectory, err := fs.Sub(webui.Dist, "dist")
// 	if err != nil {
// 		return nil, fmt.Errorf("error embedding WebUI dist/ directory: %w", err)
// 	}
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		if strings.HasPrefix(r.RequestURI, "/dashboard/") {
// 			http.StripPrefix("/dashboard/", http.FileServer(http.FS(distDirectory))).ServeHTTP(w, r)
// 			return
// 		}
// 		hdl.ServeHTTP(w, r)
// 	}), nil
// }

//go:build webui

package main

import (
	"fmt"
	"io/fs"
	"net/http"
	"strings"

	// "git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/webui"
	// "github.com/sapienzaapps/fantastic-coffee-decaffeinated/webui"
	"wasa/webui"
)

func registerWebUI(hdl http.Handler) (http.Handler, error) {
	distDirectory, err := fs.Sub(webui.Dist, "dist")
	if err != nil {
		return nil, fmt.Errorf("error embedding WebUI dist/ directory: %w", err)
	}
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasPrefix(r.RequestURI, "/dashboard/") {
			http.StripPrefix("/dashboard/", http.FileServer(http.FS(distDirectory))).ServeHTTP(w, r)
			return
		}
		hdl.ServeHTTP(w, r)
	}), nil
}