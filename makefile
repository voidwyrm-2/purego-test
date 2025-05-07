both: all
	gcc -shared -o blib.so clib.so cclib.so

all: c cc

c:
	gcc -shared -o clib.so lib.c

cc:
	g++ -shared -o cclib.so lib.cc

clean:
	rm -rf clib.so
	rm -rf cclib.so
