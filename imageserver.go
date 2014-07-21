package main

import (
	"net/http"
  "os"

	"github.com/pierrre/imageserver"
	imageserver_http "github.com/pierrre/imageserver/http"
	imageserver_http_parser_graphicsmagick "github.com/pierrre/imageserver/http/parser/graphicsmagick"
	imageserver_http_parser_list "github.com/pierrre/imageserver/http/parser/list"
	imageserver_http_parser_source "github.com/pierrre/imageserver/http/parser/source"
	imageserver_processor "github.com/pierrre/imageserver/processor"
	imageserver_processor_graphicsmagick "github.com/pierrre/imageserver/processor/graphicsmagick"
	imageserver_provider "github.com/pierrre/imageserver/provider"
	imageserver_provider_http "github.com/pierrre/imageserver/provider/http"
)

func main() {
	var imageServer imageserver.ImageServer
	imageServer = &imageserver_provider.ProviderImageServer{
		Provider: &imageserver_provider_http.HTTPProvider{},
	}
	imageServer = &imageserver_processor.ProcessorImageServer{
		ImageServer: imageServer,
		Processor: &imageserver_processor_graphicsmagick.GraphicsMagickProcessor{
			Executable: "gm",
		},
	}

	imageHTTPHandler := &imageserver_http.ImageHTTPHandler{
		Parser: &imageserver_http_parser_list.ListParser{
			&imageserver_http_parser_source.SourceParser{},
			&imageserver_http_parser_graphicsmagick.GraphicsMagickParser{},
		},
		ImageServer: imageServer,
	}

	http.Handle("/", imageHTTPHandler)
	err := http.ListenAndServe(portSetting(), nil)
	if err != nil {
		panic(err)
	}
}


func portSetting() string {
    port := os.Getenv("PORT")
    if port == "" {
        panic("No PORT env-var given!")
    }
    return ":" + port
}

