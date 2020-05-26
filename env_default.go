package main

// DefaultEnv is supposed to contain all default
// .fwd env file injected upon build time
const DefaultEnv string = `# Docker images to be used for each service
FWD_IMAGE_APP="fireworkweb/php:7.4"
FWD_IMAGE_HTTP="fireworkweb/http:latest"
FWD_IMAGE_NODE="fireworkweb/node:12"
FWD_IMAGE_CHROMEDRIVER="fireworkweb/chromedriver:latest"
FWD_IMAGE_CACHE="redis:alpine"
FWD_IMAGE_DATABASE="mysql:5.7"
FWD_IMAGE_NODE_QA="fireworkweb/node:12-qa"
FWD_IMAGE_PHP_QA="jakzal/phpqa:1.34-alpine"

# Only print out the generated commands instead of running them
FWD_DEBUG=false
# Prints out all commands and their output
FWD_VERBOSE=false
# Microseconds (defaults = 1 second)
FWD_ATTEMPTS_DELAY=1000000

# Optional, default is folder name
# FWD_NAME=custom_name

# Ports - used on docker-compose.yml
FWD_HTTP_PORT="80:80"
FWD_DATABASE_PORT="3306:3306"

# Host user for mapping execution within the containers
FWD_ASUSER=$UID

# Docker Bin
FWD_DOCKER_BIN=docker
# Docker Compose Bin
FWD_DOCKER_COMPOSE_BIN=docker-compose
# Docker Compose Exec flags
FWD_COMPOSE_EXEC_FLAGS=""
# Docker Compose Run flags
FWD_COMPOSE_RUN_FLAGS=""
# Docker Run flags
FWD_DOCKER_RUN_FLAGS="-it --init"

# SSH key for mapping within container so compose/npm
# can have access to secured sources
FWD_SSH_PATH="$HOME/.ssh"
FWD_SSH_KEY_PATH="$HOME/.ssh/id_rsa" # deprecated, use the above
FWD_CONTEXT_PATH="$PWD"
FWD_CUSTOM_PATH="$PWD/fwd"

# Start configuration
FWD_START_CHECK=true
FWD_START_DEFAULT_SERVICES="database cache app http"

# Database
DB_DATABASE="docker"
DB_USERNAME="docker"
DB_PASSWORD="secret"

# Docker Composer
COMPOSE_API_VERSION=1.25
FWD_COMPOSE_VERSION=3.7
# FWD_NETWORK="custom-network" # by default will be generated on the fly
`
