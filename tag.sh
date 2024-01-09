#!/bin/bash

# ex) sh tag.sh 1.0.0

VERSION=$1
COMMIT_ID=$(git rev-parse HEAD)
COMMIT_ID=${COMMIT_ID:0:8}
BRANCH=$(git branch --show-current)

ORIGIN=$1
if [ -z "$ORIGIN" ]; then
	ORIGIN="origin"
fi

TAG="v$VERSION-$COMMIT_ID"
LATEST="v$VERSION-latest"

# delete tag
git tag -d "$TAG"
git tag -d "$LATEST"
git push -d "$ORIGIN" "$TAG"
git push -d "$ORIGIN" "$LATEST"

# create tag
git tag "$TAG"
git tag "$LATEST"
git push --tags "$ORIGIN" "$TAG"
git push --tags "$ORIGIN" "$LATEST"
