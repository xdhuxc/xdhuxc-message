package api

import (
	"fmt"
	"net/http"

	"github.com/emicklei/go-restful"
	"github.com/emicklei/go-restful-openapi"
	"github.com/go-openapi/spec"
)

func swagger(c *restful.Container, address string) {
	config := restfulspec.Config{
		WebServices:                   c.RegisteredWebServices(),
		WebServicesURL:                fmt.Sprintf(":%s", address),
		APIPath:                       "/docs.json",
		PostBuildSwaggerObjectHandler: enrichSwaggerObject}

	c.Handle("/docs/", http.StripPrefix("/docs/", http.FileServer(http.Dir("dist"))))
	c.Add(restfulspec.NewOpenAPIService(config))
}

func enrichSwaggerObject(swo *spec.Swagger) {
	swo.Info = &spec.Info{
		InfoProps: spec.InfoProps{
			Title:       "Xdhuxc Message APIServer",
			Description: "Resource for managing xdhuxc-message api",
			Contact: &spec.ContactInfo{
				Name:  "xdhuxc",
				Email: "xdhuxc@163.com",
				URL:   "http://xdhuxc.org",
			},
			License: &spec.License{
				Name: "Xdhuxc",
				URL:  "http://xdhuxc.org",
			},
			Version: "1.0.0",
		},
	}

	swo.Tags = []spec.Tag{spec.Tag{TagProps: spec.TagProps{
		Name:        "message",
		Description: "Managing message"}}}
}
