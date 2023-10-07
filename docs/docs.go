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
        "/posts": {
            "get": {
                "description": "Return list of posts.",
                "posts": [
                    "posts"
                ],
                "summary": "Get All posts.",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "obejct"
                        }
                    }
                }
            },
            "post": {
                "description": "Save posts data in Db.",
                "produces": [
                    "application/json"
                ],
                "posts": [
                    "posts"
                ],
                "summary": "Create posts",
                "parameters": [
                    {
                        "description": "Create posts",
                        "name": "posts",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.CreatePostsRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    }
                }
            }
        },
        "/posts/{postID}": {
            "delete": {
                "description": "Remove posts data by id.",
                "produces": [
                    "application/json"
                ],
                "posts": [
                    "posts"
                ],
                "summary": "Delete posts",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    }
                }
            }
        },
        "/posts/{postId}": {
            "get": {
                "description": "Return the tahs whoes postId valu mathes id.",
                "produces": [
                    "application/json"
                ],
                "posts": [
                    "posts"
                ],
                "summary": "Get Single posts by id.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "update posts by id",
                        "name": "postId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    }
                }
            },
            "patch": {
                "description": "Update posts data.",
                "produces": [
                    "application/json"
                ],
                "posts": [
                    "posts"
                ],
                "summary": "Update posts",
                "parameters": [
                    {
                        "type": "string",
                        "description": "update posts by id",
                        "name": "postId",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Update posts",
                        "name": "posts",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.CreatePostsRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "request.CreatePostsRequest": {
            "type": "object",
            "required": [
                "name"
            ],
            "properties": {
                "name": {
                    "type": "string",
                    "maxLength": 200,
                    "minLength": 1
                }
            }
        },
        "response.Response": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "data": {},
                "status": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:3000",
	BasePath:         "/api",
	Schemes:          []string{},
	Title:            "Post Service API",
	Description:      "A Post service API in Go using Gin framework",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}