git add .
git commit -m "Ultimo Commit"
git push
go build main.go
rm main.zip
tar -a -cf main.zip main