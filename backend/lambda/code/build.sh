rm -r /lambda/build/*

for dir in */; do
  if [ $dir = 'types/' ]
  then
    echo SKIPPING $dir
    continue
  fi

  echo COMPILING $dir
  GOOS=linux GOARCH=amd64 go build -o /lambda/build/"$dir"/main "$dir"cmd/main.go
  cd /lambda/build/"$dir"
  zip -j main.zip main
  cd /lambda/code
done