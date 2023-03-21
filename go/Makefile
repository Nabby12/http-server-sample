PROJECT_PREFIX=go-http-server
IMAGE_NAME=${PROJECT_PREFIX}-image
CONTAINER_NAME=${PROJECT_PREFIX}-container

run:
	docker compose down
	docker compose up -d

build:
	docker build --target runner -t ${IMAGE_NAME} -f docker/Dockerfile .
	docker run --name ${CONTAINER_NAME} -d \
		-p 8080:8080 \
		${IMAGE_NAME}

test:
	cd app && go test -v ./...

remove:
	docker stop ${CONTAINER_NAME}
	docker rm ${CONTAINER_NAME}
	-@docker rmi ${IMAGE_NAME}
	-@docker network rm ${PROJECT_PREFIX}_default
