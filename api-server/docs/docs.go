// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag

package docs

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{.Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "Raghu",
            "email": "raghunandankst@gmail.com"
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
        "/album": {
            "get": {
                "description": "Returns dictionary object of all albums, use query string parameters to narrow down searching.\nexample: /album?name=myAlbum1\u0026name=myAlbum2",
                "tags": [
                    "Album"
                ],
                "summary": "list all the albums",
                "responses": {
                    "200": {
                        "description": "returns json data of the album.",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "404": {
                        "description": "if album is not present.",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "500": {
                        "description": "on internal server error",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            },
            "post": {
                "description": "Create album, check example for details on payload.\nExample payload:\n{\n\"name\":\"myAlbumName\"\n}",
                "tags": [
                    "Album"
                ],
                "summary": "add album",
                "parameters": [
                    {
                        "description": "refer to example",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Status Created.",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "500": {
                        "description": "on internal server error",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        },
        "/album/{albumname}": {
            "get": {
                "description": "returns dictionary object of album",
                "tags": [
                    "Album"
                ],
                "summary": "get details of an album",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Album Name",
                        "name": "albumname",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "returns json data of the album.",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "404": {
                        "description": "if album is not present.",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "500": {
                        "description": "on internal server error",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            },
            "post": {
                "tags": [
                    "Image"
                ],
                "summary": "add an image",
                "parameters": [
                    {
                        "description": "Image that needs to be uploaded",
                        "name": "image",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/multipart.Form"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Status Accepted.",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "500": {
                        "description": "on internal server error",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            },
            "delete": {
                "tags": [
                    "Album"
                ],
                "summary": "delete album",
                "parameters": [
                    {
                        "type": "string",
                        "description": "enter album name that needs to be deleted.",
                        "name": "albumname",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Status OK.",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "500": {
                        "description": "on internal server error",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        },
        "/album/{albumname}/image": {
            "get": {
                "tags": [
                    "Image"
                ],
                "summary": "get all images in album",
                "parameters": [
                    {
                        "type": "string",
                        "description": "enter album name whose images are to be returned.",
                        "name": "albumname",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Status OK.",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "500": {
                        "description": "on internal server error",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        },
        "/album/{albumname}/image/{imagename}": {
            "get": {
                "tags": [
                    "Image"
                ],
                "summary": "get Image",
                "parameters": [
                    {
                        "type": "string",
                        "description": "enter album name whose image is to be returned.",
                        "name": "albumname",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "enter image name which needs to be returned.",
                        "name": "imagename",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Status OK.",
                        "schema": {
                            "$ref": "#/definitions/multipart.Form"
                        }
                    },
                    "500": {
                        "description": "on internal server error",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            },
            "delete": {
                "tags": [
                    "Image"
                ],
                "summary": "delete an image",
                "parameters": [
                    {
                        "type": "string",
                        "description": "enter albumname in which image is present",
                        "name": "albumname",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "enter imagename which needs to be deleted",
                        "name": "imagename",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Status OK.",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "500": {
                        "description": "on internal server error",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "multipart.Form": {
            "type": "object",
            "properties": {
                "file": {
                    "type": "object",
                    "additionalProperties": {
                        "type": "array",
                        "items": {
                            "$ref": "#/definitions/multipart.FileHeader"
                        }
                    }
                },
                "value": {
                    "type": "object",
                    "additionalProperties": {
                        "type": "array",
                        "items": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "x-extension-openapi": {
        "example": "value on a json format"
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
	Version:     "0.2",
	Host:        "localhost:8080",
	BasePath:    "/",
	Schemes:     []string{},
	Title:       "Image Service API",
	Description: "Serves API requests to GET and POST images and albums",
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
