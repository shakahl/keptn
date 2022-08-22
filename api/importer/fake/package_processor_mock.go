// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package fake

import (
	"github.com/keptn/keptn/api/importer/model"
	"io"
	"sync"
)

// ImportPackageMock is a mock implementation of importer.ImportPackage.
//
// 	func TestSomethingThatUsesImportPackage(t *testing.T) {
//
// 		// make and configure a mocked importer.ImportPackage
// 		mockedImportPackage := &ImportPackageMock{
// 			CloseFunc: func() error {
// 				panic("mock out the Close method")
// 			},
// 			GetResourceFunc: func(resourceName string) (io.ReadCloser, error) {
// 				panic("mock out the GetResource method")
// 			},
// 		}
//
// 		// use mockedImportPackage in code that requires importer.ImportPackage
// 		// and then make assertions.
//
// 	}
type ImportPackageMock struct {
	// CloseFunc mocks the Close method.
	CloseFunc func() error

	// GetResourceFunc mocks the GetResource method.
	GetResourceFunc func(resourceName string) (io.ReadCloser, error)

	ResourceExistsFunc func(resourceName string) (bool, error)

	// calls tracks calls to the methods.
	calls struct {
		// Close holds details about calls to the Close method.
		Close []struct {
		}
		// GetResource holds details about calls to the GetResource method.
		GetResource []struct {
			// ResourceName is the resourceName argument value.
			ResourceName string
		}
		// ResourceExists holds details about calls to the GetResource method.
		ResourceExists []struct {
			// ResourceName is the resourceName argument value.
			ResourceName string
		}
	}
	lockClose          sync.RWMutex
	lockGetResource    sync.RWMutex
	lockResourceExists sync.RWMutex
}

// Close calls CloseFunc.
func (mock *ImportPackageMock) Close() error {
	if mock.CloseFunc == nil {
		panic("ImportPackageMock.CloseFunc: method is nil but ImportPackage.Close was just called")
	}
	callInfo := struct {
	}{}
	mock.lockClose.Lock()
	mock.calls.Close = append(mock.calls.Close, callInfo)
	mock.lockClose.Unlock()
	return mock.CloseFunc()
}

// CloseCalls gets all the calls that were made to Close.
// Check the length with:
//     len(mockedImportPackage.CloseCalls())
func (mock *ImportPackageMock) CloseCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockClose.RLock()
	calls = mock.calls.Close
	mock.lockClose.RUnlock()
	return calls
}

// GetResource calls GetResourceFunc.
func (mock *ImportPackageMock) GetResource(resourceName string) (io.ReadCloser, error) {
	if mock.GetResourceFunc == nil {
		panic("ImportPackageMock.GetResourceFunc: method is nil but ImportPackage.GetResource was just called")
	}
	callInfo := struct {
		ResourceName string
	}{
		ResourceName: resourceName,
	}
	mock.lockGetResource.Lock()
	mock.calls.GetResource = append(mock.calls.GetResource, callInfo)
	mock.lockGetResource.Unlock()
	return mock.GetResourceFunc(resourceName)
}

// ResourceExists calls ResourceExistsFunc.
func (mock *ImportPackageMock) ResourceExists(resourceName string) (bool, error) {
	if mock.ResourceExistsFunc == nil {
		panic("ImportPackageMock.ResourceExistsFunc: method is nil but ImportPackage.ResourceExists was just called")
	}
	callInfo := struct {
		ResourceName string
	}{
		ResourceName: resourceName,
	}
	mock.lockResourceExists.Lock()
	mock.calls.ResourceExists = append(mock.calls.GetResource, callInfo)
	mock.lockResourceExists.Unlock()
	return mock.ResourceExistsFunc(resourceName)
}

// GetResourceCalls gets all the calls that were made to GetResource.
// Check the length with:
//     len(mockedImportPackage.GetResourceCalls())
func (mock *ImportPackageMock) GetResourceCalls() []struct {
	ResourceName string
} {
	var calls []struct {
		ResourceName string
	}
	mock.lockGetResource.RLock()
	calls = mock.calls.GetResource
	mock.lockGetResource.RUnlock()
	return calls
}

// ManifestParserMock is a mock implementation of importer.ManifestParser.
//
// 	func TestSomethingThatUsesManifestParser(t *testing.T) {
//
// 		// make and configure a mocked importer.ManifestParser
// 		mockedManifestParser := &ManifestParserMock{
// 			ParseFunc: func(input io.Reader) (*model.ImportManifest, error) {
// 				panic("mock out the Parse method")
// 			},
// 		}
//
// 		// use mockedManifestParser in code that requires importer.ManifestParser
// 		// and then make assertions.
//
// 	}
type ManifestParserMock struct {
	// ParseFunc mocks the Parse method.
	ParseFunc func(input io.Reader) (*model.ImportManifest, error)

	// calls tracks calls to the methods.
	calls struct {
		// Parse holds details about calls to the Parse method.
		Parse []struct {
			// Input is the input argument value.
			Input io.Reader
		}
	}
	lockParse sync.RWMutex
}

// Parse calls ParseFunc.
func (mock *ManifestParserMock) Parse(input io.Reader) (*model.ImportManifest, error) {
	if mock.ParseFunc == nil {
		panic("ManifestParserMock.ParseFunc: method is nil but ManifestParser.Parse was just called")
	}
	callInfo := struct {
		Input io.Reader
	}{
		Input: input,
	}
	mock.lockParse.Lock()
	mock.calls.Parse = append(mock.calls.Parse, callInfo)
	mock.lockParse.Unlock()
	return mock.ParseFunc(input)
}

// ParseCalls gets all the calls that were made to Parse.
// Check the length with:
//     len(mockedManifestParser.ParseCalls())
func (mock *ManifestParserMock) ParseCalls() []struct {
	Input io.Reader
} {
	var calls []struct {
		Input io.Reader
	}
	mock.lockParse.RLock()
	calls = mock.calls.Parse
	mock.lockParse.RUnlock()
	return calls
}

// TaskExecutorMock is a mock implementation of importer.TaskExecutor.
//
// 	func TestSomethingThatUsesTaskExecutor(t *testing.T) {
//
// 		// make and configure a mocked importer.TaskExecutor
// 		mockedTaskExecutor := &TaskExecutorMock{
// 			ExecuteAPIFunc: func(ate model.APITaskExecution) (any, error) {
// 				panic("mock out the ExecuteAPI method")
// 			},
// 			PushResourceFunc: func(rp model.ResourcePush) (any, error) {
// 				panic("mock out the PushResource method")
// 			},
// 		}
//
// 		// use mockedTaskExecutor in code that requires importer.TaskExecutor
// 		// and then make assertions.
//
// 	}
type TaskExecutorMock struct {
	// ExecuteAPIFunc mocks the ExecuteAPI method.
	ExecuteAPIFunc func(ate model.APITaskExecution) (any, error)

	// PushResourceFunc mocks the PushResource method.
	PushResourceFunc func(rp model.ResourcePush) (any, error)

	// ActionSupportedFunc mocks the ActionSupported method.
	ActionSupportedFunc func(actionName string) bool

	// calls tracks calls to the methods.
	calls struct {
		// ExecuteAPI holds details about calls to the ExecuteAPI method.
		ExecuteAPI []struct {
			// Ate is the ate argument value.
			Ate model.APITaskExecution
		}
		// PushResource holds details about calls to the PushResource method.
		PushResource []struct {
			// Rp is the rp argument value.
			Rp model.ResourcePush
		}
		// ActionSupported holds details about calls to the ActionSupported method.
		ActionSupported []struct {
			// ActionName is the action argument value.
			ActionName string
		}
	}
	lockExecuteAPI   sync.RWMutex
	lockPushResource sync.RWMutex
}

// ExecuteAPI calls ExecuteAPIFunc.
func (mock *TaskExecutorMock) ExecuteAPI(ate model.APITaskExecution) (any, error) {
	if mock.ExecuteAPIFunc == nil {
		panic("TaskExecutorMock.ExecuteAPIFunc: method is nil but TaskExecutor.ExecuteAPI was just called")
	}
	callInfo := struct {
		Ate model.APITaskExecution
	}{
		Ate: ate,
	}
	mock.lockExecuteAPI.Lock()
	mock.calls.ExecuteAPI = append(mock.calls.ExecuteAPI, callInfo)
	mock.lockExecuteAPI.Unlock()
	return mock.ExecuteAPIFunc(ate)
}

// ExecuteAPICalls gets all the calls that were made to ExecuteAPI.
// Check the length with:
//     len(mockedTaskExecutor.ExecuteAPICalls())
func (mock *TaskExecutorMock) ExecuteAPICalls() []struct {
	Ate model.APITaskExecution
} {
	var calls []struct {
		Ate model.APITaskExecution
	}
	mock.lockExecuteAPI.RLock()
	calls = mock.calls.ExecuteAPI
	mock.lockExecuteAPI.RUnlock()
	return calls
}

// PushResource calls PushResourceFunc.
func (mock *TaskExecutorMock) PushResource(rp model.ResourcePush) (any, error) {
	if mock.PushResourceFunc == nil {
		panic("TaskExecutorMock.PushResourceFunc: method is nil but TaskExecutor.PushResource was just called")
	}
	callInfo := struct {
		Rp model.ResourcePush
	}{
		Rp: rp,
	}
	mock.lockPushResource.Lock()
	mock.calls.PushResource = append(mock.calls.PushResource, callInfo)
	mock.lockPushResource.Unlock()
	return mock.PushResourceFunc(rp)
}

// PushResourceCalls gets all the calls that were made to PushResource.
// Check the length with:
//     len(mockedTaskExecutor.PushResourceCalls())
func (mock *TaskExecutorMock) PushResourceCalls() []struct {
	Rp model.ResourcePush
} {
	var calls []struct {
		Rp model.ResourcePush
	}
	mock.lockPushResource.RLock()
	calls = mock.calls.PushResource
	mock.lockPushResource.RUnlock()
	return calls
}

// PushResource calls PushResourceFunc.
func (mock *TaskExecutorMock) ActionSupported(actionName string) bool {
	if mock.ActionSupportedFunc == nil {
		panic("TaskExecutorMock.ActionSupportedFunc: method is nil but TaskExecutor.ActionSupported was just called")
	}
	callInfo := struct {
		ActionName string
	}{
		ActionName: actionName,
	}
	mock.lockPushResource.Lock()
	mock.calls.ActionSupported = append(mock.calls.ActionSupported, callInfo)
	mock.lockPushResource.Unlock()
	return mock.ActionSupportedFunc(actionName)
}

// PushResourceCalls gets all the calls that were made to PushResource.
// Check the length with:
//     len(mockedTaskExecutor.PushResourceCalls())
func (mock *TaskExecutorMock) ActionSupportedCalls() []struct {
	ActionName string
} {
	var calls []struct {
		ActionName string
	}
	mock.lockPushResource.RLock()
	calls = mock.calls.ActionSupported
	mock.lockPushResource.RUnlock()
	return calls
}
