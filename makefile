up:
	docker compose up -d

down:
	docker compose down --remove-orphans --volumes

api:
	cd ./server && go run main.go

graf:
	docker exec -it grafana /bin/bash
