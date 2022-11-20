// Package golangdocumentation GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag
package golangdocumentation

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

// const docTemplate = `{
//     "schemes": {{ marshal .Schemes }},
//     "swagger": "2.0",
//     "info": {
//         "description": "{{escape .Description}}",
//         "title": "{{.Title}}",
//         "contact": {},
//         "version": "{{.Version}}"
//     },
//     "host": "{{.Host}}",
//     "basePath": "{{.BasePath}}",
//     "paths": {}
// }`

// // SwaggerInfo holds exported Swagger Info so clients can modify it
// var SwaggerInfo = &swag.Spec{
// 	Version:          "1.0",
// 	Host:             "localhost:3000",
// 	BasePath:         "/",
// 	Schemes:          []string{"http"},
// 	Title:            "Golang Documentation API",
// 	Description:      "This is api documentation to hanle another web services in share hosting",
// 	InfoInstanceName: "swagger",
// 	SwaggerTemplate:  docTemplate,
// }

// func init() {
// 	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
// }


var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{.Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "API Support",
            "url": "http://www.swagger.io/support",
            "email": "support@swagger.io"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/": {
            "get": {
                "description": "get the status of server.",
                "consumes": [
                    "*/*"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "root"
                ],
                "summary": "Show the status of server.",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Schemes     []string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = swaggerInfo{
	Version:     "1.0",
	Host:        "localhost:3000",
	BasePath:    "/",
	Schemes:     []string{"http"},
	Title:       "Echo Swagger Example API",
	Description: "This is a sample server server.",
}

type s struct{}

func (s *s) ReadDoc() string {
	sInfo := SwaggerInfo
	sInfo.Description = strings.Replace(sInfo.Description, "\n", "\\n", -1)

	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface{}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, sInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}