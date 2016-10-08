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
	//Re-setting the variable after the test
	startingAssertsAreFatal := assertsAreFatal
	defer func() {
		assertsAreFatal = startingAssertsAreFatal
	}()

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
	//Re-setting the variable after the test
	startingActionOnAssert := actionOnAssert
	defer func() {
		actionOnAssert = startingActionOnAssert
	}()

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
	//Re-setting the variable after the test
	startingActionOnAssert := actionOnAssert
	defer func() {
		actionOnAssert = startingActionOnAssert
	}()

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

//*************** Test Assert Statements ***************

func TestAssertSoft(t *testing.T) {
	//Re-setting the variable after the test
	startingActionOnAssert := actionOnAssert
	defer func() {
		actionOnAssert = startingActionOnAssert
	}()

	//Test assert on true
	wasHit := false
	checkHitFunction := func(message string) {
		wasHit = true
	}
	//Set the package variable
	actionOnAssert = checkHitFunction
	AssertSoft(true, "")
	if wasHit {
		t.Error("AssertSoft fails on a true condition")
	}

	//Test assert on false
	wasHit = false
	AssertSoft(false, "")
	if !wasHit {
		t.Error("AssertSoft passes on a false condition")
	}
}

func TestAssert(t *testing.T) {
	//Re-setting the variable after the test
	startingKillProcessCommand := killProcessCommand
	defer func() {
		killProcessCommand = startingKillProcessCommand
	}()
	//Re-setting the variable after the test
	startingActionOnAssert := actionOnAssert
	defer func() {
		actionOnAssert = startingActionOnAssert
	}()
	//Re-setting the variable after the test
	startingAssertsAreFatal := assertsAreFatal
	defer func() {
		assertsAreFatal = startingAssertsAreFatal
	}()

	//Mock setup
	mockPK := mockProcessKiller{}
	killProcessCommand = mockPK.mockKillProcessCommand
	var assertActionCalled bool
	mockAssertAction := func(message string) {
		assertActionCalled = true
	}
	//Set the package variable
	actionOnAssert = mockAssertAction

	//Assert should act like SoftAssert or HardAssert depending on the value of assertsAreFatal
	cases := []struct {
		assertConditionShouldPass bool
		assertsAreFatal           bool
		expectActionOnAlertCall   bool
		exceptKillProcessCall     bool
	}{
		{true, true, false, false},
		{true, false, false, false},
		{false, true, true, true},
		{false, false, true, false},
	}
	for i, aCase := range cases {
		//Pre-set before every case
		assertActionCalled = false
		mockPK.wasHit = false

		assertsAreFatal = aCase.assertsAreFatal
		Assert(aCase.assertConditionShouldPass, "")
		//Check results of Assert
		if assertActionCalled != aCase.expectActionOnAlertCall {
			t.Errorf("Assert action called/not called incorrectly in case %d", i)
		}
		if mockPK.wasHit != aCase.exceptKillProcessCall {
			t.Errorf("Kill process called/not called incorrectly in case %d", i)
		}
	}
}

func TestAssertHard(t *testing.T) {
	//Re-setting the variable after the test
	startingKillProcessCommand := killProcessCommand
	defer func() {
		killProcessCommand = startingKillProcessCommand
	}()

	//Test assert on true
	mockPK := mockProcessKiller{wasHit: false}
	killProcessCommand = mockPK.mockKillProcessCommand
	AssertHard(true, "")
	if mockPK.wasHit {
		t.Error("AssertHard fails on a true condition")
	}

	//Test assert on false
	mockPK = mockProcessKiller{wasHit: false}
	killProcessCommand = mockPK.mockKillProcessCommand
	AssertHard(false, "")
	if !mockPK.wasHit {
		t.Error("AssertHard passes on a false condition")
	}
}

//*************** Test Process Killing ***************

func TestKillProcess(t *testing.T) {
	//Re-setting the variable after the test
	startingKillProcessCommand := killProcessCommand
	defer func() {
		killProcessCommand = startingKillProcessCommand
	}()

	mockPK := mockProcessKiller{wasHit: false}
	killProcessCommand = mockPK.mockKillProcessCommand
	killProcess("Doesn't matter")
	if !mockPK.wasHit {
		t.Error("Kill Process fails.")
	}
}
