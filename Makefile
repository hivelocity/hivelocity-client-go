VERSION=v$(shell date -u +"%Y.%m.%d").1

.PHONY: client
client:
	# Download [latest] OpenAPI spec
	rm -rf openapi.json
	wget -q -O openapi.json https://core.hivelocity.net/api/v2/swagger.json
	# Generate client
	# Checkout https://github.com/OpenAPITools/openapi-generator/releases
	@chmod +x ./openapi-generator
	@rm -rf ./client
	OPENAPI_GENERATOR_VERSION=4.3.1 ./openapi-generator generate --package-name client -i openapi.json -g go -o ./client
	rm -f client/go.mod client/go.sum
	go fmt ./...
	go build github.com/hivelocity/hivelocity-client-go/client
	go test ./...

.PHONY: clean
clean:
	@rm -rf client/
	@rm -f openapi-generator-cli-*.jar

.PHONY: tag
release:
	git tag $(VERSION)
	git push origin $(VERSION)