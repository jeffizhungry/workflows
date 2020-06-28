# --------------------------------------
#  Variables
# --------------------------------------

# --------------------------------------
#  Rules
# --------------------------------------
.PHONY: cadence
cadence:
	docker-compose -f cadence/docker-compose.yaml up