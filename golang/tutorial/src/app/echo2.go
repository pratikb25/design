// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// Echo2 prints its command-line arguments.
package main

import (
	"fmt"
	"os"
)

func main() {
	s, sep := "", ""
	for _, arg := range os.Args[1:2] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}

/*

In each iteration of the loop, range produces a pair of values:
the index and the value of the element at that index.
In this example, we don’t need the index, but the syntax of a
range loop requires that if we deal with the element, we must
deal with the index too. The solution is to use the blank identifier,
whose name is _ (that is, an underscore)

	for _, arg := range os.Args[1:2] {
		// ...
	}

As noted above, each time around the loop, the string s gets completely
new contents. The += statement makes a new string by concatenating the
old string. If the amount of data involved is large, this could be costly.

	s += sep + arg

A simpler and more efficient solution would be to use the Join function
from the strings package:

	func main() {
    	fmt.Println(strings.Join(os.Args[1:], " "))
	}

 Finally, if we don’t care about format but just want to see the values,
 perhaps for debugging, we can let Println format the results for us:

 	fmt.Println(os.Args[1:])

*/
