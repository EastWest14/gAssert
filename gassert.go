package gAssert

import (
	"fmt"
	"log"
)

//*************** Setting Assert Package Behavior ***************

var assertsAreFatal bool = true

//SetAssertFatal determines whether a failed Assert will kill the process.
func SetAssertFatal(isFatal bool) {
	assertsAreFatal = isFatal
}

var actionOnAssert func(string) = func(message string) {
	fmt.Println(message)
}

//NoActionOnAssert determines whether an alert will trigger an alert action.
//Assert action typically logs out the message, removing it may improve performance.
func NoActionOnAssert() {
	actionOnAssert = func(message string) {
		return
	}
}

//SetActionOnAssert sets the action that will happen if the assert fails.
//Assert action typically logs out the message, ignoring it may improve performance.
func SetActionOnAssert(actionFunc func(message string)) {
	actionOnAssert = actionFunc
}

//*************** Assert Statements ***************

//Assert triggers an assert action if the condition is false.
//Assert action typically involves logging out the message.
//After assert action assert kills the process or returns,
//depending on how SetAssertFatal is set.
func Assert(condition bool, message string) {

}

//AssertSoft triggers an assert action if the condition is false and returns.
//Function doesn't kill the process regardless of the value of SetAssertFatal.
func AssertSoft(condition bool, message string) {

}

//AssertSoft triggers an assert action if the condition is false and kills
//the process regardless of the value of SetAssertFatal.
func AssertHard(condition bool, message string) {

}

//*************** Killing Process ***************

//killProcessCommand exists for testing reasons only.
var killProcessCommand func(message string) = func(message string) {
	log.Fatal(message)
}

func killProcess(message string) {
	killProcessCommand(message)
}
