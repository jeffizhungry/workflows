# --------------------------------------
#  Variables
# --------------------------------------

# --------------------------------------
#  Rules
# --------------------------------------
# Install developer tools and dependencies
# - https://github.com/codeskyblue/fswatch
.PHONY: install-deps
install-deps:
	go get -u -v github.com/codeskyblue/fswatch

# Start cadence cluster
.PHONY: cadence
cadence:
	docker-compose -f cadence/docker-compose.yaml up

# Start applicationc
.PHONY: run
run:
	fswatch --config fsw.yml
