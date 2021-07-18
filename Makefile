.PHONY: generate_mocks
generate_mocks:
	@mockery --case=snake --outpkg=storagemocks --output=adapters/storagemocks --name=Repository  --dir ./domain/room/


.PHONY: openapi_http
openapi_http:
	oapi-codegen -generate types -o internal/rooms/ports/openapi_types.gen.go -package ports api/openapi/rooms.yml
	oapi-codegen -generate chi-server -o internal/rooms/ports/openapi_api.gen.go -package ports api/openapi/rooms.yml
# oapi-codegen -generate types -o internal/common/client/rooms/openapi_types.gen.go -package rooms api/openapi/rooms.yml
# oapi-codegen -generate client -o internal/common/client/rooms/openapi_client_gen.go -package rooms api/openapi/rooms.yml