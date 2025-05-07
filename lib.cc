#include <iostream>

extern "C" int cc_add(int x, int y) { return x + y; }

extern "C" void cc_putl() { std::cout << "hello\n"; }

extern "C" void cc_call_go(void (*cb)(int)) { cb(40); }

extern "C" void cc_mut(void (*cb)(int)) { cb(50); }
