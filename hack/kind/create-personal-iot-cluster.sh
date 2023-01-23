#!/usr/bin/env bash
kind delete cluster --name personal-iot
kind create cluster --config $(dirname -- "${0}")/personal-iot-cluster.yaml

