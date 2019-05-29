#!/bin/bash

export CODECOV_TOKEN="63f320c2-e299-4de1-a7a3-5383ff1a832b"

parent_path=$( cd "$(dirname "${BASH_SOURCE[0]}")" ; pwd -P )
cd $parent_path
cd ../../..
#add -race flag in the future
DB_HOST=localhost DB_PORT=5432 DB_USER=admin DB_PASS=admin go test -coverprofile=./internal/test/codecov/coverage.txt -covermode=atomic {./internal/dao,./internal/test/...} -race -count=1 -tags='memory db cli web'
wait $!

bash <(curl -s https://codecov.io/bash)
