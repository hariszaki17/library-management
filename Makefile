PROTO_SRC_DIR=proto
PROTO_GEN_DIR=$(PROTO_SRC_DIR)/gen

# Find all .proto files in the proto directory
PROTO_FILES=$(wildcard $(PROTO_SRC_DIR)/*.proto)

# Protoc command
PROTOC=protoc

# Protoc flags
PROTOC_FLAGS=\
	--go_opt=paths=source_relative \
	--go-grpc_opt=paths=source_relative

# Default target: generate .pb.go files
.PHONY: all
all: gen

# Generate .pb.go files from all .proto files
.PHONY: gen
gen:
	@mkdir -p $(PROTO_GEN_DIR)
	@for file in $(PROTO_FILES); do \
		base=$$(basename $$file .proto); \
		dest_dir=$(PROTO_GEN_DIR)/$$base; \
		echo "Processing file: $$file"; \
		echo "Base directory: $$base"; \
		echo "DEST directory: $$dest_dir"; \
		mkdir -p $(PROTO_GEN_DIR)/$$base; \
		echo "Base directory: $$base"; \
		$(PROTOC) $(PROTOC_FLAGS) \
			--go_out=$(PROTO_GEN_DIR)/$$base \
			--go-grpc_out=$(PROTO_GEN_DIR)/$$base \
			$$file; \
	done

# Clean generated files
.PHONY: clean
clean:
	rm -rf $(PROTO_GEN_DIR)
