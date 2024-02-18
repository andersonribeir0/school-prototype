IMAGE_NAME=school-prototype-app
GO_VERSION=1.22
CONTAINER_NAME=school-prototype
PORT=8080
JWT_SIGNING_KEY=$(shell cat ./requests/jwt_secret_key.example)

build:
	docker build --build-arg GO_VERSION=$(GO_VERSION) -t $(IMAGE_NAME) .

run: build
	docker run -e "JWT_SIGNING_KEY=$(JWT_SIGNING_KEY)" --name $(CONTAINER_NAME) -p $(PORT):$(PORT) --rm -d $(IMAGE_NAME)

stop:
	docker stop $(CONTAINER_NAME)

test:
	docker run --rm $(IMAGE_NAME) go test ./...

clean:
	docker rmi $(IMAGE_NAME)

generate:
	@templ generate ./...
