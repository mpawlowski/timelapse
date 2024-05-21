#!/bin/bash

git fetch --tags
latestTag=$(git tag | sort -V | tail -n 1)

# Split the version into major, minor, patch
major=$(echo $latestTag | awk -F. '{print $1}')
minor=$(echo $latestTag | awk -F. '{print $2}')
patch=0 # Always reset patch to 0 on minor tag

# Check if major is empty
if [ -z "$major" ]; then
    major="v0"
fi

# Check if minor is empty
if [ -z "$minor" ]; then
    minor=0
fi

# Increment the minor version
minor=$((minor + 1))

# Assemble the new version
newTag="$major.$minor.$patch"

# Create a new tag
git tag $newTag
git push origin $newTag
