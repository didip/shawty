BASEDIR  = $(shell dirname $(realpath $(lastword $(MAKEFILE_LIST))))

APP      = $(shell basename ${PWD})
PID      = $(BASEDIR)/${APP}.pid
BIN      = $(BASEDIR)/${APP}.bin

GO_FILES = $(wildcard $(BASEDIR)/*.go)

serve: restart
	@fswatch -o -Ee "\\.(git/|pid$$|bin$$)" . | xargs -n1 -I{}  make restart || make kill

kill:
	@kill $(shell cat $(PID)) || true

before:
	@echo "$(shell date -u +"%Y-%m-%dT%H:%M:%SZ"): restarting"

$(BIN): $(GO_FILES)
	@go build -o $@ $(GO_FILES)

restart: kill before $(BIN)
	@${BIN} & echo $$! > $(PID)

.PHONY: serve restart kill before