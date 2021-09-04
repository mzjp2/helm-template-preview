run-local:
	docker-compose -f docker-compose.local.yml up --build --detach

build-frontend-prod:
	rm -rf frontend/build && cd frontend && npm run build && cd .. && rm -rf nginx/build && cp -r frontend/build nginx/build
