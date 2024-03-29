// Package docs GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import (
	"bytes"
	"encoding/json"
	"strings"
	"text/template"

	"github.com/swaggo/swag"
)

var doc = `{
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
        "/todos/": {
            "get": {
                "description": "getall todolist",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "todos"
                ],
                "operationId": "getall todolist",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/doc_datas.GetAlldTodolistResponse"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/error_utils.MessageErrData"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/error_utils.MessageErrData"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/error_utils.MessageErrData"
                        }
                    }
                }
            },
            "post": {
                "description": "create todolist",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "todos"
                ],
                "operationId": "create todolist",
                "parameters": [
                    {
                        "description": "request body json",
                        "name": "RequestBody",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/doc_datas.CreateTodolistRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/doc_datas.CreateTodolistResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/error_utils.MessageErrData"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/error_utils.MessageErrData"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/error_utils.MessageErrData"
                        }
                    }
                }
            }
        },
        "/todos/{id}": {
            "get": {
                "description": "getbyid todolist",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "todos"
                ],
                "operationId": "getbyid todolist",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "todolist id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/doc_datas.GetByIdTodolistResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/error_utils.MessageErrData"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/error_utils.MessageErrData"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/error_utils.MessageErrData"
                        }
                    }
                }
            },
            "put": {
                "description": "update todolist",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "todos"
                ],
                "operationId": "update todolist",
                "parameters": [
                    {
                        "description": "request body json",
                        "name": "RequestBody",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/doc_datas.UpdateTodolistRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/doc_datas.UpdateTodolistResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/error_utils.MessageErrData"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/error_utils.MessageErrData"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/error_utils.MessageErrData"
                        }
                    }
                }
            },
            "delete": {
                "description": "delete todolist",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "todos"
                ],
                "operationId": "delete todolist",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "todolist id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/doc_datas.DeleteTodolistResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/error_utils.MessageErrData"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/error_utils.MessageErrData"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/error_utils.MessageErrData"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "doc_datas.CreateTodolistRequest": {
            "type": "object",
            "properties": {
                "title": {
                    "type": "string",
                    "example": "Makan"
                }
            }
        },
        "doc_datas.CreateTodolistResponse": {
            "type": "object",
            "properties": {
                "completed": {
                    "type": "boolean",
                    "example": false
                },
                "created_at": {
                    "type": "string",
                    "example": "2022-01-18T11:45:59.136128+07:00"
                },
                "id": {
                    "type": "integer",
                    "example": 1
                },
                "title": {
                    "type": "string",
                    "example": "Makan"
                }
            }
        },
        "doc_datas.DeleteTodolistResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string",
                    "example": "delete sukses"
                }
            }
        },
        "doc_datas.GetAlldTodolistResponse": {
            "type": "object",
            "properties": {
                "completed": {
                    "type": "boolean",
                    "example": true
                },
                "created_at": {
                    "type": "string",
                    "example": "2022-01-18T11:45:59.136128+07:00"
                },
                "id": {
                    "type": "integer",
                    "example": 1
                },
                "title": {
                    "type": "string",
                    "example": "Makan (Sudah)"
                }
            }
        },
        "doc_datas.GetByIdTodolistResponse": {
            "type": "object",
            "properties": {
                "completed": {
                    "type": "boolean",
                    "example": true
                },
                "created_at": {
                    "type": "string",
                    "example": "2022-01-18T11:45:59.136128+07:00"
                },
                "id": {
                    "type": "integer",
                    "example": 1
                },
                "title": {
                    "type": "string",
                    "example": "Makan (Sudah)"
                }
            }
        },
        "doc_datas.UpdateTodolistRequest": {
            "type": "object",
            "properties": {
                "completed": {
                    "type": "boolean",
                    "example": true
                },
                "title": {
                    "type": "string",
                    "example": "Makan (Sudah)"
                }
            }
        },
        "doc_datas.UpdateTodolistResponse": {
            "type": "object",
            "properties": {
                "completed": {
                    "type": "boolean",
                    "example": true
                },
                "created_at": {
                    "type": "string",
                    "example": "2022-01-18T11:45:59.136128+07:00"
                },
                "id": {
                    "type": "integer",
                    "example": 1
                },
                "title": {
                    "type": "string",
                    "example": "Makan (Sudah)"
                }
            }
        },
        "error_utils.MessageErrData": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string"
                },
                "message": {
                    "type": "string"
                },
                "status": {
                    "type": "integer"
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
	Version:     "",
	Host:        "",
	BasePath:    "",
	Schemes:     []string{},
	Title:       "",
	Description: "",
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
		"escape": func(v interface{}) string {
			// escape tabs
			str := strings.Replace(v.(string), "\t", "\\t", -1)
			// replace " with \", and if that results in \\", replace that with \\\"
			str = strings.Replace(str, "\"", "\\\"", -1)
			return strings.Replace(str, "\\\\\"", "\\\\\\\"", -1)
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
	swag.Register("swagger", &s{})
}
