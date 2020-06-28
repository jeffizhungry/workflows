# --------------------------------------
#  Variables
# --------------------------------------

# --------------------------------------
#  Rules
# --------------------------------------
.PHONY: run
run:
	go run ./app/main.go

.PHONY: cadence
cadence:
	docker-compose -f cadence/docker-compose.yaml up