# Introduction

This is the backend service to store cards into a NOSQL database

# Quickstart

Run the server `go run main.go`

The swagger documentation can be found under _localhost/swagger_

The endpoints available are listed below:

- url: localhost:8080/create-card

  - method: POST
  - data: {
    "FirstName": "Charles",
    "LastName": "Lau",
    "Role": "Waitress",
    "Company": "Kith café",
    "Country": "Singapore",
    "PhoneNumber": "09893448",
    "Website": "www.kith.org"
    }

- url: localhost:8080/ws

  - method: GET

- url: localhost:8080/upload-card
  - method: POST
  - data: multipart/file
