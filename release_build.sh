#!/bin/bash

VERSION=`cat version.txt`

rm gtodo
rm -rf releases
mkdir -p releases
echo ""

echo "Building for Linux AMD64" &&
sleep 1 &&
env GOOS=linux GOARCH=amd64 go build . &&
zip -r releases/gtodo-linux-amd64-$VERSION.zip gtodo &&
echo "Built for Linux AMD64"
echo ""
echo "Building for Linux ARM" &&
sleep 1 &&
env GOOS=linux GOARCH=arm go build . &&
zip -r releases/gtodo-linux-arm-$VERSION.zip gtodo &&
echo "Built for Linux ARM"
echo ""
echo "Building for Mac AMD64" &&
sleep 1 &&
env GOOS=darwin GOARCH=amd64 go build . &&
zip -r releases/gtodo-darwin-amd64-$VERSION.zip gtodo &&
echo "Built for Mac AMD64"
echo ""

rm gtodo

echo "Building for Windows 386 and 686" &&
sleep 1 &&
env GOOS=windows GOARCH=386 CGO_ENABLED=0 go build . &&
zip -r releases/gtodo-windows-386_686-$VERSION.zip gtodo.exe &&
rm gtodo.exe &&
echo "Built for Windows 386 and 686"
echo ""

echo "All packages release built"
