package linkwan

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

// Packet is a nested struct in linkwan response
type Packet struct {
	DevAddr                 string                   `json:"DevAddr" xml:"DevAddr"`
	Freq                    string                   `json:"Freq" xml:"Freq"`
	ClassMode               string                   `json:"ClassMode" xml:"ClassMode"`
	HasMacCommand           bool                     `json:"HasMacCommand" xml:"HasMacCommand"`
	Datr                    string                   `json:"Datr" xml:"Datr"`
	Rssi                    int                      `json:"Rssi" xml:"Rssi"`
	GwEui                   string                   `json:"GwEui" xml:"GwEui"`
	FPort                   int                      `json:"FPort" xml:"FPort"`
	NodeOwnerAliyunId       string                   `json:"NodeOwnerAliyunId" xml:"NodeOwnerAliyunId"`
	Base64EncodedMacPayload string                   `json:"Base64EncodedMacPayload" xml:"Base64EncodedMacPayload"`
	Lsnr                    float64                  `json:"Lsnr" xml:"Lsnr"`
	HasData                 bool                     `json:"HasData" xml:"HasData"`
	MacPayloadSize          int64                    `json:"MacPayloadSize" xml:"MacPayloadSize"`
	DevEui                  string                   `json:"DevEui" xml:"DevEui"`
	LogMillis               string                   `json:"LogMillis" xml:"LogMillis"`
	ProcessEvent            string                   `json:"ProcessEvent" xml:"ProcessEvent"`
	MessageType             string                   `json:"MessageType" xml:"MessageType"`
	MacCommandCIDs          []map[string]interface{} `json:"MacCommandCIDs" xml:"MacCommandCIDs"`
}
