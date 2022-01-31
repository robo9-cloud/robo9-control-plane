set GOOGLE_PROJECT_ID=zak-poc

gcloud builds submit --tag gcr.io/zak-poc/robo9-api-server:0.0.1 --project=zak-poc

gcloud run deploy robo9-api-server --image  gcr.io/zak-poc/robo9-api-server:0.0.1 --platform managed  --region us-west1 --allow-unauthenticated  --project=zak-poc
