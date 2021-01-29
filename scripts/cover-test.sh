#!/bin/bash

go test -cover -coverprofile=c.out ./internal/business/...

go tool cover -html=c.out -o cover.html

rm -f c.out

echo "------------------------------------------"
echo " Test covarega has been generated."
echo " To see cover report test: chromium cover.html"
echo "------------------------------------------"