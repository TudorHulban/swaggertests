.PHONY: swagger
swagger: 
	GO111MODULE=off swagger generate spec -o ./swagger.yaml --scan-models
	
.PHONY: swagger-valid
swagger-valid: ## Invokes swagger for validation of swagger generated documentation.
	@swagger validate ./swagger.yaml
