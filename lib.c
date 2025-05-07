#include <stdio.h>
#include <stdlib.h>

int c_add(int x, int y) { return x + y; }

int c_fac(int n) {
  int out = 1;

  for (; n != 0; n--) {
    out *= n;
  }

  return out;
}

void c_call_go(void (*cb)(int)) { cb(20); }

void c_putl(char *msg) { printf("%s", msg); }
