build:
	mkdir ./build
	go build -o ./build/scaffold
	sudo mv ./build/scaffold /usr/local/bin/scaffold
	rm -rf ./build

update: build
	cp -R ./scaffold $(HOME)/.config/

clean:
	rm -rf ./build
