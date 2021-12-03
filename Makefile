all:
	@docker build -f docker/dockerfile -t jorgepoblete/helloworld:1.0.2 .
	@docker push jorgepoblete/helloworld:1.0.2