build:
	@docker build -f docker/dockerfile -t jorgepoblete/helloworld:latest .

push: build
	@docker push jorgepoblete/helloworld:latest