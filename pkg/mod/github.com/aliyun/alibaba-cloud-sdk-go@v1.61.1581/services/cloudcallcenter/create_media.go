package cloudcallcenter

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

// CreateMedia invokes the cloudcallcenter.CreateMedia API synchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/createmedia.html
func (client *Client) CreateMedia(request *CreateMediaRequest) (response *CreateMediaResponse, err error) {
	response = CreateCreateMediaResponse()
	err = client.DoAction(request, response)
	return
}

// CreateMediaWithChan invokes the cloudcallcenter.CreateMedia API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/createmedia.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateMediaWithChan(request *CreateMediaRequest) (<-chan *CreateMediaResponse, <-chan error) {
	responseChan := make(chan *CreateMediaResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateMedia(request)
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

// CreateMediaWithCallback invokes the cloudcallcenter.CreateMedia API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/createmedia.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateMediaWithCallback(request *CreateMediaRequest, callback func(response *CreateMediaResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateMediaResponse
		var err error
		defer close(result)
		response, err = client.CreateMedia(request)
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

// CreateMediaRequest is the request struct for api CreateMedia
type CreateMediaRequest struct {
	*requests.RpcRequest
	Description  string `position:"Query" name:"Description"`
	OssFilePath  string `position:"Query" name:"OssFilePath"`
	UploadResult string `position:"Query" name:"UploadResult"`
	Type         string `position:"Query" name:"Type"`
	Content      string `position:"Query" name:"Content"`
	OssFileName  string `position:"Query" name:"OssFileName"`
	InstanceId   string `position:"Query" name:"InstanceId"`
	FileName     string `position:"Query" name:"FileName"`
	Name         string `position:"Query" name:"Name"`
}

// CreateMediaResponse is the response struct for api CreateMedia
type CreateMediaResponse struct {
	*responses.BaseResponse
	RequestId        string           `json:"RequestId" xml:"RequestId"`
	Success          bool             `json:"Success" xml:"Success"`
	Code             string           `json:"Code" xml:"Code"`
	Message          string           `json:"Message" xml:"Message"`
	HttpStatusCode   int              `json:"HttpStatusCode" xml:"HttpStatusCode"`
	MediaUploadParam MediaUploadParam `json:"MediaUploadParam" xml:"MediaUploadParam"`
}

// CreateCreateMediaRequest creates a request to invoke CreateMedia API
func CreateCreateMediaRequest() (request *CreateMediaRequest) {
	request = &CreateMediaRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CloudCallCenter", "2017-07-05", "CreateMedia", "", "")
	request.Method = requests.POST
	return
}

// CreateCreateMediaResponse creates a response to parse from CreateMedia response
func CreateCreateMediaResponse() (response *CreateMediaResponse) {
	response = &CreateMediaResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
