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

// UpdateWebhook invokes the arms.UpdateWebhook API synchronously
func (client *Client) UpdateWebhook(request *UpdateWebhookRequest) (response *UpdateWebhookResponse, err error) {
	response = CreateUpdateWebhookResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateWebhookWithChan invokes the arms.UpdateWebhook API asynchronously
func (client *Client) UpdateWebhookWithChan(request *UpdateWebhookRequest) (<-chan *UpdateWebhookResponse, <-chan error) {
	responseChan := make(chan *UpdateWebhookResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateWebhook(request)
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

// UpdateWebhookWithCallback invokes the arms.UpdateWebhook API asynchronously
func (client *Client) UpdateWebhookWithCallback(request *UpdateWebhookRequest, callback func(response *UpdateWebhookResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateWebhookResponse
		var err error
		defer close(result)
		response, err = client.UpdateWebhook(request)
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

// UpdateWebhookRequest is the request struct for api UpdateWebhook
type UpdateWebhookRequest struct {
	*requests.RpcRequest
	HttpHeaders string           `position:"Query" name:"HttpHeaders"`
	Method      string           `position:"Query" name:"Method"`
	ContactId   requests.Integer `position:"Query" name:"ContactId"`
	HttpParams  string           `position:"Query" name:"HttpParams"`
	ProxyUserId string           `position:"Query" name:"ProxyUserId"`
	Body        string           `position:"Query" name:"Body"`
	Url         string           `position:"Query" name:"Url"`
	ContactName string           `position:"Query" name:"ContactName"`
	RecoverBody string           `position:"Query" name:"RecoverBody"`
}

// UpdateWebhookResponse is the response struct for api UpdateWebhook
type UpdateWebhookResponse struct {
	*responses.BaseResponse
	UpdateWebhookIsSuccess bool   `json:"IsSuccess" xml:"IsSuccess"`
	RequestId              string `json:"RequestId" xml:"RequestId"`
}

// CreateUpdateWebhookRequest creates a request to invoke UpdateWebhook API
func CreateUpdateWebhookRequest() (request *UpdateWebhookRequest) {
	request = &UpdateWebhookRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ARMS", "2019-08-08", "UpdateWebhook", "", "")
	request.Method = requests.POST
	return
}

// CreateUpdateWebhookResponse creates a response to parse from UpdateWebhook response
func CreateUpdateWebhookResponse() (response *UpdateWebhookResponse) {
	response = &UpdateWebhookResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
