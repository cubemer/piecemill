build:
	bun run build
	cp -r dist server/dist
	cd server && GOOS=linux GOARCH=amd64 go build -o ../piecemill-linux-amd64 .
	rm -rf server/dist

clean:
	rm -f piecemill-linux-amd64
	rm -rf dist server/dist
