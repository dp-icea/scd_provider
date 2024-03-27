.PHONY: build
build:
	docker build -t deconfliction_provider .

.PHONY: run
run:
	docker rm -f deconfliction_provider
	docker run -d -p 5050:5050 --name deconfliction_provider --restart always deconfliction_provider
