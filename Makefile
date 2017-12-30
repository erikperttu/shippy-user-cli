build:
	docker build -t shippy-user-cli .
image:
	docker build -t shippy-user-cli .
run:
	docker run --net="host" \
	-e MICRO_REGISTRY=mdns \
	shippy-user-cli