package main

import (
	"fmt"
    "log"
    "net/http"

    "appstore/backend"
    "appstore/handler"
    "appstore/util"
)
func main() {
    fmt.Println("started-service")

	config, err := util.LoadApplicationConfig("conf", "deploy.yml")
    if err != nil {
        panic(err)
    }

	backend.InitElasticsearchBackend(config.ElasticsearchConfig)
    backend.InitGCSBackend(config.GCSConfig)

    log.Fatal(http.ListenAndServe(":8080", handler.InitRouter(config.TokenConfig)))
}

/* // requirements
// login
// register
// post
// search
// checkout

// FE
	Design - UI
	FE
		React / Angular / Vue
	Mobile
		iOS / Android

// API
	Rest <-
		URL
		HTTP method: READ, POST, UPDAT, DELETE
	gRPC

// BE
	Data
	 	user
		app(title, description, price, image, media file, etc)
	Store
		Database
		Storage(AWS S3, GCS)
	Go
	Code Structure
		Controller - handler
		Service - Service
		DAO(entity) - backend(model)
 */