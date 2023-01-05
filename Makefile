VERSION=v$(shell date -u +"%Y.%m.%d").1
CODEGEN_VERSION=2.4.29
CODEGEN_JAR=swagger-codegen-cli-$(CODEGEN_VERSION).jar


.PHONY: client
client: $(CODEGEN_JAR)
	# Generate client
	@rm -rf ./client
	java -jar $(CODEGEN_JAR) generate \
		-i  https://core.hivelocity.net/api/v2/swagger.json \
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

.PHONY: tag
release:
	git tag $(VERSION)
	git push origin $(VERSION)
