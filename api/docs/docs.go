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
        "/activities/": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "activity"
                ],
                "summary": "List Activities",
                "parameters": [
                    {
                        "type": "string",
                        "description": "pagination limit",
                        "name": "limit",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "active page",
                        "name": "page",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/api.PaginatedResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        " count": {
                                            "type": "integer"
                                        },
                                        "Results": {
                                            "type": "array",
                                            "items": {
                                                "$ref": "#/definitions/models.Activity"
                                            }
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/activities/{id}": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "activity"
                ],
                "summary": "Get Activity",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Activity ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Activity"
                        }
                    }
                }
            }
        },
        "/athletes/": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "athlete"
                ],
                "summary": "List Athletes",
                "parameters": [
                    {
                        "type": "string",
                        "description": "pagination limit",
                        "name": "limit",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "active page",
                        "name": "page",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/api.PaginatedResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        " count": {
                                            "type": "integer"
                                        },
                                        "Results": {
                                            "type": "array",
                                            "items": {
                                                "$ref": "#/definitions/models.Athlete"
                                            }
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/athletes/{id}": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "athlete"
                ],
                "summary": "Get Athlete",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Athlete ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Athlete"
                        }
                    }
                }
            }
        },
        "/config/": {
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
                            "$ref": "#/definitions/models.Config"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Config"
                        }
                    }
                }
            }
        },
        "/gears/": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "gear"
                ],
                "summary": "List Gears",
                "parameters": [
                    {
                        "type": "string",
                        "description": "pagination limit",
                        "name": "limit",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "active page",
                        "name": "page",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/api.PaginatedResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        " count": {
                                            "type": "integer"
                                        },
                                        "Results": {
                                            "type": "array",
                                            "items": {
                                                "$ref": "#/definitions/models.Gear"
                                            }
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/gears/{id}": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "gear"
                ],
                "summary": "Get Gear",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Gear ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Gear"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "api.PaginatedResponse": {
            "type": "object",
            "properties": {
                "count": {
                    "type": "integer"
                },
                "results": {}
            }
        },
        "models.Activity": {
            "type": "object",
            "properties": {
                "achievement_count": {
                    "type": "integer"
                },
                "athlete": {
                    "$ref": "#/definitions/models.Athlete"
                },
                "athleteID": {
                    "type": "integer"
                },
                "athlete_count": {
                    "type": "integer"
                },
                "average_cadence": {
                    "type": "number"
                },
                "average_heart_rate": {
                    "type": "number"
                },
                "average_speed": {
                    "type": "number"
                },
                "average_temp": {
                    "type": "number"
                },
                "average_watts": {
                    "type": "number"
                },
                "comment_count": {
                    "type": "integer"
                },
                "commute": {
                    "type": "boolean"
                },
                "device_watts": {
                    "type": "boolean"
                },
                "distance": {
                    "type": "number"
                },
                "elapsed_time": {
                    "type": "integer"
                },
                "end_lat_lng": {
                    "type": "string"
                },
                "external_id": {
                    "type": "string"
                },
                "flagged": {
                    "type": "boolean"
                },
                "gear_id": {
                    "description": "bike or pair of shoes",
                    "type": "string"
                },
                "has_kudos": {
                    "type": "boolean"
                },
                "id": {
                    "type": "integer"
                },
                "kilojoules": {
                    "type": "number"
                },
                "kudos_count": {
                    "type": "integer"
                },
                "location_city": {
                    "type": "string"
                },
                "location_country": {
                    "type": "string"
                },
                "location_state": {
                    "type": "string"
                },
                "manual": {
                    "type": "boolean"
                },
                "map": {
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                },
                "max_heart_rate": {
                    "type": "number"
                },
                "max_speed": {
                    "type": "number"
                },
                "moving_time": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "photo_count": {
                    "type": "integer"
                },
                "private": {
                    "type": "boolean"
                },
                "start_date": {
                    "type": "string"
                },
                "start_date_local": {
                    "type": "string"
                },
                "start_lat_lng": {
                    "type": "string"
                },
                "time_zone": {
                    "type": "string"
                },
                "total_elevation_gain": {
                    "type": "number"
                },
                "trainer": {
                    "type": "boolean"
                },
                "truncated": {
                    "description": "only present if activity is owned by authenticated athlete, returns 0 if not truncated by privacy zones",
                    "type": "integer"
                },
                "type": {
                    "type": "string"
                },
                "upload_id": {
                    "type": "integer"
                },
                "weighted_average_watts": {
                    "type": "integer"
                }
            }
        },
        "models.Athlete": {
            "type": "object",
            "properties": {
                "approve_followers": {
                    "description": "if has enhanced privacy enabled",
                    "type": "boolean"
                },
                "badge_type_id": {
                    "type": "integer"
                },
                "city": {
                    "type": "string"
                },
                "country": {
                    "type": "string"
                },
                "created_at": {
                    "type": "string"
                },
                "firstname": {
                    "type": "string"
                },
                "follower": {
                    "description": "this athlete’s following status of the authenticated athlete",
                    "type": "string"
                },
                "friend": {
                    "description": "‘pending’, ‘accepted’, ‘blocked’ or ‘null’, the authenticated athlete’s following status of this athlete",
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "lastname": {
                    "type": "string"
                },
                "premium": {
                    "type": "boolean"
                },
                "profile": {
                    "description": "URL to a 124x124 pixel profile picture",
                    "type": "string"
                },
                "profile_medium": {
                    "description": "URL to a 62x62 pixel profile picture",
                    "type": "string"
                },
                "sex": {
                    "type": "string"
                },
                "state": {
                    "type": "string"
                },
                "tx": {
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                },
                "updated_at": {
                    "type": "string"
                }
            }
        },
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
                }
            }
        },
        "models.Gear": {
            "type": "object",
            "properties": {
                "athlete": {
                    "$ref": "#/definitions/models.Athlete"
                },
                "athleteID": {
                    "type": "integer"
                },
                "brand_name": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "distance": {
                    "type": "number"
                },
                "frame_type": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "model_name": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "primary": {
                    "type": "boolean"
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