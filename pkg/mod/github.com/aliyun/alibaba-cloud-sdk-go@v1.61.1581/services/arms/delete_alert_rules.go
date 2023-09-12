package arms

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

// DeleteAlertRules invokes the arms.DeleteAlertRules API synchronously
func (client *Client) DeleteAlertRules(request *DeleteAlertRulesRequest) (response *DeleteAlertRulesResponse, err error) {
	response = CreateDeleteAlertRulesResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteAlertRulesWithChan invokes the arms.DeleteAlertRules API asynchronously
func (client *Client) DeleteAlertRulesWithChan(request *DeleteAlertRulesRequest) (<-chan *DeleteAlertRulesResponse, <-chan error) {
	responseChan := make(chan *DeleteAlertRulesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteAlertRules(request)
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

// DeleteAlertRulesWithCallback invokes the arms.DeleteAlertRules API asynchronously
func (client *Client) DeleteAlertRulesWithCallback(request *DeleteAlertRulesRequest, callback func(response *DeleteAlertRulesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteAlertRulesResponse
		var err error
		defer close(result)
		response, err = client.DeleteAlertRules(request)
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

// DeleteAlertRulesRequest is the request struct for api DeleteAlertRules
type DeleteAlertRulesRequest struct {
	*requests.RpcRequest
	AlertIds    string `position:"Query" name:"AlertIds"`
	ProxyUserId string `position:"Query" name:"ProxyUserId"`
}

// DeleteAlertRulesResponse is the response struct for api DeleteAlertRules
type DeleteAlertRulesResponse struct {
	*responses.BaseResponse
	DeleteAlertRulesIsSuccess bool   `json:"IsSuccess" xml:"IsSuccess"`
	RequestId                 string `json:"RequestId" xml:"RequestId"`
}

// CreateDeleteAlertRulesRequest creates a request to invoke DeleteAlertRules API
func CreateDeleteAlertRulesRequest() (request *DeleteAlertRulesRequest) {
	request = &DeleteAlertRulesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ARMS", "2019-08-08", "DeleteAlertRules", "", "")
	request.Method = requests.POST
	return
}

// CreateDeleteAlertRulesResponse creates a response to parse from DeleteAlertRules response
func CreateDeleteAlertRulesResponse() (response *DeleteAlertRulesResponse) {
	response = &DeleteAlertRulesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
