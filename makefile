NAME := $(shell basename "$(PWD)")

docker: $(NAME)
	docker build --build-arg NAME=$(NAME) -f dockerfile -t byuoitav/pi-credentials-microservice:latest .
	docker login -u $(DOCKER_USERNAME) -p $(DOCKER_PASSWORD) 
	docker push byuoitav/pi-credentials-microservice:latest

$(NAME):
	make build

build:
	env GOOS=linux CGO_ENABLED=0 go build -o $(NAME) -v

clean:
	go clean
	rm -f $(NAME)
