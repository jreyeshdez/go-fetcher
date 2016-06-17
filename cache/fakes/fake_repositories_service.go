// This file was generated by counterfeiter
package fakes

import (
	"sync"

	"github.com/cloudfoundry/go-fetcher/cache"
	"github.com/google/go-github/github"
)

type FakeRepositoriesService struct {
	ListByOrgStub        func(org string, opt *github.RepositoryListByOrgOptions) ([]github.Repository, *github.Response, error)
	listByOrgMutex       sync.RWMutex
	listByOrgArgsForCall []struct {
		org string
		opt *github.RepositoryListByOrgOptions
	}
	listByOrgReturns struct {
		result1 []github.Repository
		result2 *github.Response
		result3 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeRepositoriesService) ListByOrg(org string, opt *github.RepositoryListByOrgOptions) ([]github.Repository, *github.Response, error) {
	fake.listByOrgMutex.Lock()
	fake.listByOrgArgsForCall = append(fake.listByOrgArgsForCall, struct {
		org string
		opt *github.RepositoryListByOrgOptions
	}{org, opt})
	fake.recordInvocation("ListByOrg", []interface{}{org, opt})
	fake.listByOrgMutex.Unlock()
	if fake.ListByOrgStub != nil {
		return fake.ListByOrgStub(org, opt)
	} else {
		return fake.listByOrgReturns.result1, fake.listByOrgReturns.result2, fake.listByOrgReturns.result3
	}
}

func (fake *FakeRepositoriesService) ListByOrgCallCount() int {
	fake.listByOrgMutex.RLock()
	defer fake.listByOrgMutex.RUnlock()
	return len(fake.listByOrgArgsForCall)
}

func (fake *FakeRepositoriesService) ListByOrgArgsForCall(i int) (string, *github.RepositoryListByOrgOptions) {
	fake.listByOrgMutex.RLock()
	defer fake.listByOrgMutex.RUnlock()
	return fake.listByOrgArgsForCall[i].org, fake.listByOrgArgsForCall[i].opt
}

func (fake *FakeRepositoriesService) ListByOrgReturns(result1 []github.Repository, result2 *github.Response, result3 error) {
	fake.ListByOrgStub = nil
	fake.listByOrgReturns = struct {
		result1 []github.Repository
		result2 *github.Response
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeRepositoriesService) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.listByOrgMutex.RLock()
	defer fake.listByOrgMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeRepositoriesService) recordInvocation(key string, args []interface{}) {
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

var _ cache.RepositoriesService = new(FakeRepositoriesService)
