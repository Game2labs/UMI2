# Builds an image using Buildx. Usage:
# build <name> <tag> <dockerfile> <context>
function build() {
    echo "Building $1."
    echo "Tag: $2"
    echo "Dockerfile: $3"
    echo "Context: $4"
    docker buildx build \
        --tag "$2" \
        --cache-from "type=local,src=/tmp/.buildx-cache/$1" \
        --cache-to="type=local,dest=/tmp/.buildx-cache-new/$1" \
        --file "$3" \
        --load "$4" \
        &
}

mkdir -p /tmp/.buildx-cache-new
build l2geth "ethereumoptimism/l2geth:latest" "./l2geth/Dockerfile" .
build l1chain "ethereumoptimism/hardhat-node:latest" "./umi2/docker/hardhat/Dockerfile" ./umi2/docker/hardhat

wait

build deployer "ethereumoptimism/deployer:latest" "./umi2/docker/Dockerfile.deployer" .
build dtl "ethereumoptimism/data-transport-layer:latest" "./umi2/docker/Dockerfile.data-transport-layer" .
build relayer "ethereumoptimism/message-relayer:latest" "./umi2/docker/Dockerfile.message-relayer" .
build relayer "ethereumoptimism/fault-detector:latest" "./umi2/docker/Dockerfile.fault-detector" .
build integration-tests "ethereumoptimism/integration-tests:latest" "./umi2/docker/Dockerfile.integration-tests" .

wait
