// Code generated by go-swagger; DO NOT EDIT.

// This file is part of MinIO Kubernetes Cloud
// Copyright (c) 2020 MinIO, Inc.
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.
//

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
    "title": "MinIO For Kubernetes",
    "version": "0.1.0"
  },
  "basePath": "/api/v1",
  "paths": {
    "/namespaces/{namespace}/storage-pools-info": {
      "get": {
        "tags": [
          "AdminAPI"
        ],
        "summary": "Get info of all StoragePools by type",
        "operationId": "GetStoragePoolsInfo",
        "parameters": [
          {
            "type": "string",
            "name": "namespace",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/storagePoolsInfo"
            }
          },
          "default": {
            "description": "Generic error response.",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/namespaces/{namespace}/tenants": {
      "get": {
        "tags": [
          "AdminAPI"
        ],
        "summary": "List Tenants by Namespace",
        "operationId": "ListTenants",
        "parameters": [
          {
            "type": "string",
            "name": "namespace",
            "in": "path",
            "required": true
          },
          {
            "type": "string",
            "name": "sort_by",
            "in": "query"
          },
          {
            "type": "integer",
            "format": "int32",
            "name": "offset",
            "in": "query"
          },
          {
            "type": "integer",
            "format": "int32",
            "name": "limit",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/listTenantsResponse"
            }
          },
          "default": {
            "description": "Generic error response.",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/namespaces/{namespace}/tenants/{tenant}": {
      "get": {
        "tags": [
          "AdminAPI"
        ],
        "summary": "Tenant Info",
        "operationId": "TenantInfo",
        "parameters": [
          {
            "type": "string",
            "name": "namespace",
            "in": "path",
            "required": true
          },
          {
            "type": "string",
            "name": "tenant",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/tenant"
            }
          },
          "default": {
            "description": "Generic error response.",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      },
      "put": {
        "tags": [
          "AdminAPI"
        ],
        "summary": "Update Tenant",
        "operationId": "UpdateTenant",
        "parameters": [
          {
            "type": "string",
            "name": "namespace",
            "in": "path",
            "required": true
          },
          {
            "type": "string",
            "name": "tenant",
            "in": "path",
            "required": true
          },
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/updateTenantRequest"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "A successful response."
          },
          "default": {
            "description": "Generic error response.",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      },
      "delete": {
        "tags": [
          "AdminAPI"
        ],
        "summary": "Delete Tenant",
        "operationId": "DeleteTenant",
        "parameters": [
          {
            "type": "string",
            "name": "namespace",
            "in": "path",
            "required": true
          },
          {
            "type": "string",
            "name": "tenant",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "204": {
            "description": "A successful response."
          },
          "default": {
            "description": "Generic error response.",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/storage-classes": {
      "get": {
        "tags": [
          "AdminAPI"
        ],
        "summary": "List Storage Classes",
        "operationId": "ListStorageClasses",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/storageClasses"
            }
          },
          "default": {
            "description": "Generic error response.",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/tenants": {
      "get": {
        "tags": [
          "AdminAPI"
        ],
        "summary": "List Tenant of All Namespaces",
        "operationId": "ListAllTenants",
        "parameters": [
          {
            "type": "string",
            "name": "sort_by",
            "in": "query"
          },
          {
            "type": "integer",
            "format": "int32",
            "name": "offset",
            "in": "query"
          },
          {
            "type": "integer",
            "format": "int32",
            "name": "limit",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/listTenantsResponse"
            }
          },
          "default": {
            "description": "Generic error response.",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      },
      "post": {
        "tags": [
          "AdminAPI"
        ],
        "summary": "Create Tenant",
        "operationId": "CreateTenant",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/createTenantRequest"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "A successful response."
          },
          "default": {
            "description": "Generic error response.",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "createTenantRequest": {
      "type": "object",
      "required": [
        "name",
        "volume_configuration",
        "namespace"
      ],
      "properties": {
        "access_key": {
          "type": "string"
        },
        "enable_mcs": {
          "type": "boolean",
          "default": true
        },
        "enable_ssl": {
          "type": "boolean",
          "default": true
        },
        "image": {
          "type": "string"
        },
        "mounth_path": {
          "type": "string"
        },
        "name": {
          "type": "string",
          "pattern": "^[a-z0-9-]{3,63}$"
        },
        "namespace": {
          "type": "string"
        },
        "secret_key": {
          "type": "string"
        },
        "service_name": {
          "type": "string"
        },
        "volume_configuration": {
          "type": "object",
          "required": [
            "size"
          ],
          "properties": {
            "size": {
              "type": "string"
            },
            "storage_class": {
              "type": "string"
            }
          }
        },
        "volumes_per_server": {
          "type": "integer"
        },
        "zones": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/zone"
          }
        }
      }
    },
    "error": {
      "type": "object",
      "required": [
        "message"
      ],
      "properties": {
        "code": {
          "type": "integer",
          "format": "int64"
        },
        "message": {
          "type": "string"
        }
      }
    },
    "listTenantsResponse": {
      "type": "object",
      "properties": {
        "tenants": {
          "type": "array",
          "title": "list of resulting tenants",
          "items": {
            "$ref": "#/definitions/tenantList"
          }
        },
        "total": {
          "type": "integer",
          "format": "int64",
          "title": "number of tenants accessible to tenant user"
        }
      }
    },
    "principal": {
      "type": "string"
    },
    "storageClasses": {
      "type": "array",
      "items": {
        "type": "string"
      }
    },
    "storagePool": {
      "type": "object",
      "properties": {
        "freeSpace": {
          "type": "integer",
          "format": "int64"
        },
        "name": {
          "type": "string"
        },
        "nodes": {
          "type": "array",
          "items": {
            "type": "string"
          }
        },
        "storageClasses": {
          "type": "object",
          "$ref": "#/definitions/storageClasses"
        },
        "totalSpace": {
          "type": "integer",
          "format": "int64"
        }
      }
    },
    "storagePoolType": {
      "type": "object",
      "properties": {
        "name": {
          "type": "string"
        },
        "storagePools": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/storagePool"
          }
        }
      }
    },
    "storagePoolsInfo": {
      "type": "object",
      "properties": {
        "types": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/storagePoolType"
          }
        }
      }
    },
    "tenant": {
      "type": "object",
      "properties": {
        "creation_date": {
          "type": "string"
        },
        "currentState": {
          "type": "string"
        },
        "instance_count": {
          "type": "integer"
        },
        "name": {
          "type": "string"
        },
        "namespace": {
          "type": "string"
        },
        "volume_count": {
          "type": "integer"
        },
        "volume_size": {
          "type": "integer"
        },
        "volumes_per_server": {
          "type": "integer"
        },
        "zone_count": {
          "type": "integer"
        },
        "zones": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/zone"
          }
        }
      }
    },
    "tenantList": {
      "type": "object",
      "properties": {
        "creation_date": {
          "type": "string"
        },
        "currentState": {
          "type": "string"
        },
        "instance_count": {
          "type": "integer"
        },
        "name": {
          "type": "string"
        },
        "namespace": {
          "type": "string"
        },
        "volume_count": {
          "type": "integer"
        },
        "volume_size": {
          "type": "integer"
        },
        "zone_count": {
          "type": "integer"
        }
      }
    },
    "updateTenantRequest": {
      "type": "object",
      "properties": {
        "image": {
          "type": "string",
          "pattern": "^((.*?)/(.*?):(.+))$"
        }
      }
    },
    "zone": {
      "type": "object",
      "properties": {
        "name": {
          "type": "string"
        },
        "servers": {
          "type": "integer"
        }
      }
    }
  },
  "securityDefinitions": {
    "key": {
      "type": "oauth2",
      "flow": "accessCode",
      "authorizationUrl": "http://min.io",
      "tokenUrl": "http://min.io"
    }
  },
  "security": [
    {
      "key": []
    }
  ]
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
    "title": "MinIO For Kubernetes",
    "version": "0.1.0"
  },
  "basePath": "/api/v1",
  "paths": {
    "/namespaces/{namespace}/storage-pools-info": {
      "get": {
        "tags": [
          "AdminAPI"
        ],
        "summary": "Get info of all StoragePools by type",
        "operationId": "GetStoragePoolsInfo",
        "parameters": [
          {
            "type": "string",
            "name": "namespace",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/storagePoolsInfo"
            }
          },
          "default": {
            "description": "Generic error response.",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/namespaces/{namespace}/tenants": {
      "get": {
        "tags": [
          "AdminAPI"
        ],
        "summary": "List Tenants by Namespace",
        "operationId": "ListTenants",
        "parameters": [
          {
            "type": "string",
            "name": "namespace",
            "in": "path",
            "required": true
          },
          {
            "type": "string",
            "name": "sort_by",
            "in": "query"
          },
          {
            "type": "integer",
            "format": "int32",
            "name": "offset",
            "in": "query"
          },
          {
            "type": "integer",
            "format": "int32",
            "name": "limit",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/listTenantsResponse"
            }
          },
          "default": {
            "description": "Generic error response.",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/namespaces/{namespace}/tenants/{tenant}": {
      "get": {
        "tags": [
          "AdminAPI"
        ],
        "summary": "Tenant Info",
        "operationId": "TenantInfo",
        "parameters": [
          {
            "type": "string",
            "name": "namespace",
            "in": "path",
            "required": true
          },
          {
            "type": "string",
            "name": "tenant",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/tenant"
            }
          },
          "default": {
            "description": "Generic error response.",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      },
      "put": {
        "tags": [
          "AdminAPI"
        ],
        "summary": "Update Tenant",
        "operationId": "UpdateTenant",
        "parameters": [
          {
            "type": "string",
            "name": "namespace",
            "in": "path",
            "required": true
          },
          {
            "type": "string",
            "name": "tenant",
            "in": "path",
            "required": true
          },
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/updateTenantRequest"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "A successful response."
          },
          "default": {
            "description": "Generic error response.",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      },
      "delete": {
        "tags": [
          "AdminAPI"
        ],
        "summary": "Delete Tenant",
        "operationId": "DeleteTenant",
        "parameters": [
          {
            "type": "string",
            "name": "namespace",
            "in": "path",
            "required": true
          },
          {
            "type": "string",
            "name": "tenant",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "204": {
            "description": "A successful response."
          },
          "default": {
            "description": "Generic error response.",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/storage-classes": {
      "get": {
        "tags": [
          "AdminAPI"
        ],
        "summary": "List Storage Classes",
        "operationId": "ListStorageClasses",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/storageClasses"
            }
          },
          "default": {
            "description": "Generic error response.",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/tenants": {
      "get": {
        "tags": [
          "AdminAPI"
        ],
        "summary": "List Tenant of All Namespaces",
        "operationId": "ListAllTenants",
        "parameters": [
          {
            "type": "string",
            "name": "sort_by",
            "in": "query"
          },
          {
            "type": "integer",
            "format": "int32",
            "name": "offset",
            "in": "query"
          },
          {
            "type": "integer",
            "format": "int32",
            "name": "limit",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/listTenantsResponse"
            }
          },
          "default": {
            "description": "Generic error response.",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      },
      "post": {
        "tags": [
          "AdminAPI"
        ],
        "summary": "Create Tenant",
        "operationId": "CreateTenant",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/createTenantRequest"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "A successful response."
          },
          "default": {
            "description": "Generic error response.",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "CreateTenantRequestVolumeConfiguration": {
      "type": "object",
      "required": [
        "size"
      ],
      "properties": {
        "size": {
          "type": "string"
        },
        "storage_class": {
          "type": "string"
        }
      }
    },
    "createTenantRequest": {
      "type": "object",
      "required": [
        "name",
        "volume_configuration",
        "namespace"
      ],
      "properties": {
        "access_key": {
          "type": "string"
        },
        "enable_mcs": {
          "type": "boolean",
          "default": true
        },
        "enable_ssl": {
          "type": "boolean",
          "default": true
        },
        "image": {
          "type": "string"
        },
        "mounth_path": {
          "type": "string"
        },
        "name": {
          "type": "string",
          "pattern": "^[a-z0-9-]{3,63}$"
        },
        "namespace": {
          "type": "string"
        },
        "secret_key": {
          "type": "string"
        },
        "service_name": {
          "type": "string"
        },
        "volume_configuration": {
          "type": "object",
          "required": [
            "size"
          ],
          "properties": {
            "size": {
              "type": "string"
            },
            "storage_class": {
              "type": "string"
            }
          }
        },
        "volumes_per_server": {
          "type": "integer"
        },
        "zones": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/zone"
          }
        }
      }
    },
    "error": {
      "type": "object",
      "required": [
        "message"
      ],
      "properties": {
        "code": {
          "type": "integer",
          "format": "int64"
        },
        "message": {
          "type": "string"
        }
      }
    },
    "listTenantsResponse": {
      "type": "object",
      "properties": {
        "tenants": {
          "type": "array",
          "title": "list of resulting tenants",
          "items": {
            "$ref": "#/definitions/tenantList"
          }
        },
        "total": {
          "type": "integer",
          "format": "int64",
          "title": "number of tenants accessible to tenant user"
        }
      }
    },
    "principal": {
      "type": "string"
    },
    "storageClasses": {
      "type": "array",
      "items": {
        "type": "string"
      }
    },
    "storagePool": {
      "type": "object",
      "properties": {
        "freeSpace": {
          "type": "integer",
          "format": "int64"
        },
        "name": {
          "type": "string"
        },
        "nodes": {
          "type": "array",
          "items": {
            "type": "string"
          }
        },
        "storageClasses": {
          "type": "object",
          "$ref": "#/definitions/storageClasses"
        },
        "totalSpace": {
          "type": "integer",
          "format": "int64"
        }
      }
    },
    "storagePoolType": {
      "type": "object",
      "properties": {
        "name": {
          "type": "string"
        },
        "storagePools": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/storagePool"
          }
        }
      }
    },
    "storagePoolsInfo": {
      "type": "object",
      "properties": {
        "types": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/storagePoolType"
          }
        }
      }
    },
    "tenant": {
      "type": "object",
      "properties": {
        "creation_date": {
          "type": "string"
        },
        "currentState": {
          "type": "string"
        },
        "instance_count": {
          "type": "integer"
        },
        "name": {
          "type": "string"
        },
        "namespace": {
          "type": "string"
        },
        "volume_count": {
          "type": "integer"
        },
        "volume_size": {
          "type": "integer"
        },
        "volumes_per_server": {
          "type": "integer"
        },
        "zone_count": {
          "type": "integer"
        },
        "zones": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/zone"
          }
        }
      }
    },
    "tenantList": {
      "type": "object",
      "properties": {
        "creation_date": {
          "type": "string"
        },
        "currentState": {
          "type": "string"
        },
        "instance_count": {
          "type": "integer"
        },
        "name": {
          "type": "string"
        },
        "namespace": {
          "type": "string"
        },
        "volume_count": {
          "type": "integer"
        },
        "volume_size": {
          "type": "integer"
        },
        "zone_count": {
          "type": "integer"
        }
      }
    },
    "updateTenantRequest": {
      "type": "object",
      "properties": {
        "image": {
          "type": "string",
          "pattern": "^((.*?)/(.*?):(.+))$"
        }
      }
    },
    "zone": {
      "type": "object",
      "properties": {
        "name": {
          "type": "string"
        },
        "servers": {
          "type": "integer"
        }
      }
    }
  },
  "securityDefinitions": {
    "key": {
      "type": "oauth2",
      "flow": "accessCode",
      "authorizationUrl": "http://min.io",
      "tokenUrl": "http://min.io"
    }
  },
  "security": [
    {
      "key": []
    }
  ]
}`))
}
