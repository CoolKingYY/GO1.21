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

// BookkeepingStatistics is a nested struct in companyreg response
type BookkeepingStatistics struct {
	TaxCjrbzj     float64     `json:"TaxCjrbzj" xml:"TaxCjrbzj"`
	TaxYhs        float64     `json:"TaxYhs" xml:"TaxYhs"`
	Income        float64     `json:"Income" xml:"Income"`
	TaxAmountNote string      `json:"TaxAmountNote" xml:"TaxAmountNote"`
	TaxFjsNote    string      `json:"TaxFjsNote" xml:"TaxFjsNote"`
	Year          int         `json:"Year" xml:"Year"`
	SubjectCount  int         `json:"SubjectCount" xml:"SubjectCount"`
	TaxOther      float64     `json:"TaxOther" xml:"TaxOther"`
	TaxGhjfNote   string      `json:"TaxGhjfNote" xml:"TaxGhjfNote"`
	TaxYhsNote    string      `json:"TaxYhsNote" xml:"TaxYhsNote"`
	Expend        float64     `json:"Expend" xml:"Expend"`
	Month         int         `json:"Month" xml:"Month"`
	TaxQysds      float64     `json:"TaxQysds" xml:"TaxQysds"`
	TaxSljjNote   string      `json:"TaxSljjNote" xml:"TaxSljjNote"`
	TaxGhjf       float64     `json:"TaxGhjf" xml:"TaxGhjf"`
	TaxQysdsNote  string      `json:"TaxQysdsNote" xml:"TaxQysdsNote"`
	TaxOtherNote  string      `json:"TaxOtherNote" xml:"TaxOtherNote"`
	TaxZzsNote    string      `json:"TaxZzsNote" xml:"TaxZzsNote"`
	TaxZzs        float64     `json:"TaxZzs" xml:"TaxZzs"`
	TaxAmount     float64     `json:"TaxAmount" xml:"TaxAmount"`
	TaxFjs        float64     `json:"TaxFjs" xml:"TaxFjs"`
	VoucherCount  int         `json:"VoucherCount" xml:"VoucherCount"`
	ProduceBizId  string      `json:"ProduceBizId" xml:"ProduceBizId"`
	Profit        float64     `json:"Profit" xml:"Profit"`
	TaxSljj       float64     `json:"TaxSljj" xml:"TaxSljj"`
	TaxCjrbzjNote string      `json:"TaxCjrbzjNote" xml:"TaxCjrbzjNote"`
	TaxDetails    []TaxDetail `json:"TaxDetails" xml:"TaxDetails"`
}