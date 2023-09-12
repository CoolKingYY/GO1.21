package green

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

// DeletePerson invokes the green.DeletePerson API synchronously
func (client *Client) DeletePerson(request *DeletePersonRequest) (response *DeletePersonResponse, err error) {
	response = CreateDeletePersonResponse()
	err = client.DoAction(request, response)
	return
}

// DeletePersonWithChan invokes the green.DeletePerson API asynchronously
func (client *Client) DeletePersonWithChan(request *DeletePersonRequest) (<-chan *DeletePersonResponse, <-chan error) {
	responseChan := make(chan *DeletePersonResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeletePerson(request)
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

// DeletePersonWithCallback invokes the green.DeletePerson API asynchronously
func (client *Client) DeletePersonWithCallback(request *DeletePersonRequest, callback func(response *DeletePersonResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeletePersonResponse
		var err error
		defer close(result)
		response, err = client.DeletePerson(request)
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

// DeletePersonRequest is the request struct for api DeletePerson
type DeletePersonRequest struct {
	*requests.RoaRequest
	ClientInfo string `position:"Query" name:"ClientInfo"`
}

// DeletePersonResponse is the response struct for api DeletePerson
type DeletePersonResponse struct {
	*responses.BaseResponse
}

// CreateDeletePersonRequest creates a request to invoke DeletePerson API
func CreateDeletePersonRequest() (request *DeletePersonRequest) {
	request = &DeletePersonRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("Green", "2018-05-09", "DeletePerson", "/green/sface/person/delete", "", "")
	request.Method = requests.POST
	return
}

// CreateDeletePersonResponse creates a response to parse from DeletePerson response
func CreateDeletePersonResponse() (response *DeletePersonResponse) {
	response = &DeletePersonResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
