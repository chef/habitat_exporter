#!/bin/bash
VERSION=$(grep Version version.go | awk -F'"' '{print $2}')
echo "=> Version: $VERSION"

PROGNAME="${PWD##*/}"

echo "=> Tagging current release"
git tag "$VERSION"
echo "=> Pushing tag"
git push origin "$VERSION"

goreleaser --rm-dist
