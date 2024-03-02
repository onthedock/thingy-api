#!/bin/bash

curl -X PUT -d '{"Id": "01HQZ880A3N6Z6AZCJTF7JYPQH", "Name": "newThingy"}' \
  -s http://localhost:8080/api/v1/thingy > /dev/null
curl -s http://localhost:8080/api/v1/thingy/01HQZ880A3N6Z6AZCJTF7JYPQH

echo