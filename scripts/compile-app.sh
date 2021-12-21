
echo "Compiling app in 3"
sleep 1
echo "2"
sleep 1
echo "1"
sleep 1
mv .env ../
GOOS=linux GOARCH=amd64 go build -o bin/application
git add bin/application
git commit -m "Compile app"
pushd ../
echo ".env has been moved to `pwd`"
popd
