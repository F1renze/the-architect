

compile_proto: ## compile protobuf defintions in all services
	@tput setaf 3;echo 'tip: plz execute next commands(for fish) when failed:\nset PATH $$HOME/go/bin $$PATH\nset PATH /usr/local/go/bin $$PATH\n'; tput setaf 2; \
	for d in srv; do \
		for f in $$d/**/proto/*.proto; do \
			protoc -I=. --micro_out=. --go_out=. $$f; \
			echo compiled: $$f; \
		done \
	done; \


.DEFAULT_GOAL :=
help: ## show help info
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)



