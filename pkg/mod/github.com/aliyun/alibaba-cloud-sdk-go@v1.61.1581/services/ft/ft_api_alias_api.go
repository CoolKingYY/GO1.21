package ft

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

// FTApiAliasApi invokes the ft.FTApiAliasApi API synchronously
func (client *Client) FTApiAliasApi(request *FTApiAliasApiRequest) (response *FTApiAliasApiResponse, err error) {
	response = CreateFTApiAliasApiResponse()
	err = client.DoAction(request, response)
	return
}

// FTApiAliasApiWithChan invokes the ft.FTApiAliasApi API asynchronously
func (client *Client) FTApiAliasApiWithChan(request *FTApiAliasApiRequest) (<-chan *FTApiAliasApiResponse, <-chan error) {
	responseChan := make(chan *FTApiAliasApiResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.FTApiAliasApi(request)
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

// FTApiAliasApiWithCallback invokes the ft.FTApiAliasApi API asynchronously
func (client *Client) FTApiAliasApiWithCallback(request *FTApiAliasApiRequest, callback func(response *FTApiAliasApiResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *FTApiAliasApiResponse
		var err error
		defer close(result)
		response, err = client.FTApiAliasApi(request)
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

// FTApiAliasApiRequest is the request struct for api FTApiAliasApi
type FTApiAliasApiRequest struct {
	*requests.RpcRequest
	Name string `position:"Query" name:"Name"`
}

// FTApiAliasApiResponse is the response struct for api FTApiAliasApi
type FTApiAliasApiResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Name      string `json:"Name" xml:"Name"`
}

// CreateFTApiAliasApiRequest creates a request to invoke FTApiAliasApi API
func CreateFTApiAliasApiRequest() (request *FTApiAliasApiRequest) {
	request = &FTApiAliasApiRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ft", "2018-07-13", "FTApiAliasApi", "", "")
	request.Method = requests.POST
	return
}

// CreateFTApiAliasApiResponse creates a response to parse from FTApiAliasApi response
func CreateFTApiAliasApiResponse() (response *FTApiAliasApiResponse) {
	response = &FTApiAliasApiResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}