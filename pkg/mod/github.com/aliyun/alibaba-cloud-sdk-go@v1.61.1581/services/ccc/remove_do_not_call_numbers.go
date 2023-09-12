package ccc

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

// RemoveDoNotCallNumbers invokes the ccc.RemoveDoNotCallNumbers API synchronously
func (client *Client) RemoveDoNotCallNumbers(request *RemoveDoNotCallNumbersRequest) (response *RemoveDoNotCallNumbersResponse, err error) {
	response = CreateRemoveDoNotCallNumbersResponse()
	err = client.DoAction(request, response)
	return
}

// RemoveDoNotCallNumbersWithChan invokes the ccc.RemoveDoNotCallNumbers API asynchronously
func (client *Client) RemoveDoNotCallNumbersWithChan(request *RemoveDoNotCallNumbersRequest) (<-chan *RemoveDoNotCallNumbersResponse, <-chan error) {
	responseChan := make(chan *RemoveDoNotCallNumbersResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.RemoveDoNotCallNumbers(request)
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

// RemoveDoNotCallNumbersWithCallback invokes the ccc.RemoveDoNotCallNumbers API asynchronously
func (client *Client) RemoveDoNotCallNumbersWithCallback(request *RemoveDoNotCallNumbersRequest, callback func(response *RemoveDoNotCallNumbersResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *RemoveDoNotCallNumbersResponse
		var err error
		defer close(result)
		response, err = client.RemoveDoNotCallNumbers(request)
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

// RemoveDoNotCallNumbersRequest is the request struct for api RemoveDoNotCallNumbers
type RemoveDoNotCallNumbersRequest struct {
	*requests.RpcRequest
	NumberList string `position:"Query" name:"NumberList"`
	InstanceId string `position:"Query" name:"InstanceId"`
}

// RemoveDoNotCallNumbersResponse is the response struct for api RemoveDoNotCallNumbers
type RemoveDoNotCallNumbersResponse struct {
	*responses.BaseResponse
	HttpStatusCode int      `json:"HttpStatusCode" xml:"HttpStatusCode"`
	Code           string   `json:"Code" xml:"Code"`
	Message        string   `json:"Message" xml:"Message"`
	Data           string   `json:"Data" xml:"Data"`
	RequestId      string   `json:"RequestId" xml:"RequestId"`
	Params         []string `json:"Params" xml:"Params"`
}

// CreateRemoveDoNotCallNumbersRequest creates a request to invoke RemoveDoNotCallNumbers API
func CreateRemoveDoNotCallNumbersRequest() (request *RemoveDoNotCallNumbersRequest) {
	request = &RemoveDoNotCallNumbersRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CCC", "2020-07-01", "RemoveDoNotCallNumbers", "", "")
	request.Method = requests.POST
	return
}

// CreateRemoveDoNotCallNumbersResponse creates a response to parse from RemoveDoNotCallNumbers response
func CreateRemoveDoNotCallNumbersResponse() (response *RemoveDoNotCallNumbersResponse) {
	response = &RemoveDoNotCallNumbersResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
