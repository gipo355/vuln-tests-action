docker-build:
	@docker buildx build -t halamert/hello-world-docker-go-action -f Dockerfile . --load 
docker-push:
	@docker buildx build -t halamert/hello-world-docker-go-action -f Dockerfile . --push
