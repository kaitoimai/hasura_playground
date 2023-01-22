LOG_LEVEL := INFO
HASURA_CLI_VERSION := v2.16.1

.PHONY: build
build:
	@docker compose build --no-cache

.PHONY: upf
upf:
	@docker compose up ${SERVICES} 

.PHONY: console
console:
	@hasura --project hasura console --admin-secret secret --skip-update-check

.PHONY: apply
apply:
	@hasura --project hasura deploy --admin-secret secret --skip-update-check

.PHONY: gqlgen
gqlgen:
	@docker compose exec rs gqlgen generate
