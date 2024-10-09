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
        "license": {
            "name": "Apache 2.0",
            "url": "https://github.com/thirdweb-dev/indexer/blob/main/LICENSE"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/{chainId}/events": {
            "get": {
                "security": [
                    {
                        "BasicAuth": []
                    }
                ],
                "description": "Retrieve all logs across all contracts",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "events"
                ],
                "summary": "Get all logs",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Chain ID",
                        "name": "chainId",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Filter parameters",
                        "name": "filter",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Field to group results by",
                        "name": "group_by",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Field to sort results by",
                        "name": "sort_by",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Sort order (asc or desc)",
                        "name": "sort_order",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Page number for pagination",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Number of items per page",
                        "name": "limit",
                        "in": "query"
                    },
                    {
                        "type": "array",
                        "items": {
                            "type": "string"
                        },
                        "collectionFormat": "csv",
                        "description": "List of aggregate functions to apply",
                        "name": "aggregate",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/api.QueryResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "array",
                                            "items": {
                                                "$ref": "#/definitions/handlers.LogModel"
                                            }
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/api.Error"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/api.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/api.Error"
                        }
                    }
                }
            }
        },
        "/{chainId}/events/{contract}": {
            "get": {
                "security": [
                    {
                        "BasicAuth": []
                    }
                ],
                "description": "Retrieve logs for a specific contract",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "events"
                ],
                "summary": "Get logs by contract",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Chain ID",
                        "name": "chainId",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Contract address",
                        "name": "contract",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Filter parameters",
                        "name": "filter",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Field to group results by",
                        "name": "group_by",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Field to sort results by",
                        "name": "sort_by",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Sort order (asc or desc)",
                        "name": "sort_order",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Page number for pagination",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Number of items per page",
                        "name": "limit",
                        "in": "query"
                    },
                    {
                        "type": "array",
                        "items": {
                            "type": "string"
                        },
                        "collectionFormat": "csv",
                        "description": "List of aggregate functions to apply",
                        "name": "aggregate",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/api.QueryResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "array",
                                            "items": {
                                                "$ref": "#/definitions/handlers.LogModel"
                                            }
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/api.Error"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/api.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/api.Error"
                        }
                    }
                }
            }
        },
        "/{chainId}/events/{contract}/{signature}": {
            "get": {
                "security": [
                    {
                        "BasicAuth": []
                    }
                ],
                "description": "Retrieve logs for a specific contract and event signature",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "events"
                ],
                "summary": "Get logs by contract and event signature",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Chain ID",
                        "name": "chainId",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Contract address",
                        "name": "contract",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Event signature",
                        "name": "signature",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Filter parameters",
                        "name": "filter",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Field to group results by",
                        "name": "group_by",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Field to sort results by",
                        "name": "sort_by",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Sort order (asc or desc)",
                        "name": "sort_order",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Page number for pagination",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Number of items per page",
                        "name": "limit",
                        "in": "query"
                    },
                    {
                        "type": "array",
                        "items": {
                            "type": "string"
                        },
                        "collectionFormat": "csv",
                        "description": "List of aggregate functions to apply",
                        "name": "aggregate",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/api.QueryResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "array",
                                            "items": {
                                                "$ref": "#/definitions/handlers.LogModel"
                                            }
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/api.Error"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/api.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/api.Error"
                        }
                    }
                }
            }
        },
        "/{chainId}/transactions": {
            "get": {
                "security": [
                    {
                        "BasicAuth": []
                    }
                ],
                "description": "Retrieve all transactions across all contracts",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "transactions"
                ],
                "summary": "Get all transactions",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Chain ID",
                        "name": "chainId",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Filter parameters",
                        "name": "filter",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Field to group results by",
                        "name": "group_by",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Field to sort results by",
                        "name": "sort_by",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Sort order (asc or desc)",
                        "name": "sort_order",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Page number for pagination",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Number of items per page",
                        "name": "limit",
                        "in": "query"
                    },
                    {
                        "type": "array",
                        "items": {
                            "type": "string"
                        },
                        "collectionFormat": "csv",
                        "description": "List of aggregate functions to apply",
                        "name": "aggregate",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/api.QueryResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "array",
                                            "items": {
                                                "$ref": "#/definitions/handlers.TransactionModel"
                                            }
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/api.Error"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/api.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/api.Error"
                        }
                    }
                }
            }
        },
        "/{chainId}/transactions/{to}": {
            "get": {
                "security": [
                    {
                        "BasicAuth": []
                    }
                ],
                "description": "Retrieve transactions for a specific contract",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "transactions"
                ],
                "summary": "Get transactions by contract",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Chain ID",
                        "name": "chainId",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Contract address",
                        "name": "to",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Filter parameters",
                        "name": "filter",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Field to group results by",
                        "name": "group_by",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Field to sort results by",
                        "name": "sort_by",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Sort order (asc or desc)",
                        "name": "sort_order",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Page number for pagination",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Number of items per page",
                        "name": "limit",
                        "in": "query"
                    },
                    {
                        "type": "array",
                        "items": {
                            "type": "string"
                        },
                        "collectionFormat": "csv",
                        "description": "List of aggregate functions to apply",
                        "name": "aggregate",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/api.QueryResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "array",
                                            "items": {
                                                "$ref": "#/definitions/handlers.TransactionModel"
                                            }
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/api.Error"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/api.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/api.Error"
                        }
                    }
                }
            }
        },
        "/{chainId}/transactions/{to}/{signature}": {
            "get": {
                "security": [
                    {
                        "BasicAuth": []
                    }
                ],
                "description": "Retrieve transactions for a specific contract and signature (Not implemented yet)",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "transactions"
                ],
                "summary": "Get transactions by contract and signature",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Chain ID",
                        "name": "chainId",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Contract address",
                        "name": "to",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Function signature",
                        "name": "signature",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Filter parameters",
                        "name": "filter",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Field to group results by",
                        "name": "group_by",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Field to sort results by",
                        "name": "sort_by",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Sort order (asc or desc)",
                        "name": "sort_order",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Page number for pagination",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Number of items per page",
                        "name": "limit",
                        "in": "query"
                    },
                    {
                        "type": "array",
                        "items": {
                            "type": "string"
                        },
                        "collectionFormat": "csv",
                        "description": "List of aggregate functions to apply",
                        "name": "aggregate",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/api.QueryResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "array",
                                            "items": {
                                                "$ref": "#/definitions/handlers.TransactionModel"
                                            }
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/api.Error"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/api.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/api.Error"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "api.Error": {
            "description": "Error represents an API error response",
            "type": "object",
            "properties": {
                "code": {
                    "description": "@Description HTTP status code",
                    "type": "integer"
                },
                "message": {
                    "description": "@Description Error message",
                    "type": "string"
                },
                "support_id": {
                    "description": "@Description Support ID for tracking the error",
                    "type": "string"
                }
            }
        },
        "api.Meta": {
            "description": "Meta represents metadata for a query response",
            "type": "object",
            "properties": {
                "address": {
                    "description": "@Description Contract address",
                    "type": "string"
                },
                "chain_id": {
                    "description": "@Description Chain ID of the blockchain",
                    "type": "integer"
                },
                "limit": {
                    "description": "@Description Number of items per page",
                    "type": "integer"
                },
                "page": {
                    "description": "@Description Current page number",
                    "type": "integer"
                },
                "signature": {
                    "description": "@Description Function or event signature",
                    "type": "string"
                },
                "total_items": {
                    "description": "@Description Total number of items",
                    "type": "integer"
                },
                "total_pages": {
                    "description": "@Description Total number of pages",
                    "type": "integer"
                }
            }
        },
        "api.QueryResponse": {
            "description": "QueryResponse represents the response structure for a query",
            "type": "object",
            "properties": {
                "aggregations": {
                    "description": "@Description Aggregation results",
                    "type": "object",
                    "additionalProperties": {
                        "type": "string"
                    }
                },
                "data": {
                    "description": "@Description Query result data"
                },
                "meta": {
                    "description": "@Description Metadata for the query response",
                    "allOf": [
                        {
                            "$ref": "#/definitions/api.Meta"
                        }
                    ]
                }
            }
        },
        "handlers.LogModel": {
            "type": "object",
            "properties": {
                "address": {
                    "type": "string"
                },
                "block_hash": {
                    "type": "string"
                },
                "block_number": {
                    "type": "string"
                },
                "block_timestamp": {
                    "type": "integer"
                },
                "chain_id": {
                    "type": "string"
                },
                "data": {
                    "type": "string"
                },
                "log_index": {
                    "type": "integer"
                },
                "topics": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "transaction_hash": {
                    "type": "string"
                },
                "transaction_index": {
                    "type": "integer"
                }
            }
        },
        "handlers.TransactionModel": {
            "type": "object",
            "properties": {
                "access_list_json": {
                    "type": "string"
                },
                "block_hash": {
                    "type": "string"
                },
                "block_number": {
                    "type": "string"
                },
                "block_timestamp": {
                    "type": "integer"
                },
                "chain_id": {
                    "type": "string"
                },
                "data": {
                    "type": "string"
                },
                "from_address": {
                    "type": "string"
                },
                "gas": {
                    "type": "integer"
                },
                "gas_price": {
                    "type": "string"
                },
                "hash": {
                    "type": "string"
                },
                "max_fee_per_gas": {
                    "type": "string"
                },
                "max_priority_fee_per_gas": {
                    "type": "string"
                },
                "nonce": {
                    "type": "integer"
                },
                "r": {
                    "type": "string"
                },
                "s": {
                    "type": "string"
                },
                "to_address": {
                    "type": "string"
                },
                "transaction_index": {
                    "type": "integer"
                },
                "transaction_type": {
                    "type": "integer"
                },
                "v": {
                    "type": "string"
                },
                "value": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "BasicAuth": {
            "type": "basic"
        }
    },
    "security": [
        {
            "BasicAuth": []
        }
    ]
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "v0.0.1-beta",
	Host:             "",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Thirdweb Insight",
	Description:      "API for querying blockchain transactions and events",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
