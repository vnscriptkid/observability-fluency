up:
	docker compose up -d

down:
	docker compose down --remove-orphans --volumes

api:
	cd ./server && go run main.go

graf:
	docker exec -it grafana /bin/bash

grafui:
	open http://localhost:3000

alertui:
	open http://localhost:9093

promui:
	open http://localhost:9090