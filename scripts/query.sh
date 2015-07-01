#!/bin/bash

curl http://127.0.0.1:8500/v1/catalog/service/$1 | python -m json.tool
