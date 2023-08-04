git add .
git commit -m "Ultimo Commit"
git push
go build -o main
rm main.zip
tar -a -cf main.zip main