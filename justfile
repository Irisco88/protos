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
    # rm -rf ./gen
    buf generate ./





gopath := `go env | grep GOPATH | cut -d "=" -f2 | tr -d '"'`
# builds permission plugin
plugin:
    @echo "build plugin..."
    @go build -ldflags "-s -w" -o "./bin/protoc-gen-permission" ./plugins/permission && echo "permission generated successfully."
    @echo "copy compiled permission plugin to {{gopath}}/bin"
    @cp "./bin/protoc-gen-permission" "{{gopath}}/bin"