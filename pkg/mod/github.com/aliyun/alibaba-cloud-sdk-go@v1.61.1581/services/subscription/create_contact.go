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

// CreateContact invokes the subscription.CreateContact API synchronously
// api document: https://help.aliyun.com/api/subscription/createcontact.html
func (client *Client) CreateContact(request *CreateContactRequest) (response *CreateContactResponse, err error) {
	response = CreateCreateContactResponse()
	err = client.DoAction(request, response)
	return
}

// CreateContactWithChan invokes the subscription.CreateContact API asynchronously
// api document: https://help.aliyun.com/api/subscription/createcontact.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateContactWithChan(request *CreateContactRequest) (<-chan *CreateContactResponse, <-chan error) {
	responseChan := make(chan *CreateContactResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateContact(request)
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

// CreateContactWithCallback invokes the subscription.CreateContact API asynchronously
// api document: https://help.aliyun.com/api/subscription/createcontact.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateContactWithCallback(request *CreateContactRequest, callback func(response *CreateContactResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateContactResponse
		var err error
		defer close(result)
		response, err = client.CreateContact(request)
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

// CreateContactRequest is the request struct for api CreateContact
type CreateContactRequest struct {
	*requests.RpcRequest
	ClientToken string `position:"Query" name:"ClientToken"`
	Mobile      string `position:"Body" name:"Mobile"`
	Locale      string `position:"Query" name:"Locale"`
	Name        string `position:"Body" name:"Name"`
	Position    string `position:"Body" name:"Position"`
	Email       string `position:"Body" name:"Email"`
}

// CreateContactResponse is the response struct for api CreateContact
type CreateContactResponse struct {
	*responses.BaseResponse
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	ContactId int64  `json:"ContactId" xml:"ContactId"`
	Code      string `json:"Code" xml:"Code"`
	Success   bool   `json:"Success" xml:"Success"`
}

// CreateCreateContactRequest creates a request to invoke CreateContact API
func CreateCreateContactRequest() (request *CreateContactRequest) {
	request = &CreateContactRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Subscription", "2021-01-15", "CreateContact", "", "")
	return
}

// CreateCreateContactResponse creates a response to parse from CreateContact response
func CreateCreateContactResponse() (response *CreateContactResponse) {
	response = &CreateContactResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
