IMAGE_NAME ?= telegram-gateway
IMAGE_TAG ?= loc
HOST=telegram-gateway.loc
# Keep in sync with .env
PORT ?= 8080

all: run

build:
	docker build -t $(IMAGE_NAME):$(IMAGE_TAG) .

run: build
	docker run --rm \
		--env-file .env \
		-e VIRTUAL_HOST=$(HOST) \
		-e VIRTUAL_PORT=$(PORT) \
		-v $(PWD)/config.yaml:/etc/config.yaml \
	$(IMAGE_NAME):$(IMAGE_TAG)

send-message:
	if [ -z "$(CHAT_ID)" ]; then echo "CHAT_ID is not set"; exit 1; fi
	if [ -z "$(MESSAGE)" ]; then echo "MESSAGE is not set"; exit 1; fi
	curl -X POST http://$(HOST)/sendmessage --data '{"chat_id":$(CHAT_ID),"text":"$(MESSAGE)"}'
