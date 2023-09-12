package imm

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

// ImagesItem is a nested struct in imm response
type ImagesItem struct {
	AddressModifyTime            string                   `json:"AddressModifyTime" xml:"AddressModifyTime"`
	FacesStatus                  string                   `json:"FacesStatus" xml:"FacesStatus"`
	ModifyTime                   string                   `json:"ModifyTime" xml:"ModifyTime"`
	AddressFailReason            string                   `json:"AddressFailReason" xml:"AddressFailReason"`
	CroppingSuggestionFailReason string                   `json:"CroppingSuggestionFailReason" xml:"CroppingSuggestionFailReason"`
	CreateTime                   string                   `json:"CreateTime" xml:"CreateTime"`
	OCRModifyTime                string                   `json:"OCRModifyTime" xml:"OCRModifyTime"`
	ImageFormat                  string                   `json:"ImageFormat" xml:"ImageFormat"`
	FacesFailReason              string                   `json:"FacesFailReason" xml:"FacesFailReason"`
	AddressStatus                string                   `json:"AddressStatus" xml:"AddressStatus"`
	CroppingSuggestionModifyTime string                   `json:"CroppingSuggestionModifyTime" xml:"CroppingSuggestionModifyTime"`
	Exif                         string                   `json:"Exif" xml:"Exif"`
	CroppingSuggestionStatus     string                   `json:"CroppingSuggestionStatus" xml:"CroppingSuggestionStatus"`
	ImageHeight                  int                      `json:"ImageHeight" xml:"ImageHeight"`
	SourceType                   string                   `json:"SourceType" xml:"SourceType"`
	RemarksB                     string                   `json:"RemarksB" xml:"RemarksB"`
	FileSize                     int                      `json:"FileSize" xml:"FileSize"`
	RemarksArrayA                string                   `json:"RemarksArrayA" xml:"RemarksArrayA"`
	Orientation                  string                   `json:"Orientation" xml:"Orientation"`
	TagsStatus                   string                   `json:"TagsStatus" xml:"TagsStatus"`
	FacesModifyTime              string                   `json:"FacesModifyTime" xml:"FacesModifyTime"`
	RemarksD                     string                   `json:"RemarksD" xml:"RemarksD"`
	TagsFailReason               string                   `json:"TagsFailReason" xml:"TagsFailReason"`
	OCRStatus                    string                   `json:"OCRStatus" xml:"OCRStatus"`
	ImageUri                     string                   `json:"ImageUri" xml:"ImageUri"`
	ExternalId                   string                   `json:"ExternalId" xml:"ExternalId"`
	ImageTime                    string                   `json:"ImageTime" xml:"ImageTime"`
	Location                     string                   `json:"Location" xml:"Location"`
	ImageQualityStatus           string                   `json:"ImageQualityStatus" xml:"ImageQualityStatus"`
	TagsModifyTime               string                   `json:"TagsModifyTime" xml:"TagsModifyTime"`
	RemarksArrayB                string                   `json:"RemarksArrayB" xml:"RemarksArrayB"`
	RemarksA                     string                   `json:"RemarksA" xml:"RemarksA"`
	ImageQualityModifyTime       string                   `json:"ImageQualityModifyTime" xml:"ImageQualityModifyTime"`
	ImageQualityFailReason       string                   `json:"ImageQualityFailReason" xml:"ImageQualityFailReason"`
	SourcePosition               string                   `json:"SourcePosition" xml:"SourcePosition"`
	OCRFailReason                string                   `json:"OCRFailReason" xml:"OCRFailReason"`
	RemarksC                     string                   `json:"RemarksC" xml:"RemarksC"`
	SourceUri                    string                   `json:"SourceUri" xml:"SourceUri"`
	ImageWidth                   int                      `json:"ImageWidth" xml:"ImageWidth"`
	ImageQuality                 ImageQuality             `json:"ImageQuality" xml:"ImageQuality"`
	Address                      Address                  `json:"Address" xml:"Address"`
	CroppingSuggestion           []CroppingSuggestionItem `json:"CroppingSuggestion" xml:"CroppingSuggestion"`
	Tags                         []TagsItem               `json:"Tags" xml:"Tags"`
	OCR                          []OCRItem                `json:"OCR" xml:"OCR"`
	Faces                        []FacesItem              `json:"Faces" xml:"Faces"`
}
