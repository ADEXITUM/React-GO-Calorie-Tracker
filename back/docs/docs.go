// Code generated by swaggo/swag. DO NOT EDIT.

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
        "/calculate/calories": {
            "post": {
                "description": "Calculate calories based on Harris-Benedicts formula",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "calculation"
                ],
                "summary": "Calculate Calories",
                "operationId": "calculate-calories",
                "parameters": [
                    {
                        "description": "calculation info",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/github_com_ADEXITUM_calorie-counter_pkg_entities.Calculation"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "number"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/http.ProtocolError"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/http.ProtocolError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/http.ProtocolError"
                        }
                    },
                    "default": {
                        "description": "",
                        "schema": {
                            "$ref": "#/definitions/http.ProtocolError"
                        }
                    }
                }
            }
        },
        "/calculate/massindex": {
            "post": {
                "description": "Calculate Body Mass Index based on Adolf Ketles formula",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "calculation"
                ],
                "summary": "Calculate Body Mass Index",
                "operationId": "calculate-massidex",
                "parameters": [
                    {
                        "description": "calculation info",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/github_com_ADEXITUM_calorie-counter_pkg_entities.Calculation"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "number"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/http.ProtocolError"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/http.ProtocolError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/http.ProtocolError"
                        }
                    },
                    "default": {
                        "description": "",
                        "schema": {
                            "$ref": "#/definitions/http.ProtocolError"
                        }
                    }
                }
            }
        },
        "/entries": {
            "get": {
                "description": "Get all previously added entries.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "entries"
                ],
                "summary": "Get All Entries",
                "operationId": "get-all-entries",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {}
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/http.ProtocolError"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/http.ProtocolError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/http.ProtocolError"
                        }
                    },
                    "default": {
                        "description": "",
                        "schema": {
                            "$ref": "#/definitions/http.ProtocolError"
                        }
                    }
                }
            }
        },
        "/entries/create": {
            "post": {
                "description": "Add entry of comsumed dish, its macroes and calories.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "entries"
                ],
                "summary": "Add Entry",
                "operationId": "add-entry",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {}
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/http.ProtocolError"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/http.ProtocolError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/http.ProtocolError"
                        }
                    },
                    "default": {
                        "description": "",
                        "schema": {
                            "$ref": "#/definitions/http.ProtocolError"
                        }
                    }
                }
            }
        },
        "/entries/delete/all": {
            "delete": {
                "description": "Delete all entries in a \"calories\" collection.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "entries"
                ],
                "summary": "Delete All Entries",
                "operationId": "delete-all-entries",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {}
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/http.ProtocolError"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/http.ProtocolError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/http.ProtocolError"
                        }
                    },
                    "default": {
                        "description": "",
                        "schema": {
                            "$ref": "#/definitions/http.ProtocolError"
                        }
                    }
                }
            }
        },
        "/entries/delete/{id}": {
            "delete": {
                "description": "Delete entry by its ID.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "entries"
                ],
                "summary": "Delete Entry",
                "operationId": "delete-entry",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {}
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/http.ProtocolError"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/http.ProtocolError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/http.ProtocolError"
                        }
                    },
                    "default": {
                        "description": "",
                        "schema": {
                            "$ref": "#/definitions/http.ProtocolError"
                        }
                    }
                }
            }
        },
        "/entries/update/:id": {
            "put": {
                "description": "Update entry with new information.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "entries"
                ],
                "summary": "Update Entry",
                "operationId": "update-entry",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {}
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/http.ProtocolError"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/http.ProtocolError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/http.ProtocolError"
                        }
                    },
                    "default": {
                        "description": "",
                        "schema": {
                            "$ref": "#/definitions/http.ProtocolError"
                        }
                    }
                }
            }
        },
        "/entries/{id}": {
            "get": {
                "description": "Get entry by its ID.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "entries"
                ],
                "summary": "Get Entry by ID",
                "operationId": "get-entry-by-id",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {}
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/http.ProtocolError"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/http.ProtocolError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/http.ProtocolError"
                        }
                    },
                    "default": {
                        "description": "",
                        "schema": {
                            "$ref": "#/definitions/http.ProtocolError"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "github_com_ADEXITUM_calorie-counter_pkg_entities.Calculation": {
            "type": "object",
            "properties": {
                "age": {
                    "type": "number"
                },
                "gender": {
                    "type": "string"
                },
                "height": {
                    "type": "number"
                },
                "weight": {
                    "type": "number"
                }
            }
        },
        "http.ProtocolError": {
            "type": "object",
            "properties": {
                "errorString": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0.0",
	Host:             "localhost:8080",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Calorie Tracker",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
