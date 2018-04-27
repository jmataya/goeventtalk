#!/bin/bash

echo "Creating cart..."

curl -X POST \
     http://localhost:21337/cart \
     -H 'content-type: application/json' \
     -d '{
           "customerId": 3
         }' | jq .
