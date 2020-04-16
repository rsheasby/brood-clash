// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag at
// 2020-04-16 07:53:57.994605 +0200 SAST m=+0.043026873

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
        "contact": {},
        "license": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/answers/{id}/reveal": {
            "post": {
                "security": [
                    {
                        "CodeAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Reveal answer",
                "operationId": "reveal-answer",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Answer ID, must be UUID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "Success"
                    },
                    "400": {
                        "description": "Couldn't parse ID param into UUID"
                    },
                    "403": {
                        "description": "Unauthorised"
                    },
                    "404": {
                        "description": "Answer doesn't exist"
                    },
                    "418": {
                        "description": "Answer already revealed",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/currentQuestion": {
            "get": {
                "security": [
                    {
                        "CodeAuth": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Get Current Question",
                "operationId": "get-current-question",
                "responses": {
                    "200": {
                        "description": "Success",
                        "schema": {
                            "$ref": "#/definitions/models.Question"
                        }
                    },
                    "401": {
                        "description": "Unauthorised"
                    },
                    "404": {
                        "description": "No current question"
                    }
                }
            }
        },
        "/incorrectAnswer": {
            "post": {
                "security": [
                    {
                        "CodeAuth": []
                    }
                ],
                "summary": "Incorrect Answer",
                "operationId": "incorrect-answer",
                "responses": {
                    "204": {
                        "description": "Success"
                    },
                    "401": {
                        "description": "Unauthorised"
                    }
                }
            }
        },
        "/questions": {
            "get": {
                "security": [
                    {
                        "CodeAuth": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Get All Questions",
                "operationId": "get-all-questions",
                "responses": {
                    "200": {
                        "description": "Success",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Question"
                            }
                        }
                    },
                    "401": {
                        "description": "Unauthorised"
                    },
                    "500": {
                        "description": "Failed to retrieve questions"
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "CodeAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "summary": "Post Questions",
                "operationId": "post-questions",
                "parameters": [
                    {
                        "description": "Questions to be created",
                        "name": "questions",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Question"
                            }
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Questions created"
                    },
                    "202": {
                        "description": "Some questions couldn't be created"
                    },
                    "401": {
                        "description": "Unauthorised"
                    }
                }
            }
        },
        "/questions/{id}": {
            "delete": {
                "security": [
                    {
                        "CodeAuth": []
                    }
                ],
                "summary": "Delete Question",
                "operationId": "delete-question",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Question ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "Success"
                    },
                    "400": {
                        "description": "Couldn't parse ID param into UUID"
                    },
                    "401": {
                        "description": "Unauthorised"
                    },
                    "404": {
                        "description": "Question doesn't exist"
                    },
                    "500": {
                        "description": "Unknown error"
                    }
                }
            }
        },
        "/questions/{id}/select": {
            "post": {
                "security": [
                    {
                        "CodeAuth": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Select Question",
                "operationId": "select-question",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Question ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Success",
                        "schema": {
                            "$ref": "#/definitions/models.Question"
                        }
                    },
                    "401": {
                        "description": "Unauthorised"
                    },
                    "404": {
                        "description": "Question not found"
                    },
                    "418": {
                        "description": "Question has already been shown before"
                    },
                    "500": {
                        "description": "Unknown error"
                    }
                }
            }
        },
        "/reset": {
            "post": {
                "security": [
                    {
                        "CodeAuth": []
                    }
                ],
                "summary": "Reset Game State",
                "operationId": "reset-game-state",
                "responses": {
                    "204": {
                        "description": "Successfully reset the game state"
                    },
                    "401": {
                        "description": "Unauthorised"
                    },
                    "500": {
                        "description": "Failed to reset game state"
                    }
                }
            }
        }
    },
    "definitions": {
        "models.Answer": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string",
                    "readOnly": true
                },
                "points": {
                    "type": "integer",
                    "maximum": 100,
                    "minimum": 1
                },
                "revealed": {
                    "type": "boolean",
                    "readOnly": true
                },
                "text": {
                    "type": "string"
                }
            }
        },
        "models.Question": {
            "type": "object",
            "properties": {
                "answers": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Answer"
                    }
                },
                "hasBeenShown": {
                    "type": "boolean",
                    "readOnly": true
                },
                "id": {
                    "type": "string",
                    "readOnly": true
                },
                "text": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "CodeAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
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
	Title:       "Brood-Clash API",
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
