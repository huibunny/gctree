// Package docs GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag
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
        "/user/delete": {
            "post": {
                "description": "Delete system",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Delete"
                ],
                "summary": "Delete",
                "operationId": "Delete",
                "parameters": [
                    {
                        "description": "Delete System",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/v1.doDeleteRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/v1.appResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/v1.response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/v1.response"
                        }
                    }
                }
            }
        },
        "/user/detail": {
            "post": {
                "description": "Detail system",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Detail"
                ],
                "summary": "Detail",
                "operationId": "Detail",
                "parameters": [
                    {
                        "description": "Detail System",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/v1.doDetailRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/v1.appResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/v1.response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/v1.response"
                        }
                    }
                }
            }
        },
        "/user/list": {
            "post": {
                "description": "List system",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "List"
                ],
                "summary": "List",
                "operationId": "List",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/v1.appResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/v1.response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/v1.response"
                        }
                    }
                }
            }
        },
        "/user/save": {
            "post": {
                "description": "Save system",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "save"
                ],
                "summary": "Save",
                "operationId": "Save",
                "parameters": [
                    {
                        "description": "Save System",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/v1.doSaveRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/v1.appResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/v1.response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/v1.response"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "entity.TreeNode": {
            "type": "object",
            "properties": {
                "app_id": {
                    "type": "string",
                    "example": "sdfasdflksajflksadjflksja"
                },
                "app_name": {
                    "type": "string",
                    "example": "las"
                },
                "app_type_id": {
                    "type": "string",
                    "example": "fsdfdjklfsdjlkafjslakjfs"
                },
                "brief": {
                    "type": "string",
                    "example": "https://las.com/las.png"
                },
                "cover": {
                    "type": "string",
                    "example": "https://las.com/las.png"
                },
                "create_time": {
                    "type": "string",
                    "example": "2022-10-11 20:20:20"
                },
                "icon": {
                    "type": "string",
                    "example": "https://las.com/las.ico"
                },
                "is_del": {
                    "type": "boolean",
                    "example": false
                },
                "project_id": {
                    "type": "string",
                    "example": "sdfasdflksajflksadjflk111"
                },
                "tenant_id": {
                    "type": "string",
                    "example": "sdfasdflksajflksadjflk111"
                },
                "update_time": {
                    "type": "string",
                    "example": "2022-10-11 20:20:20"
                },
                "user_id": {
                    "type": "string",
                    "example": "sdfasdflksajflksadjflk111"
                }
            }
        },
        "v1.appResponse": {
            "type": "object",
            "properties": {
                "data": {},
                "errcode": {
                    "type": "integer",
                    "example": 0
                },
                "message": {
                    "type": "string",
                    "example": "success"
                }
            }
        },
        "v1.doDeleteRequest": {
            "type": "object",
            "required": [
                "app_id_list"
            ],
            "properties": {
                "app_id_list": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                }
            }
        },
        "v1.doDetailRequest": {
            "type": "object",
            "required": [
                "app_id"
            ],
            "properties": {
                "app_id": {
                    "type": "string",
                    "example": "ksjdflksdjflksdjf"
                }
            }
        },
        "v1.doSaveRequest": {
            "type": "object",
            "properties": {
                "app_list": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/entity.TreeNode"
                    }
                }
            }
        },
        "v1.response": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string",
                    "example": "success"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8820",
	BasePath:         "/v1",
	Schemes:          []string{},
	Title:            "gctree API",
	Description:      "Using a translation service as an example",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
