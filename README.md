# account-toolkit

## About
Responsible for sharing common functionalities between Account services.

## Technologies
* Golang 1.18

## Development requirements
* Visual Studio Code

## Directory Structure
- `api`
    - OpenAPI/Swagger specs, JSON schema files, protocol definition files.
- `build`
    - It has all cloud package, container (Docker), operating system (deb, rpm, pkg) and scripts settings.
- `cmd`
    - It has the `main` function that imports and invokes code from the `/internal` and `/pkg` directories.
- `internal`
    - It has all the code that is not available for import.
- `local-dev`
    - Possui toda configuração do docker.

## Running
- `Docker`
    - Run the following command: docker-compose -f local-dev/docker-compose.yaml --profile infra up -d
- `Create subject kafka`
    - Create subject
    curl --location --request POST 'http://localhost:8081/subjects/com.account.create/versions' \
    --header 'Content-Type: application/json' \
    --data-raw '{
        "schema": "{\"type\":\"record\",\"name\":\"Account_Create\",\"namespace\":\"com.account.create\",\"fields\":[{\"name\":\"alias\",\"type\":\"string\"},{\"name\":\"city\",\"type\":\"string\"},{\"name\":\"district\",\"type\":\"string\"},{\"name\":\"email\",\"type\":\"string\"},{\"name\":\"full_number\",\"type\":\"string\"},{\"name\":\"name\",\"type\":\"string\"},{\"name\":\"public_place\",\"type\":\"string\"},{\"name\":\"zip_code\",\"type\":\"string\"}]}"
    }'
- `Run the project`
    - Run -> start debugging -> to allow -> perform account creation via Postman -> stop
    curl --location --request POST 'http://localhost:1002/v1/accounts' \
    --header 'Content-Type: application/json' \
    --data-raw '{
        "email": "van3@email.com'\''",
        "full_number": "5511964127228",
        "name": "Van",
        "zip_code": "01001-000"  
    }'
- `View messages sent to Kafka`
    - Access the control center in docker compose -> select the cluster -> select the topic (account_create) -> select the messages -> insert 0 in partition -> Enter

## Stop running
- `Stop docker`
    - docker-compose -f local-dev/docker-compose.yaml --profile infra down