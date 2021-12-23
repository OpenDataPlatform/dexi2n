#! /bin/bash

VERSION=v0.1.0

UNCOMMITTED="no"

if [ "x$GITHUB_TOKEN" == "x" ]; then
  echo "Please export GITHUB_TOKEN=...."
  exit 1
fi

MYDIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"

TEST=$(git status --porcelain|wc -l)
if [ 0 -ne $TEST -a $UNCOMMITTED != "yes" ]; then
   echo "Please, commit before releasing"
   exit 1
fi

echo "Let's go"
git tag $VERSION

#goreleaser --rm-dist release --skip-validate
goreleaser release --rm-dist --skip-validate



