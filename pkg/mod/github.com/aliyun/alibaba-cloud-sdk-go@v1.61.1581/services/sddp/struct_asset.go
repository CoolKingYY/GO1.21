package sddp

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

// Asset is a nested struct in sddp response
type Asset struct {
	Acl               string `json:"Acl" xml:"Acl"`
	CreationTime      int64  `json:"CreationTime" xml:"CreationTime"`
	DataType          string `json:"DataType" xml:"DataType"`
	Owner             string `json:"Owner" xml:"Owner"`
	SensitiveRatio    string `json:"SensitiveRatio" xml:"SensitiveRatio"`
	Protection        bool   `json:"Protection" xml:"Protection"`
	DepartName        string `json:"DepartName" xml:"DepartName"`
	Labelsec          bool   `json:"Labelsec" xml:"Labelsec"`
	TotalCount        int    `json:"TotalCount" xml:"TotalCount"`
	RiskLevelId       int64  `json:"RiskLevelId" xml:"RiskLevelId"`
	RuleName          string `json:"RuleName" xml:"RuleName"`
	Sensitive         bool   `json:"Sensitive" xml:"Sensitive"`
	ObjectKey         string `json:"ObjectKey" xml:"ObjectKey"`
	RiskLevelName     string `json:"RiskLevelName" xml:"RiskLevelName"`
	OdpsRiskLevelName string `json:"OdpsRiskLevelName" xml:"OdpsRiskLevelName"`
	ProductId         string `json:"ProductId" xml:"ProductId"`
	Name              string `json:"Name" xml:"Name"`
	SensitiveCount    int    `json:"SensitiveCount" xml:"SensitiveCount"`
	Id                string `json:"Id" xml:"Id"`
	ProductCode       string `json:"ProductCode" xml:"ProductCode"`
}