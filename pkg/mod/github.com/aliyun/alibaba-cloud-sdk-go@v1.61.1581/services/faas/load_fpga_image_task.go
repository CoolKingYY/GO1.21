package faas

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

// LoadFpgaImageTask invokes the faas.LoadFpgaImageTask API synchronously
// api document: https://help.aliyun.com/api/faas/loadfpgaimagetask.html
func (client *Client) LoadFpgaImageTask(request *LoadFpgaImageTaskRequest) (response *LoadFpgaImageTaskResponse, err error) {
	response = CreateLoadFpgaImageTaskResponse()
	err = client.DoAction(request, response)
	return
}

// LoadFpgaImageTaskWithChan invokes the faas.LoadFpgaImageTask API asynchronously
// api document: https://help.aliyun.com/api/faas/loadfpgaimagetask.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) LoadFpgaImageTaskWithChan(request *LoadFpgaImageTaskRequest) (<-chan *LoadFpgaImageTaskResponse, <-chan error) {
	responseChan := make(chan *LoadFpgaImageTaskResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.LoadFpgaImageTask(request)
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

// LoadFpgaImageTaskWithCallback invokes the faas.LoadFpgaImageTask API asynchronously
// api document: https://help.aliyun.com/api/faas/loadfpgaimagetask.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) LoadFpgaImageTaskWithCallback(request *LoadFpgaImageTaskRequest, callback func(response *LoadFpgaImageTaskResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *LoadFpgaImageTaskResponse
		var err error
		defer close(result)
		response, err = client.LoadFpgaImageTask(request)
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

// LoadFpgaImageTaskRequest is the request struct for api LoadFpgaImageTask
type LoadFpgaImageTaskRequest struct {
	*requests.RpcRequest
	InstanceId        string `position:"Query" name:"InstanceId"`
	Bdf               string `position:"Query" name:"Bdf"`
	FpgaImageUniqueId string `position:"Query" name:"FpgaImageUniqueId"`
}

// LoadFpgaImageTaskResponse is the response struct for api LoadFpgaImageTask
type LoadFpgaImageTaskResponse struct {
	*responses.BaseResponse
	RequestId         string `json:"RequestId" xml:"RequestId"`
	BDFInECS          string `json:"BDFInECS" xml:"BDFInECS"`
	FpgaUniqueId      string `json:"FpgaUniqueId" xml:"FpgaUniqueId"`
	FpgaImageUniqueId string `json:"FpgaImageUniqueId" xml:"FpgaImageUniqueId"`
	InstanceId        string `json:"InstanceId" xml:"InstanceId"`
	TaskId            string `json:"TaskId" xml:"TaskId"`
	TaskStatus        string `json:"TaskStatus" xml:"TaskStatus"`
}

// CreateLoadFpgaImageTaskRequest creates a request to invoke LoadFpgaImageTask API
func CreateLoadFpgaImageTaskRequest() (request *LoadFpgaImageTaskRequest) {
	request = &LoadFpgaImageTaskRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("faas", "2020-02-17", "LoadFpgaImageTask", "faas", "openAPI")
	request.Method = requests.POST
	return
}

// CreateLoadFpgaImageTaskResponse creates a response to parse from LoadFpgaImageTask response
func CreateLoadFpgaImageTaskResponse() (response *LoadFpgaImageTaskResponse) {
	response = &LoadFpgaImageTaskResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}