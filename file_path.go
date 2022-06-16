package main

import (
	"fmt"
	"path/filepath"
	"strings"
)

func main() {

	// Join should be used to construct paths in a portable way.
	// It takes any number of arguments and constructs a hierarchical path from them.

	p := filepath.Join("dir1", "dir2", "filename")
	fmt.Println("p:", p)

	// You should always use Join instead of concatenating /s or \s manually.
	// In addition to providing portability, Join will also normalize paths by removing superfluous separators and directory changes.
	fmt.Println(filepath.Join("dir1//", "filename"))
	fmt.Println(filepath.Join("dir1/../dir1", "filename"))

	fmt.Println("Dir(p):", filepath.Dir(p))
	fmt.Println("Base(p):", filepath.Base(p))
	// We can check whether a path is absolute.
	fmt.Println(filepath.IsAbs("dir/file"))
	fmt.Println(filepath.IsAbs("/dir/file"))

	// Dir and Base can be used to split a path to the directory and the file. Alternatively, Split will return both in the same call.
	filename := "config.json"

	ext := filepath.Ext(filename)
	fmt.Println(ext)

	fmt.Println(strings.TrimSuffix(filename, ext))

	rel, err := filepath.Rel("a/b", "a/b/t/file")
	if err != nil {
		panic(err)
	}
	fmt.Println(rel)

	rel, err = filepath.Rel("a/b", "a/c/t/file")
	if err != nil {
		panic(err)
	}
	fmt.Println(rel)
}

// Some file names have extensions following a dot. We can split the extension out of such names with Ext.

// Rel finds a relative path between a base and a target. It returns an error if the target cannot be made relative to base.
// output :

// p: dir1/dir2/filename
// dir1/filename
// dir1/filename
// Dir(p): dir1/dir2
// Base(p): filename
// false
// true
// .json
// config
// t/file
// ../c/t/file
