terraform {
  required_providers {
    prismacloudcompute = {
      source  = "github.com/terraform-providers/prismacloudcompute"
      version = "~> 1.0"
    }
  }
}

provider "prismacloudcompute" {
  url      = var.url
  username = var.username
  password = var.password
  port = var.port
}

resource "prismacloudcompute_collections" "example" {
    name = "CollectionName"
}

resource "prismacloudcompute_policiesruntimecontainer" "example" {
    name = "My Policy"
    policy_type = "network"
    rule {
        name = "my rule"
        criteria = "savedSearchId"
        parameters = {
            "savedSearch": "true",
            "withIac": "false",
        }
        rule_type = "Network"
    }
}

/*
resource "google_endpoints_service" "endpoints_service" {
  service_name = "echo-api.endpoints.${google_project.endpoints_project.project_id}.cloud.goog"
  project      = google_project.endpoints_project.project_id

  openapi_config = <<EOF
swagger: "2.0"
info:
  description: "A simple Google Cloud Endpoints API example."
  title: "Endpoints Example"
  version: "1.0.0"
host: "echo-api.endpoints.${google_project.endpoints_project.project_id}.cloud.goog"
basePath: "/"
consumes:
- "application/json"
produces:
- "application/json"
schemes:
- "https"
paths:
  "/echo":
    post:
      description: "Echo back a given message."
      operationId: "echo"
      produces:
      - "application/json"
      responses:
        200:
          description: "Echo"
          schema:
            $ref: "#/definitions/echoMessage"
      parameters:
      - description: "Message to echo"
        in: body
        name: message
        required: true
        schema:
          $ref: "#/definitions/echoMessage"
      security:
      - api_key: []
definitions:
  echoMessage:
    properties:
      message:
        type: "string"
EOF


  depends_on = [google_project_service.endpoints_project_sm]
}
*/
