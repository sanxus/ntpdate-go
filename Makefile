build:
	go build -buildmode=pie -trimpath -ldflags="-s -w" .

clean:
	rm ntpdate
