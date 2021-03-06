#!/bin/bash

go test -cover -v -coverprofile=c.out ./internal/business/...

go tool cover -html=c.out -o cover.html

rm -f c.out

echo "------------------------------------------"
echo " Test coverage has been generated."
echo " To see cover report test: chromium cover.html"
echo "------------------------------------------"