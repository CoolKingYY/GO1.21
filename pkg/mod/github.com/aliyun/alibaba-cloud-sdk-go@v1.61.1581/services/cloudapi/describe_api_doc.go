package cloudapi

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

// DescribeApiDoc invokes the cloudapi.DescribeApiDoc API synchronously
// api document: https://help.aliyun.com/api/cloudapi/describeapidoc.html
func (client *Client) DescribeApiDoc(request *DescribeApiDocRequest) (response *DescribeApiDocResponse, err error) {
	response = CreateDescribeApiDocResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeApiDocWithChan invokes the cloudapi.DescribeApiDoc API asynchronously
// api document: https://help.aliyun.com/api/cloudapi/describeapidoc.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeApiDocWithChan(request *DescribeApiDocRequest) (<-chan *DescribeApiDocResponse, <-chan error) {
	responseChan := make(chan *DescribeApiDocResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeApiDoc(request)
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

// DescribeApiDocWithCallback invokes the cloudapi.DescribeApiDoc API asynchronously
// api document: https://help.aliyun.com/api/cloudapi/describeapidoc.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeApiDocWithCallback(request *DescribeApiDocRequest, callback func(response *DescribeApiDocResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeApiDocResponse
		var err error
		defer close(result)
		response, err = client.DescribeApiDoc(request)
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

// DescribeApiDocRequest is the request struct for api DescribeApiDoc
type DescribeApiDocRequest struct {
	*requests.RpcRequest
	StageName     string `position:"Query" name:"StageName"`
	GroupId       string `position:"Query" name:"GroupId"`
	SecurityToken string `position:"Query" name:"SecurityToken"`
	ApiId         string `position:"Query" name:"ApiId"`
}

// DescribeApiDocResponse is the response struct for api DescribeApiDoc
type DescribeApiDocResponse struct {
	*responses.BaseResponse
	RequestId          string                             `json:"RequestId" xml:"RequestId"`
	RegionId           string                             `json:"RegionId" xml:"RegionId"`
	GroupId            string                             `json:"GroupId" xml:"GroupId"`
	GroupName          string                             `json:"GroupName" xml:"GroupName"`
	StageName          string                             `json:"StageName" xml:"StageName"`
	ApiId              string                             `json:"ApiId" xml:"ApiId"`
	ApiName            string                             `json:"ApiName" xml:"ApiName"`
	Description        string                             `json:"Description" xml:"Description"`
	Visibility         string                             `json:"Visibility" xml:"Visibility"`
	AuthType           string                             `json:"AuthType" xml:"AuthType"`
	ResultType         string                             `json:"ResultType" xml:"ResultType"`
	ResultSample       string                             `json:"ResultSample" xml:"ResultSample"`
	FailResultSample   string                             `json:"FailResultSample" xml:"FailResultSample"`
	DeployedTime       string                             `json:"DeployedTime" xml:"DeployedTime"`
	ForceNonceCheck    bool                               `json:"ForceNonceCheck" xml:"ForceNonceCheck"`
	DisableInternet    bool                               `json:"DisableInternet" xml:"DisableInternet"`
	RequestConfig      RequestConfig                      `json:"RequestConfig" xml:"RequestConfig"`
	ErrorCodeSamples   ErrorCodeSamplesInDescribeApiDoc   `json:"ErrorCodeSamples" xml:"ErrorCodeSamples"`
	ResultDescriptions ResultDescriptionsInDescribeApiDoc `json:"ResultDescriptions" xml:"ResultDescriptions"`
	RequestParameters  RequestParametersInDescribeApiDoc  `json:"RequestParameters" xml:"RequestParameters"`
}

// CreateDescribeApiDocRequest creates a request to invoke DescribeApiDoc API
func CreateDescribeApiDocRequest() (request *DescribeApiDocRequest) {
	request = &DescribeApiDocRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CloudAPI", "2016-07-14", "DescribeApiDoc", "apigateway", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeApiDocResponse creates a response to parse from DescribeApiDoc response
func CreateDescribeApiDocResponse() (response *DescribeApiDocResponse) {
	response = &DescribeApiDocResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
