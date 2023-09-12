package facebody

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

// DeleteFaceImageTemplate invokes the facebody.DeleteFaceImageTemplate API synchronously
func (client *Client) DeleteFaceImageTemplate(request *DeleteFaceImageTemplateRequest) (response *DeleteFaceImageTemplateResponse, err error) {
	response = CreateDeleteFaceImageTemplateResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteFaceImageTemplateWithChan invokes the facebody.DeleteFaceImageTemplate API asynchronously
func (client *Client) DeleteFaceImageTemplateWithChan(request *DeleteFaceImageTemplateRequest) (<-chan *DeleteFaceImageTemplateResponse, <-chan error) {
	responseChan := make(chan *DeleteFaceImageTemplateResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteFaceImageTemplate(request)
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

// DeleteFaceImageTemplateWithCallback invokes the facebody.DeleteFaceImageTemplate API asynchronously
func (client *Client) DeleteFaceImageTemplateWithCallback(request *DeleteFaceImageTemplateRequest, callback func(response *DeleteFaceImageTemplateResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteFaceImageTemplateResponse
		var err error
		defer close(result)
		response, err = client.DeleteFaceImageTemplate(request)
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

// DeleteFaceImageTemplateRequest is the request struct for api DeleteFaceImageTemplate
type DeleteFaceImageTemplateRequest struct {
	*requests.RpcRequest
	FormatResultToJson requests.Boolean `position:"Query" name:"FormatResultToJson"`
	UserId             string           `position:"Body" name:"UserId"`
	OssFile            string           `position:"Query" name:"OssFile"`
	TemplateId         string           `position:"Body" name:"TemplateId"`
	RequestProxyBy     string           `position:"Query" name:"RequestProxyBy"`
}

// DeleteFaceImageTemplateResponse is the response struct for api DeleteFaceImageTemplate
type DeleteFaceImageTemplateResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
}

// CreateDeleteFaceImageTemplateRequest creates a request to invoke DeleteFaceImageTemplate API
func CreateDeleteFaceImageTemplateRequest() (request *DeleteFaceImageTemplateRequest) {
	request = &DeleteFaceImageTemplateRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("facebody", "2019-12-30", "DeleteFaceImageTemplate", "facebody", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDeleteFaceImageTemplateResponse creates a response to parse from DeleteFaceImageTemplate response
func CreateDeleteFaceImageTemplateResponse() (response *DeleteFaceImageTemplateResponse) {
	response = &DeleteFaceImageTemplateResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
