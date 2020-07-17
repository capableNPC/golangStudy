package main

import "fmt"

/*
*fmt:
func Print(a ...interface{}) (n int, err error)
func Printf(format string, a ...interface{}) (n int, err error)
func Println(a ...interface{}) (n int, err error)
	func Fprint(w io.Writer, a ...interface{}) (n int, err error)
	func Fprintf(w io.Writer, format string, a ...interface{}) (n int, err error)
	func Fprintln(w io.Writer, a ...interface{}) (n int, err error)
	func Sprint(a ...interface{}) string
func Sprintf(format string, a ...interface{}) string
func Sprintln(a ...interface{}) string
func Errorf(format string, a ...interface{}) error
func Scan(a ...interface{}) (n int, err error)
func Scanf(format string, a ...interface{}) (n int, err error)
func Scanln(a ...interface{}) (n int, err error)
func Fscan(r io.Reader, a ...interface{}) (n int, err error)
func Fscanln(r io.Reader, a ...interface{}) (n int, err error)
func Fscanf(r io.Reader, format string, a ...interface{}) (n int, err error)

*/

func main() {
	var n = 100
	fmt.Printf("%T,%v\n", n, n)
	fmt.Printf("%b,%o,%d,%x\n", n, n, n, n)
	str := "asdaaa"
	fmt.Printf("%s\n", str)
	fmt.Printf("%v\n", str)
	fmt.Printf("%#v\n", str)
}
