.PHONY: generate_mocks
generate_mocks:
	@mockery --case=snake --outpkg=storagemocks --output=internal/rooms/adapters/storagemocks --name=Repository  --dir ./internal/rooms/domain/room/
	@mockery --case=snake --outpkg=commandmocks --output=internal/common/adapters/commandmocks --name=CommandBus --dir ./internal/common/domain/
	@mockery --case=snake --outpkg=querymocks --output=internal/common/adapters/querymocks --name=QueryBus --dir ./internal/common/domain/


