dev:
	docker-compose -f docker-compose.yml -f docker-compose.development.yml up --build

prod:
	docker-compose up --build

logs-frontend:
	docker logs --follow notes-frontend

logs-api:
	docker logs --follow notes-api

exec-api:
	docker exec -it notes-api bash

exec-frontend:
	docker exec -it notes-frontend bash