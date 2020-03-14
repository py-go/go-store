.PHONY:
	local
	run

ENTRY_POINT=./cmd/app
TEST_PATH=./cmd/...

build_ci: # Build executable
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o go-store ${ENTRY_POINT}

image: build_ci # Build executable and docker image
	docker build -t go-store .

integration_test_image:
	make -C integration_tests image

local: # Run go application locally
	PORT=3000 PROBE_PORT=8085 go run ${ENTRY_POINT}

loaddata:
	PORT=3000 PROBE_PORT=8085 go run ${ENTRY_POINT} "loaddata"
run: image # Run docker container in foreground
	docker run -p 3000:3000 -p 8085:8085 go-store

coverage:
	go clean -testcache ${TEST_PATH}
	go test ${TEST_PATH} -coverprofile=coverage.html
	go tool cover -html=coverage.html

unit_test:
	go test ${TEST_PATH}
