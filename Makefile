.PHONY: generate_mocks
generate_mocks:
	@mockery --case=snake --outpkg=storagemocks --output=adapters/storagemocks --name=Repository  --dir ./domain/room/


