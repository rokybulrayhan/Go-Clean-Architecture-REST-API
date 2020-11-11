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
        "contact": {
            "name": "Alexander Bryksin",
            "url": "https://github.com/AleksK1NG",
            "email": "alexander.bryksin@yandex.ru"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/auth/avatar": {
            "post": {
                "description": "Post user avatar image",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Post avatar",
                "parameters": [
                    {
                        "type": "file",
                        "description": "Body with image file",
                        "name": "file",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/httpErrors.RestError"
                        }
                    }
                }
            }
        },
        "/auth/find": {
            "get": {
                "description": "Get the list of all users",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Get users",
                "parameters": [
                    {
                        "type": "integer",
                        "format": "page",
                        "description": "page",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "format": "size",
                        "description": "size",
                        "name": "size",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "format": "orderBy",
                        "description": "order by",
                        "name": "orderBy",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.UsersList"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/httpErrors.RestError"
                        }
                    }
                }
            }
        },
        "/auth/login": {
            "post": {
                "description": "login user, returns user and set session",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Login new user",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.User"
                        }
                    }
                }
            }
        },
        "/auth/logout": {
            "post": {
                "description": "logout user removing session",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Logout user",
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/auth/me": {
            "get": {
                "description": "Get current user by id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Get user by id",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.User"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/httpErrors.RestError"
                        }
                    }
                }
            }
        },
        "/auth/register": {
            "post": {
                "description": "register new user, returns user and token",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Register new user",
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/models.User"
                        }
                    }
                }
            }
        },
        "/auth/{id}": {
            "get": {
                "description": "get string by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "get user by id",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "user_id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.User"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/httpErrors.RestError"
                        }
                    }
                }
            },
            "put": {
                "description": "update existing user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Update user",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "user_id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.User"
                        }
                    }
                }
            },
            "delete": {
                "description": "some description",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Delete",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "user_id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/httpErrors.RestError"
                        }
                    }
                }
            }
        },
        "/comments": {
            "post": {
                "description": "create new comment",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Create new comment",
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/models.Comment"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/httpErrors.RestErr"
                        }
                    }
                }
            }
        },
        "/comments/byNewsId/{id}": {
            "get": {
                "description": "Get all comment by news id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Get comments by news",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "news_id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "format": "page",
                        "description": "page",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "format": "size",
                        "description": "size",
                        "name": "size",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "format": "orderBy",
                        "description": "order by",
                        "name": "orderBy",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.CommentsList"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/httpErrors.RestErr"
                        }
                    }
                }
            }
        },
        "/comments/{id}": {
            "get": {
                "description": "Get comment by id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Get comment",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "comment_id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Comment"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/httpErrors.RestErr"
                        }
                    }
                }
            },
            "put": {
                "description": "update new comment",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Update comment",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "comment_id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Comment"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/httpErrors.RestErr"
                        }
                    }
                }
            },
            "delete": {
                "description": "delete comment",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Delete comment",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "comment_id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/httpErrors.RestErr"
                        }
                    }
                }
            }
        },
        "/news": {
            "get": {
                "description": "Get all news with pagination",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Get all news",
                "parameters": [
                    {
                        "type": "integer",
                        "format": "page",
                        "description": "page",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "format": "size",
                        "description": "size",
                        "name": "size",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "format": "orderBy",
                        "description": "order by",
                        "name": "orderBy",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.NewsList"
                        }
                    }
                }
            }
        },
        "/news/create": {
            "post": {
                "description": "Create news handler",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Create news",
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/models.News"
                        }
                    }
                }
            }
        },
        "/news/search": {
            "get": {
                "description": "Search news by title",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Search by title",
                "parameters": [
                    {
                        "type": "integer",
                        "format": "page",
                        "description": "page",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "format": "size",
                        "description": "size",
                        "name": "size",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "format": "orderBy",
                        "description": "order by",
                        "name": "orderBy",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.NewsList"
                        }
                    }
                }
            }
        },
        "/news/{id}": {
            "get": {
                "description": "Get by id news handler",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Get by id news",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "news_id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.News"
                        }
                    }
                }
            },
            "put": {
                "description": "Update news handler",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Update news",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "news_id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.News"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete by id news handler",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Delete news",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "news_id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "httpErrors.RestErr": {
            "type": "object"
        },
        "httpErrors.RestError": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string"
                },
                "status": {
                    "type": "integer"
                }
            }
        },
        "models.Comment": {
            "type": "object",
            "required": [
                "author_id",
                "message",
                "news_id"
            ],
            "properties": {
                "author_id": {
                    "type": "string"
                },
                "comment_id": {
                    "type": "string"
                },
                "created_at": {
                    "type": "string"
                },
                "likes": {
                    "type": "integer"
                },
                "message": {
                    "type": "string"
                },
                "news_id": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                }
            }
        },
        "models.CommentBase": {
            "type": "object",
            "required": [
                "author",
                "author_id",
                "message"
            ],
            "properties": {
                "author": {
                    "type": "string"
                },
                "author_id": {
                    "type": "string"
                },
                "avatar_url": {
                    "type": "string"
                },
                "comment_id": {
                    "type": "string"
                },
                "likes": {
                    "type": "integer"
                },
                "message": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                }
            }
        },
        "models.CommentsList": {
            "type": "object",
            "properties": {
                "comments": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.CommentBase"
                    }
                },
                "has_more": {
                    "type": "boolean"
                },
                "page": {
                    "type": "integer"
                },
                "size": {
                    "type": "integer"
                },
                "total_count": {
                    "type": "integer"
                },
                "total_pages": {
                    "type": "integer"
                }
            }
        },
        "models.News": {
            "type": "object",
            "required": [
                "author_id",
                "content",
                "title"
            ],
            "properties": {
                "author_id": {
                    "type": "string"
                },
                "category": {
                    "type": "string"
                },
                "content": {
                    "type": "string"
                },
                "created_at": {
                    "type": "string"
                },
                "image_url": {
                    "type": "string"
                },
                "news_id": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                }
            }
        },
        "models.NewsList": {
            "type": "object",
            "properties": {
                "has_more": {
                    "type": "boolean"
                },
                "news": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.News"
                    }
                },
                "page": {
                    "type": "integer"
                },
                "size": {
                    "type": "integer"
                },
                "total_count": {
                    "type": "integer"
                },
                "total_pages": {
                    "type": "integer"
                }
            }
        },
        "models.User": {
            "type": "object",
            "required": [
                "first_name",
                "last_name",
                "password"
            ],
            "properties": {
                "about": {
                    "type": "string"
                },
                "address": {
                    "type": "string"
                },
                "avatar": {
                    "type": "string"
                },
                "birthday": {
                    "type": "string"
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
                "email": {
                    "type": "string"
                },
                "first_name": {
                    "type": "string"
                },
                "gender": {
                    "type": "string"
                },
                "last_name": {
                    "type": "string"
                },
                "login_date": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "phone_number": {
                    "type": "string"
                },
                "postcode": {
                    "type": "integer"
                },
                "role": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                },
                "user_id": {
                    "type": "string"
                }
            }
        },
        "models.UsersList": {
            "type": "object",
            "properties": {
                "has_more": {
                    "type": "boolean"
                },
                "page": {
                    "type": "integer"
                },
                "size": {
                    "type": "integer"
                },
                "total_count": {
                    "type": "integer"
                },
                "total_pages": {
                    "type": "integer"
                },
                "users": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.User"
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
	Host:        "",
	BasePath:    "/api/v1",
	Schemes:     []string{},
	Title:       "Go Example API",
	Description: "This is a example Golang REST API.",
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
