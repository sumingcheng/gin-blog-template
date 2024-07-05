IMAGE_NAME := gin-blog:v1.0.1
CONTAINER_NAME := gin-blog
INTERNAL_PORT := 5678
EXTERNAL_PORT ?= 5678

build:
	@docker build --no-cache -t $(IMAGE_NAME) .

run:
	@docker run -d --name $(CONTAINER_NAME) -p $(EXTERNAL_PORT):$(INTERNAL_PORT) $(IMAGE_NAME)

stop:
	@docker stop $(CONTAINER_NAME)
	@docker rm $(CONTAINER_NAME)

clean:
	@docker rmi $(IMAGE_NAME)
