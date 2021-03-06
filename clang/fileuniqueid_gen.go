package clang

// #include "./clang-c/Index.h"
// #include "go-clang.h"
import "C"

// FileUniqueID uniquely identifies a CXFile, that refers to the same underlying file, across an indexing session.
type FileUniqueID struct {
	c C.CXFileUniqueID
}
