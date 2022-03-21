#!/bin/bash
HELM_LINT="/usr/local/bin/helm"

if [ -d  manifests ]; then
    rm -rf /tmp/manifests
    cp -r manifests /tmp
    if [ "$(uname)" == "Linux" ]; then
        find /tmp/manifests -type f | grep 'Chart.yaml' | xargs sed -i 's/\[\[\ TagOfRepo\ \]\]/v0.0.0/g'
    elif [ "$(uname)" == "Darwin" ]; then
        find /tmp/manifests -type f | grep 'Chart.yaml' | xargs sed -i '' 's/\[\[\ TagOfRepo\ \]\]/v0.0.0/g'
    fi
    find /tmp/manifests -type f | sed -n 's/Chart.yaml//p' | xargs $HELM_LINT lint
    rm -rf /tmp/manifests
fi
