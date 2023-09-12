package subscription

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

// UpdateSubscriptionItem invokes the subscription.UpdateSubscriptionItem API synchronously
// api document: https://help.aliyun.com/api/subscription/updatesubscriptionitem.html
func (client *Client) UpdateSubscriptionItem(request *UpdateSubscriptionItemRequest) (response *UpdateSubscriptionItemResponse, err error) {
	response = CreateUpdateSubscriptionItemResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateSubscriptionItemWithChan invokes the subscription.UpdateSubscriptionItem API asynchronously
// api document: https://help.aliyun.com/api/subscription/updatesubscriptionitem.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) UpdateSubscriptionItemWithChan(request *UpdateSubscriptionItemRequest) (<-chan *UpdateSubscriptionItemResponse, <-chan error) {
	responseChan := make(chan *UpdateSubscriptionItemResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateSubscriptionItem(request)
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

// UpdateSubscriptionItemWithCallback invokes the subscription.UpdateSubscriptionItem API asynchronously
// api document: https://help.aliyun.com/api/subscription/updatesubscriptionitem.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) UpdateSubscriptionItemWithCallback(request *UpdateSubscriptionItemRequest, callback func(response *UpdateSubscriptionItemResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateSubscriptionItemResponse
		var err error
		defer close(result)
		response, err = client.UpdateSubscriptionItem(request)
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

// UpdateSubscriptionItemRequest is the request struct for api UpdateSubscriptionItem
type UpdateSubscriptionItemRequest struct {
	*requests.RpcRequest
	ClientToken   string           `position:"Query" name:"ClientToken"`
	Locale        string           `position:"Query" name:"Locale"`
	ContactIds    string           `position:"Body" name:"ContactIds"`
	ItemId        requests.Integer `position:"Body" name:"ItemId"`
	SmsStatus     requests.Integer `position:"Body" name:"SmsStatus"`
	PmsgStatus    requests.Integer `position:"Body" name:"PmsgStatus"`
	WebhookStatus requests.Integer `position:"Body" name:"WebhookStatus"`
	TtsStatus     requests.Integer `position:"Body" name:"TtsStatus"`
	WebhookIds    string           `position:"Body" name:"WebhookIds"`
	EmailStatus   requests.Integer `position:"Body" name:"EmailStatus"`
}

// UpdateSubscriptionItemResponse is the response struct for api UpdateSubscriptionItem
type UpdateSubscriptionItemResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Message   string `json:"Message" xml:"Message"`
	Code      string `json:"Code" xml:"Code"`
	Success   bool   `json:"Success" xml:"Success"`
	Result    bool   `json:"Result" xml:"Result"`
}

// CreateUpdateSubscriptionItemRequest creates a request to invoke UpdateSubscriptionItem API
func CreateUpdateSubscriptionItemRequest() (request *UpdateSubscriptionItemRequest) {
	request = &UpdateSubscriptionItemRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Subscription", "2021-01-15", "UpdateSubscriptionItem", "", "")
	return
}

// CreateUpdateSubscriptionItemResponse creates a response to parse from UpdateSubscriptionItem response
func CreateUpdateSubscriptionItemResponse() (response *UpdateSubscriptionItemResponse) {
	response = &UpdateSubscriptionItemResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
