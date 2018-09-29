#!/bin/sh

REGISTRY=asia.gcr.io/planet-pluto-dev
IMAGE=odpt

# build and push latest
gcloud builds submit --config=_cloudbuild.yaml .

# delete untag
digest=`gcloud container images list-tags $REGISTRY/$IMAGE --filter='-tags:*' --format='get(digest)'`
if [ "$digest" != "" ]; then
  echo digest: $digest
  gcloud container images delete --quiet $REGISTRY/$IMAGE@$digest
fi
