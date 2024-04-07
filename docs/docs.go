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
        "/cars/delete": {
            "delete": {
                "description": "delete car",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "DeleteCar",
                "parameters": [
                    {
                        "description": "h",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/http.deleteCarRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "integer"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "integer"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "integer"
                        }
                    }
                }
            }
        },
        "/cars/get": {
            "get": {
                "description": "get cars",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "GetCars",
                "parameters": [
                    {
                        "type": "string",
                        "description": "h",
                        "name": "regNum",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "h",
                        "name": "model",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "h",
                        "name": "mark",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "h",
                        "name": "name",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "h",
                        "name": "surname",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "h",
                        "name": "patronymic",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/http.getCarsResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "integer"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "integer"
                        }
                    }
                }
            }
        },
        "/cars/new": {
            "post": {
                "description": "new cars",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "NewCars",
                "parameters": [
                    {
                        "description": "h",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/http.newCarRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "integer"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "integer"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "integer"
                        }
                    }
                }
            }
        },
        "/cars/update": {
            "put": {
                "description": "update car",
                "consumes": [
                    "application/json"
                ],
                "summary": "UpdateCar",
                "parameters": [
                    {
                        "description": "h",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/http.updateCarRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "integer"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "integer"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "integer"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "http.deleteCarRequest": {
            "type": "object",
            "properties": {
                "regNum": {
                    "type": "string"
                }
            }
        },
        "http.getCarsResponse": {
            "type": "object",
            "properties": {
                "cars": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Car"
                    }
                }
            }
        },
        "http.newCarRequest": {
            "type": "object",
            "properties": {
                "regNums": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                }
            }
        },
        "http.updateCarRequest": {
            "type": "object",
            "properties": {
                "regNum": {
                    "type": "string"
                },
                "updates": {
                    "$ref": "#/definitions/models.Car"
                }
            }
        },
        "models.Car": {
            "type": "object",
            "properties": {
                "mark": {
                    "type": "string"
                },
                "model": {
                    "type": "string"
                },
                "owner": {
                    "$ref": "#/definitions/models.People"
                },
                "regNum": {
                    "type": "string"
                }
            }
        },
        "models.People": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string"
                },
                "patronymic": {
                    "type": "string"
                },
                "surname": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8087",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Car App API",
	Description:      "API Server For Car's Catalog",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
