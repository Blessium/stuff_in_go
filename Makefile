GOCMD = go
BIN_NAME_SAMPLE_SERVER = sample_prom

OUTPUT_DIR = out/bin
CONFIG_DIR = config

BINARIES = ${BIN_NAME_SAMPLE_SERVER}

.PHONY: docker-build build docker-run run docker-clean clean

docker-build:
	docker compose -f ${CONFIG_DIR}/docker-compose.yml build

docker-run: docker-build
	docker compose -f ${CONFIG_DIR}/docker-compose.yml up -d

docker-clean:
	docker compose -f ${CONFIG_DIR}/docker-compose.yml down --volumes --remove-orphans



build:
	mkdir -p ${OUTPUT_DIR}
	@for bin in $(BINARIES); do \
		echo "Building $$bin"; \
		go build -o ${OUTPUT_DIR}/$$bin cmd/$$bin/main.go; \
	done

run: build
	./${OUTPUT_DIR}/${BIN_NAME_SAMPLE_SERVER}

clean:
	rm ./${OUTPUT_DIR}/${BIN_NAME_SAMPLE_SERVER}
	rmdir ${OUTPUT_DIR}
	rmdir out
