package live

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

// DescribeLiveDomainStagingConfig invokes the live.DescribeLiveDomainStagingConfig API synchronously
func (client *Client) DescribeLiveDomainStagingConfig(request *DescribeLiveDomainStagingConfigRequest) (response *DescribeLiveDomainStagingConfigResponse, err error) {
	response = CreateDescribeLiveDomainStagingConfigResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeLiveDomainStagingConfigWithChan invokes the live.DescribeLiveDomainStagingConfig API asynchronously
func (client *Client) DescribeLiveDomainStagingConfigWithChan(request *DescribeLiveDomainStagingConfigRequest) (<-chan *DescribeLiveDomainStagingConfigResponse, <-chan error) {
	responseChan := make(chan *DescribeLiveDomainStagingConfigResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeLiveDomainStagingConfig(request)
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

// DescribeLiveDomainStagingConfigWithCallback invokes the live.DescribeLiveDomainStagingConfig API asynchronously
func (client *Client) DescribeLiveDomainStagingConfigWithCallback(request *DescribeLiveDomainStagingConfigRequest, callback func(response *DescribeLiveDomainStagingConfigResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeLiveDomainStagingConfigResponse
		var err error
		defer close(result)
		response, err = client.DescribeLiveDomainStagingConfig(request)
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

// DescribeLiveDomainStagingConfigRequest is the request struct for api DescribeLiveDomainStagingConfig
type DescribeLiveDomainStagingConfigRequest struct {
	*requests.RpcRequest
	FunctionNames string           `position:"Query" name:"FunctionNames"`
	DomainName    string           `position:"Query" name:"DomainName"`
	OwnerId       requests.Integer `position:"Query" name:"OwnerId"`
}

// DescribeLiveDomainStagingConfigResponse is the response struct for api DescribeLiveDomainStagingConfig
type DescribeLiveDomainStagingConfigResponse struct {
	*responses.BaseResponse
	RequestId     string         `json:"RequestId" xml:"RequestId"`
	DomainConfigs []DomainConfig `json:"DomainConfigs" xml:"DomainConfigs"`
}

// CreateDescribeLiveDomainStagingConfigRequest creates a request to invoke DescribeLiveDomainStagingConfig API
func CreateDescribeLiveDomainStagingConfigRequest() (request *DescribeLiveDomainStagingConfigRequest) {
	request = &DescribeLiveDomainStagingConfigRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("live", "2016-11-01", "DescribeLiveDomainStagingConfig", "live", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeLiveDomainStagingConfigResponse creates a response to parse from DescribeLiveDomainStagingConfig response
func CreateDescribeLiveDomainStagingConfigResponse() (response *DescribeLiveDomainStagingConfigResponse) {
	response = &DescribeLiveDomainStagingConfigResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
