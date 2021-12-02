all:
	@docker build -f docker/dockerfile -t jorgepoblete/helloworld:latest .
	@docker push jorgepoblete/helloworld:latest