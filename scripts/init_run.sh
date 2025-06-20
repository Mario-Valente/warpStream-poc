#!/usr/bin/env bash

set -euo pipefail

kind create cluster --name warpstream-test --config ./infra/kind/kind-config.yaml

kubectl config use-context kind-warpstream-test


helm repo add warpstream https://warpstreamlabs.github.io/charts
helm repo update
helm upgrade --install warpstream-agent warpstream/warpstream-agent --values ./infra/agents/charts/values.yaml \
    --namespace default \
    --set config.bucketURL="" \
    --set config.apiKey="" \
    --set config.region="us-east-1" \
    --set config.virtualClusterID=""