all: tard

.PHONY: tard
tard:
	@./scripts/build.sh tard

package-tard: cleanpackage
	@scripts/package.sh tard Linux   linux   amd64

package: cleanpackage package-tard

package-macos: cleanpackage
	@scripts/package.sh tard Mac   darwin   amd64

cleanpackage:
	@rm -rf packages/

.PHONY: regen
regen: regen-dtag

regen-dtag:
	@echo "dtag protoc regen..."
	@cd mods/dtag && protoc \
		--go_out=. --go_opt=paths=source_relative \
		ingestible.proto
	@protoc-go-inject-tag -input="mods/dtag/*.pb.go"

test: all
	@go test -v -count=1 \
		./main/tard
