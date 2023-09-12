package sgw

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

// DescribeGatewayLDAPInfo invokes the sgw.DescribeGatewayLDAPInfo API synchronously
func (client *Client) DescribeGatewayLDAPInfo(request *DescribeGatewayLDAPInfoRequest) (response *DescribeGatewayLDAPInfoResponse, err error) {
	response = CreateDescribeGatewayLDAPInfoResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeGatewayLDAPInfoWithChan invokes the sgw.DescribeGatewayLDAPInfo API asynchronously
func (client *Client) DescribeGatewayLDAPInfoWithChan(request *DescribeGatewayLDAPInfoRequest) (<-chan *DescribeGatewayLDAPInfoResponse, <-chan error) {
	responseChan := make(chan *DescribeGatewayLDAPInfoResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeGatewayLDAPInfo(request)
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

// DescribeGatewayLDAPInfoWithCallback invokes the sgw.DescribeGatewayLDAPInfo API asynchronously
func (client *Client) DescribeGatewayLDAPInfoWithCallback(request *DescribeGatewayLDAPInfoRequest, callback func(response *DescribeGatewayLDAPInfoResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeGatewayLDAPInfoResponse
		var err error
		defer close(result)
		response, err = client.DescribeGatewayLDAPInfo(request)
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

// DescribeGatewayLDAPInfoRequest is the request struct for api DescribeGatewayLDAPInfo
type DescribeGatewayLDAPInfoRequest struct {
	*requests.RpcRequest
	SecurityToken string `position:"Query" name:"SecurityToken"`
	GatewayId     string `position:"Query" name:"GatewayId"`
}

// DescribeGatewayLDAPInfoResponse is the response struct for api DescribeGatewayLDAPInfo
type DescribeGatewayLDAPInfoResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	BaseDN    string `json:"BaseDN" xml:"BaseDN"`
	RootDN    string `json:"RootDN" xml:"RootDN"`
	ServerIp  string `json:"ServerIp" xml:"ServerIp"`
	IsTls     bool   `json:"IsTls" xml:"IsTls"`
	IsEnabled bool   `json:"IsEnabled" xml:"IsEnabled"`
}

// CreateDescribeGatewayLDAPInfoRequest creates a request to invoke DescribeGatewayLDAPInfo API
func CreateDescribeGatewayLDAPInfoRequest() (request *DescribeGatewayLDAPInfoRequest) {
	request = &DescribeGatewayLDAPInfoRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("sgw", "2018-05-11", "DescribeGatewayLDAPInfo", "hcs_sgw", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeGatewayLDAPInfoResponse creates a response to parse from DescribeGatewayLDAPInfo response
func CreateDescribeGatewayLDAPInfoResponse() (response *DescribeGatewayLDAPInfoResponse) {
	response = &DescribeGatewayLDAPInfoResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
