# pull dependencies
pull:
    wget -O "pb/google/api/http.proto" "https://github.com/googleapis/googleapis/raw/master/google/api/http.proto"
    wget -O "pb/google/api/annotations.proto" "https://github.com/googleapis/googleapis/raw/master/google/api/annotations.proto"

# format protos
format:
    buf format -w -d

# run buf linter
lint: format
    buf lint

# generate protos
generate: lint
    buf generate