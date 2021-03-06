// Code generated by go-swagger; DO NOT EDIT.

package restapi

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
)

var (
	// SwaggerJSON embedded version of the swagger document used at generation time
	SwaggerJSON json.RawMessage
	// FlatSwaggerJSON embedded flattened version of the swagger document used at generation time
	FlatSwaggerJSON json.RawMessage
)

func init() {
	SwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "Run AWS CLI commands in direktiv",
    "title": "aws-cli",
    "version": "1.0.0",
    "x-direktiv": {
      "category": "cloud",
      "container": "direktiv/aws-cli",
      "long-description": "This is a longer description for the application aws-cli"
    }
  },
  "paths": {
    "/": {
      "post": {
        "parameters": [
          {
            "type": "string",
            "description": "direktiv action id is an UUID. \nFor development it can be set to 'development'\n",
            "name": "Direktiv-ActionID",
            "in": "header"
          },
          {
            "type": "string",
            "description": "direktiv temp dir is the working directory for that request\nFor development it can be set to e.g. '/tmp'\n",
            "name": "Direktiv-TempDir",
            "in": "header"
          },
          {
            "name": "body",
            "in": "body",
            "schema": {
              "type": "object",
              "required": [
                "access-key",
                "secret-key"
              ],
              "properties": {
                "access-key": {
                  "description": "Amazon access key ID",
                  "type": "string",
                  "example": "BBQJTV4BBQJTNWBBQJY5"
                },
                "commands": {
                  "type": "array",
                  "items": {
                    "type": "string"
                  },
                  "example": [
                    "ec2 create-security-group --group-name secgroup1 --description mySecGroup",
                    "ec2 authorize-security-group-ingress --group-name secgroup1 --cidr 0.0.0.0/0 --protocol tcp --port 443"
                  ]
                },
                "region": {
                  "description": "Region in which to execute command",
                  "type": "string",
                  "default": "us-east-1",
                  "example": "us-east-1"
                },
                "secret-key": {
                  "description": "Amazon secret access key",
                  "type": "string",
                  "example": "ezaEezaEs0q5WezaEs0q5ezaEs0q5aEs0ezaEs0q5"
                }
              }
            }
          }
        ],
        "responses": {
          "200": {
            "description": "nice greeting",
            "schema": {
              "type": "object",
              "additionalProperties": false,
              "example": {
                "greeting": "Hello YourName"
              }
            }
          },
          "default": {
            "description": "generic error response",
            "schema": {
              "$ref": "#/definitions/error"
            },
            "headers": {
              "Direktiv-ErrorCode": {
                "type": "string"
              },
              "Direktiv-ErrorMessage": {
                "type": "string"
              }
            }
          }
        },
        "x-direktiv": {
          "cmds": [
            {
              "action": "exec",
              "exec": "echo '{{ fileExists \"jens\" }}'"
            }
          ],
          "output": "{\n  \"greeting\": \"{{ index (index . \"cmd0\") \"result\" }}\"\n}\n"
        }
      },
      "delete": {
        "parameters": [
          {
            "type": "string",
            "name": "Direktiv-ActionID",
            "in": "header"
          }
        ],
        "responses": {
          "200": {
            "description": ""
          }
        },
        "x-direktiv": {
          "cancel": "echo 'cancelling action id {{ .DirektivActionID }}'"
        }
      }
    }
  },
  "definitions": {
    "direktivFile": {
      "type": "object",
      "x-go-type": {
        "import": {
          "package": "github.com/direktiv/apps/go/pkg/apps"
        },
        "type": "DirektivFile"
      }
    },
    "error": {
      "type": "object",
      "required": [
        "errorCode",
        "errorMessage"
      ],
      "properties": {
        "errorCode": {
          "type": "string"
        },
        "errorMessage": {
          "type": "string"
        }
      }
    }
  }
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "Run AWS CLI commands in direktiv",
    "title": "aws-cli",
    "version": "1.0.0",
    "x-direktiv": {
      "category": "cloud",
      "container": "direktiv/aws-cli",
      "long-description": "This is a longer description for the application aws-cli"
    }
  },
  "paths": {
    "/": {
      "post": {
        "parameters": [
          {
            "type": "string",
            "description": "direktiv action id is an UUID. \nFor development it can be set to 'development'\n",
            "name": "Direktiv-ActionID",
            "in": "header"
          },
          {
            "type": "string",
            "description": "direktiv temp dir is the working directory for that request\nFor development it can be set to e.g. '/tmp'\n",
            "name": "Direktiv-TempDir",
            "in": "header"
          },
          {
            "name": "body",
            "in": "body",
            "schema": {
              "type": "object",
              "required": [
                "access-key",
                "secret-key"
              ],
              "properties": {
                "access-key": {
                  "description": "Amazon access key ID",
                  "type": "string",
                  "example": "BBQJTV4BBQJTNWBBQJY5"
                },
                "commands": {
                  "type": "array",
                  "items": {
                    "type": "string"
                  },
                  "example": [
                    "ec2 create-security-group --group-name secgroup1 --description mySecGroup",
                    "ec2 authorize-security-group-ingress --group-name secgroup1 --cidr 0.0.0.0/0 --protocol tcp --port 443"
                  ]
                },
                "region": {
                  "description": "Region in which to execute command",
                  "type": "string",
                  "default": "us-east-1",
                  "example": "us-east-1"
                },
                "secret-key": {
                  "description": "Amazon secret access key",
                  "type": "string",
                  "example": "ezaEezaEs0q5WezaEs0q5ezaEs0q5aEs0ezaEs0q5"
                }
              }
            }
          }
        ],
        "responses": {
          "200": {
            "description": "nice greeting",
            "schema": {
              "type": "object",
              "additionalProperties": false,
              "example": {
                "greeting": "Hello YourName"
              }
            }
          },
          "default": {
            "description": "generic error response",
            "schema": {
              "$ref": "#/definitions/error"
            },
            "headers": {
              "Direktiv-ErrorCode": {
                "type": "string"
              },
              "Direktiv-ErrorMessage": {
                "type": "string"
              }
            }
          }
        },
        "x-direktiv": {
          "cmds": [
            {
              "action": "exec",
              "exec": "echo '{{ fileExists \"jens\" }}'"
            }
          ],
          "output": "{\n  \"greeting\": \"{{ index (index . \"cmd0\") \"result\" }}\"\n}\n"
        }
      },
      "delete": {
        "parameters": [
          {
            "type": "string",
            "name": "Direktiv-ActionID",
            "in": "header"
          }
        ],
        "responses": {
          "200": {
            "description": ""
          }
        },
        "x-direktiv": {
          "cancel": "echo 'cancelling action id {{ .DirektivActionID }}'"
        }
      }
    }
  },
  "definitions": {
    "direktivFile": {
      "type": "object",
      "x-go-type": {
        "import": {
          "package": "github.com/direktiv/apps/go/pkg/apps"
        },
        "type": "DirektivFile"
      }
    },
    "error": {
      "type": "object",
      "required": [
        "errorCode",
        "errorMessage"
      ],
      "properties": {
        "errorCode": {
          "type": "string"
        },
        "errorMessage": {
          "type": "string"
        }
      }
    }
  }
}`))
}
