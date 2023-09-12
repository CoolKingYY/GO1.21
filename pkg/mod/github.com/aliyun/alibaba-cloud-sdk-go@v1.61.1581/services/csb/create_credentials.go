package csb

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

// CreateCredentials invokes the csb.CreateCredentials API synchronously
// api document: https://help.aliyun.com/api/csb/createcredentials.html
func (client *Client) CreateCredentials(request *CreateCredentialsRequest) (response *CreateCredentialsResponse, err error) {
	response = CreateCreateCredentialsResponse()
	err = client.DoAction(request, response)
	return
}

// CreateCredentialsWithChan invokes the csb.CreateCredentials API asynchronously
// api document: https://help.aliyun.com/api/csb/createcredentials.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateCredentialsWithChan(request *CreateCredentialsRequest) (<-chan *CreateCredentialsResponse, <-chan error) {
	responseChan := make(chan *CreateCredentialsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateCredentials(request)
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

// CreateCredentialsWithCallback invokes the csb.CreateCredentials API asynchronously
// api document: https://help.aliyun.com/api/csb/createcredentials.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateCredentialsWithCallback(request *CreateCredentialsRequest, callback func(response *CreateCredentialsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateCredentialsResponse
		var err error
		defer close(result)
		response, err = client.CreateCredentials(request)
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

// CreateCredentialsRequest is the request struct for api CreateCredentials
type CreateCredentialsRequest struct {
	*requests.RpcRequest
	Data  string           `position:"Body" name:"Data"`
	CsbId requests.Integer `position:"Query" name:"CsbId"`
}

// CreateCredentialsResponse is the response struct for api CreateCredentials
type CreateCredentialsResponse struct {
	*responses.BaseResponse
	Code      int    `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateCreateCredentialsRequest creates a request to invoke CreateCredentials API
func CreateCreateCredentialsRequest() (request *CreateCredentialsRequest) {
	request = &CreateCredentialsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CSB", "2017-11-18", "CreateCredentials", "", "")
	request.Method = requests.POST
	return
}

// CreateCreateCredentialsResponse creates a response to parse from CreateCredentials response
func CreateCreateCredentialsResponse() (response *CreateCredentialsResponse) {
	response = &CreateCredentialsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
