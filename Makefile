# REQUIRED SECTION
include ./.golang.mk
# END OF REQUIRED SECTION

# Run 'make help' for the list of default targets

# Example of overriding of default target
#test: ## run test with coverage using the vendor directory
#	go test -mod vendor -v -cover ./... -coverprofile cover.out

# Threshold increased from default
coverage: test
	goverreport -coverprofile=cover.out -sort=block -order=desc -threshold=95

cdktf: 
	npm install -g cdktf-cli

STACK:=cdk
LD_FLAGS := -s -w -X $(GOMODULE)/cmd.version=$(LD_VERSION) -X $(GOMODULE)/cmd.commit=$(LD_COMMIT) -X $(GOMODULE)/cmd.date=$(LD_DATE) -X $(GOMODULE)/cmd.stack=$(STACK)

synthesize: install
synthesize: ## Synthesize Terraform resources to cdktf.out/
	cdktf synth $(STACK) 

diff: ## Perform a diff (terraform plan) for the given stack
	cdktf diff $(STACK)   

deploy: ## Deploy the given stack
	cdktf deploy $(STACK) 

destroy: ## Destroy the given stack
	cdktf destroy $(STACK)

get: ## Generates provider-specific bindings 
	TERRAFORM_BINARY_NAME=`asdf which terraform` cdktf get


lint: 
	@$(GOBIN)/golangci-lint run cmd