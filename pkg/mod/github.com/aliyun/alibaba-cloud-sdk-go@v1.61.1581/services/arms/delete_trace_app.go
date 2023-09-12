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

// DeleteTraceApp invokes the arms.DeleteTraceApp API synchronously
func (client *Client) DeleteTraceApp(request *DeleteTraceAppRequest) (response *DeleteTraceAppResponse, err error) {
	response = CreateDeleteTraceAppResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteTraceAppWithChan invokes the arms.DeleteTraceApp API asynchronously
func (client *Client) DeleteTraceAppWithChan(request *DeleteTraceAppRequest) (<-chan *DeleteTraceAppResponse, <-chan error) {
	responseChan := make(chan *DeleteTraceAppResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteTraceApp(request)
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

// DeleteTraceAppWithCallback invokes the arms.DeleteTraceApp API asynchronously
func (client *Client) DeleteTraceAppWithCallback(request *DeleteTraceAppRequest, callback func(response *DeleteTraceAppResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteTraceAppResponse
		var err error
		defer close(result)
		response, err = client.DeleteTraceApp(request)
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

// DeleteTraceAppRequest is the request struct for api DeleteTraceApp
type DeleteTraceAppRequest struct {
	*requests.RpcRequest
	AppId string `position:"Query" name:"AppId"`
	Pid   string `position:"Query" name:"Pid"`
	Type  string `position:"Query" name:"Type"`
}

// DeleteTraceAppResponse is the response struct for api DeleteTraceApp
type DeleteTraceAppResponse struct {
	*responses.BaseResponse
	Data      string `json:"Data" xml:"Data"`
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateDeleteTraceAppRequest creates a request to invoke DeleteTraceApp API
func CreateDeleteTraceAppRequest() (request *DeleteTraceAppRequest) {
	request = &DeleteTraceAppRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ARMS", "2019-08-08", "DeleteTraceApp", "", "")
	request.Method = requests.POST
	return
}

// CreateDeleteTraceAppResponse creates a response to parse from DeleteTraceApp response
func CreateDeleteTraceAppResponse() (response *DeleteTraceAppResponse) {
	response = &DeleteTraceAppResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
