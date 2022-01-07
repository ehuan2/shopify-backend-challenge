build: backend-build frontend-build

# tag with the latest build, use the appropriate dockerfile and start in right place
# we follow a similar structure to building both - we specify the file we want, then we use that folder as the entrypoint to copy
backend-build:
	docker build \
		-t ehuan2/shopify-backend-challenge-backend:latest \
		-f ./src/backend/Dockerfile-golang \
	./src/backend/

frontend-build:
	docker build \
		-t ehuan2/shopify-backend-challenge-frontend:latest \
		-f ./src/frontend/Dockerfile-frontend \
	./src/frontend/

run:
	docker-compose up

push:
	docker push ehuan2/shopify-backend-challenge-backend:latest
	docker push ehuan2/shopify-backend-challenge-frontend:latest
