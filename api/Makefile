

.PHONY: codegen-type
codegen-type:
	oapi-codegen -generate "types" openapi.yaml > ./openapi/types.gen.go

.PHONY: codegen-client
codegen-client:
	oapi-codegen -generate "client" openapi.yaml > ./openapi/client.gen.go

.PHONY: codegen-server
codegen-server:
	oapi-codegen -generate "server" openapi.yaml > ./openapi/server.gen.go

.PHONY: codegen-spec
codegen-spec:
	oapi-codegen -generate "spec" openapi.yaml > ./openapi/spec.gen.go
