// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"sync"

	"github.com/weaveworks/profiles/api/v1alpha1"
	"github.com/weaveworks/profiles/pkg/api"
)

type FakeCatalog struct {
	GetStub        func(string, string) *v1alpha1.ProfileDescription
	getMutex       sync.RWMutex
	getArgsForCall []struct {
		arg1 string
		arg2 string
	}
	getReturns struct {
		result1 *v1alpha1.ProfileDescription
	}
	getReturnsOnCall map[int]struct {
		result1 *v1alpha1.ProfileDescription
	}
	GetWithVersionStub        func(string, string, string) *v1alpha1.ProfileDescription
	getWithVersionMutex       sync.RWMutex
	getWithVersionArgsForCall []struct {
		arg1 string
		arg2 string
		arg3 string
	}
	getWithVersionReturns struct {
		result1 *v1alpha1.ProfileDescription
	}
	getWithVersionReturnsOnCall map[int]struct {
		result1 *v1alpha1.ProfileDescription
	}
	ProfilesGreaterThanVersionStub        func(string, string, string) []v1alpha1.ProfileDescription
	profilesGreaterThanVersionMutex       sync.RWMutex
	profilesGreaterThanVersionArgsForCall []struct {
		arg1 string
		arg2 string
		arg3 string
	}
	profilesGreaterThanVersionReturns struct {
		result1 []v1alpha1.ProfileDescription
	}
	profilesGreaterThanVersionReturnsOnCall map[int]struct {
		result1 []v1alpha1.ProfileDescription
	}
	SearchStub        func(string) []v1alpha1.ProfileDescription
	searchMutex       sync.RWMutex
	searchArgsForCall []struct {
		arg1 string
	}
	searchReturns struct {
		result1 []v1alpha1.ProfileDescription
	}
	searchReturnsOnCall map[int]struct {
		result1 []v1alpha1.ProfileDescription
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeCatalog) Get(arg1 string, arg2 string) *v1alpha1.ProfileDescription {
	fake.getMutex.Lock()
	ret, specificReturn := fake.getReturnsOnCall[len(fake.getArgsForCall)]
	fake.getArgsForCall = append(fake.getArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	stub := fake.GetStub
	fakeReturns := fake.getReturns
	fake.recordInvocation("Get", []interface{}{arg1, arg2})
	fake.getMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeCatalog) GetCallCount() int {
	fake.getMutex.RLock()
	defer fake.getMutex.RUnlock()
	return len(fake.getArgsForCall)
}

func (fake *FakeCatalog) GetCalls(stub func(string, string) *v1alpha1.ProfileDescription) {
	fake.getMutex.Lock()
	defer fake.getMutex.Unlock()
	fake.GetStub = stub
}

func (fake *FakeCatalog) GetArgsForCall(i int) (string, string) {
	fake.getMutex.RLock()
	defer fake.getMutex.RUnlock()
	argsForCall := fake.getArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeCatalog) GetReturns(result1 *v1alpha1.ProfileDescription) {
	fake.getMutex.Lock()
	defer fake.getMutex.Unlock()
	fake.GetStub = nil
	fake.getReturns = struct {
		result1 *v1alpha1.ProfileDescription
	}{result1}
}

func (fake *FakeCatalog) GetReturnsOnCall(i int, result1 *v1alpha1.ProfileDescription) {
	fake.getMutex.Lock()
	defer fake.getMutex.Unlock()
	fake.GetStub = nil
	if fake.getReturnsOnCall == nil {
		fake.getReturnsOnCall = make(map[int]struct {
			result1 *v1alpha1.ProfileDescription
		})
	}
	fake.getReturnsOnCall[i] = struct {
		result1 *v1alpha1.ProfileDescription
	}{result1}
}

func (fake *FakeCatalog) GetWithVersion(arg1 string, arg2 string, arg3 string) *v1alpha1.ProfileDescription {
	fake.getWithVersionMutex.Lock()
	ret, specificReturn := fake.getWithVersionReturnsOnCall[len(fake.getWithVersionArgsForCall)]
	fake.getWithVersionArgsForCall = append(fake.getWithVersionArgsForCall, struct {
		arg1 string
		arg2 string
		arg3 string
	}{arg1, arg2, arg3})
	stub := fake.GetWithVersionStub
	fakeReturns := fake.getWithVersionReturns
	fake.recordInvocation("GetWithVersion", []interface{}{arg1, arg2, arg3})
	fake.getWithVersionMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeCatalog) GetWithVersionCallCount() int {
	fake.getWithVersionMutex.RLock()
	defer fake.getWithVersionMutex.RUnlock()
	return len(fake.getWithVersionArgsForCall)
}

func (fake *FakeCatalog) GetWithVersionCalls(stub func(string, string, string) *v1alpha1.ProfileDescription) {
	fake.getWithVersionMutex.Lock()
	defer fake.getWithVersionMutex.Unlock()
	fake.GetWithVersionStub = stub
}

func (fake *FakeCatalog) GetWithVersionArgsForCall(i int) (string, string, string) {
	fake.getWithVersionMutex.RLock()
	defer fake.getWithVersionMutex.RUnlock()
	argsForCall := fake.getWithVersionArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeCatalog) GetWithVersionReturns(result1 *v1alpha1.ProfileDescription) {
	fake.getWithVersionMutex.Lock()
	defer fake.getWithVersionMutex.Unlock()
	fake.GetWithVersionStub = nil
	fake.getWithVersionReturns = struct {
		result1 *v1alpha1.ProfileDescription
	}{result1}
}

func (fake *FakeCatalog) GetWithVersionReturnsOnCall(i int, result1 *v1alpha1.ProfileDescription) {
	fake.getWithVersionMutex.Lock()
	defer fake.getWithVersionMutex.Unlock()
	fake.GetWithVersionStub = nil
	if fake.getWithVersionReturnsOnCall == nil {
		fake.getWithVersionReturnsOnCall = make(map[int]struct {
			result1 *v1alpha1.ProfileDescription
		})
	}
	fake.getWithVersionReturnsOnCall[i] = struct {
		result1 *v1alpha1.ProfileDescription
	}{result1}
}

func (fake *FakeCatalog) ProfilesGreaterThanVersion(arg1 string, arg2 string, arg3 string) []v1alpha1.ProfileDescription {
	fake.profilesGreaterThanVersionMutex.Lock()
	ret, specificReturn := fake.profilesGreaterThanVersionReturnsOnCall[len(fake.profilesGreaterThanVersionArgsForCall)]
	fake.profilesGreaterThanVersionArgsForCall = append(fake.profilesGreaterThanVersionArgsForCall, struct {
		arg1 string
		arg2 string
		arg3 string
	}{arg1, arg2, arg3})
	stub := fake.ProfilesGreaterThanVersionStub
	fakeReturns := fake.profilesGreaterThanVersionReturns
	fake.recordInvocation("ProfilesGreaterThanVersion", []interface{}{arg1, arg2, arg3})
	fake.profilesGreaterThanVersionMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeCatalog) ProfilesGreaterThanVersionCallCount() int {
	fake.profilesGreaterThanVersionMutex.RLock()
	defer fake.profilesGreaterThanVersionMutex.RUnlock()
	return len(fake.profilesGreaterThanVersionArgsForCall)
}

func (fake *FakeCatalog) ProfilesGreaterThanVersionCalls(stub func(string, string, string) []v1alpha1.ProfileDescription) {
	fake.profilesGreaterThanVersionMutex.Lock()
	defer fake.profilesGreaterThanVersionMutex.Unlock()
	fake.ProfilesGreaterThanVersionStub = stub
}

func (fake *FakeCatalog) ProfilesGreaterThanVersionArgsForCall(i int) (string, string, string) {
	fake.profilesGreaterThanVersionMutex.RLock()
	defer fake.profilesGreaterThanVersionMutex.RUnlock()
	argsForCall := fake.profilesGreaterThanVersionArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeCatalog) ProfilesGreaterThanVersionReturns(result1 []v1alpha1.ProfileDescription) {
	fake.profilesGreaterThanVersionMutex.Lock()
	defer fake.profilesGreaterThanVersionMutex.Unlock()
	fake.ProfilesGreaterThanVersionStub = nil
	fake.profilesGreaterThanVersionReturns = struct {
		result1 []v1alpha1.ProfileDescription
	}{result1}
}

func (fake *FakeCatalog) ProfilesGreaterThanVersionReturnsOnCall(i int, result1 []v1alpha1.ProfileDescription) {
	fake.profilesGreaterThanVersionMutex.Lock()
	defer fake.profilesGreaterThanVersionMutex.Unlock()
	fake.ProfilesGreaterThanVersionStub = nil
	if fake.profilesGreaterThanVersionReturnsOnCall == nil {
		fake.profilesGreaterThanVersionReturnsOnCall = make(map[int]struct {
			result1 []v1alpha1.ProfileDescription
		})
	}
	fake.profilesGreaterThanVersionReturnsOnCall[i] = struct {
		result1 []v1alpha1.ProfileDescription
	}{result1}
}

func (fake *FakeCatalog) Search(arg1 string) []v1alpha1.ProfileDescription {
	fake.searchMutex.Lock()
	ret, specificReturn := fake.searchReturnsOnCall[len(fake.searchArgsForCall)]
	fake.searchArgsForCall = append(fake.searchArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.SearchStub
	fakeReturns := fake.searchReturns
	fake.recordInvocation("Search", []interface{}{arg1})
	fake.searchMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeCatalog) SearchCallCount() int {
	fake.searchMutex.RLock()
	defer fake.searchMutex.RUnlock()
	return len(fake.searchArgsForCall)
}

func (fake *FakeCatalog) SearchCalls(stub func(string) []v1alpha1.ProfileDescription) {
	fake.searchMutex.Lock()
	defer fake.searchMutex.Unlock()
	fake.SearchStub = stub
}

func (fake *FakeCatalog) SearchArgsForCall(i int) string {
	fake.searchMutex.RLock()
	defer fake.searchMutex.RUnlock()
	argsForCall := fake.searchArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeCatalog) SearchReturns(result1 []v1alpha1.ProfileDescription) {
	fake.searchMutex.Lock()
	defer fake.searchMutex.Unlock()
	fake.SearchStub = nil
	fake.searchReturns = struct {
		result1 []v1alpha1.ProfileDescription
	}{result1}
}

func (fake *FakeCatalog) SearchReturnsOnCall(i int, result1 []v1alpha1.ProfileDescription) {
	fake.searchMutex.Lock()
	defer fake.searchMutex.Unlock()
	fake.SearchStub = nil
	if fake.searchReturnsOnCall == nil {
		fake.searchReturnsOnCall = make(map[int]struct {
			result1 []v1alpha1.ProfileDescription
		})
	}
	fake.searchReturnsOnCall[i] = struct {
		result1 []v1alpha1.ProfileDescription
	}{result1}
}

func (fake *FakeCatalog) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getMutex.RLock()
	defer fake.getMutex.RUnlock()
	fake.getWithVersionMutex.RLock()
	defer fake.getWithVersionMutex.RUnlock()
	fake.profilesGreaterThanVersionMutex.RLock()
	defer fake.profilesGreaterThanVersionMutex.RUnlock()
	fake.searchMutex.RLock()
	defer fake.searchMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeCatalog) recordInvocation(key string, args []interface{}) {
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

var _ api.Catalog = new(FakeCatalog)
