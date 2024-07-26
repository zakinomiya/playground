gcc_options = -std=c++20 -Wall --pedantic-errors 

program: main.cpp all.h all.h.gch
	g++ $(gcc_options) $< -o $@

all.h.gch: all.h # preconmpile header 
	g++ $(gcc_options) -x c++-header $< -o $@

clean:
	rm program *.gch 
	
run: program
	./program

.PHONY: clean run
