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

// AddStudioLayout invokes the live.AddStudioLayout API synchronously
func (client *Client) AddStudioLayout(request *AddStudioLayoutRequest) (response *AddStudioLayoutResponse, err error) {
	response = CreateAddStudioLayoutResponse()
	err = client.DoAction(request, response)
	return
}

// AddStudioLayoutWithChan invokes the live.AddStudioLayout API asynchronously
func (client *Client) AddStudioLayoutWithChan(request *AddStudioLayoutRequest) (<-chan *AddStudioLayoutResponse, <-chan error) {
	responseChan := make(chan *AddStudioLayoutResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.AddStudioLayout(request)
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

// AddStudioLayoutWithCallback invokes the live.AddStudioLayout API asynchronously
func (client *Client) AddStudioLayoutWithCallback(request *AddStudioLayoutRequest, callback func(response *AddStudioLayoutResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *AddStudioLayoutResponse
		var err error
		defer close(result)
		response, err = client.AddStudioLayout(request)
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

// AddStudioLayoutRequest is the request struct for api AddStudioLayout
type AddStudioLayoutRequest struct {
	*requests.RpcRequest
	ScreenInputConfigList string           `position:"Query" name:"ScreenInputConfigList"`
	LayoutType            string           `position:"Query" name:"LayoutType"`
	LayoutName            string           `position:"Query" name:"LayoutName"`
	LayerOrderConfigList  string           `position:"Query" name:"LayerOrderConfigList"`
	MediaInputConfigList  string           `position:"Query" name:"MediaInputConfigList"`
	CasterId              string           `position:"Query" name:"CasterId"`
	BgImageConfig         string           `position:"Query" name:"BgImageConfig"`
	OwnerId               requests.Integer `position:"Query" name:"OwnerId"`
	CommonConfig          string           `position:"Query" name:"CommonConfig"`
}

// AddStudioLayoutResponse is the response struct for api AddStudioLayout
type AddStudioLayoutResponse struct {
	*responses.BaseResponse
	LayoutId  string `json:"LayoutId" xml:"LayoutId"`
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateAddStudioLayoutRequest creates a request to invoke AddStudioLayout API
func CreateAddStudioLayoutRequest() (request *AddStudioLayoutRequest) {
	request = &AddStudioLayoutRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("live", "2016-11-01", "AddStudioLayout", "live", "openAPI")
	request.Method = requests.POST
	return
}

// CreateAddStudioLayoutResponse creates a response to parse from AddStudioLayout response
func CreateAddStudioLayoutResponse() (response *AddStudioLayoutResponse) {
	response = &AddStudioLayoutResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
