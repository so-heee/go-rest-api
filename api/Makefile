
.PHONY: codegen-all
codegen-all:
	$(MAKE) codegen-server
	$(MAKE) codegen-client
	$(MAKE) codegen-type
	$(MAKE) codegen-spec

.PHONY: codegen-type
codegen-type:
	@go run github.com/deepmap/oapi-codegen/cmd/oapi-codegen@latest -generate "types" openapi.yaml > ./openapi/types.gen.go

.PHONY: codegen-client
codegen-client:
	@go run github.com/deepmap/oapi-codegen/cmd/oapi-codegen@latest -generate "client" openapi.yaml > ./openapi/client.gen.go

.PHONY: codegen-server
codegen-server:
	@go run github.com/deepmap/oapi-codegen/cmd/oapi-codegen@latest -generate "server" openapi.yaml > ./openapi/server.gen.go

.PHONY: codegen-spec
codegen-spec:
	@go run github.com/deepmap/oapi-codegen/cmd/oapi-codegen@latest -generate "spec" openapi.yaml > ./openapi/spec.gen.go
