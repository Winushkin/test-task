start:
	mkdir ./logs ./reports

pg-up:
	docker compose --env-file .env up -d --build postgres

up:
	docker compose --env-file .env up -d --build


.PHONY: backup-from-container
backup-from-container:
	mkdir -p ./backup/logs
	mkdir -p ./backup/reports

	docker cp app:/app/logs ./backup/

	docker cp app:/app/reports ./backup/

	@echo "Backup completed successfully."