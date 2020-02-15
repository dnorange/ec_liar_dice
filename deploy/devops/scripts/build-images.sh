#!/bin/bash

# Back to the root of the project
cd $(dirname $0)
cd ../../..

echo "Building images..."
docker build -t dice -t dnorange/dice:latest .
echo "Built successfully"