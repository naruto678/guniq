build: 
	go build -o guniq . 
args_test: build
	./guniq test.txt 
test: build
	./guniq countries.txt | wc -l 
