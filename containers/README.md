# Container Build

When you are building these containers make sure you are running inside of linux or using wsl on windows to run linux natively

This will not work on windows. You will get the following error:

```bash 

BUILD FAILURE: Build step failure: build step 2 "gcr.io/$PROJECT/terraform:1.0.9" failed: starting step container failed: Error response from daemon: OCI runtime create failed: container_linux.go:380: starting container process caused: exec: "/builder/entrypoint.bash": permission denied: unknown
ERROR: (gcloud.builds.submit) build b00c6748-238d-4939-8da4-492bf5befc10 completed with status "FAILURE"

```