// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"context"
	"sync"

	"github.com/alphagov/paas-billing/eventio"
)

type FakeEventFetcher struct {
	FetchEventsStub        func(ctx context.Context, lastKnownEvent *eventio.RawEvent) ([]eventio.RawEvent, error)
	fetchEventsMutex       sync.RWMutex
	fetchEventsArgsForCall []struct {
		ctx            context.Context
		lastKnownEvent *eventio.RawEvent
	}
	fetchEventsReturns struct {
		result1 []eventio.RawEvent
		result2 error
	}
	fetchEventsReturnsOnCall map[int]struct {
		result1 []eventio.RawEvent
		result2 error
	}
	KindStub        func() string
	kindMutex       sync.RWMutex
	kindArgsForCall []struct{}
	kindReturns     struct {
		result1 string
	}
	kindReturnsOnCall map[int]struct {
		result1 string
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeEventFetcher) FetchEvents(ctx context.Context, lastKnownEvent *eventio.RawEvent) ([]eventio.RawEvent, error) {
	fake.fetchEventsMutex.Lock()
	ret, specificReturn := fake.fetchEventsReturnsOnCall[len(fake.fetchEventsArgsForCall)]
	fake.fetchEventsArgsForCall = append(fake.fetchEventsArgsForCall, struct {
		ctx            context.Context
		lastKnownEvent *eventio.RawEvent
	}{ctx, lastKnownEvent})
	fake.recordInvocation("FetchEvents", []interface{}{ctx, lastKnownEvent})
	fake.fetchEventsMutex.Unlock()
	if fake.FetchEventsStub != nil {
		return fake.FetchEventsStub(ctx, lastKnownEvent)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.fetchEventsReturns.result1, fake.fetchEventsReturns.result2
}

func (fake *FakeEventFetcher) FetchEventsCallCount() int {
	fake.fetchEventsMutex.RLock()
	defer fake.fetchEventsMutex.RUnlock()
	return len(fake.fetchEventsArgsForCall)
}

func (fake *FakeEventFetcher) FetchEventsArgsForCall(i int) (context.Context, *eventio.RawEvent) {
	fake.fetchEventsMutex.RLock()
	defer fake.fetchEventsMutex.RUnlock()
	return fake.fetchEventsArgsForCall[i].ctx, fake.fetchEventsArgsForCall[i].lastKnownEvent
}

func (fake *FakeEventFetcher) FetchEventsReturns(result1 []eventio.RawEvent, result2 error) {
	fake.FetchEventsStub = nil
	fake.fetchEventsReturns = struct {
		result1 []eventio.RawEvent
		result2 error
	}{result1, result2}
}

func (fake *FakeEventFetcher) FetchEventsReturnsOnCall(i int, result1 []eventio.RawEvent, result2 error) {
	fake.FetchEventsStub = nil
	if fake.fetchEventsReturnsOnCall == nil {
		fake.fetchEventsReturnsOnCall = make(map[int]struct {
			result1 []eventio.RawEvent
			result2 error
		})
	}
	fake.fetchEventsReturnsOnCall[i] = struct {
		result1 []eventio.RawEvent
		result2 error
	}{result1, result2}
}

func (fake *FakeEventFetcher) Kind() string {
	fake.kindMutex.Lock()
	ret, specificReturn := fake.kindReturnsOnCall[len(fake.kindArgsForCall)]
	fake.kindArgsForCall = append(fake.kindArgsForCall, struct{}{})
	fake.recordInvocation("Kind", []interface{}{})
	fake.kindMutex.Unlock()
	if fake.KindStub != nil {
		return fake.KindStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.kindReturns.result1
}

func (fake *FakeEventFetcher) KindCallCount() int {
	fake.kindMutex.RLock()
	defer fake.kindMutex.RUnlock()
	return len(fake.kindArgsForCall)
}

func (fake *FakeEventFetcher) KindReturns(result1 string) {
	fake.KindStub = nil
	fake.kindReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeEventFetcher) KindReturnsOnCall(i int, result1 string) {
	fake.KindStub = nil
	if fake.kindReturnsOnCall == nil {
		fake.kindReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.kindReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeEventFetcher) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.fetchEventsMutex.RLock()
	defer fake.fetchEventsMutex.RUnlock()
	fake.kindMutex.RLock()
	defer fake.kindMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeEventFetcher) recordInvocation(key string, args []interface{}) {
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

var _ eventio.EventFetcher = new(FakeEventFetcher)