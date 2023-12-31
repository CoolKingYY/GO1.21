package vod

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

// LiveRecordVideo is a nested struct in vod response
type LiveRecordVideo struct {
	AppName         string                            `json:"AppName" xml:"AppName"`
	PlaylistId      string                            `json:"PlaylistId" xml:"PlaylistId"`
	StreamName      string                            `json:"StreamName" xml:"StreamName"`
	RecordEndTime   string                            `json:"RecordEndTime" xml:"RecordEndTime"`
	RecordStartTime string                            `json:"RecordStartTime" xml:"RecordStartTime"`
	DomainName      string                            `json:"DomainName" xml:"DomainName"`
	Video           VideoInListLiveRecordVideo        `json:"Video" xml:"Video"`
	PlayInfoList    PlayInfoListInListLiveRecordVideo `json:"PlayInfoList" xml:"PlayInfoList"`
}
