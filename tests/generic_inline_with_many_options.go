// Code generated by http://github.com/gojuno/minimock (dev). DO NOT EDIT.

package tests

//go:generate minimock -i github.com/gojuno/minimock/v3/tests.genericInlineUnionWithManyTypes -o generic_inline_with_many_options.go -n GenericInlineUnionWithManyTypesMock -p tests

import (
	"sync"
	mm_atomic "sync/atomic"
	mm_time "time"

	"github.com/gojuno/minimock/v3"
)

// GenericInlineUnionWithManyTypesMock implements genericInlineUnionWithManyTypes
type GenericInlineUnionWithManyTypesMock[T int | float64 | string] struct {
	t          minimock.Tester
	finishOnce sync.Once

	funcName          func(t1 T)
	inspectFuncName   func(t1 T)
	afterNameCounter  uint64
	beforeNameCounter uint64
	NameMock          mGenericInlineUnionWithManyTypesMockName[T]
}

// NewGenericInlineUnionWithManyTypesMock returns a mock for genericInlineUnionWithManyTypes
func NewGenericInlineUnionWithManyTypesMock[T int | float64 | string](t minimock.Tester) *GenericInlineUnionWithManyTypesMock[T] {
	m := &GenericInlineUnionWithManyTypesMock[T]{t: t}

	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.NameMock = mGenericInlineUnionWithManyTypesMockName[T]{mock: m}
	m.NameMock.callArgs = []*GenericInlineUnionWithManyTypesMockNameParams[T]{}

	t.Cleanup(m.MinimockFinish)

	return m
}

type mGenericInlineUnionWithManyTypesMockName[T int | float64 | string] struct {
	optional           bool
	mock               *GenericInlineUnionWithManyTypesMock[T]
	defaultExpectation *GenericInlineUnionWithManyTypesMockNameExpectation[T]
	expectations       []*GenericInlineUnionWithManyTypesMockNameExpectation[T]

	callArgs []*GenericInlineUnionWithManyTypesMockNameParams[T]
	mutex    sync.RWMutex

	expectedInvocations uint64
}

// GenericInlineUnionWithManyTypesMockNameExpectation specifies expectation struct of the genericInlineUnionWithManyTypes.Name
type GenericInlineUnionWithManyTypesMockNameExpectation[T int | float64 | string] struct {
	mock      *GenericInlineUnionWithManyTypesMock[T]
	params    *GenericInlineUnionWithManyTypesMockNameParams[T]
	paramPtrs *GenericInlineUnionWithManyTypesMockNameParamPtrs[T]

	Counter uint64
}

// GenericInlineUnionWithManyTypesMockNameParams contains parameters of the genericInlineUnionWithManyTypes.Name
type GenericInlineUnionWithManyTypesMockNameParams[T int | float64 | string] struct {
	t1 T
}

// GenericInlineUnionWithManyTypesMockNameParamPtrs contains pointers to parameters of the genericInlineUnionWithManyTypes.Name
type GenericInlineUnionWithManyTypesMockNameParamPtrs[T int | float64 | string] struct {
	t1 *T
}

// Marks this method to be optional. The default behavior of any method with Return() is '1 or more', meaning
// the test will fail minimock's automatic final call check if the mocked method was not called at least once.
// Optional() makes method check to work in '0 or more' mode.
// It is NOT RECOMMENDED to use this option unless you really need it, as default behaviour helps to
// catch the problems when the expected method call is totally skipped during test run.
func (mmName *mGenericInlineUnionWithManyTypesMockName[T]) Optional() *mGenericInlineUnionWithManyTypesMockName[T] {
	mmName.optional = true
	return mmName
}

// Expect sets up expected params for genericInlineUnionWithManyTypes.Name
func (mmName *mGenericInlineUnionWithManyTypesMockName[T]) Expect(t1 T) *mGenericInlineUnionWithManyTypesMockName[T] {
	if mmName.mock.funcName != nil {
		mmName.mock.t.Fatalf("GenericInlineUnionWithManyTypesMock.Name mock is already set by Set")
	}

	if mmName.defaultExpectation == nil {
		mmName.defaultExpectation = &GenericInlineUnionWithManyTypesMockNameExpectation[T]{}
	}

	if mmName.defaultExpectation.paramPtrs != nil {
		mmName.mock.t.Fatalf("GenericInlineUnionWithManyTypesMock.Name mock is already set by ExpectParams functions")
	}

	mmName.defaultExpectation.params = &GenericInlineUnionWithManyTypesMockNameParams[T]{t1}
	for _, e := range mmName.expectations {
		if minimock.Equal(e.params, mmName.defaultExpectation.params) {
			mmName.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmName.defaultExpectation.params)
		}
	}

	return mmName
}

// ExpectT1Param1 sets up expected param t1 for genericInlineUnionWithManyTypes.Name
func (mmName *mGenericInlineUnionWithManyTypesMockName[T]) ExpectT1Param1(t1 T) *mGenericInlineUnionWithManyTypesMockName[T] {
	if mmName.mock.funcName != nil {
		mmName.mock.t.Fatalf("GenericInlineUnionWithManyTypesMock.Name mock is already set by Set")
	}

	if mmName.defaultExpectation == nil {
		mmName.defaultExpectation = &GenericInlineUnionWithManyTypesMockNameExpectation[T]{}
	}

	if mmName.defaultExpectation.params != nil {
		mmName.mock.t.Fatalf("GenericInlineUnionWithManyTypesMock.Name mock is already set by Expect")
	}

	if mmName.defaultExpectation.paramPtrs == nil {
		mmName.defaultExpectation.paramPtrs = &GenericInlineUnionWithManyTypesMockNameParamPtrs[T]{}
	}
	mmName.defaultExpectation.paramPtrs.t1 = &t1

	return mmName
}

// Inspect accepts an inspector function that has same arguments as the genericInlineUnionWithManyTypes.Name
func (mmName *mGenericInlineUnionWithManyTypesMockName[T]) Inspect(f func(t1 T)) *mGenericInlineUnionWithManyTypesMockName[T] {
	if mmName.mock.inspectFuncName != nil {
		mmName.mock.t.Fatalf("Inspect function is already set for GenericInlineUnionWithManyTypesMock.Name")
	}

	mmName.mock.inspectFuncName = f

	return mmName
}

// Return sets up results that will be returned by genericInlineUnionWithManyTypes.Name
func (mmName *mGenericInlineUnionWithManyTypesMockName[T]) Return() *GenericInlineUnionWithManyTypesMock[T] {
	if mmName.mock.funcName != nil {
		mmName.mock.t.Fatalf("GenericInlineUnionWithManyTypesMock.Name mock is already set by Set")
	}

	if mmName.defaultExpectation == nil {
		mmName.defaultExpectation = &GenericInlineUnionWithManyTypesMockNameExpectation[T]{mock: mmName.mock}
	}

	return mmName.mock
}

// Set uses given function f to mock the genericInlineUnionWithManyTypes.Name method
func (mmName *mGenericInlineUnionWithManyTypesMockName[T]) Set(f func(t1 T)) *GenericInlineUnionWithManyTypesMock[T] {
	if mmName.defaultExpectation != nil {
		mmName.mock.t.Fatalf("Default expectation is already set for the genericInlineUnionWithManyTypes.Name method")
	}

	if len(mmName.expectations) > 0 {
		mmName.mock.t.Fatalf("Some expectations are already set for the genericInlineUnionWithManyTypes.Name method")
	}

	mmName.mock.funcName = f
	return mmName.mock
}

// Times sets number of times genericInlineUnionWithManyTypes.Name should be invoked
func (mmName *mGenericInlineUnionWithManyTypesMockName[T]) Times(n uint64) *mGenericInlineUnionWithManyTypesMockName[T] {
	if n == 0 {
		mmName.mock.t.Fatalf("Times of GenericInlineUnionWithManyTypesMock.Name mock can not be zero")
	}
	mm_atomic.StoreUint64(&mmName.expectedInvocations, n)
	return mmName
}

func (mmName *mGenericInlineUnionWithManyTypesMockName[T]) invocationsDone() bool {
	if len(mmName.expectations) == 0 && mmName.defaultExpectation == nil && mmName.mock.funcName == nil {
		return true
	}

	totalInvocations := mm_atomic.LoadUint64(&mmName.mock.afterNameCounter)
	expectedInvocations := mm_atomic.LoadUint64(&mmName.expectedInvocations)

	return totalInvocations > 0 && (expectedInvocations == 0 || expectedInvocations == totalInvocations)
}

// Name implements genericInlineUnionWithManyTypes
func (mmName *GenericInlineUnionWithManyTypesMock[T]) Name(t1 T) {
	mm_atomic.AddUint64(&mmName.beforeNameCounter, 1)
	defer mm_atomic.AddUint64(&mmName.afterNameCounter, 1)

	if helper, ok := mmName.t.(interface{ Helper() }); ok {
		helper.Helper()
	}

	if mmName.inspectFuncName != nil {
		mmName.inspectFuncName(t1)
	}

	mm_params := GenericInlineUnionWithManyTypesMockNameParams[T]{t1}

	// Record call args
	mmName.NameMock.mutex.Lock()
	mmName.NameMock.callArgs = append(mmName.NameMock.callArgs, &mm_params)
	mmName.NameMock.mutex.Unlock()

	for _, e := range mmName.NameMock.expectations {
		if minimock.Equal(*e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return
		}
	}

	if mmName.NameMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmName.NameMock.defaultExpectation.Counter, 1)
		mm_want := mmName.NameMock.defaultExpectation.params
		mm_want_ptrs := mmName.NameMock.defaultExpectation.paramPtrs

		mm_got := GenericInlineUnionWithManyTypesMockNameParams[T]{t1}

		if mm_want_ptrs != nil {

			if mm_want_ptrs.t1 != nil && !minimock.Equal(*mm_want_ptrs.t1, mm_got.t1) {
				mmName.t.Errorf("GenericInlineUnionWithManyTypesMock.Name got unexpected parameter t1, want: %#v, got: %#v%s\n", *mm_want_ptrs.t1, mm_got.t1, minimock.Diff(*mm_want_ptrs.t1, mm_got.t1))
			}

		} else if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmName.t.Errorf("GenericInlineUnionWithManyTypesMock.Name got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		return

	}
	if mmName.funcName != nil {
		mmName.funcName(t1)
		return
	}
	mmName.t.Fatalf("Unexpected call to GenericInlineUnionWithManyTypesMock.Name. %v", t1)

}

// NameAfterCounter returns a count of finished GenericInlineUnionWithManyTypesMock.Name invocations
func (mmName *GenericInlineUnionWithManyTypesMock[T]) NameAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmName.afterNameCounter)
}

// NameBeforeCounter returns a count of GenericInlineUnionWithManyTypesMock.Name invocations
func (mmName *GenericInlineUnionWithManyTypesMock[T]) NameBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmName.beforeNameCounter)
}

// Calls returns a list of arguments used in each call to GenericInlineUnionWithManyTypesMock.Name.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmName *mGenericInlineUnionWithManyTypesMockName[T]) Calls() []*GenericInlineUnionWithManyTypesMockNameParams[T] {
	mmName.mutex.RLock()

	argCopy := make([]*GenericInlineUnionWithManyTypesMockNameParams[T], len(mmName.callArgs))
	copy(argCopy, mmName.callArgs)

	mmName.mutex.RUnlock()

	return argCopy
}

// MinimockNameDone returns true if the count of the Name invocations corresponds
// the number of defined expectations
func (m *GenericInlineUnionWithManyTypesMock[T]) MinimockNameDone() bool {
	if m.NameMock.optional {
		// Optional methods provide '0 or more' call count restriction.
		return true
	}

	for _, e := range m.NameMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	return m.NameMock.invocationsDone()
}

// MinimockNameInspect logs each unmet expectation
func (m *GenericInlineUnionWithManyTypesMock[T]) MinimockNameInspect() {
	for _, e := range m.NameMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to GenericInlineUnionWithManyTypesMock.Name with params: %#v", *e.params)
		}
	}

	afterNameCounter := mm_atomic.LoadUint64(&m.afterNameCounter)
	// if default expectation was set then invocations count should be greater than zero
	if m.NameMock.defaultExpectation != nil && afterNameCounter < 1 {
		if m.NameMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to GenericInlineUnionWithManyTypesMock.Name")
		} else {
			m.t.Errorf("Expected call to GenericInlineUnionWithManyTypesMock.Name with params: %#v", *m.NameMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcName != nil && afterNameCounter < 1 {
		m.t.Error("Expected call to GenericInlineUnionWithManyTypesMock.Name")
	}

	if !m.NameMock.invocationsDone() && afterNameCounter > 0 {
		m.t.Errorf("Expected %d calls to GenericInlineUnionWithManyTypesMock.Name but found %d calls",
			mm_atomic.LoadUint64(&m.NameMock.expectedInvocations), afterNameCounter)
	}
}

// MinimockFinish checks that all mocked methods have been called the expected number of times
func (m *GenericInlineUnionWithManyTypesMock[T]) MinimockFinish() {
	m.finishOnce.Do(func() {
		if !m.minimockDone() {
			m.MinimockNameInspect()
		}
	})
}

// MinimockWait waits for all mocked methods to be called the expected number of times
func (m *GenericInlineUnionWithManyTypesMock[T]) MinimockWait(timeout mm_time.Duration) {
	timeoutCh := mm_time.After(timeout)
	for {
		if m.minimockDone() {
			return
		}
		select {
		case <-timeoutCh:
			m.MinimockFinish()
			return
		case <-mm_time.After(10 * mm_time.Millisecond):
		}
	}
}

func (m *GenericInlineUnionWithManyTypesMock[T]) minimockDone() bool {
	done := true
	return done &&
		m.MinimockNameDone()
}
