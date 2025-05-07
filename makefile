both: all
	gcc -shared -o blib.so clib.so cclib.so
	mv clib.so cclib.so run/

all: c cc

c:
	gcc -shared -o clib.so lib.c

cc:
	g++ -shared -o cclib.so lib.cc

clean:
	rm -rf *.so
	rm -rf *.so
	rm -rf run/*.so
