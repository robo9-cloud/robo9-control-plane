
build-container-local:
	docker build  . -t api-server

build-container-gcp:
	gcloud builds submit --tag gcr.io/$(GOOGLE_PROJECT)/robo9-api-server:0.0.1 --project=$(GOOGLE_PROJECT)

run-container-gcp:
	gcloud run deploy robo9-api-server --image  gcr.io/$(GOOGLE_PROJECT)/robo9-api-server:0.0.1 --platform managed  --region $(GOOGLE_DEFAULT_REGION) --allow-unauthenticated  --project=$(GOOGLE_PROJECT) --port=18080

run-container-local:
	docker-compose up api-server

run-database-local:
	docker-compose up postgres pgadmin

