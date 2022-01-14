/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by counterfeiter. DO NOT EDIT.
package fastforwardfakes

import (
	"sync"

	"sigs.k8s.io/release-sdk/git"
)

type FakeImpl struct {
	AskStub        func(string, string, int) (string, bool, error)
	askMutex       sync.RWMutex
	askArgsForCall []struct {
		arg1 string
		arg2 string
		arg3 int
	}
	askReturns struct {
		result1 string
		result2 bool
		result3 error
	}
	askReturnsOnCall map[int]struct {
		result1 string
		result2 bool
		result3 error
	}
	CloneOrOpenDefaultGitHubRepoSSHStub        func(string) (*git.Repo, error)
	cloneOrOpenDefaultGitHubRepoSSHMutex       sync.RWMutex
	cloneOrOpenDefaultGitHubRepoSSHArgsForCall []struct {
		arg1 string
	}
	cloneOrOpenDefaultGitHubRepoSSHReturns struct {
		result1 *git.Repo
		result2 error
	}
	cloneOrOpenDefaultGitHubRepoSSHReturnsOnCall map[int]struct {
		result1 *git.Repo
		result2 error
	}
	IsReleaseBranchStub        func(string) bool
	isReleaseBranchMutex       sync.RWMutex
	isReleaseBranchArgsForCall []struct {
		arg1 string
	}
	isReleaseBranchReturns struct {
		result1 bool
	}
	isReleaseBranchReturnsOnCall map[int]struct {
		result1 bool
	}
	RepoCheckoutStub        func(*git.Repo, string, ...string) error
	repoCheckoutMutex       sync.RWMutex
	repoCheckoutArgsForCall []struct {
		arg1 *git.Repo
		arg2 string
		arg3 []string
	}
	repoCheckoutReturns struct {
		result1 error
	}
	repoCheckoutReturnsOnCall map[int]struct {
		result1 error
	}
	RepoCleanupStub        func(*git.Repo) error
	repoCleanupMutex       sync.RWMutex
	repoCleanupArgsForCall []struct {
		arg1 *git.Repo
	}
	repoCleanupReturns struct {
		result1 error
	}
	repoCleanupReturnsOnCall map[int]struct {
		result1 error
	}
	RepoCurrentBranchStub        func(*git.Repo) (string, error)
	repoCurrentBranchMutex       sync.RWMutex
	repoCurrentBranchArgsForCall []struct {
		arg1 *git.Repo
	}
	repoCurrentBranchReturns struct {
		result1 string
		result2 error
	}
	repoCurrentBranchReturnsOnCall map[int]struct {
		result1 string
		result2 error
	}
	RepoDescribeStub        func(*git.Repo, *git.DescribeOptions) (string, error)
	repoDescribeMutex       sync.RWMutex
	repoDescribeArgsForCall []struct {
		arg1 *git.Repo
		arg2 *git.DescribeOptions
	}
	repoDescribeReturns struct {
		result1 string
		result2 error
	}
	repoDescribeReturnsOnCall map[int]struct {
		result1 string
		result2 error
	}
	RepoDirStub        func(*git.Repo) string
	repoDirMutex       sync.RWMutex
	repoDirArgsForCall []struct {
		arg1 *git.Repo
	}
	repoDirReturns struct {
		result1 string
	}
	repoDirReturnsOnCall map[int]struct {
		result1 string
	}
	RepoHasRemoteBranchStub        func(*git.Repo, string) (bool, error)
	repoHasRemoteBranchMutex       sync.RWMutex
	repoHasRemoteBranchArgsForCall []struct {
		arg1 *git.Repo
		arg2 string
	}
	repoHasRemoteBranchReturns struct {
		result1 bool
		result2 error
	}
	repoHasRemoteBranchReturnsOnCall map[int]struct {
		result1 bool
		result2 error
	}
	RepoHeadStub        func(*git.Repo) (string, error)
	repoHeadMutex       sync.RWMutex
	repoHeadArgsForCall []struct {
		arg1 *git.Repo
	}
	repoHeadReturns struct {
		result1 string
		result2 error
	}
	repoHeadReturnsOnCall map[int]struct {
		result1 string
		result2 error
	}
	RepoMergeStub        func(*git.Repo, string) error
	repoMergeMutex       sync.RWMutex
	repoMergeArgsForCall []struct {
		arg1 *git.Repo
		arg2 string
	}
	repoMergeReturns struct {
		result1 error
	}
	repoMergeReturnsOnCall map[int]struct {
		result1 error
	}
	RepoMergeBaseStub        func(*git.Repo, string, string) (string, error)
	repoMergeBaseMutex       sync.RWMutex
	repoMergeBaseArgsForCall []struct {
		arg1 *git.Repo
		arg2 string
		arg3 string
	}
	repoMergeBaseReturns struct {
		result1 string
		result2 error
	}
	repoMergeBaseReturnsOnCall map[int]struct {
		result1 string
		result2 error
	}
	RepoPushStub        func(*git.Repo, string) error
	repoPushMutex       sync.RWMutex
	repoPushArgsForCall []struct {
		arg1 *git.Repo
		arg2 string
	}
	repoPushReturns struct {
		result1 error
	}
	repoPushReturnsOnCall map[int]struct {
		result1 error
	}
	RepoSetDryStub        func(*git.Repo)
	repoSetDryMutex       sync.RWMutex
	repoSetDryArgsForCall []struct {
		arg1 *git.Repo
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeImpl) Ask(arg1 string, arg2 string, arg3 int) (string, bool, error) {
	fake.askMutex.Lock()
	ret, specificReturn := fake.askReturnsOnCall[len(fake.askArgsForCall)]
	fake.askArgsForCall = append(fake.askArgsForCall, struct {
		arg1 string
		arg2 string
		arg3 int
	}{arg1, arg2, arg3})
	stub := fake.AskStub
	fakeReturns := fake.askReturns
	fake.recordInvocation("Ask", []interface{}{arg1, arg2, arg3})
	fake.askMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3
}

func (fake *FakeImpl) AskCallCount() int {
	fake.askMutex.RLock()
	defer fake.askMutex.RUnlock()
	return len(fake.askArgsForCall)
}

func (fake *FakeImpl) AskCalls(stub func(string, string, int) (string, bool, error)) {
	fake.askMutex.Lock()
	defer fake.askMutex.Unlock()
	fake.AskStub = stub
}

func (fake *FakeImpl) AskArgsForCall(i int) (string, string, int) {
	fake.askMutex.RLock()
	defer fake.askMutex.RUnlock()
	argsForCall := fake.askArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeImpl) AskReturns(result1 string, result2 bool, result3 error) {
	fake.askMutex.Lock()
	defer fake.askMutex.Unlock()
	fake.AskStub = nil
	fake.askReturns = struct {
		result1 string
		result2 bool
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeImpl) AskReturnsOnCall(i int, result1 string, result2 bool, result3 error) {
	fake.askMutex.Lock()
	defer fake.askMutex.Unlock()
	fake.AskStub = nil
	if fake.askReturnsOnCall == nil {
		fake.askReturnsOnCall = make(map[int]struct {
			result1 string
			result2 bool
			result3 error
		})
	}
	fake.askReturnsOnCall[i] = struct {
		result1 string
		result2 bool
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeImpl) CloneOrOpenDefaultGitHubRepoSSH(arg1 string) (*git.Repo, error) {
	fake.cloneOrOpenDefaultGitHubRepoSSHMutex.Lock()
	ret, specificReturn := fake.cloneOrOpenDefaultGitHubRepoSSHReturnsOnCall[len(fake.cloneOrOpenDefaultGitHubRepoSSHArgsForCall)]
	fake.cloneOrOpenDefaultGitHubRepoSSHArgsForCall = append(fake.cloneOrOpenDefaultGitHubRepoSSHArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.CloneOrOpenDefaultGitHubRepoSSHStub
	fakeReturns := fake.cloneOrOpenDefaultGitHubRepoSSHReturns
	fake.recordInvocation("CloneOrOpenDefaultGitHubRepoSSH", []interface{}{arg1})
	fake.cloneOrOpenDefaultGitHubRepoSSHMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeImpl) CloneOrOpenDefaultGitHubRepoSSHCallCount() int {
	fake.cloneOrOpenDefaultGitHubRepoSSHMutex.RLock()
	defer fake.cloneOrOpenDefaultGitHubRepoSSHMutex.RUnlock()
	return len(fake.cloneOrOpenDefaultGitHubRepoSSHArgsForCall)
}

func (fake *FakeImpl) CloneOrOpenDefaultGitHubRepoSSHCalls(stub func(string) (*git.Repo, error)) {
	fake.cloneOrOpenDefaultGitHubRepoSSHMutex.Lock()
	defer fake.cloneOrOpenDefaultGitHubRepoSSHMutex.Unlock()
	fake.CloneOrOpenDefaultGitHubRepoSSHStub = stub
}

func (fake *FakeImpl) CloneOrOpenDefaultGitHubRepoSSHArgsForCall(i int) string {
	fake.cloneOrOpenDefaultGitHubRepoSSHMutex.RLock()
	defer fake.cloneOrOpenDefaultGitHubRepoSSHMutex.RUnlock()
	argsForCall := fake.cloneOrOpenDefaultGitHubRepoSSHArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeImpl) CloneOrOpenDefaultGitHubRepoSSHReturns(result1 *git.Repo, result2 error) {
	fake.cloneOrOpenDefaultGitHubRepoSSHMutex.Lock()
	defer fake.cloneOrOpenDefaultGitHubRepoSSHMutex.Unlock()
	fake.CloneOrOpenDefaultGitHubRepoSSHStub = nil
	fake.cloneOrOpenDefaultGitHubRepoSSHReturns = struct {
		result1 *git.Repo
		result2 error
	}{result1, result2}
}

func (fake *FakeImpl) CloneOrOpenDefaultGitHubRepoSSHReturnsOnCall(i int, result1 *git.Repo, result2 error) {
	fake.cloneOrOpenDefaultGitHubRepoSSHMutex.Lock()
	defer fake.cloneOrOpenDefaultGitHubRepoSSHMutex.Unlock()
	fake.CloneOrOpenDefaultGitHubRepoSSHStub = nil
	if fake.cloneOrOpenDefaultGitHubRepoSSHReturnsOnCall == nil {
		fake.cloneOrOpenDefaultGitHubRepoSSHReturnsOnCall = make(map[int]struct {
			result1 *git.Repo
			result2 error
		})
	}
	fake.cloneOrOpenDefaultGitHubRepoSSHReturnsOnCall[i] = struct {
		result1 *git.Repo
		result2 error
	}{result1, result2}
}

func (fake *FakeImpl) IsReleaseBranch(arg1 string) bool {
	fake.isReleaseBranchMutex.Lock()
	ret, specificReturn := fake.isReleaseBranchReturnsOnCall[len(fake.isReleaseBranchArgsForCall)]
	fake.isReleaseBranchArgsForCall = append(fake.isReleaseBranchArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.IsReleaseBranchStub
	fakeReturns := fake.isReleaseBranchReturns
	fake.recordInvocation("IsReleaseBranch", []interface{}{arg1})
	fake.isReleaseBranchMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeImpl) IsReleaseBranchCallCount() int {
	fake.isReleaseBranchMutex.RLock()
	defer fake.isReleaseBranchMutex.RUnlock()
	return len(fake.isReleaseBranchArgsForCall)
}

func (fake *FakeImpl) IsReleaseBranchCalls(stub func(string) bool) {
	fake.isReleaseBranchMutex.Lock()
	defer fake.isReleaseBranchMutex.Unlock()
	fake.IsReleaseBranchStub = stub
}

func (fake *FakeImpl) IsReleaseBranchArgsForCall(i int) string {
	fake.isReleaseBranchMutex.RLock()
	defer fake.isReleaseBranchMutex.RUnlock()
	argsForCall := fake.isReleaseBranchArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeImpl) IsReleaseBranchReturns(result1 bool) {
	fake.isReleaseBranchMutex.Lock()
	defer fake.isReleaseBranchMutex.Unlock()
	fake.IsReleaseBranchStub = nil
	fake.isReleaseBranchReturns = struct {
		result1 bool
	}{result1}
}

func (fake *FakeImpl) IsReleaseBranchReturnsOnCall(i int, result1 bool) {
	fake.isReleaseBranchMutex.Lock()
	defer fake.isReleaseBranchMutex.Unlock()
	fake.IsReleaseBranchStub = nil
	if fake.isReleaseBranchReturnsOnCall == nil {
		fake.isReleaseBranchReturnsOnCall = make(map[int]struct {
			result1 bool
		})
	}
	fake.isReleaseBranchReturnsOnCall[i] = struct {
		result1 bool
	}{result1}
}

func (fake *FakeImpl) RepoCheckout(arg1 *git.Repo, arg2 string, arg3 ...string) error {
	fake.repoCheckoutMutex.Lock()
	ret, specificReturn := fake.repoCheckoutReturnsOnCall[len(fake.repoCheckoutArgsForCall)]
	fake.repoCheckoutArgsForCall = append(fake.repoCheckoutArgsForCall, struct {
		arg1 *git.Repo
		arg2 string
		arg3 []string
	}{arg1, arg2, arg3})
	stub := fake.RepoCheckoutStub
	fakeReturns := fake.repoCheckoutReturns
	fake.recordInvocation("RepoCheckout", []interface{}{arg1, arg2, arg3})
	fake.repoCheckoutMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3...)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeImpl) RepoCheckoutCallCount() int {
	fake.repoCheckoutMutex.RLock()
	defer fake.repoCheckoutMutex.RUnlock()
	return len(fake.repoCheckoutArgsForCall)
}

func (fake *FakeImpl) RepoCheckoutCalls(stub func(*git.Repo, string, ...string) error) {
	fake.repoCheckoutMutex.Lock()
	defer fake.repoCheckoutMutex.Unlock()
	fake.RepoCheckoutStub = stub
}

func (fake *FakeImpl) RepoCheckoutArgsForCall(i int) (*git.Repo, string, []string) {
	fake.repoCheckoutMutex.RLock()
	defer fake.repoCheckoutMutex.RUnlock()
	argsForCall := fake.repoCheckoutArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeImpl) RepoCheckoutReturns(result1 error) {
	fake.repoCheckoutMutex.Lock()
	defer fake.repoCheckoutMutex.Unlock()
	fake.RepoCheckoutStub = nil
	fake.repoCheckoutReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeImpl) RepoCheckoutReturnsOnCall(i int, result1 error) {
	fake.repoCheckoutMutex.Lock()
	defer fake.repoCheckoutMutex.Unlock()
	fake.RepoCheckoutStub = nil
	if fake.repoCheckoutReturnsOnCall == nil {
		fake.repoCheckoutReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.repoCheckoutReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeImpl) RepoCleanup(arg1 *git.Repo) error {
	fake.repoCleanupMutex.Lock()
	ret, specificReturn := fake.repoCleanupReturnsOnCall[len(fake.repoCleanupArgsForCall)]
	fake.repoCleanupArgsForCall = append(fake.repoCleanupArgsForCall, struct {
		arg1 *git.Repo
	}{arg1})
	stub := fake.RepoCleanupStub
	fakeReturns := fake.repoCleanupReturns
	fake.recordInvocation("RepoCleanup", []interface{}{arg1})
	fake.repoCleanupMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeImpl) RepoCleanupCallCount() int {
	fake.repoCleanupMutex.RLock()
	defer fake.repoCleanupMutex.RUnlock()
	return len(fake.repoCleanupArgsForCall)
}

func (fake *FakeImpl) RepoCleanupCalls(stub func(*git.Repo) error) {
	fake.repoCleanupMutex.Lock()
	defer fake.repoCleanupMutex.Unlock()
	fake.RepoCleanupStub = stub
}

func (fake *FakeImpl) RepoCleanupArgsForCall(i int) *git.Repo {
	fake.repoCleanupMutex.RLock()
	defer fake.repoCleanupMutex.RUnlock()
	argsForCall := fake.repoCleanupArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeImpl) RepoCleanupReturns(result1 error) {
	fake.repoCleanupMutex.Lock()
	defer fake.repoCleanupMutex.Unlock()
	fake.RepoCleanupStub = nil
	fake.repoCleanupReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeImpl) RepoCleanupReturnsOnCall(i int, result1 error) {
	fake.repoCleanupMutex.Lock()
	defer fake.repoCleanupMutex.Unlock()
	fake.RepoCleanupStub = nil
	if fake.repoCleanupReturnsOnCall == nil {
		fake.repoCleanupReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.repoCleanupReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeImpl) RepoCurrentBranch(arg1 *git.Repo) (string, error) {
	fake.repoCurrentBranchMutex.Lock()
	ret, specificReturn := fake.repoCurrentBranchReturnsOnCall[len(fake.repoCurrentBranchArgsForCall)]
	fake.repoCurrentBranchArgsForCall = append(fake.repoCurrentBranchArgsForCall, struct {
		arg1 *git.Repo
	}{arg1})
	stub := fake.RepoCurrentBranchStub
	fakeReturns := fake.repoCurrentBranchReturns
	fake.recordInvocation("RepoCurrentBranch", []interface{}{arg1})
	fake.repoCurrentBranchMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeImpl) RepoCurrentBranchCallCount() int {
	fake.repoCurrentBranchMutex.RLock()
	defer fake.repoCurrentBranchMutex.RUnlock()
	return len(fake.repoCurrentBranchArgsForCall)
}

func (fake *FakeImpl) RepoCurrentBranchCalls(stub func(*git.Repo) (string, error)) {
	fake.repoCurrentBranchMutex.Lock()
	defer fake.repoCurrentBranchMutex.Unlock()
	fake.RepoCurrentBranchStub = stub
}

func (fake *FakeImpl) RepoCurrentBranchArgsForCall(i int) *git.Repo {
	fake.repoCurrentBranchMutex.RLock()
	defer fake.repoCurrentBranchMutex.RUnlock()
	argsForCall := fake.repoCurrentBranchArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeImpl) RepoCurrentBranchReturns(result1 string, result2 error) {
	fake.repoCurrentBranchMutex.Lock()
	defer fake.repoCurrentBranchMutex.Unlock()
	fake.RepoCurrentBranchStub = nil
	fake.repoCurrentBranchReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeImpl) RepoCurrentBranchReturnsOnCall(i int, result1 string, result2 error) {
	fake.repoCurrentBranchMutex.Lock()
	defer fake.repoCurrentBranchMutex.Unlock()
	fake.RepoCurrentBranchStub = nil
	if fake.repoCurrentBranchReturnsOnCall == nil {
		fake.repoCurrentBranchReturnsOnCall = make(map[int]struct {
			result1 string
			result2 error
		})
	}
	fake.repoCurrentBranchReturnsOnCall[i] = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeImpl) RepoDescribe(arg1 *git.Repo, arg2 *git.DescribeOptions) (string, error) {
	fake.repoDescribeMutex.Lock()
	ret, specificReturn := fake.repoDescribeReturnsOnCall[len(fake.repoDescribeArgsForCall)]
	fake.repoDescribeArgsForCall = append(fake.repoDescribeArgsForCall, struct {
		arg1 *git.Repo
		arg2 *git.DescribeOptions
	}{arg1, arg2})
	stub := fake.RepoDescribeStub
	fakeReturns := fake.repoDescribeReturns
	fake.recordInvocation("RepoDescribe", []interface{}{arg1, arg2})
	fake.repoDescribeMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeImpl) RepoDescribeCallCount() int {
	fake.repoDescribeMutex.RLock()
	defer fake.repoDescribeMutex.RUnlock()
	return len(fake.repoDescribeArgsForCall)
}

func (fake *FakeImpl) RepoDescribeCalls(stub func(*git.Repo, *git.DescribeOptions) (string, error)) {
	fake.repoDescribeMutex.Lock()
	defer fake.repoDescribeMutex.Unlock()
	fake.RepoDescribeStub = stub
}

func (fake *FakeImpl) RepoDescribeArgsForCall(i int) (*git.Repo, *git.DescribeOptions) {
	fake.repoDescribeMutex.RLock()
	defer fake.repoDescribeMutex.RUnlock()
	argsForCall := fake.repoDescribeArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeImpl) RepoDescribeReturns(result1 string, result2 error) {
	fake.repoDescribeMutex.Lock()
	defer fake.repoDescribeMutex.Unlock()
	fake.RepoDescribeStub = nil
	fake.repoDescribeReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeImpl) RepoDescribeReturnsOnCall(i int, result1 string, result2 error) {
	fake.repoDescribeMutex.Lock()
	defer fake.repoDescribeMutex.Unlock()
	fake.RepoDescribeStub = nil
	if fake.repoDescribeReturnsOnCall == nil {
		fake.repoDescribeReturnsOnCall = make(map[int]struct {
			result1 string
			result2 error
		})
	}
	fake.repoDescribeReturnsOnCall[i] = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeImpl) RepoDir(arg1 *git.Repo) string {
	fake.repoDirMutex.Lock()
	ret, specificReturn := fake.repoDirReturnsOnCall[len(fake.repoDirArgsForCall)]
	fake.repoDirArgsForCall = append(fake.repoDirArgsForCall, struct {
		arg1 *git.Repo
	}{arg1})
	stub := fake.RepoDirStub
	fakeReturns := fake.repoDirReturns
	fake.recordInvocation("RepoDir", []interface{}{arg1})
	fake.repoDirMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeImpl) RepoDirCallCount() int {
	fake.repoDirMutex.RLock()
	defer fake.repoDirMutex.RUnlock()
	return len(fake.repoDirArgsForCall)
}

func (fake *FakeImpl) RepoDirCalls(stub func(*git.Repo) string) {
	fake.repoDirMutex.Lock()
	defer fake.repoDirMutex.Unlock()
	fake.RepoDirStub = stub
}

func (fake *FakeImpl) RepoDirArgsForCall(i int) *git.Repo {
	fake.repoDirMutex.RLock()
	defer fake.repoDirMutex.RUnlock()
	argsForCall := fake.repoDirArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeImpl) RepoDirReturns(result1 string) {
	fake.repoDirMutex.Lock()
	defer fake.repoDirMutex.Unlock()
	fake.RepoDirStub = nil
	fake.repoDirReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeImpl) RepoDirReturnsOnCall(i int, result1 string) {
	fake.repoDirMutex.Lock()
	defer fake.repoDirMutex.Unlock()
	fake.RepoDirStub = nil
	if fake.repoDirReturnsOnCall == nil {
		fake.repoDirReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.repoDirReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeImpl) RepoHasRemoteBranch(arg1 *git.Repo, arg2 string) (bool, error) {
	fake.repoHasRemoteBranchMutex.Lock()
	ret, specificReturn := fake.repoHasRemoteBranchReturnsOnCall[len(fake.repoHasRemoteBranchArgsForCall)]
	fake.repoHasRemoteBranchArgsForCall = append(fake.repoHasRemoteBranchArgsForCall, struct {
		arg1 *git.Repo
		arg2 string
	}{arg1, arg2})
	stub := fake.RepoHasRemoteBranchStub
	fakeReturns := fake.repoHasRemoteBranchReturns
	fake.recordInvocation("RepoHasRemoteBranch", []interface{}{arg1, arg2})
	fake.repoHasRemoteBranchMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeImpl) RepoHasRemoteBranchCallCount() int {
	fake.repoHasRemoteBranchMutex.RLock()
	defer fake.repoHasRemoteBranchMutex.RUnlock()
	return len(fake.repoHasRemoteBranchArgsForCall)
}

func (fake *FakeImpl) RepoHasRemoteBranchCalls(stub func(*git.Repo, string) (bool, error)) {
	fake.repoHasRemoteBranchMutex.Lock()
	defer fake.repoHasRemoteBranchMutex.Unlock()
	fake.RepoHasRemoteBranchStub = stub
}

func (fake *FakeImpl) RepoHasRemoteBranchArgsForCall(i int) (*git.Repo, string) {
	fake.repoHasRemoteBranchMutex.RLock()
	defer fake.repoHasRemoteBranchMutex.RUnlock()
	argsForCall := fake.repoHasRemoteBranchArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeImpl) RepoHasRemoteBranchReturns(result1 bool, result2 error) {
	fake.repoHasRemoteBranchMutex.Lock()
	defer fake.repoHasRemoteBranchMutex.Unlock()
	fake.RepoHasRemoteBranchStub = nil
	fake.repoHasRemoteBranchReturns = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *FakeImpl) RepoHasRemoteBranchReturnsOnCall(i int, result1 bool, result2 error) {
	fake.repoHasRemoteBranchMutex.Lock()
	defer fake.repoHasRemoteBranchMutex.Unlock()
	fake.RepoHasRemoteBranchStub = nil
	if fake.repoHasRemoteBranchReturnsOnCall == nil {
		fake.repoHasRemoteBranchReturnsOnCall = make(map[int]struct {
			result1 bool
			result2 error
		})
	}
	fake.repoHasRemoteBranchReturnsOnCall[i] = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *FakeImpl) RepoHead(arg1 *git.Repo) (string, error) {
	fake.repoHeadMutex.Lock()
	ret, specificReturn := fake.repoHeadReturnsOnCall[len(fake.repoHeadArgsForCall)]
	fake.repoHeadArgsForCall = append(fake.repoHeadArgsForCall, struct {
		arg1 *git.Repo
	}{arg1})
	stub := fake.RepoHeadStub
	fakeReturns := fake.repoHeadReturns
	fake.recordInvocation("RepoHead", []interface{}{arg1})
	fake.repoHeadMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeImpl) RepoHeadCallCount() int {
	fake.repoHeadMutex.RLock()
	defer fake.repoHeadMutex.RUnlock()
	return len(fake.repoHeadArgsForCall)
}

func (fake *FakeImpl) RepoHeadCalls(stub func(*git.Repo) (string, error)) {
	fake.repoHeadMutex.Lock()
	defer fake.repoHeadMutex.Unlock()
	fake.RepoHeadStub = stub
}

func (fake *FakeImpl) RepoHeadArgsForCall(i int) *git.Repo {
	fake.repoHeadMutex.RLock()
	defer fake.repoHeadMutex.RUnlock()
	argsForCall := fake.repoHeadArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeImpl) RepoHeadReturns(result1 string, result2 error) {
	fake.repoHeadMutex.Lock()
	defer fake.repoHeadMutex.Unlock()
	fake.RepoHeadStub = nil
	fake.repoHeadReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeImpl) RepoHeadReturnsOnCall(i int, result1 string, result2 error) {
	fake.repoHeadMutex.Lock()
	defer fake.repoHeadMutex.Unlock()
	fake.RepoHeadStub = nil
	if fake.repoHeadReturnsOnCall == nil {
		fake.repoHeadReturnsOnCall = make(map[int]struct {
			result1 string
			result2 error
		})
	}
	fake.repoHeadReturnsOnCall[i] = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeImpl) RepoMerge(arg1 *git.Repo, arg2 string) error {
	fake.repoMergeMutex.Lock()
	ret, specificReturn := fake.repoMergeReturnsOnCall[len(fake.repoMergeArgsForCall)]
	fake.repoMergeArgsForCall = append(fake.repoMergeArgsForCall, struct {
		arg1 *git.Repo
		arg2 string
	}{arg1, arg2})
	stub := fake.RepoMergeStub
	fakeReturns := fake.repoMergeReturns
	fake.recordInvocation("RepoMerge", []interface{}{arg1, arg2})
	fake.repoMergeMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeImpl) RepoMergeCallCount() int {
	fake.repoMergeMutex.RLock()
	defer fake.repoMergeMutex.RUnlock()
	return len(fake.repoMergeArgsForCall)
}

func (fake *FakeImpl) RepoMergeCalls(stub func(*git.Repo, string) error) {
	fake.repoMergeMutex.Lock()
	defer fake.repoMergeMutex.Unlock()
	fake.RepoMergeStub = stub
}

func (fake *FakeImpl) RepoMergeArgsForCall(i int) (*git.Repo, string) {
	fake.repoMergeMutex.RLock()
	defer fake.repoMergeMutex.RUnlock()
	argsForCall := fake.repoMergeArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeImpl) RepoMergeReturns(result1 error) {
	fake.repoMergeMutex.Lock()
	defer fake.repoMergeMutex.Unlock()
	fake.RepoMergeStub = nil
	fake.repoMergeReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeImpl) RepoMergeReturnsOnCall(i int, result1 error) {
	fake.repoMergeMutex.Lock()
	defer fake.repoMergeMutex.Unlock()
	fake.RepoMergeStub = nil
	if fake.repoMergeReturnsOnCall == nil {
		fake.repoMergeReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.repoMergeReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeImpl) RepoMergeBase(arg1 *git.Repo, arg2 string, arg3 string) (string, error) {
	fake.repoMergeBaseMutex.Lock()
	ret, specificReturn := fake.repoMergeBaseReturnsOnCall[len(fake.repoMergeBaseArgsForCall)]
	fake.repoMergeBaseArgsForCall = append(fake.repoMergeBaseArgsForCall, struct {
		arg1 *git.Repo
		arg2 string
		arg3 string
	}{arg1, arg2, arg3})
	stub := fake.RepoMergeBaseStub
	fakeReturns := fake.repoMergeBaseReturns
	fake.recordInvocation("RepoMergeBase", []interface{}{arg1, arg2, arg3})
	fake.repoMergeBaseMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeImpl) RepoMergeBaseCallCount() int {
	fake.repoMergeBaseMutex.RLock()
	defer fake.repoMergeBaseMutex.RUnlock()
	return len(fake.repoMergeBaseArgsForCall)
}

func (fake *FakeImpl) RepoMergeBaseCalls(stub func(*git.Repo, string, string) (string, error)) {
	fake.repoMergeBaseMutex.Lock()
	defer fake.repoMergeBaseMutex.Unlock()
	fake.RepoMergeBaseStub = stub
}

func (fake *FakeImpl) RepoMergeBaseArgsForCall(i int) (*git.Repo, string, string) {
	fake.repoMergeBaseMutex.RLock()
	defer fake.repoMergeBaseMutex.RUnlock()
	argsForCall := fake.repoMergeBaseArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeImpl) RepoMergeBaseReturns(result1 string, result2 error) {
	fake.repoMergeBaseMutex.Lock()
	defer fake.repoMergeBaseMutex.Unlock()
	fake.RepoMergeBaseStub = nil
	fake.repoMergeBaseReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeImpl) RepoMergeBaseReturnsOnCall(i int, result1 string, result2 error) {
	fake.repoMergeBaseMutex.Lock()
	defer fake.repoMergeBaseMutex.Unlock()
	fake.RepoMergeBaseStub = nil
	if fake.repoMergeBaseReturnsOnCall == nil {
		fake.repoMergeBaseReturnsOnCall = make(map[int]struct {
			result1 string
			result2 error
		})
	}
	fake.repoMergeBaseReturnsOnCall[i] = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeImpl) RepoPush(arg1 *git.Repo, arg2 string) error {
	fake.repoPushMutex.Lock()
	ret, specificReturn := fake.repoPushReturnsOnCall[len(fake.repoPushArgsForCall)]
	fake.repoPushArgsForCall = append(fake.repoPushArgsForCall, struct {
		arg1 *git.Repo
		arg2 string
	}{arg1, arg2})
	stub := fake.RepoPushStub
	fakeReturns := fake.repoPushReturns
	fake.recordInvocation("RepoPush", []interface{}{arg1, arg2})
	fake.repoPushMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeImpl) RepoPushCallCount() int {
	fake.repoPushMutex.RLock()
	defer fake.repoPushMutex.RUnlock()
	return len(fake.repoPushArgsForCall)
}

func (fake *FakeImpl) RepoPushCalls(stub func(*git.Repo, string) error) {
	fake.repoPushMutex.Lock()
	defer fake.repoPushMutex.Unlock()
	fake.RepoPushStub = stub
}

func (fake *FakeImpl) RepoPushArgsForCall(i int) (*git.Repo, string) {
	fake.repoPushMutex.RLock()
	defer fake.repoPushMutex.RUnlock()
	argsForCall := fake.repoPushArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeImpl) RepoPushReturns(result1 error) {
	fake.repoPushMutex.Lock()
	defer fake.repoPushMutex.Unlock()
	fake.RepoPushStub = nil
	fake.repoPushReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeImpl) RepoPushReturnsOnCall(i int, result1 error) {
	fake.repoPushMutex.Lock()
	defer fake.repoPushMutex.Unlock()
	fake.RepoPushStub = nil
	if fake.repoPushReturnsOnCall == nil {
		fake.repoPushReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.repoPushReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeImpl) RepoSetDry(arg1 *git.Repo) {
	fake.repoSetDryMutex.Lock()
	fake.repoSetDryArgsForCall = append(fake.repoSetDryArgsForCall, struct {
		arg1 *git.Repo
	}{arg1})
	stub := fake.RepoSetDryStub
	fake.recordInvocation("RepoSetDry", []interface{}{arg1})
	fake.repoSetDryMutex.Unlock()
	if stub != nil {
		fake.RepoSetDryStub(arg1)
	}
}

func (fake *FakeImpl) RepoSetDryCallCount() int {
	fake.repoSetDryMutex.RLock()
	defer fake.repoSetDryMutex.RUnlock()
	return len(fake.repoSetDryArgsForCall)
}

func (fake *FakeImpl) RepoSetDryCalls(stub func(*git.Repo)) {
	fake.repoSetDryMutex.Lock()
	defer fake.repoSetDryMutex.Unlock()
	fake.RepoSetDryStub = stub
}

func (fake *FakeImpl) RepoSetDryArgsForCall(i int) *git.Repo {
	fake.repoSetDryMutex.RLock()
	defer fake.repoSetDryMutex.RUnlock()
	argsForCall := fake.repoSetDryArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeImpl) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.askMutex.RLock()
	defer fake.askMutex.RUnlock()
	fake.cloneOrOpenDefaultGitHubRepoSSHMutex.RLock()
	defer fake.cloneOrOpenDefaultGitHubRepoSSHMutex.RUnlock()
	fake.isReleaseBranchMutex.RLock()
	defer fake.isReleaseBranchMutex.RUnlock()
	fake.repoCheckoutMutex.RLock()
	defer fake.repoCheckoutMutex.RUnlock()
	fake.repoCleanupMutex.RLock()
	defer fake.repoCleanupMutex.RUnlock()
	fake.repoCurrentBranchMutex.RLock()
	defer fake.repoCurrentBranchMutex.RUnlock()
	fake.repoDescribeMutex.RLock()
	defer fake.repoDescribeMutex.RUnlock()
	fake.repoDirMutex.RLock()
	defer fake.repoDirMutex.RUnlock()
	fake.repoHasRemoteBranchMutex.RLock()
	defer fake.repoHasRemoteBranchMutex.RUnlock()
	fake.repoHeadMutex.RLock()
	defer fake.repoHeadMutex.RUnlock()
	fake.repoMergeMutex.RLock()
	defer fake.repoMergeMutex.RUnlock()
	fake.repoMergeBaseMutex.RLock()
	defer fake.repoMergeBaseMutex.RUnlock()
	fake.repoPushMutex.RLock()
	defer fake.repoPushMutex.RUnlock()
	fake.repoSetDryMutex.RLock()
	defer fake.repoSetDryMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeImpl) recordInvocation(key string, args []interface{}) {
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
