.PHONY: build
build:
	docker build  -t app-irrigator -f containers/Dockerfile .

.PHONY: run
run: build
	docker run app-irrigator

## Testing

.PHONY: build_unit
build_unit:
	docker build  -t app-irrigator-unit-test -f containers/Dockerfile.test_unit .

.PHONY: test_unit
test_unit: build_unit
	docker run -t app-irrigator-unit-test 