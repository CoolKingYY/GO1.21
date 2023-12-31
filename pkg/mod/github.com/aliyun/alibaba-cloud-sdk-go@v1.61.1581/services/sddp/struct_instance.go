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

// Instance is a nested struct in sddp response
type Instance struct {
	CreationTime        int64   `json:"CreationTime" xml:"CreationTime"`
	Acl                 string  `json:"Acl" xml:"Acl"`
	LastFinishTime      int64   `json:"LastFinishTime" xml:"LastFinishTime"`
	Owner               string  `json:"Owner" xml:"Owner"`
	CountDetails        string  `json:"CountDetails" xml:"CountDetails"`
	FileCountDetails    string  `json:"FileCountDetails" xml:"FileCountDetails"`
	TenantName          string  `json:"TenantName" xml:"TenantName"`
	Protection          bool    `json:"Protection" xml:"Protection"`
	DepartName          string  `json:"DepartName" xml:"DepartName"`
	Labelsec            bool    `json:"Labelsec" xml:"Labelsec"`
	RiskScore           float64 `json:"RiskScore" xml:"RiskScore"`
	RiskLevelId         int64   `json:"RiskLevelId" xml:"RiskLevelId"`
	S3Count             int     `json:"S3Count" xml:"S3Count"`
	S1Count             int     `json:"S1Count" xml:"S1Count"`
	ProductId           string  `json:"ProductId" xml:"ProductId"`
	Name                string  `json:"Name" xml:"Name"`
	S2Count             int     `json:"S2Count" xml:"S2Count"`
	EngineType          string  `json:"EngineType" xml:"EngineType"`
	TotalCount          int     `json:"TotalCount" xml:"TotalCount"`
	InstanceDescription string  `json:"InstanceDescription" xml:"InstanceDescription"`
	RuleName            string  `json:"RuleName" xml:"RuleName"`
	RegionId            string  `json:"RegionId" xml:"RegionId"`
	Sensitive           bool    `json:"Sensitive" xml:"Sensitive"`
	SensLevelName       string  `json:"SensLevelName" xml:"SensLevelName"`
	RegionName          string  `json:"RegionName" xml:"RegionName"`
	LastRiskScore       float64 `json:"LastRiskScore" xml:"LastRiskScore"`
	RiskLevelName       string  `json:"RiskLevelName" xml:"RiskLevelName"`
	OdpsRiskLevelName   string  `json:"OdpsRiskLevelName" xml:"OdpsRiskLevelName"`
	SensitiveCount      int     `json:"SensitiveCount" xml:"SensitiveCount"`
	Id                  int64   `json:"Id" xml:"Id"`
	ProductCode         string  `json:"ProductCode" xml:"ProductCode"`
}
