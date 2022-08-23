PWD := ${CURDIR}
CONTAINER_LABEL := "wwwperipherals"
CONTAINER_REMOVE_FLAG := "--rm"
CONTAINER_NAME := "klakegg/hugo"
CONTAINER_VERSION := latest
BASE_URL := "https://tamada.github.io/peripherals"

build:
	docker run ${CONTAINER_REMOVE_FLAG} --name ${CONTAINER_LABEL}_build -v "${PWD}":/src -v ${PWD}/public:/output -e HUGO_THEME=cayman-hugo-theme -e HUGO_BASEURL=${BASE_URL} ${CONTAINER_NAME}:${CONTAINER_VERSION}

start:
	docker run ${CONTAINER_REMOVE_FLAG} -d --name ${CONTAINER_LABEL} -p 1313:1313 -v "${PWD}":/src -v "$(PWD)"/public:/output -e HUGO_THEME=cayman-hugo-theme -e HUGO_WATCH="true" -e HUGO_BASEURL=${BASE_URL} ${CONTAINER_NAME}:${CONTAINER_VERSION}

stop:
	docker stop ${CONTAINER_LABEL}
