title: /etc/shadow passwords with Go
date: 2013-04-29 08:50:00
url: etc-shadow-passwords-with-go

Here's how you can create hashed + salted passwords for use in /etc/shadow. I'm using libcrypt.

	package main

	/*
	#cgo LDFLAGS: -lcrypt
	#include 
	*/
	import "C"

	import (
		"fmt"
	)

	func main() {
		fmt.Println("Hashed:", C.GoString(C.crypt(C.CString("password!!!"), C.CString("$6$Vi.DuMQS"))));
	}

The output should be...

	Hashed: $6$Vi.DuMQS$3hoKGTZ4ym8W3VHhLith2rGnChBtEobC3h07MVfdzk/0GxnWlkAUZ7/msJ1t93ekA8qc7jzVfP./8fnkfk/e6/

