# --------------------------------------
#  Variables
# --------------------------------------

# --------------------------------------
#  Rules - Cadence Cluster
# --------------------------------------

# Start cadence cluster (simple)
.PHONY: cadence
cadence: clean-cadence register start-cadence

# Start cadence cluster
.PHONY: start-cadence
start-cadence:
	docker-compose -f cadence/docker-compose.yaml up

# Clean cadence cluster
# NOTE: Run this if you see "Incompatible versionsversion mismatch for keyspace/database"
# This is due a db migration issue
.PHONY: clean-cadence
clean-cadence:
	docker-compose -f cadence/docker-compose.yaml down

# Register domains
.PHONY: register
register:
	docker run --network=host --rm ubercadence/cli:master --do simple-domain domain register --rd 1

# --------------------------------------
#  Rules - App
# --------------------------------------

# Install developer tools and dependencies
.PHONY: install-deps
install-deps:
	go get -u -v github.com/codeskyblue/fswatch

# Start applicationc
.PHONY: run
run:
	fswatch --config fsw.yml
