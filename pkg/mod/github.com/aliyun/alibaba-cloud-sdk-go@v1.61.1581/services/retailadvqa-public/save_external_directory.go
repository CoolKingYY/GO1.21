package retailadvqa_public

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// SaveExternalDirectory invokes the retailadvqa_public.SaveExternalDirectory API synchronously
func (client *Client) SaveExternalDirectory(request *SaveExternalDirectoryRequest) (response *SaveExternalDirectoryResponse, err error) {
	response = CreateSaveExternalDirectoryResponse()
	err = client.DoAction(request, response)
	return
}

// SaveExternalDirectoryWithChan invokes the retailadvqa_public.SaveExternalDirectory API asynchronously
func (client *Client) SaveExternalDirectoryWithChan(request *SaveExternalDirectoryRequest) (<-chan *SaveExternalDirectoryResponse, <-chan error) {
	responseChan := make(chan *SaveExternalDirectoryResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SaveExternalDirectory(request)
		if err != nil {
			errChan <- err
		} else {
			responseChan <- response
		}
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

// SaveExternalDirectoryWithCallback invokes the retailadvqa_public.SaveExternalDirectory API asynchronously
func (client *Client) SaveExternalDirectoryWithCallback(request *SaveExternalDirectoryRequest, callback func(response *SaveExternalDirectoryResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SaveExternalDirectoryResponse
		var err error
		defer close(result)
		response, err = client.SaveExternalDirectory(request)
		callback(response, err)
		result <- 1
	})
	if err != nil {
		defer close(result)
		callback(nil, err)
		result <- 0
	}
	return result
}

// SaveExternalDirectoryRequest is the request struct for api SaveExternalDirectory
type SaveExternalDirectoryRequest struct {
	*requests.RpcRequest
	AccessId          string           `position:"Query" name:"AccessId"`
	ParentDirectoryId string           `position:"Query" name:"ParentDirectoryId"`
	DirectoryId       string           `position:"Query" name:"DirectoryId"`
	DirectoryName     string           `position:"Query" name:"DirectoryName"`
	Type              requests.Integer `position:"Query" name:"Type"`
	WorkspaceId       string           `position:"Query" name:"WorkspaceId"`
}

// SaveExternalDirectoryResponse is the response struct for api SaveExternalDirectory
type SaveExternalDirectoryResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	ErrorDesc string `json:"ErrorDesc" xml:"ErrorDesc"`
	TraceId   string `json:"TraceId" xml:"TraceId"`
	Data      string `json:"Data" xml:"Data"`
	ErrorCode string `json:"ErrorCode" xml:"ErrorCode"`
	Success   bool   `json:"Success" xml:"Success"`
}

// CreateSaveExternalDirectoryRequest creates a request to invoke SaveExternalDirectory API
func CreateSaveExternalDirectoryRequest() (request *SaveExternalDirectoryRequest) {
	request = &SaveExternalDirectoryRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("retailadvqa-public", "2020-05-15", "SaveExternalDirectory", "", "")
	request.Method = requests.POST
	return
}

// CreateSaveExternalDirectoryResponse creates a response to parse from SaveExternalDirectory response
func CreateSaveExternalDirectoryResponse() (response *SaveExternalDirectoryResponse) {
	response = &SaveExternalDirectoryResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}