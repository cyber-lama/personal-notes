dev:
	docker-compose -f docker-compose.yml -f docker-compose.development.yml up --build

prod:
	docker-compose up --build

logs-frontend:
	docker logs --follow cocash-frontend

logs-api:
	docker logs --follow cocash-api

exec-api:
	docker exec -it cocash-api /bin/sh

exec-frontend:
	docker exec -it cocash-frontend /bin/sh