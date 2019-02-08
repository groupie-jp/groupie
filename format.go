package main

import (
	"fmt"
	"golang.org/x/text/width"
)

//var (
//	// Fold is a transform that maps all runes to their canonical width.
//	//
//	// Note that the NFKC and NFKD transforms in golang.org/x/text/unicode/norm
//	// provide a more generic folding mechanism.
//	Fold Transformer = Transformer{foldTransform{}}
//
//	// Widen is a transform that maps runes to their wide variant, if
//	// available.
//	Widen Transformer = Transformer{wideTransform{}}
//
//	// Narrow is a transform that maps runes to their narrow variant, if
//	// available.
//	Narrow Transformer = Transformer{narrowTransform{}}
//)

func main() {
	fmt.Println(width.Narrow.String("ゆにばーす"))
}
