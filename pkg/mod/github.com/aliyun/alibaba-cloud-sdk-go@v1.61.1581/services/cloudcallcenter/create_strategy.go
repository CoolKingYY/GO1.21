package cloudcallcenter

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

// CreateStrategy invokes the cloudcallcenter.CreateStrategy API synchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/createstrategy.html
func (client *Client) CreateStrategy(request *CreateStrategyRequest) (response *CreateStrategyResponse, err error) {
	response = CreateCreateStrategyResponse()
	err = client.DoAction(request, response)
	return
}

// CreateStrategyWithChan invokes the cloudcallcenter.CreateStrategy API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/createstrategy.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateStrategyWithChan(request *CreateStrategyRequest) (<-chan *CreateStrategyResponse, <-chan error) {
	responseChan := make(chan *CreateStrategyResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateStrategy(request)
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

// CreateStrategyWithCallback invokes the cloudcallcenter.CreateStrategy API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/createstrategy.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateStrategyWithCallback(request *CreateStrategyRequest, callback func(response *CreateStrategyResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateStrategyResponse
		var err error
		defer close(result)
		response, err = client.CreateStrategy(request)
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

// CreateStrategyRequest is the request struct for api CreateStrategy
type CreateStrategyRequest struct {
	*requests.RpcRequest
	InstanceId   string `position:"Query" name:"InstanceId"`
	StrategyJson string `position:"Query" name:"StrategyJson"`
}

// CreateStrategyResponse is the response struct for api CreateStrategy
type CreateStrategyResponse struct {
	*responses.BaseResponse
	RequestId      string   `json:"RequestId" xml:"RequestId"`
	Success        bool     `json:"Success" xml:"Success"`
	Code           string   `json:"Code" xml:"Code"`
	Message        string   `json:"Message" xml:"Message"`
	HttpStatusCode int      `json:"HttpStatusCode" xml:"HttpStatusCode"`
	Strategy       Strategy `json:"Strategy" xml:"Strategy"`
}

// CreateCreateStrategyRequest creates a request to invoke CreateStrategy API
func CreateCreateStrategyRequest() (request *CreateStrategyRequest) {
	request = &CreateStrategyRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CloudCallCenter", "2017-07-05", "CreateStrategy", "", "")
	request.Method = requests.POST
	return
}

// CreateCreateStrategyResponse creates a response to parse from CreateStrategy response
func CreateCreateStrategyResponse() (response *CreateStrategyResponse) {
	response = &CreateStrategyResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
