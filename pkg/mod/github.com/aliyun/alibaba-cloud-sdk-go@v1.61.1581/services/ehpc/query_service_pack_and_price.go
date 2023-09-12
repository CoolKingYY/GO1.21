package ehpc

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

// QueryServicePackAndPrice invokes the ehpc.QueryServicePackAndPrice API synchronously
func (client *Client) QueryServicePackAndPrice(request *QueryServicePackAndPriceRequest) (response *QueryServicePackAndPriceResponse, err error) {
	response = CreateQueryServicePackAndPriceResponse()
	err = client.DoAction(request, response)
	return
}

// QueryServicePackAndPriceWithChan invokes the ehpc.QueryServicePackAndPrice API asynchronously
func (client *Client) QueryServicePackAndPriceWithChan(request *QueryServicePackAndPriceRequest) (<-chan *QueryServicePackAndPriceResponse, <-chan error) {
	responseChan := make(chan *QueryServicePackAndPriceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryServicePackAndPrice(request)
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

// QueryServicePackAndPriceWithCallback invokes the ehpc.QueryServicePackAndPrice API asynchronously
func (client *Client) QueryServicePackAndPriceWithCallback(request *QueryServicePackAndPriceRequest, callback func(response *QueryServicePackAndPriceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryServicePackAndPriceResponse
		var err error
		defer close(result)
		response, err = client.QueryServicePackAndPrice(request)
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

// QueryServicePackAndPriceRequest is the request struct for api QueryServicePackAndPrice
type QueryServicePackAndPriceRequest struct {
	*requests.RpcRequest
}

// QueryServicePackAndPriceResponse is the response struct for api QueryServicePackAndPrice
type QueryServicePackAndPriceResponse struct {
	*responses.BaseResponse
	RequestId      string      `json:"RequestId" xml:"RequestId"`
	RegionId       string      `json:"RegionId" xml:"RegionId"`
	TradePrice     float64     `json:"TradePrice" xml:"TradePrice"`
	OriginalPrice  float64     `json:"OriginalPrice" xml:"OriginalPrice"`
	DiscountPrice  float64     `json:"DiscountPrice" xml:"DiscountPrice"`
	Currency       string      `json:"Currency" xml:"Currency"`
	OriginalAmount int         `json:"OriginalAmount" xml:"OriginalAmount"`
	ChargeAmount   int         `json:"ChargeAmount" xml:"ChargeAmount"`
	ServicePack    ServicePack `json:"ServicePack" xml:"ServicePack"`
}

// CreateQueryServicePackAndPriceRequest creates a request to invoke QueryServicePackAndPrice API
func CreateQueryServicePackAndPriceRequest() (request *QueryServicePackAndPriceRequest) {
	request = &QueryServicePackAndPriceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("EHPC", "2018-04-12", "QueryServicePackAndPrice", "", "")
	request.Method = requests.GET
	return
}

// CreateQueryServicePackAndPriceResponse creates a response to parse from QueryServicePackAndPrice response
func CreateQueryServicePackAndPriceResponse() (response *QueryServicePackAndPriceResponse) {
	response = &QueryServicePackAndPriceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
