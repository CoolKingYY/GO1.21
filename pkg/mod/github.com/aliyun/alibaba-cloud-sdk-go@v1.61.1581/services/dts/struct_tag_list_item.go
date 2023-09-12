package dts

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

// TagListItem is a nested struct in dts response
type TagListItem struct {
	ResourceId   string `json:"ResourceId" xml:"ResourceId"`
	SrcRegion    string `json:"SrcRegion" xml:"SrcRegion"`
	GmtModified  string `json:"GmtModified" xml:"GmtModified"`
	TagCategory  string `json:"TagCategory" xml:"TagCategory"`
	Id           int64  `json:"Id" xml:"Id"`
	Creator      int64  `json:"Creator" xml:"Creator"`
	GmtCreate    string `json:"GmtCreate" xml:"GmtCreate"`
	ResourceType string `json:"ResourceType" xml:"ResourceType"`
	RegionId     string `json:"RegionId" xml:"RegionId"`
	AliUid       int64  `json:"AliUid" xml:"AliUid"`
	TagValue     string `json:"TagValue" xml:"TagValue"`
	Scope        string `json:"Scope" xml:"Scope"`
	TagKey       string `json:"TagKey" xml:"TagKey"`
}
