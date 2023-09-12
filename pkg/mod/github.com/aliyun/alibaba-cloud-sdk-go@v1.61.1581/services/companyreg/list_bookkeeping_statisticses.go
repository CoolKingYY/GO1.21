package companyreg

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

// ListBookkeepingStatisticses invokes the companyreg.ListBookkeepingStatisticses API synchronously
func (client *Client) ListBookkeepingStatisticses(request *ListBookkeepingStatisticsesRequest) (response *ListBookkeepingStatisticsesResponse, err error) {
	response = CreateListBookkeepingStatisticsesResponse()
	err = client.DoAction(request, response)
	return
}

// ListBookkeepingStatisticsesWithChan invokes the companyreg.ListBookkeepingStatisticses API asynchronously
func (client *Client) ListBookkeepingStatisticsesWithChan(request *ListBookkeepingStatisticsesRequest) (<-chan *ListBookkeepingStatisticsesResponse, <-chan error) {
	responseChan := make(chan *ListBookkeepingStatisticsesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListBookkeepingStatisticses(request)
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

// ListBookkeepingStatisticsesWithCallback invokes the companyreg.ListBookkeepingStatisticses API asynchronously
func (client *Client) ListBookkeepingStatisticsesWithCallback(request *ListBookkeepingStatisticsesRequest, callback func(response *ListBookkeepingStatisticsesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListBookkeepingStatisticsesResponse
		var err error
		defer close(result)
		response, err = client.ListBookkeepingStatisticses(request)
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

// ListBookkeepingStatisticsesRequest is the request struct for api ListBookkeepingStatisticses
type ListBookkeepingStatisticsesRequest struct {
	*requests.RpcRequest
	Year         requests.Integer `position:"Query" name:"Year"`
	ProduceBizId string           `position:"Query" name:"ProduceBizId"`
	PageNumber   requests.Integer `position:"Query" name:"PageNumber"`
	Month        requests.Integer `position:"Query" name:"Month"`
	PageSize     requests.Integer `position:"Query" name:"PageSize"`
}

// ListBookkeepingStatisticsesResponse is the response struct for api ListBookkeepingStatisticses
type ListBookkeepingStatisticsesResponse struct {
	*responses.BaseResponse
	TaxAmountSum            float64                 `json:"TaxAmountSum" xml:"TaxAmountSum"`
	ExpendSum               float64                 `json:"ExpendSum" xml:"ExpendSum"`
	ProfitSum               float64                 `json:"ProfitSum" xml:"ProfitSum"`
	RequestId               string                  `json:"RequestId" xml:"RequestId"`
	PageSize                int                     `json:"PageSize" xml:"PageSize"`
	PageNumber              int                     `json:"PageNumber" xml:"PageNumber"`
	IncomeSum               float64                 `json:"IncomeSum" xml:"IncomeSum"`
	TotalCount              int                     `json:"TotalCount" xml:"TotalCount"`
	BookkeepingStatisticses []BookkeepingStatistics `json:"BookkeepingStatisticses" xml:"BookkeepingStatisticses"`
}

// CreateListBookkeepingStatisticsesRequest creates a request to invoke ListBookkeepingStatisticses API
func CreateListBookkeepingStatisticsesRequest() (request *ListBookkeepingStatisticsesRequest) {
	request = &ListBookkeepingStatisticsesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("companyreg", "2019-05-08", "ListBookkeepingStatisticses", "", "")
	request.Method = requests.POST
	return
}

// CreateListBookkeepingStatisticsesResponse creates a response to parse from ListBookkeepingStatisticses response
func CreateListBookkeepingStatisticsesResponse() (response *ListBookkeepingStatisticsesResponse) {
	response = &ListBookkeepingStatisticsesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
