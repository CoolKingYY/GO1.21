package mopen

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

// Data is a nested struct in mopen response
type Data struct {
	DeviceName   string   `json:"DeviceName" xml:"DeviceName"`
	Product      string   `json:"Product" xml:"Product"`
	Result       string   `json:"Result" xml:"Result"`
	DeviceSecret string   `json:"DeviceSecret" xml:"DeviceSecret"`
	CanvasId     int      `json:"CanvasId" xml:"CanvasId"`
	Creator      string   `json:"Creator" xml:"Creator"`
	GroupId      string   `json:"GroupId" xml:"GroupId"`
	ResultType   string   `json:"ResultType" xml:"ResultType"`
	CanvasList   []Canvas `json:"CanvasList" xml:"CanvasList"`
}
