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

// AddShowIntoShowList invokes the live.AddShowIntoShowList API synchronously
func (client *Client) AddShowIntoShowList(request *AddShowIntoShowListRequest) (response *AddShowIntoShowListResponse, err error) {
	response = CreateAddShowIntoShowListResponse()
	err = client.DoAction(request, response)
	return
}

// AddShowIntoShowListWithChan invokes the live.AddShowIntoShowList API asynchronously
func (client *Client) AddShowIntoShowListWithChan(request *AddShowIntoShowListRequest) (<-chan *AddShowIntoShowListResponse, <-chan error) {
	responseChan := make(chan *AddShowIntoShowListResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.AddShowIntoShowList(request)
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

// AddShowIntoShowListWithCallback invokes the live.AddShowIntoShowList API asynchronously
func (client *Client) AddShowIntoShowListWithCallback(request *AddShowIntoShowListRequest, callback func(response *AddShowIntoShowListResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *AddShowIntoShowListResponse
		var err error
		defer close(result)
		response, err = client.AddShowIntoShowList(request)
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

// AddShowIntoShowListRequest is the request struct for api AddShowIntoShowList
type AddShowIntoShowListRequest struct {
	*requests.RpcRequest
	LiveInputType requests.Integer `position:"Query" name:"LiveInputType"`
	Duration      requests.Integer `position:"Query" name:"Duration"`
	RepeatTimes   requests.Integer `position:"Query" name:"RepeatTimes"`
	ShowName      string           `position:"Query" name:"ShowName"`
	ResourceId    string           `position:"Query" name:"ResourceId"`
	CasterId      string           `position:"Query" name:"CasterId"`
	OwnerId       requests.Integer `position:"Query" name:"OwnerId"`
	ResourceType  string           `position:"Query" name:"ResourceType"`
	ResourceUrl   string           `position:"Query" name:"ResourceUrl"`
	Spot          requests.Integer `position:"Query" name:"Spot"`
}

// AddShowIntoShowListResponse is the response struct for api AddShowIntoShowList
type AddShowIntoShowListResponse struct {
	*responses.BaseResponse
	ShowId    string `json:"ShowId" xml:"ShowId"`
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateAddShowIntoShowListRequest creates a request to invoke AddShowIntoShowList API
func CreateAddShowIntoShowListRequest() (request *AddShowIntoShowListRequest) {
	request = &AddShowIntoShowListRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("live", "2016-11-01", "AddShowIntoShowList", "live", "openAPI")
	request.Method = requests.POST
	return
}

// CreateAddShowIntoShowListResponse creates a response to parse from AddShowIntoShowList response
func CreateAddShowIntoShowListResponse() (response *AddShowIntoShowListResponse) {
	response = &AddShowIntoShowListResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}