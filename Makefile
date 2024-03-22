.PHONY: build
build:
	docker build -t deconfliction_provider .

.PHONY: run
run:
	docker run -d -p 5050:5050 --name deconfliction_provider deconfliction_provider
