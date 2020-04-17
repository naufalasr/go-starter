# first is default task when running "make" without args
build: sql generate format gobuild lint

build-pgserve: format gobuild-pgserve lint

build-apiserver: format gobuild-apiserver lint

generate: sqlboiler pgserve-swagger

format:
	go fmt

gobuild: 
	go build -o bin/app

gobuild-pgserve:
	go build -o bin/integresql ./pgserve

gobuild-apiserver:
	go build -o bin/apiserver ./cmd/api

lint:
	golangci-lint run --fast

sqlboiler:
	sqlboiler --wipe --no-hooks psql

# https://github.com/golang/go/issues/24573
# w/o cache - see "go help testflag"
# use https://github.com/kyoh86/richgo to color
# note that these tests should not run verbose by default (e.g. use your IDE for this)
# TODO: add test shuffling/seeding when landed in go v1.15 (https://github.com/golang/go/issues/28592)
test:
	richgo test -cover -race -count=1 ./...

init: modules tools tidy
	@go version

# cache go modules (locally into .pkg)
modules:
	go mod download

# https://marcofranssen.nl/manage-go-tools-via-go-modules/
tools:
	cat tools.go | grep _ | awk -F'"' '{print $$2}' | xargs -tI % go install %

tidy:
	go mod tidy

clean:
	rm -rf bin

sql: sql-format sql-lint sql-check-migrations

sql-format:
	@echo "make sql-format"
	@find ${PWD} -name ".*" -prune -o -type f -iname "*.sql" -print \
		| xargs -i pg_format {} -o {}

# check syntax via the real database
# https://stackoverflow.com/questions/8271606/postgresql-syntax-check-without-running-the-query
sql-lint:
	@echo "make sql-lint"
	@find ${PWD} -name ".*" -prune -o -type f -iname "*.sql" -print \
		| xargs -i sed '1s#^#DO $$SYNTAX_CHECK$$ BEGIN RETURN;#; $$aEND; $$SYNTAX_CHECK$$;' {} \
		| psql --quiet -v ON_ERROR_STOP=1

sql-check-migrations:
	@echo "make sql-check-migrations"
	@(grep -R " NULL" ./migrations/ | grep --invert "DEFAULT NULL" | grep --invert "NOT") && (echo "Unnecessary use of NULL keyword" && exit 1) || exit 0

reset:
	@echo "DROP & CREATE database:"
	@echo "  PGHOST=${PGHOST} PGDATABASE=${PGDATABASE}" PGUSER=${PGUSER}
	@echo -n "Are you sure? [y/N] " && read ans && [ $${ans:-N} = y ]
	psql -d postgres -c 'DROP DATABASE IF EXISTS "${PGDATABASE}";'
	psql -d postgres -c 'CREATE DATABASE "${PGDATABASE}" WITH OWNER ${PGUSER} TEMPLATE "template0"'

# https://www.gnu.org/software/make/manual/html_node/Phony-Targets.html
# ignore matching file/make rule combinations in working-dir
.PHONY: test

pgserve-swagger-spec: 
	@echo "make pgserve-swagger-spec"
	@swagger generate spec \
		-i pgservetypes/swagger/swagger.yml \
		--include=allaboutapps.at/aw/go-mranftl-sample/pgservetypes \
		-o pgservetypes/swagger/swagger.json \
		--scan-models \
		-q

pgserve-swagger-models:
	@echo "make pgserve-swagger-models"
	@swagger generate model \
		--allow-template-override \
		--template-dir=pgservetypes/swagger \
		--spec=pgservetypes/swagger/swagger.json \
		--existing-models=allaboutapps.at/aw/go-mranftl-sample/pgservetypes \
		--model-package=pgservetypes \
		--all-definitions \
		-q

pgserve-swagger-validate:
	@echo "make pgserve-swagger-validate"
	@swagger validate pgservetypes/swagger/swagger.json \
		--stop-on-error \
		-q

pgserve-swagger: pgserve-swagger-spec pgserve-swagger-validate pgserve-swagger-models
