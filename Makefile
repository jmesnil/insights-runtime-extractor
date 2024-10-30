.DEFAULT_GOAL = build-image
.PHONY: runtime-samples

build-rust-dev-image:
	./scripts/build-rust-dev.sh

rust-dev:
	./scripts/rust-dev.sh

build-image:
	./scripts/build.sh

build-hash-image:
	./scripts/build-hash.sh

e2e-test: build-image
	cd exporter && make e2e-test

runtime-samples:
	cd runtime-samples && make

rust-unit-tests:
	cd extractor && make unit-tests

unit-tests:
	cd fingerprints && make unit-tests
	cd exporter && make unit-tests
