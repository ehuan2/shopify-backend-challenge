# tag with the latest build, use the appropriate dockerfile and start in right place
backend-build:
	docker build \
		-t ehuan2/shopify-backend-challenge-backend:latest \
		-f ./src/goose-counter/Dockerfile-goose-counter \
	./src/goose-counter/

frontend-build:
	docker build \
		-t ehuan2/shopify-backend-challenge-frontend:latest \
		-f ./src/frontend/Dockerfile-frontend \
	./src/frontend/

build: backend-build frontend-build

run:
	docker-compose up

push:
	docker push ehuan2/shopify-backend-challenge-backend:latest
	docker push ehuan2/shopify-backend-challenge-frontend:latest
