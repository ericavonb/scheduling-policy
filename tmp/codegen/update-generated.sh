#!/usr/bin/env bash

set -o errexit
set -o nounset
set -o pipefail

vendor/k8s.io/code-generator/generate-groups.sh \
deepcopy \
github.com/ericavonb/scheduling-policy/pkg/generated \
github.com/ericavonb/scheduling-policy/pkg/apis \
policy:v1alpha1 \
--go-header-file "./tmp/codegen/boilerplate.go.txt"
