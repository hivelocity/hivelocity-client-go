CODEGEN_VERSION=2.4.39
CODEGEN_JAR=swagger-codegen-cli-$(CODEGEN_VERSION).jar

swagger.yaml:
	# We want to keep swagger.yaml in version control, so that we can see what changed.
	curl -sSL https://core.hivelocity.net/api/v2/swagger.json | yq -P > swagger.yaml

.PHONY: client
client: $(CODEGEN_JAR) swagger.yaml
	# Generate client
	@rm -rf ./client
	java -jar $(CODEGEN_JAR) generate \
		-i swagger.yaml \
		-l go \
		-o ./client
	go fmt ./...
	go build github.com/hivelocity/hivelocity-client-go/client
	go test ./...

$(CODEGEN_JAR):
	wget https://repo1.maven.org/maven2/io/swagger/swagger-codegen-cli/$(CODEGEN_VERSION)/$(CODEGEN_JAR)

.PHONY: clean
clean:
	@rm -rf client/
	@rm -f swagger-codegen-cli*jar

generate-modules:
	./hack/golang-modules-update.sh
