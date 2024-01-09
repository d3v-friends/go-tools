#!/bin/bash

# ex) sh tag.sh 1.0.0

VERSION=$1
COMMIT_ID=$(git rev-parse HEAD)
COMMIT_ID=${COMMIT_ID:0:8}

ORIGIN=$2
if [ -z "$ORIGIN" ]; then
	ORIGIN="origin"
fi

TAG="v$VERSION-$COMMIT_ID"

# delete tag
git tag -d "$TAG"
git push -d "$ORIGIN" "$TAG"

# create tag
git tag "$TAG"
git push --tags "$ORIGIN" "$TAG"
