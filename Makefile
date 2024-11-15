.DEFAULT_GOAL := buildall

IMAGE-PREFIX ?= ssingh3339

export DOCKER_CLI_EXPERIMENTAL=enabled

.PHONY: build-rsocket-ping
build-rsocket-ping:
	@docker buildx create --use --name=crossplat --node=crossplat && \
	docker buildx build \
		--platform linux/amd64,linux/arm64 \
		--output "type=docker,push=false" \
		--tag $(IMAGE-PREFIX)/rsocket-ping:latest \
		rsocket-ping/.

.PHONY: publish-rsocket-ping
publish-rsocket-ping:
	@docker buildx create --use --name=crossplat --node=crossplat && \
	docker buildx build \
		--platform linux/amd64,linux/arm64 \
		--output "type=image,push=true" \
		--tag $(IMAGE-PREFIX)/rsocket-ping:latest \
		rsocket-ping/.

.PHONY: build-udp-server
build-udp-server:
	@docker buildx create --use --name=crossplat --node=crossplat && \
	docker buildx build \
		--platform linux/amd64,linux/arm64 \
		--output "type=docker,push=false" \
		--tag $(IMAGE-PREFIX)/udp-server:latest \
		udp-server/.

.PHONY: publish-udp-server
publish-udp-server:
	@docker buildx create --use --name=crossplat --node=crossplat && \
	docker buildx build \
		--platform linux/amd64,linux/arm64 \
		--output "type=image,push=true" \
		--tag $(IMAGE-PREFIX)/udp-server:latest \
		udp-server/.

.PHONY: build-echo-graphql
build-echo-graphql:
	@docker buildx create --use --name=crossplat --node=crossplat && \
	docker buildx build \
		--platform linux/amd64,linux/arm64 \
		--output "type=docker,push=false" \
		--tag $(IMAGE-PREFIX)/echo-graphql:latest \
		echo-graphql/.

.PHONY: publish-echo-graphql
publish-echo-graphql:
	@docker buildx create --use --name=crossplat --node=crossplat && \
	docker buildx build \
		--platform linux/amd64,linux/arm64 \
		--output "type=image,push=true" \
		--tag $(IMAGE-PREFIX)/echo-graphql:latest \
		echo-graphql/.

.PHONY: build-graphql-stream
build-graphql-stream:
	@docker buildx create --use --name=crossplat --node=crossplat && \
	docker buildx build \
		--platform linux/amd64,linux/arm64 \
		--output "type=docker,push=false" \
		--tag $(IMAGE-PREFIX)/graphql-stream:latest \
		graphql-stream/.

.PHONY: publish-graphql-stream
publish-graphql-stream:
	@docker buildx create --use --name=crossplat --node=crossplat && \
	docker buildx build \
		--platform linux/amd64,linux/arm64 \
		--output "type=image,push=true" \
		--tag $(IMAGE-PREFIX)/graphql-stream:latest \
		graphql-stream/.

.PHONY: build-http-ping
build-http-ping:
	@docker buildx create --use --name=crossplat --node=crossplat && \
	docker buildx build \
		--platform linux/amd64,linux/arm64 \
		--output "type=docker,push=false" \
		--tag $(IMAGE-PREFIX)/http-ping:latest \
		http-ping/.

.PHONY: publish-http-ping
publish-http-ping:
	@docker buildx create --use --name=crossplat --node=crossplat && \
	docker buildx build \
		--platform linux/amd64,linux/arm64 \
		--output "type=image,push=true" \
		--tag $(IMAGE-PREFIX)/http-ping:latest \
		http-ping/.

.PHONY: build-grpc-ping
build-grpc-ping:
	@docker buildx create --use --name=crossplat --node=crossplat && \
	docker buildx build \
		--platform linux/amd64,linux/arm64 \
		--output "type=docker,push=false" \
		--tag $(IMAGE-PREFIX)/grpc-ping:latest \
		grpc-ping/.

.PHONY: publish-grpc-ping
publish-grpc-ping:
	@docker buildx create --use --name=crossplat --node=crossplat && \
	docker buildx build \
		--platform linux/amd64,linux/arm64 \
		--output "type=image,push=true" \
		--tag $(IMAGE-PREFIX)/grpc-ping:latest \
		grpc-ping/.

.PHONY: build-ws-echo
build-ws-echo:
	@docker buildx create --use --name=crossplat --node=crossplat && \
	docker buildx build \
		--platform linux/amd64,linux/arm64 \
		--output "type=docker,push=false" \
		--tag $(IMAGE-PREFIX)/ws-echo:latest \
		ws-echo/.

.PHONY: publish-ws-echo
publish-ws-echo:
	@docker buildx create --use --name=crossplat --node=crossplat && \
	docker buildx build \
		--platform linux/amd64,linux/arm64 \
		--output "type=image,push=true" \
		--tag $(IMAGE-PREFIX)/ws-echo:latest \
		ws-echo/.

.PHONY: build-graphql-rest
build-graphql-rest:
	@docker buildx create --use --name=crossplat --node=crossplat && \
	docker buildx build \
		--platform linux/amd64,linux/arm64 \
		--output "type=docker,push=false" \
		--tag $(IMAGE-PREFIX)/graphql-rest:latest \
		graphql-rest/.

.PHONY: publish-graphql-rest
publish-graphql-rest:
	@docker buildx create --use --name=crossplat --node=crossplat && \
	docker buildx build \
		--platform linux/amd64,linux/arm64 \
		--output "type=image,push=true" \
		--tag $(IMAGE-PREFIX)/graphql-rest:latest \
		graphql-rest/.

.PHONY: buildall
buildall: build-rsocket-ping build-udp-server build-echo-graphql build-graphql-stream build-http-ping build-grpc-ping build-ws-echo build-graphql-rest

.PHONY: publishall
publishall: publish-rsocket-ping publish-udp-server publish-echo-graphql publish-graphql-stream publish-http-ping publish-grpc-ping publish-ws-echo publish-graphql-rest
