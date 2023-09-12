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

// Video is a nested struct in vod response
type Video struct {
	AuditManualStatus    string                      `json:"AuditManualStatus" xml:"AuditManualStatus"`
	MediaSource          string                      `json:"MediaSource" xml:"MediaSource"`
	ModifyTime           string                      `json:"ModifyTime" xml:"ModifyTime"`
	Title                string                      `json:"Title" xml:"Title"`
	CreateTime           string                      `json:"CreateTime" xml:"CreateTime"`
	TemplateGroupId      string                      `json:"TemplateGroupId" xml:"TemplateGroupId"`
	MediaType            string                      `json:"MediaType" xml:"MediaType"`
	AuditAbnormalModules string                      `json:"AuditAbnormalModules" xml:"AuditAbnormalModules"`
	CustomMediaInfo      string                      `json:"CustomMediaInfo" xml:"CustomMediaInfo"`
	AuditAIStatus        string                      `json:"AuditAIStatus" xml:"AuditAIStatus"`
	RegionId             string                      `json:"RegionId" xml:"RegionId"`
	Duration             float64                     `json:"Duration" xml:"Duration"`
	CateName             string                      `json:"CateName" xml:"CateName"`
	Size                 int64                       `json:"Size" xml:"Size"`
	Description          string                      `json:"Description" xml:"Description"`
	CateId               int64                       `json:"CateId" xml:"CateId"`
	Tags                 string                      `json:"Tags" xml:"Tags"`
	AuditLabel           string                      `json:"AuditLabel" xml:"AuditLabel"`
	PreprocessStatus     string                      `json:"PreprocessStatus" xml:"PreprocessStatus"`
	TranscodeMode        string                      `json:"TranscodeMode" xml:"TranscodeMode"`
	ModificationTime     string                      `json:"ModificationTime" xml:"ModificationTime"`
	StorageLocation      string                      `json:"StorageLocation" xml:"StorageLocation"`
	CreationTime         string                      `json:"CreationTime" xml:"CreationTime"`
	AuditAIResult        string                      `json:"AuditAIResult" xml:"AuditAIResult"`
	CoverURL             string                      `json:"CoverURL" xml:"CoverURL"`
	AppId                string                      `json:"AppId" xml:"AppId"`
	AuditTemplateId      string                      `json:"AuditTemplateId" xml:"AuditTemplateId"`
	Status               string                      `json:"Status" xml:"Status"`
	AuditStatus          string                      `json:"AuditStatus" xml:"AuditStatus"`
	DownloadSwitch       string                      `json:"DownloadSwitch" xml:"DownloadSwitch"`
	VideoId              string                      `json:"VideoId" xml:"VideoId"`
	SpriteSnapshots      []string                    `json:"SpriteSnapshots" xml:"SpriteSnapshots"`
	Snapshots            []string                    `json:"Snapshots" xml:"Snapshots"`
	PlayInfoList         []PlayInfoInSearchMedia     `json:"PlayInfoList" xml:"PlayInfoList"`
	ThumbnailList        ThumbnailListInGetVideoInfo `json:"ThumbnailList" xml:"ThumbnailList"`
}