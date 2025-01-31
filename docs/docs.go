// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://localhost:8080/swagger/index.html",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/address/geocode": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "GEO_Data"
                ],
                "summary": "Search for an GEO",
                "operationId": "GEO_address",
                "parameters": [
                    {
                        "description": "Geocode request",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.GeocodeRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "get data",
                        "schema": {
                            "$ref": "#/definitions/models.AddressGeo"
                        }
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        },
        "/address/search": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "GEO_Data"
                ],
                "summary": "Search for an address",
                "operationId": "search_address",
                "parameters": [
                    {
                        "description": "Search request",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.SearchRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "get data",
                        "schema": {
                            "$ref": "#/definitions/models.AddressSearch"
                        }
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        },
        "/login": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "authorization"
                ],
                "summary": "SingIn a user",
                "operationId": "SingIn",
                "parameters": [
                    {
                        "description": "User",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.User"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "JWT token"
                    },
                    "400": {
                        "description": "Invalid request format"
                    },
                    "500": {
                        "description": "Response writer error on write"
                    }
                }
            }
        },
        "/register": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "authorization"
                ],
                "summary": "Register a user",
                "operationId": "SingUp",
                "parameters": [
                    {
                        "description": "User",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.User"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "User registered successfully"
                    },
                    "400": {
                        "description": "Invalid request format"
                    },
                    "500": {
                        "description": "Response writer error on write"
                    }
                }
            }
        }
    },
    "definitions": {
        "models.AddressGeo": {
            "description": "AddressGeo represents the geocode result for an address.",
            "type": "object",
            "properties": {
                "suggestions": {
                    "type": "array",
                    "items": {
                        "type": "object",
                        "properties": {
                            "data": {
                                "type": "object",
                                "properties": {
                                    "country": {
                                        "type": "string"
                                    },
                                    "postal_code": {
                                        "type": "string"
                                    }
                                }
                            },
                            "unrestricted_value": {
                                "type": "string"
                            },
                            "value": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        },
        "models.AddressSearch": {
            "description": "AddressSearch represents the search result for an address.",
            "type": "object",
            "properties": {
                "area": {},
                "area_fias_id": {},
                "area_kladr_id": {},
                "area_type": {},
                "area_type_full": {},
                "area_with_type": {},
                "beltway_distance": {},
                "beltway_hit": {
                    "type": "string"
                },
                "block": {},
                "block_type": {},
                "block_type_full": {},
                "capital_marker": {
                    "type": "string"
                },
                "city": {},
                "city_area": {
                    "type": "string"
                },
                "city_district": {
                    "type": "string"
                },
                "city_district_fias_id": {},
                "city_district_kladr_id": {},
                "city_district_type": {
                    "type": "string"
                },
                "city_district_type_full": {
                    "type": "string"
                },
                "city_district_with_type": {
                    "type": "string"
                },
                "city_fias_id": {},
                "city_kladr_id": {},
                "city_type": {},
                "city_type_full": {},
                "city_with_type": {},
                "country": {
                    "type": "string"
                },
                "country_iso_code": {
                    "type": "string"
                },
                "entrance": {},
                "federal_district": {
                    "type": "string"
                },
                "fias_actuality_state": {
                    "type": "string"
                },
                "fias_code": {
                    "type": "string"
                },
                "fias_id": {
                    "type": "string"
                },
                "fias_level": {
                    "type": "string"
                },
                "flat": {
                    "type": "string"
                },
                "flat_area": {
                    "type": "string"
                },
                "flat_cadnum": {
                    "type": "string"
                },
                "flat_fias_id": {
                    "type": "string"
                },
                "flat_price": {
                    "type": "string"
                },
                "flat_type": {
                    "type": "string"
                },
                "flat_type_full": {
                    "type": "string"
                },
                "floor": {},
                "geo_lat": {
                    "type": "string"
                },
                "geo_lon": {
                    "type": "string"
                },
                "house": {
                    "type": "string"
                },
                "house_cadnum": {
                    "type": "string"
                },
                "house_fias_id": {
                    "type": "string"
                },
                "house_kladr_id": {
                    "type": "string"
                },
                "house_type": {
                    "type": "string"
                },
                "house_type_full": {
                    "type": "string"
                },
                "kladr_id": {
                    "type": "string"
                },
                "metro": {
                    "type": "array",
                    "items": {
                        "type": "object",
                        "properties": {
                            "distance": {
                                "type": "number"
                            },
                            "line": {
                                "type": "string"
                            },
                            "name": {
                                "type": "string"
                            }
                        }
                    }
                },
                "okato": {
                    "type": "string"
                },
                "oktmo": {
                    "type": "string"
                },
                "postal_box": {},
                "postal_code": {
                    "type": "string"
                },
                "qc": {
                    "type": "integer"
                },
                "qc_complete": {
                    "type": "integer"
                },
                "qc_geo": {
                    "type": "integer"
                },
                "qc_house": {
                    "type": "integer"
                },
                "region": {
                    "type": "string"
                },
                "region_fias_id": {
                    "type": "string"
                },
                "region_iso_code": {
                    "type": "string"
                },
                "region_kladr_id": {
                    "type": "string"
                },
                "region_type": {
                    "type": "string"
                },
                "region_type_full": {
                    "type": "string"
                },
                "region_with_type": {
                    "type": "string"
                },
                "result": {
                    "type": "string"
                },
                "settlement": {},
                "settlement_fias_id": {},
                "settlement_kladr_id": {},
                "settlement_type": {},
                "settlement_type_full": {},
                "settlement_with_type": {},
                "source": {
                    "type": "string"
                },
                "square_meter_price": {
                    "type": "string"
                },
                "stead": {},
                "stead_cadnum": {},
                "stead_fias_id": {},
                "stead_kladr_id": {},
                "stead_type": {},
                "stead_type_full": {},
                "street": {
                    "type": "string"
                },
                "street_fias_id": {
                    "type": "string"
                },
                "street_kladr_id": {
                    "type": "string"
                },
                "street_type": {
                    "type": "string"
                },
                "street_type_full": {
                    "type": "string"
                },
                "street_with_type": {
                    "type": "string"
                },
                "tax_office": {
                    "type": "string"
                },
                "tax_office_legal": {
                    "type": "string"
                },
                "timezone": {
                    "type": "string"
                },
                "unparsed_parts": {}
            }
        },
        "models.GeocodeRequest": {
            "description": "GeocodeRequest represents the request body for address geocoding.",
            "type": "object",
            "properties": {
                "lat": {
                    "type": "string"
                },
                "lon": {
                    "type": "string"
                }
            }
        },
        "models.SearchRequest": {
            "description": "SearchRequest represents the request body for address search",
            "type": "object",
            "properties": {
                "query": {
                    "type": "string"
                }
            }
        },
        "models.User": {
            "type": "object",
            "properties": {
                "login": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "ApiKeyAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "2.0",
	Host:             "localhost:8080",
	BasePath:         "/api",
	Schemes:          []string{},
	Title:            "GEO API",
	Description:      "This is a sample API for address searching and geocoding using Dadata API.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
