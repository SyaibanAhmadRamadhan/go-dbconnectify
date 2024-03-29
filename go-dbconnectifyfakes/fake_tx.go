// Code generated by counterfeiter. DO NOT EDIT.
package godbconnectifyfakes

import (
	"context"
	"sync"

	go_dbconnectify "github.com/SyaibanAhmadRamadhan/go-dbconnectify"
)

type FakeTx struct {
	DoTransactionStub        func(context.Context, *go_dbconnectify.TxOption, func(c context.Context) (commit bool, err error)) error
	doTransactionMutex       sync.RWMutex
	doTransactionArgsForCall []struct {
		arg1 context.Context
		arg2 *go_dbconnectify.TxOption
		arg3 func(c context.Context) (commit bool, err error)
	}
	doTransactionReturns struct {
		result1 error
	}
	doTransactionReturnsOnCall map[int]struct {
		result1 error
	}
	DoTransactionxStub        func(context.Context, *go_dbconnectify.TxOption, func(c context.Context) (err error)) error
	doTransactionxMutex       sync.RWMutex
	doTransactionxArgsForCall []struct {
		arg1 context.Context
		arg2 *go_dbconnectify.TxOption
		arg3 func(c context.Context) (err error)
	}
	doTransactionxReturns struct {
		result1 error
	}
	doTransactionxReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeTx) DoTransaction(arg1 context.Context, arg2 *go_dbconnectify.TxOption, arg3 func(c context.Context) (commit bool, err error)) error {
	fake.doTransactionMutex.Lock()
	ret, specificReturn := fake.doTransactionReturnsOnCall[len(fake.doTransactionArgsForCall)]
	fake.doTransactionArgsForCall = append(fake.doTransactionArgsForCall, struct {
		arg1 context.Context
		arg2 *go_dbconnectify.TxOption
		arg3 func(c context.Context) (commit bool, err error)
	}{arg1, arg2, arg3})
	stub := fake.DoTransactionStub
	fakeReturns := fake.doTransactionReturns
	fake.recordInvocation("DoTransaction", []interface{}{arg1, arg2, arg3})
	fake.doTransactionMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeTx) DoTransactionCallCount() int {
	fake.doTransactionMutex.RLock()
	defer fake.doTransactionMutex.RUnlock()
	return len(fake.doTransactionArgsForCall)
}

func (fake *FakeTx) DoTransactionCalls(stub func(context.Context, *go_dbconnectify.TxOption, func(c context.Context) (commit bool, err error)) error) {
	fake.doTransactionMutex.Lock()
	defer fake.doTransactionMutex.Unlock()
	fake.DoTransactionStub = stub
}

func (fake *FakeTx) DoTransactionArgsForCall(i int) (context.Context, *go_dbconnectify.TxOption, func(c context.Context) (commit bool, err error)) {
	fake.doTransactionMutex.RLock()
	defer fake.doTransactionMutex.RUnlock()
	argsForCall := fake.doTransactionArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeTx) DoTransactionReturns(result1 error) {
	fake.doTransactionMutex.Lock()
	defer fake.doTransactionMutex.Unlock()
	fake.DoTransactionStub = nil
	fake.doTransactionReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeTx) DoTransactionReturnsOnCall(i int, result1 error) {
	fake.doTransactionMutex.Lock()
	defer fake.doTransactionMutex.Unlock()
	fake.DoTransactionStub = nil
	if fake.doTransactionReturnsOnCall == nil {
		fake.doTransactionReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.doTransactionReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeTx) DoTransactionx(arg1 context.Context, arg2 *go_dbconnectify.TxOption, arg3 func(c context.Context) (err error)) error {
	fake.doTransactionxMutex.Lock()
	ret, specificReturn := fake.doTransactionxReturnsOnCall[len(fake.doTransactionxArgsForCall)]
	fake.doTransactionxArgsForCall = append(fake.doTransactionxArgsForCall, struct {
		arg1 context.Context
		arg2 *go_dbconnectify.TxOption
		arg3 func(c context.Context) (err error)
	}{arg1, arg2, arg3})
	stub := fake.DoTransactionxStub
	fakeReturns := fake.doTransactionxReturns
	fake.recordInvocation("DoTransactionx", []interface{}{arg1, arg2, arg3})
	fake.doTransactionxMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeTx) DoTransactionxCallCount() int {
	fake.doTransactionxMutex.RLock()
	defer fake.doTransactionxMutex.RUnlock()
	return len(fake.doTransactionxArgsForCall)
}

func (fake *FakeTx) DoTransactionxCalls(stub func(context.Context, *go_dbconnectify.TxOption, func(c context.Context) (err error)) error) {
	fake.doTransactionxMutex.Lock()
	defer fake.doTransactionxMutex.Unlock()
	fake.DoTransactionxStub = stub
}

func (fake *FakeTx) DoTransactionxArgsForCall(i int) (context.Context, *go_dbconnectify.TxOption, func(c context.Context) (err error)) {
	fake.doTransactionxMutex.RLock()
	defer fake.doTransactionxMutex.RUnlock()
	argsForCall := fake.doTransactionxArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeTx) DoTransactionxReturns(result1 error) {
	fake.doTransactionxMutex.Lock()
	defer fake.doTransactionxMutex.Unlock()
	fake.DoTransactionxStub = nil
	fake.doTransactionxReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeTx) DoTransactionxReturnsOnCall(i int, result1 error) {
	fake.doTransactionxMutex.Lock()
	defer fake.doTransactionxMutex.Unlock()
	fake.DoTransactionxStub = nil
	if fake.doTransactionxReturnsOnCall == nil {
		fake.doTransactionxReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.doTransactionxReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeTx) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.doTransactionMutex.RLock()
	defer fake.doTransactionMutex.RUnlock()
	fake.doTransactionxMutex.RLock()
	defer fake.doTransactionxMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeTx) recordInvocation(key string, args []interface{}) {
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

var _ go_dbconnectify.Tx = new(FakeTx)
