package mse

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

// CreateNacosService invokes the mse.CreateNacosService API synchronously
func (client *Client) CreateNacosService(request *CreateNacosServiceRequest) (response *CreateNacosServiceResponse, err error) {
	response = CreateCreateNacosServiceResponse()
	err = client.DoAction(request, response)
	return
}

// CreateNacosServiceWithChan invokes the mse.CreateNacosService API asynchronously
func (client *Client) CreateNacosServiceWithChan(request *CreateNacosServiceRequest) (<-chan *CreateNacosServiceResponse, <-chan error) {
	responseChan := make(chan *CreateNacosServiceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateNacosService(request)
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

// CreateNacosServiceWithCallback invokes the mse.CreateNacosService API asynchronously
func (client *Client) CreateNacosServiceWithCallback(request *CreateNacosServiceRequest, callback func(response *CreateNacosServiceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateNacosServiceResponse
		var err error
		defer close(result)
		response, err = client.CreateNacosService(request)
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

// CreateNacosServiceRequest is the request struct for api CreateNacosService
type CreateNacosServiceRequest struct {
	*requests.RpcRequest
	Ephemeral        requests.Boolean `position:"Query" name:"Ephemeral"`
	ClusterId        string           `position:"Query" name:"ClusterId"`
	GroupName        string           `position:"Query" name:"GroupName"`
	InstanceId       string           `position:"Query" name:"InstanceId"`
	NamespaceId      string           `position:"Query" name:"NamespaceId"`
	AcceptLanguage   string           `position:"Query" name:"AcceptLanguage"`
	ServiceName      string           `position:"Query" name:"ServiceName"`
	ProtectThreshold string           `position:"Query" name:"ProtectThreshold"`
}

// CreateNacosServiceResponse is the response struct for api CreateNacosService
type CreateNacosServiceResponse struct {
	*responses.BaseResponse
	RequestId      string `json:"RequestId" xml:"RequestId"`
	HttpStatusCode int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
	Data           string `json:"Data" xml:"Data"`
	Code           int    `json:"Code" xml:"Code"`
	Success        bool   `json:"Success" xml:"Success"`
	Message        string `json:"Message" xml:"Message"`
}

// CreateCreateNacosServiceRequest creates a request to invoke CreateNacosService API
func CreateCreateNacosServiceRequest() (request *CreateNacosServiceRequest) {
	request = &CreateNacosServiceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("mse", "2019-05-31", "CreateNacosService", "mse", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateNacosServiceResponse creates a response to parse from CreateNacosService response
func CreateCreateNacosServiceResponse() (response *CreateNacosServiceResponse) {
	response = &CreateNacosServiceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
