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
    "description": "Description for http-request",
    "title": "http-request",
    "version": "1.0.0",
    "x-direktiv": {
      "category": "Unknown",
      "container": "direktiv/http-request",
      "license": "[https://www.apache.org/licenses/LICENSE-2.0](https://www.apache.org/licenses/LICENSE-2.0)",
      "long-description": "This is a longer description for the application http-request",
      "maintainer": null
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
                "url"
              ],
              "properties": {
                "body": {
                  "description": "Defines the payload of the request. The ` + "`" + `kind` + "`" + ` value can have three different values: \n- string: Plain string payload, e.g. JSON\n- base64: Will be \"converted\" to binary and attached\n- file: File payload",
                  "type": "object",
                  "properties": {
                    "kind": {
                      "type": "string",
                      "enum": [
                        "string",
                        "file",
                        "base64"
                      ]
                    },
                    "value": {
                      "type": "string"
                    }
                  },
                  "example": {
                    "kind": "string",
                    "value": "This is the payload"
                  }
                },
                "error200": {
                  "description": "If set to ` + "`" + `true` + "`" + ` responses with status above 299 will be treated as errors.",
                  "type": "boolean",
                  "example": true
                },
                "headers": {
                  "description": "List of key/values send as headers with the request.",
                  "type": "object",
                  "additionalProperties": {
                    "type": "string"
                  },
                  "example": {
                    "myheader": "value"
                  }
                },
                "insecure": {
                  "description": "Skips the verification the server certificate chain and host name.",
                  "type": "boolean",
                  "example": true
                },
                "method": {
                  "description": "HTTP method. Defaults to GET.",
                  "type": "string",
                  "enum": [
                    "GET",
                    "HEAD",
                    "POST",
                    "PUT",
                    "DELETE",
                    "OPTIONS",
                    "CONNECT",
                    "TRACE",
                    "PATCH"
                  ],
                  "example": "POST"
                },
                "params": {
                  "description": "List of key/values appended to URL as query parameters.",
                  "type": "object",
                  "additionalProperties": {
                    "type": "string"
                  },
                  "example": {
                    "query1": "queryvalue"
                  }
                },
                "password": {
                  "description": "If username and password are set, it will be used for basic authenitcation for the request. \nThis should be passed in as Direktiv secret.",
                  "type": "string",
                  "example": "mypassword"
                },
                "url": {
                  "description": "URL for the request.",
                  "type": "string",
                  "example": "http://www.direktiv.io"
                },
                "username": {
                  "description": "If username and password are set, it will be used for basic authenitcation for the request.",
                  "type": "string",
                  "example": "myuser"
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
              "action": "http",
              "data": {
                "kind": "string",
                "value": "HELLWORLD {{ .Headers }}"
              },
              "headers": [
                {
                  "jens": "gerke1"
                },
                {
                  "content-type": "text/plain"
                }
              ],
              "method": "post",
              "url": "https://webhook.site/38c796f4-bb9e-4d81-aad5-9e15e3cd5f4f"
            }
          ],
          "debug": true
        }
      },
      "delete": {
        "parameters": [
          {
            "type": "string",
            "description": "On cancel Direktiv sends a DELETE request to\nthe action with id in the header\n",
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
          "cancel": "echo 'cancel {{ .DirektivActionID }}'"
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
    "description": "Description for http-request",
    "title": "http-request",
    "version": "1.0.0",
    "x-direktiv": {
      "category": "Unknown",
      "container": "direktiv/http-request",
      "license": "[https://www.apache.org/licenses/LICENSE-2.0](https://www.apache.org/licenses/LICENSE-2.0)",
      "long-description": "This is a longer description for the application http-request",
      "maintainer": null
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
                "url"
              ],
              "properties": {
                "body": {
                  "description": "Defines the payload of the request. The ` + "`" + `kind` + "`" + ` value can have three different values: \n- string: Plain string payload, e.g. JSON\n- base64: Will be \"converted\" to binary and attached\n- file: File payload",
                  "type": "object",
                  "properties": {
                    "kind": {
                      "type": "string",
                      "enum": [
                        "string",
                        "file",
                        "base64"
                      ]
                    },
                    "value": {
                      "type": "string"
                    }
                  },
                  "example": {
                    "kind": "string",
                    "value": "This is the payload"
                  }
                },
                "error200": {
                  "description": "If set to ` + "`" + `true` + "`" + ` responses with status above 299 will be treated as errors.",
                  "type": "boolean",
                  "example": true
                },
                "headers": {
                  "description": "List of key/values send as headers with the request.",
                  "type": "object",
                  "additionalProperties": {
                    "type": "string"
                  },
                  "example": {
                    "myheader": "value"
                  }
                },
                "insecure": {
                  "description": "Skips the verification the server certificate chain and host name.",
                  "type": "boolean",
                  "example": true
                },
                "method": {
                  "description": "HTTP method. Defaults to GET.",
                  "type": "string",
                  "enum": [
                    "GET",
                    "HEAD",
                    "POST",
                    "PUT",
                    "DELETE",
                    "OPTIONS",
                    "CONNECT",
                    "TRACE",
                    "PATCH"
                  ],
                  "example": "POST"
                },
                "params": {
                  "description": "List of key/values appended to URL as query parameters.",
                  "type": "object",
                  "additionalProperties": {
                    "type": "string"
                  },
                  "example": {
                    "query1": "queryvalue"
                  }
                },
                "password": {
                  "description": "If username and password are set, it will be used for basic authenitcation for the request. \nThis should be passed in as Direktiv secret.",
                  "type": "string",
                  "example": "mypassword"
                },
                "url": {
                  "description": "URL for the request.",
                  "type": "string",
                  "example": "http://www.direktiv.io"
                },
                "username": {
                  "description": "If username and password are set, it will be used for basic authenitcation for the request.",
                  "type": "string",
                  "example": "myuser"
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
              "action": "http",
              "data": {
                "kind": "string",
                "value": "HELLWORLD {{ .Headers }}"
              },
              "headers": [
                {
                  "jens": "gerke1"
                },
                {
                  "content-type": "text/plain"
                }
              ],
              "method": "post",
              "url": "https://webhook.site/38c796f4-bb9e-4d81-aad5-9e15e3cd5f4f"
            }
          ],
          "debug": true
        }
      },
      "delete": {
        "parameters": [
          {
            "type": "string",
            "description": "On cancel Direktiv sends a DELETE request to\nthe action with id in the header\n",
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
          "cancel": "echo 'cancel {{ .DirektivActionID }}'"
        }
      }
    }
  },
  "definitions": {
    "PostParamsBodyBody": {
      "description": "Defines the payload of the request. The ` + "`" + `kind` + "`" + ` value can have three different values: \n- string: Plain string payload, e.g. JSON\n- base64: Will be \"converted\" to binary and attached\n- file: File payload",
      "type": "object",
      "properties": {
        "kind": {
          "type": "string",
          "enum": [
            "string",
            "file",
            "base64"
          ]
        },
        "value": {
          "type": "string"
        }
      },
      "example": {
        "kind": "string",
        "value": "This is the payload"
      }
    },
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
