package companyreg

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

// GetBalanceCash invokes the companyreg.GetBalanceCash API synchronously
func (client *Client) GetBalanceCash(request *GetBalanceCashRequest) (response *GetBalanceCashResponse, err error) {
	response = CreateGetBalanceCashResponse()
	err = client.DoAction(request, response)
	return
}

// GetBalanceCashWithChan invokes the companyreg.GetBalanceCash API asynchronously
func (client *Client) GetBalanceCashWithChan(request *GetBalanceCashRequest) (<-chan *GetBalanceCashResponse, <-chan error) {
	responseChan := make(chan *GetBalanceCashResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetBalanceCash(request)
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

// GetBalanceCashWithCallback invokes the companyreg.GetBalanceCash API asynchronously
func (client *Client) GetBalanceCashWithCallback(request *GetBalanceCashRequest, callback func(response *GetBalanceCashResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetBalanceCashResponse
		var err error
		defer close(result)
		response, err = client.GetBalanceCash(request)
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

// GetBalanceCashRequest is the request struct for api GetBalanceCash
type GetBalanceCashRequest struct {
	*requests.RpcRequest
	Period string `position:"Query" name:"Period"`
	BizId  string `position:"Query" name:"BizId"`
}

// GetBalanceCashResponse is the response struct for api GetBalanceCash
type GetBalanceCashResponse struct {
	*responses.BaseResponse
	Name      string `json:"Name" xml:"Name"`
	Balance   string `json:"Balance" xml:"Balance"`
	Code      string `json:"Code" xml:"Code"`
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateGetBalanceCashRequest creates a request to invoke GetBalanceCash API
func CreateGetBalanceCashRequest() (request *GetBalanceCashRequest) {
	request = &GetBalanceCashRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("companyreg", "2020-10-22", "GetBalanceCash", "", "")
	request.Method = requests.GET
	return
}

// CreateGetBalanceCashResponse creates a response to parse from GetBalanceCash response
func CreateGetBalanceCashResponse() (response *GetBalanceCashResponse) {
	response = &GetBalanceCashResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
