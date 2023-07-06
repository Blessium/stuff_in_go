GOCMD = go
BIN_NAME_SAMPLE_SERVER = sample_prom
OUTPUT_DIR= out/bin

build:
	mkdir -p ${OUTPUT_DIR}
	go build -o ${OUTPUT_DIR}/${BIN_NAME_SAMPLE_SERVER} cmd/${BIN_NAME_SAMPLE_SERVER}/main.go	

run: build
	./${OUTPUT_DIR}/${BIN_NAME_SAMPLE_SERVER}

clean:
	rm ./${OUTPUT_DIR}/${BIN_NAME_SAMPLE_SERVER}
	rmdir ${OUTPUT_DIR}
	rmdir out
