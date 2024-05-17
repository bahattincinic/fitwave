// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/config": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "config"
                ],
                "summary": "Get Application Config",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Config"
                        }
                    }
                }
            },
            "put": {
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "config"
                ],
                "summary": "Upsert Application Config",
                "parameters": [
                    {
                        "description": "Config Input",
                        "name": "config",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.ConfigInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.ConfigInput"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.Config": {
            "type": "object",
            "properties": {
                "access_token": {
                    "type": "string"
                },
                "athlete_id": {
                    "type": "integer"
                },
                "client_id": {
                    "type": "string"
                },
                "client_secret": {
                    "type": "string"
                },
                "created_at": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "updated_at": {
                    "type": "string"
                }
            }
        },
        "models.ConfigInput": {
            "type": "object",
            "properties": {
                "access_token": {
                    "type": "string"
                },
                "athlete_id": {
                    "type": "integer"
                },
                "client_id": {
                    "type": "string"
                },
                "client_secret": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
