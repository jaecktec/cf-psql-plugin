// This file was generated by counterfeiter
package cfmysqlfakes

import (
	"sync"

	"github.com/andreasf/cf-mysql-plugin/cfmysql"
)

type FakeMysqlRunner struct {
	RunMysqlStub        func(hostname string, port int, dbName string, username string, password string, args ...string) error
	runMysqlMutex       sync.RWMutex
	runMysqlArgsForCall []struct {
		hostname string
		port     int
		dbName   string
		username string
		password string
		args     []string
	}
	runMysqlReturns struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeMysqlRunner) RunMysql(hostname string, port int, dbName string, username string, password string, args ...string) error {
	fake.runMysqlMutex.Lock()
	fake.runMysqlArgsForCall = append(fake.runMysqlArgsForCall, struct {
		hostname string
		port     int
		dbName   string
		username string
		password string
		args     []string
	}{hostname, port, dbName, username, password, args})
	fake.recordInvocation("RunMysql", []interface{}{hostname, port, dbName, username, password, args})
	fake.runMysqlMutex.Unlock()
	if fake.RunMysqlStub != nil {
		return fake.RunMysqlStub(hostname, port, dbName, username, password, args...)
	} else {
		return fake.runMysqlReturns.result1
	}
}

func (fake *FakeMysqlRunner) RunMysqlCallCount() int {
	fake.runMysqlMutex.RLock()
	defer fake.runMysqlMutex.RUnlock()
	return len(fake.runMysqlArgsForCall)
}

func (fake *FakeMysqlRunner) RunMysqlArgsForCall(i int) (string, int, string, string, string, []string) {
	fake.runMysqlMutex.RLock()
	defer fake.runMysqlMutex.RUnlock()
	return fake.runMysqlArgsForCall[i].hostname, fake.runMysqlArgsForCall[i].port, fake.runMysqlArgsForCall[i].dbName, fake.runMysqlArgsForCall[i].username, fake.runMysqlArgsForCall[i].password, fake.runMysqlArgsForCall[i].args
}

func (fake *FakeMysqlRunner) RunMysqlReturns(result1 error) {
	fake.RunMysqlStub = nil
	fake.runMysqlReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeMysqlRunner) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.runMysqlMutex.RLock()
	defer fake.runMysqlMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeMysqlRunner) recordInvocation(key string, args []interface{}) {
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

var _ cfmysql.MysqlRunner = new(FakeMysqlRunner)
