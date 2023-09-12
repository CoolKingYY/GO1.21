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

// Application is a nested struct in companyreg response
type Application struct {
	Type                     int     `json:"Type" xml:"Type"`
	OrderPrice               float64 `json:"OrderPrice" xml:"OrderPrice"`
	UpdateTime               int64   `json:"UpdateTime" xml:"UpdateTime"`
	Domain                   string  `json:"Domain" xml:"Domain"`
	ProduceVersion           string  `json:"ProduceVersion" xml:"ProduceVersion"`
	ActionType               string  `json:"ActionType" xml:"ActionType"`
	UserId                   string  `json:"UserId" xml:"UserId"`
	ApplicationType          int     `json:"ApplicationType" xml:"ApplicationType"`
	BizId                    string  `json:"BizId" xml:"BizId"`
	PartnerCode              string  `json:"PartnerCode" xml:"PartnerCode"`
	IncludeForeignInvestment bool    `json:"IncludeForeignInvestment" xml:"IncludeForeignInvestment"`
	IntentionBizId           string  `json:"IntentionBizId" xml:"IntentionBizId"`
	CompanyAddress           string  `json:"CompanyAddress" xml:"CompanyAddress"`
	LegalPersonName          string  `json:"LegalPersonName" xml:"LegalPersonName"`
	Version                  string  `json:"Version" xml:"Version"`
	ApplicationStatus        int     `json:"ApplicationStatus" xml:"ApplicationStatus"`
	CompanyArea              string  `json:"CompanyArea" xml:"CompanyArea"`
	OrderId                  string  `json:"OrderId" xml:"OrderId"`
	CompanyName              string  `json:"CompanyName" xml:"CompanyName"`
}
