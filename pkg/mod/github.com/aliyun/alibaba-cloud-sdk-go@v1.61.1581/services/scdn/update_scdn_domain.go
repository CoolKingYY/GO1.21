package scdn

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

// UpdateScdnDomain invokes the scdn.UpdateScdnDomain API synchronously
func (client *Client) UpdateScdnDomain(request *UpdateScdnDomainRequest) (response *UpdateScdnDomainResponse, err error) {
	response = CreateUpdateScdnDomainResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateScdnDomainWithChan invokes the scdn.UpdateScdnDomain API asynchronously
func (client *Client) UpdateScdnDomainWithChan(request *UpdateScdnDomainRequest) (<-chan *UpdateScdnDomainResponse, <-chan error) {
	responseChan := make(chan *UpdateScdnDomainResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateScdnDomain(request)
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

// UpdateScdnDomainWithCallback invokes the scdn.UpdateScdnDomain API asynchronously
func (client *Client) UpdateScdnDomainWithCallback(request *UpdateScdnDomainRequest, callback func(response *UpdateScdnDomainResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateScdnDomainResponse
		var err error
		defer close(result)
		response, err = client.UpdateScdnDomain(request)
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

// UpdateScdnDomainRequest is the request struct for api UpdateScdnDomain
type UpdateScdnDomainRequest struct {
	*requests.RpcRequest
	Sources         string           `position:"Query" name:"Sources"`
	ResourceGroupId string           `position:"Query" name:"ResourceGroupId"`
	SecurityToken   string           `position:"Query" name:"SecurityToken"`
	DomainName      string           `position:"Query" name:"DomainName"`
	OwnerId         requests.Integer `position:"Query" name:"OwnerId"`
}

// UpdateScdnDomainResponse is the response struct for api UpdateScdnDomain
type UpdateScdnDomainResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateUpdateScdnDomainRequest creates a request to invoke UpdateScdnDomain API
func CreateUpdateScdnDomainRequest() (request *UpdateScdnDomainRequest) {
	request = &UpdateScdnDomainRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("scdn", "2017-11-15", "UpdateScdnDomain", "", "")
	request.Method = requests.POST
	return
}

// CreateUpdateScdnDomainResponse creates a response to parse from UpdateScdnDomain response
func CreateUpdateScdnDomainResponse() (response *UpdateScdnDomainResponse) {
	response = &UpdateScdnDomainResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
