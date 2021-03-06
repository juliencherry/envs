// Code generated by counterfeiter. DO NOT EDIT.
package commandfakes

import (
	"sync"

	"github.com/juliencherry/envs/command"
)

type FakeCFCLIWrapper struct {
	LoginStub        func(api string, username string, password string, skipSSLValidation bool) error
	loginMutex       sync.RWMutex
	loginArgsForCall []struct {
		api               string
		username          string
		password          string
		skipSSLValidation bool
	}
	loginReturns struct {
		result1 error
	}
	loginReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeCFCLIWrapper) Login(api string, username string, password string, skipSSLValidation bool) error {
	fake.loginMutex.Lock()
	ret, specificReturn := fake.loginReturnsOnCall[len(fake.loginArgsForCall)]
	fake.loginArgsForCall = append(fake.loginArgsForCall, struct {
		api               string
		username          string
		password          string
		skipSSLValidation bool
	}{api, username, password, skipSSLValidation})
	fake.recordInvocation("Login", []interface{}{api, username, password, skipSSLValidation})
	fake.loginMutex.Unlock()
	if fake.LoginStub != nil {
		return fake.LoginStub(api, username, password, skipSSLValidation)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.loginReturns.result1
}

func (fake *FakeCFCLIWrapper) LoginCallCount() int {
	fake.loginMutex.RLock()
	defer fake.loginMutex.RUnlock()
	return len(fake.loginArgsForCall)
}

func (fake *FakeCFCLIWrapper) LoginArgsForCall(i int) (string, string, string, bool) {
	fake.loginMutex.RLock()
	defer fake.loginMutex.RUnlock()
	return fake.loginArgsForCall[i].api, fake.loginArgsForCall[i].username, fake.loginArgsForCall[i].password, fake.loginArgsForCall[i].skipSSLValidation
}

func (fake *FakeCFCLIWrapper) LoginReturns(result1 error) {
	fake.LoginStub = nil
	fake.loginReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeCFCLIWrapper) LoginReturnsOnCall(i int, result1 error) {
	fake.LoginStub = nil
	if fake.loginReturnsOnCall == nil {
		fake.loginReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.loginReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeCFCLIWrapper) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.loginMutex.RLock()
	defer fake.loginMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeCFCLIWrapper) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ command.CFCLIWrapper = new(FakeCFCLIWrapper)
