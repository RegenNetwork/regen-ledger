#!/usr/bin/env bash

set -eo pipefail

mkdir -p ./tmp-swagger-gen
proto_dirs=$(find ./proto -path -prune -o -name '*.proto' -print0 | xargs -0 -n1 dirname | sort | uniq)
for dir in $proto_dirs; do

  # generate swagger files (filter query files)
  query_file=$(find "${dir}" -maxdepth 2 -name 'query.proto')
  if [[ ! -z "$query_file" ]]; then
    buf protoc  \
      -I "proto" \
      -I "third_party/proto" \
      "$query_file" \
      --swagger_out=./tmp-swagger-gen \
      --swagger_opt=logtostderr=true --swagger_opt=fqn_for_swagger_name=true --swagger_opt=simple_operation_ids=true
  fi
done


rm -f ./tmp-swagger-gen/swagger.yaml

# download Cosmos SDK swagger doc
SDK_VERSION=$(go list -m -f '{{ .Version }}' github.com/cosmos/cosmos-sdk)
echo "SDK version ${SDK_VERSION}"
wget "https://raw.githubusercontent.com/cosmos/cosmos-sdk/${SDK_VERSION}/client/docs/swagger-ui/swagger.yaml" -P ./tmp-swagger-gen
mv ./tmp-swagger-gen/swagger.yaml ./tmp-swagger-gen/swagger-sdk.yaml

# # download IBC swagger doc
IBC_VERSION=$(go list -m -f '{{ .Version }}' github.com/cosmos/ibc-go/v2)
echo "IBC version ${IBC_VERSION}"
wget "https://raw.githubusercontent.com/cosmos/ibc-go/${IBC_VERSION}/docs/client/swagger-ui/swagger.yaml" -P ./tmp-swagger-gen
mv ./tmp-swagger-gen/swagger.yaml ./tmp-swagger-gen/swagger-ibc.yaml

# combine swagger files
# uses nodejs package `swagger-combine`.
# all the individual swagger files need to be configured in `config.json` for merging
swagger-combine ./client/docs/config.json -o ./tmp-swagger-gen/swagger.yaml -f yaml --continueOnConflictingPaths true --includeDefinitions true

# move generated swagger file to swagger-ui directory
cp ./tmp-swagger-gen/swagger.yaml ./client/docs/swagger-ui/

# clean swagger files
rm -rf ./tmp-swagger-gen
