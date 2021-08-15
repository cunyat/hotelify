.PHONY: generate_mocks
generate_mocks:
	@bash scripts/mocks.sh

test:
	@./scripts/test.sh rooms .test.env
