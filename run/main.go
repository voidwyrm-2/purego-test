package main

// copied and changed from https://github.com/ebitengine/purego/blob/main/examples/libc/main.go

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"unsafe"

	"github.com/ebitengine/purego"
	"github.com/voidwyrm-2/purego-test/run/openlib"
)

var dynFuncs = map[string]func(register func(fptr any)){
	"c_add": func(register func(fptr any)) {
		var fn func(int, int) int
		register(&fn)
		fmt.Println(fn(20, 30))
	},
	"c_fac": func(register func(fptr any)) {
		var fn func(int) int
		register(&fn)
		fmt.Println(fn(30))
	},
	"c_call_go": func(register func(fptr any)) {
		var fn func(func(int))
		register(&fn)
		fn(func(i int) {
			fmt.Println(i)
		})
	},
	"c_putl": func(register func(fptr any)) {
		var fn func(*byte)
		register(&fn)

		msg := []byte("it works")

		ptr := unsafe.Pointer(&msg)

		bptr := (*byte)(ptr)

		fn(bptr)
	},
	"cc_add": func(register func(fptr any)) {
		var fn func(int, int) int
		register(&fn)
		fmt.Println(fn(20, 30))
	},
	"cc_putl": func(register func(fptr any)) {
		var fn func()
		register(&fn)
		// fn("hello!")
		fn()
	},
	"cc_call_go": func(register func(fptr any)) {
		var fn func(func(int))
		register(&fn)
		fn(func(i int) {
			fmt.Println(i)
		})
	},
	"cc_mut": func(register func(fptr any)) {
		var fn func(func(int))
		register(&fn)
		l := []int{}
		fmt.Println(l)
		fn(func(i int) {
			l = append(l, i)
		})
		fmt.Println(l)
	},
	"zig_add": func(register func(fptr any)) {
		var fn func(int, int) int
		register(&fn)
		fmt.Println(fn(30, 30))
	},
	"zig_putl": func(register func(fptr any)) {
		var fn func(string)
		register(&fn)
		fn("ziggy!")
	},
}

func main() {
	list := flag.Bool("l", false, "List the available functions")
	fname := flag.String("c", "", "The function to call")
	file := flag.String("f", "clib", "The library to load")

	flag.Parse()

	if *list {
		for k := range dynFuncs {
			fmt.Println(k)
		}

		return
	}

	ext := ".so"
	if strings.HasPrefix(*file, "z") {
		ext = ".a"
	}

	lib, err := openlib.OpenLibrary("../" + *file + ext)
	if err != nil {
		fmt.Println("opening error:", err.Error())
		os.Exit(1)
	}

	if cb, ok := dynFuncs[*fname]; !ok {
		fmt.Printf("no function named '%s'\n", *fname)
		os.Exit(1)
	} else {
		cb(func(fptr any) {
			purego.RegisterLibFunc(fptr, lib, *fname)
		})
	}
}
