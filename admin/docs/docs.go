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
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "这里写联系人信息",
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
        "/login": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "根据用户名和密码登录",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "login 登录接口",
                "parameters": [
                    {
                        "type": "string",
                        "name": "password",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "name": "username",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.LoginResp"
                        }
                    }
                }
            }
        },
        "/posts2": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "根据用户名密码角色增加用户",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "增加用户信息",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer 用户令牌",
                        "name": "Authorization",
                        "in": "header"
                    },
                    {
                        "type": "string",
                        "name": "password",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "name": "role",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "name": "username",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        }
    },
    "definitions": {
        "model.LoginResp": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "token": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost",
	BasePath:         "/my-blog",
	Schemes:          []string{},
	Title:            "myBlog",
	Description:      "--",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}