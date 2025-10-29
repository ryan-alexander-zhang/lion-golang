package main

import (
	"errors"
	"fmt"
	"lion-golang/apperr"
)

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

func main() {
	//TIP <p>Press <shortcut actionId="ShowIntentionActions"/> when your caret is at the underlined text
	// to see how GoLand suggests fixing the warning.</p><p>Alternatively, if available, click the lightbulb to view possible fixes.</p>
	s := "gopher"
	fmt.Printf("Hello and welcome, %s!\n", s)

	for i := 1; i <= 5; i++ {
		//TIP <p>To start your debugging session, right-click your code in the editor and select the Debug option.</p> <p>We have set one <icon src="AllIcons.Debugger.Db_set_breakpoint"/> breakpoint
		// for you, but you can always add more by pressing <shortcut actionId="ToggleLineBreakpoint"/>.</p>
		fmt.Println("i =", 100/i)
	}

	var returnedError error

	newError := apperr.NewError(apperr.UserNotFound, "user not found error example")
	if newError != nil {
		fmt.Println("Error occurred:", newError.Error())
		returnedError = apperr.WrapError(apperr.ResourceNotFound, "wrapping error example", newError)
	}
	if returnedError != nil {
		fmt.Println("Returned Error:", returnedError.Error())
	}
	stardardError := errors.New("a standard error")
	if appErr, ok := apperr.As(returnedError); ok {
		fmt.Println("AppError Code:", appErr.Code.CodeString())
	}

	// errors.as
	if appErr, ok := apperr.As(stardardError); ok {
		fmt.Println("AppError Code:", appErr.Code.CodeString())
	} else {
		fmt.Println("stardardError is not an AppError")
	}

	// errors.is
	if errors.Is(returnedError, newError) {
		fmt.Println("returnedError is an AppError")
	}

	println(apperr.IsType(returnedError, apperr.ErrorTypeUser))
	println(apperr.IsType(returnedError, apperr.ErrorTypeBiz))
}
