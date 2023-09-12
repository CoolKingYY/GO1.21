package dataworks_public

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

// ResumeInstance invokes the dataworks_public.ResumeInstance API synchronously
func (client *Client) ResumeInstance(request *ResumeInstanceRequest) (response *ResumeInstanceResponse, err error) {
	response = CreateResumeInstanceResponse()
	err = client.DoAction(request, response)
	return
}

// ResumeInstanceWithChan invokes the dataworks_public.ResumeInstance API asynchronously
func (client *Client) ResumeInstanceWithChan(request *ResumeInstanceRequest) (<-chan *ResumeInstanceResponse, <-chan error) {
	responseChan := make(chan *ResumeInstanceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ResumeInstance(request)
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

// ResumeInstanceWithCallback invokes the dataworks_public.ResumeInstance API asynchronously
func (client *Client) ResumeInstanceWithCallback(request *ResumeInstanceRequest, callback func(response *ResumeInstanceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ResumeInstanceResponse
		var err error
		defer close(result)
		response, err = client.ResumeInstance(request)
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

// ResumeInstanceRequest is the request struct for api ResumeInstance
type ResumeInstanceRequest struct {
	*requests.RpcRequest
	ProjectEnv string           `position:"Body" name:"ProjectEnv"`
	InstanceId requests.Integer `position:"Body" name:"InstanceId"`
}

// ResumeInstanceResponse is the response struct for api ResumeInstance
type ResumeInstanceResponse struct {
	*responses.BaseResponse
	ErrorCode      string `json:"ErrorCode" xml:"ErrorCode"`
	ErrorMessage   string `json:"ErrorMessage" xml:"ErrorMessage"`
	HttpStatusCode int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
	RequestId      string `json:"RequestId" xml:"RequestId"`
	Success        bool   `json:"Success" xml:"Success"`
	Data           bool   `json:"Data" xml:"Data"`
}

// CreateResumeInstanceRequest creates a request to invoke ResumeInstance API
func CreateResumeInstanceRequest() (request *ResumeInstanceRequest) {
	request = &ResumeInstanceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dataworks-public", "2020-05-18", "ResumeInstance", "", "")
	request.Method = requests.POST
	return
}

// CreateResumeInstanceResponse creates a response to parse from ResumeInstance response
func CreateResumeInstanceResponse() (response *ResumeInstanceResponse) {
	response = &ResumeInstanceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
