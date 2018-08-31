#!/bin/bash

TAG_ID=$1
GIT_TOKEN=$2

echo "Tagging ..."
git tag -a $TAG -m "$TAG release"

echo "Releasing $TAG_ID ..."
JSON='{"tag_name": "'"$TAG"'","target_commitish": "master","name": "'"$TAG"'","body": "'"$TAG"'-release","draft": false,"prerelease": false}'
curl -H "$AUTH" \
    -H "Content-Type: application/json" \
    -d "$JSON" \
    $GH_REPO/releases