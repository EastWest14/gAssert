package gAssert

import (
	"testing"
)

//*************** Mock Structures ***************

type mockProcessKiller struct {
	wasHit bool
}

func (pk *mockProcessKiller) mockKillProcessCommand(message string) {
	pk.wasHit = true
}

//*************** Test Setting Assert Package Behavior ***************

func TestSetAssertFatal(t *testing.T) {
	//Test enable
	assertsAreFatal = false
	SetAssertFatal(true)
	if !assertsAreFatal {
		t.Error("Asserts were suppose to be set as fatal")
	}

	//Test dissable
	assertsAreFatal = true
	SetAssertFatal(false)
	if assertsAreFatal {
		t.Error("Assers were suppose to be set as not fatal")
	}
}

func TestNoActionOnAssert(t *testing.T) {
	wasHit := false
	checkHitFunction := func(message string) {
		wasHit = true
	}
	//Set the package variable
	actionOnAssert = checkHitFunction
	//Suppose to un-set the variable
	NoActionOnAssert()

	//Check if the assert action was un-set
	actionOnAssert("")
	if wasHit {
		t.Error("Assert action executed despite being disabled")
	}
}

func TestSetActionOnAssert(t *testing.T) {
	actionOnAssert = func(message string) {
		return
	}
	wasHit := false
	checkHitFunction := func(message string) {
		wasHit = true
	}
	SetActionOnAssert(checkHitFunction)
	actionOnAssert("")

	//Check correct function was called
	if !wasHit {
		t.Error("Failed to set function on assert")
	}
}

//*************** Test Process Killing ***************

func TestKillProcess(t *testing.T) {
	mockPK := mockProcessKiller{wasHit: false}
	killProcessCommand = mockPK.mockKillProcessCommand
	killProcess("Doesn't matter")
	if !mockPK.wasHit {
		t.Error("Kill Process fails.")
	}
}
