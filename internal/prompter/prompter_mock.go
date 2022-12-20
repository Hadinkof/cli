// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package prompter

import (
	"sync"
)

// Ensure, that PrompterMock does implement Prompter.
// If this is not the case, regenerate this file with moq.
var _ Prompter = &PrompterMock{}

// PrompterMock is a mock implementation of Prompter.
//
//	func TestSomethingThatUsesPrompter(t *testing.T) {
//
//		// make and configure a mocked Prompter
//		mockedPrompter := &PrompterMock{
//			AuthTokenFunc: func() (string, error) {
//				panic("mock out the AuthToken method")
//			},
//			ConfirmFunc: func(s string, b bool) (bool, error) {
//				panic("mock out the Confirm method")
//			},
//			ConfirmDeletionFunc: func(s string) error {
//				panic("mock out the ConfirmDeletion method")
//			},
//			InputFunc: func(s1 string, s2 string) (string, error) {
//				panic("mock out the Input method")
//			},
//			InputHostnameFunc: func() (string, error) {
//				panic("mock out the InputHostname method")
//			},
//			MarkdownEditorFunc: func(s1 string, s2 string, b bool) (string, error) {
//				panic("mock out the MarkdownEditor method")
//			},
//			MultiSelectFunc: func(s1 string, s2 string, strings []string) (int, error) {
//				panic("mock out the MultiSelect method")
//			},
//			PasswordFunc: func(s string) (string, error) {
//				panic("mock out the Password method")
//			},
//			SelectFunc: func(s1 string, s2 string, strings []string) (int, error) {
//				panic("mock out the Select method")
//			},
//		}
//
//		// use mockedPrompter in code that requires Prompter
//		// and then make assertions.
//
//	}
type PrompterMock struct {
	// AuthTokenFunc mocks the AuthToken method.
	AuthTokenFunc func() (string, error)

	// ConfirmFunc mocks the Confirm method.
	ConfirmFunc func(s string, b bool) (bool, error)

	// ConfirmDeletionFunc mocks the ConfirmDeletion method.
	ConfirmDeletionFunc func(s string) error

	// InputFunc mocks the Input method.
	InputFunc func(s1 string, s2 string) (string, error)

	// InputHostnameFunc mocks the InputHostname method.
	InputHostnameFunc func() (string, error)

	// MarkdownEditorFunc mocks the MarkdownEditor method.
	MarkdownEditorFunc func(s1 string, s2 string, b bool) (string, error)

	// MultiSelectFunc mocks the MultiSelect method.
	MultiSelectFunc func(s1 string, s2 string, strings []string) (int, error)

	// PasswordFunc mocks the Password method.
	PasswordFunc func(s string) (string, error)

	// SelectFunc mocks the Select method.
	SelectFunc func(s1 string, s2 string, strings []string) (int, error)

	// calls tracks calls to the methods.
	calls struct {
		// AuthToken holds details about calls to the AuthToken method.
		AuthToken []struct {
		}
		// Confirm holds details about calls to the Confirm method.
		Confirm []struct {
			// S is the s argument value.
			S string
			// B is the b argument value.
			B bool
		}
		// ConfirmDeletion holds details about calls to the ConfirmDeletion method.
		ConfirmDeletion []struct {
			// S is the s argument value.
			S string
		}
		// Input holds details about calls to the Input method.
		Input []struct {
			// S1 is the s1 argument value.
			S1 string
			// S2 is the s2 argument value.
			S2 string
		}
		// InputHostname holds details about calls to the InputHostname method.
		InputHostname []struct {
		}
		// MarkdownEditor holds details about calls to the MarkdownEditor method.
		MarkdownEditor []struct {
			// S1 is the s1 argument value.
			S1 string
			// S2 is the s2 argument value.
			S2 string
			// B is the b argument value.
			B bool
		}
		// MultiSelect holds details about calls to the MultiSelect method.
		MultiSelect []struct {
			// S1 is the s1 argument value.
			S1 string
			// S2 is the s2 argument value.
			S2 string
			// Strings is the strings argument value.
			Strings []string
		}
		// Password holds details about calls to the Password method.
		Password []struct {
			// S is the s argument value.
			S string
		}
		// Select holds details about calls to the Select method.
		Select []struct {
			// S1 is the s1 argument value.
			S1 string
			// S2 is the s2 argument value.
			S2 string
			// Strings is the strings argument value.
			Strings []string
		}
	}
	lockAuthToken       sync.RWMutex
	lockConfirm         sync.RWMutex
	lockConfirmDeletion sync.RWMutex
	lockInput           sync.RWMutex
	lockInputHostname   sync.RWMutex
	lockMarkdownEditor  sync.RWMutex
	lockMultiSelect     sync.RWMutex
	lockPassword        sync.RWMutex
	lockSelect          sync.RWMutex
}

// AuthToken calls AuthTokenFunc.
func (mock *PrompterMock) AuthToken() (string, error) {
	if mock.AuthTokenFunc == nil {
		panic("PrompterMock.AuthTokenFunc: method is nil but Prompter.AuthToken was just called")
	}
	callInfo := struct {
	}{}
	mock.lockAuthToken.Lock()
	mock.calls.AuthToken = append(mock.calls.AuthToken, callInfo)
	mock.lockAuthToken.Unlock()
	return mock.AuthTokenFunc()
}

// AuthTokenCalls gets all the calls that were made to AuthToken.
// Check the length with:
//
//	len(mockedPrompter.AuthTokenCalls())
func (mock *PrompterMock) AuthTokenCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockAuthToken.RLock()
	calls = mock.calls.AuthToken
	mock.lockAuthToken.RUnlock()
	return calls
}

// Confirm calls ConfirmFunc.
func (mock *PrompterMock) Confirm(s string, b bool) (bool, error) {
	if mock.ConfirmFunc == nil {
		panic("PrompterMock.ConfirmFunc: method is nil but Prompter.Confirm was just called")
	}
	callInfo := struct {
		S string
		B bool
	}{
		S: s,
		B: b,
	}
	mock.lockConfirm.Lock()
	mock.calls.Confirm = append(mock.calls.Confirm, callInfo)
	mock.lockConfirm.Unlock()
	return mock.ConfirmFunc(s, b)
}

// ConfirmCalls gets all the calls that were made to Confirm.
// Check the length with:
//
//	len(mockedPrompter.ConfirmCalls())
func (mock *PrompterMock) ConfirmCalls() []struct {
	S string
	B bool
} {
	var calls []struct {
		S string
		B bool
	}
	mock.lockConfirm.RLock()
	calls = mock.calls.Confirm
	mock.lockConfirm.RUnlock()
	return calls
}

// ConfirmDeletion calls ConfirmDeletionFunc.
func (mock *PrompterMock) ConfirmDeletion(s string) error {
	if mock.ConfirmDeletionFunc == nil {
		panic("PrompterMock.ConfirmDeletionFunc: method is nil but Prompter.ConfirmDeletion was just called")
	}
	callInfo := struct {
		S string
	}{
		S: s,
	}
	mock.lockConfirmDeletion.Lock()
	mock.calls.ConfirmDeletion = append(mock.calls.ConfirmDeletion, callInfo)
	mock.lockConfirmDeletion.Unlock()
	return mock.ConfirmDeletionFunc(s)
}

// ConfirmDeletionCalls gets all the calls that were made to ConfirmDeletion.
// Check the length with:
//
//	len(mockedPrompter.ConfirmDeletionCalls())
func (mock *PrompterMock) ConfirmDeletionCalls() []struct {
	S string
} {
	var calls []struct {
		S string
	}
	mock.lockConfirmDeletion.RLock()
	calls = mock.calls.ConfirmDeletion
	mock.lockConfirmDeletion.RUnlock()
	return calls
}

// Input calls InputFunc.
func (mock *PrompterMock) Input(s1 string, s2 string) (string, error) {
	if mock.InputFunc == nil {
		panic("PrompterMock.InputFunc: method is nil but Prompter.Input was just called")
	}
	callInfo := struct {
		S1 string
		S2 string
	}{
		S1: s1,
		S2: s2,
	}
	mock.lockInput.Lock()
	mock.calls.Input = append(mock.calls.Input, callInfo)
	mock.lockInput.Unlock()
	return mock.InputFunc(s1, s2)
}

// InputCalls gets all the calls that were made to Input.
// Check the length with:
//
//	len(mockedPrompter.InputCalls())
func (mock *PrompterMock) InputCalls() []struct {
	S1 string
	S2 string
} {
	var calls []struct {
		S1 string
		S2 string
	}
	mock.lockInput.RLock()
	calls = mock.calls.Input
	mock.lockInput.RUnlock()
	return calls
}

// InputHostname calls InputHostnameFunc.
func (mock *PrompterMock) InputHostname() (string, error) {
	if mock.InputHostnameFunc == nil {
		panic("PrompterMock.InputHostnameFunc: method is nil but Prompter.InputHostname was just called")
	}
	callInfo := struct {
	}{}
	mock.lockInputHostname.Lock()
	mock.calls.InputHostname = append(mock.calls.InputHostname, callInfo)
	mock.lockInputHostname.Unlock()
	return mock.InputHostnameFunc()
}

// InputHostnameCalls gets all the calls that were made to InputHostname.
// Check the length with:
//
//	len(mockedPrompter.InputHostnameCalls())
func (mock *PrompterMock) InputHostnameCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockInputHostname.RLock()
	calls = mock.calls.InputHostname
	mock.lockInputHostname.RUnlock()
	return calls
}

// MarkdownEditor calls MarkdownEditorFunc.
func (mock *PrompterMock) MarkdownEditor(s1 string, s2 string, b bool) (string, error) {
	if mock.MarkdownEditorFunc == nil {
		panic("PrompterMock.MarkdownEditorFunc: method is nil but Prompter.MarkdownEditor was just called")
	}
	callInfo := struct {
		S1 string
		S2 string
		B  bool
	}{
		S1: s1,
		S2: s2,
		B:  b,
	}
	mock.lockMarkdownEditor.Lock()
	mock.calls.MarkdownEditor = append(mock.calls.MarkdownEditor, callInfo)
	mock.lockMarkdownEditor.Unlock()
	return mock.MarkdownEditorFunc(s1, s2, b)
}

// MarkdownEditorCalls gets all the calls that were made to MarkdownEditor.
// Check the length with:
//
//	len(mockedPrompter.MarkdownEditorCalls())
func (mock *PrompterMock) MarkdownEditorCalls() []struct {
	S1 string
	S2 string
	B  bool
} {
	var calls []struct {
		S1 string
		S2 string
		B  bool
	}
	mock.lockMarkdownEditor.RLock()
	calls = mock.calls.MarkdownEditor
	mock.lockMarkdownEditor.RUnlock()
	return calls
}

// MultiSelect calls MultiSelectFunc.
func (mock *PrompterMock) MultiSelect(s1 string, s2 string, strings []string) (int, error) {
	if mock.MultiSelectFunc == nil {
		panic("PrompterMock.MultiSelectFunc: method is nil but Prompter.MultiSelect was just called")
	}
	callInfo := struct {
		S1      string
		S2      string
		Strings []string
	}{
		S1:      s1,
		S2:      s2,
		Strings: strings,
	}
	mock.lockMultiSelect.Lock()
	mock.calls.MultiSelect = append(mock.calls.MultiSelect, callInfo)
	mock.lockMultiSelect.Unlock()
	return mock.MultiSelectFunc(s1, s2, strings)
}

// MultiSelectCalls gets all the calls that were made to MultiSelect.
// Check the length with:
//
//	len(mockedPrompter.MultiSelectCalls())
func (mock *PrompterMock) MultiSelectCalls() []struct {
	S1      string
	S2      string
	Strings []string
} {
	var calls []struct {
		S1      string
		S2      string
		Strings []string
	}
	mock.lockMultiSelect.RLock()
	calls = mock.calls.MultiSelect
	mock.lockMultiSelect.RUnlock()
	return calls
}

// Password calls PasswordFunc.
func (mock *PrompterMock) Password(s string) (string, error) {
	if mock.PasswordFunc == nil {
		panic("PrompterMock.PasswordFunc: method is nil but Prompter.Password was just called")
	}
	callInfo := struct {
		S string
	}{
		S: s,
	}
	mock.lockPassword.Lock()
	mock.calls.Password = append(mock.calls.Password, callInfo)
	mock.lockPassword.Unlock()
	return mock.PasswordFunc(s)
}

// PasswordCalls gets all the calls that were made to Password.
// Check the length with:
//
//	len(mockedPrompter.PasswordCalls())
func (mock *PrompterMock) PasswordCalls() []struct {
	S string
} {
	var calls []struct {
		S string
	}
	mock.lockPassword.RLock()
	calls = mock.calls.Password
	mock.lockPassword.RUnlock()
	return calls
}

// Select calls SelectFunc.
func (mock *PrompterMock) Select(s1 string, s2 string, strings []string) (int, error) {
	if mock.SelectFunc == nil {
		panic("PrompterMock.SelectFunc: method is nil but Prompter.Select was just called")
	}
	callInfo := struct {
		S1      string
		S2      string
		Strings []string
	}{
		S1:      s1,
		S2:      s2,
		Strings: strings,
	}
	mock.lockSelect.Lock()
	mock.calls.Select = append(mock.calls.Select, callInfo)
	mock.lockSelect.Unlock()
	return mock.SelectFunc(s1, s2, strings)
}

// SelectCalls gets all the calls that were made to Select.
// Check the length with:
//
//	len(mockedPrompter.SelectCalls())
func (mock *PrompterMock) SelectCalls() []struct {
	S1      string
	S2      string
	Strings []string
} {
	var calls []struct {
		S1      string
		S2      string
		Strings []string
	}
	mock.lockSelect.RLock()
	calls = mock.calls.Select
	mock.lockSelect.RUnlock()
	return calls
}
