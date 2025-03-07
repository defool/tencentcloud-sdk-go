// Copyright (c) 2017-2018 THL A29 Limited, a Tencent company. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package v20181201

import (
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2018-12-01"

type Client struct {
    common.Client
}

// Deprecated
func NewClientWithSecretId(secretId, secretKey, region string) (client *Client, err error) {
    cpf := profile.NewClientProfile()
    client = &Client{}
    client.Init(region).WithSecretId(secretId, secretKey).WithProfile(cpf)
    return
}

func NewClient(credential *common.Credential, region string, clientProfile *profile.ClientProfile) (client *Client, err error) {
    client = &Client{}
    client.Init(region).
        WithCredential(credential).
        WithProfile(clientProfile)
    return
}


func NewDescribeMaterialListRequest() (request *DescribeMaterialListRequest) {
    request = &DescribeMaterialListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("facefusion", APIVersion, "DescribeMaterialList")
    return
}

func NewDescribeMaterialListResponse() (response *DescribeMaterialListResponse) {
    response = &DescribeMaterialListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 通常通过腾讯云人脸融合的控制台可以查看到素材相关的参数数据，可以满足使用。本接口返回活动的素材数据，包括素材状态等。用于用户通过Api查看素材相关数据，方便使用。
func (c *Client) DescribeMaterialList(request *DescribeMaterialListRequest) (response *DescribeMaterialListResponse, err error) {
    if request == nil {
        request = NewDescribeMaterialListRequest()
    }
    response = NewDescribeMaterialListResponse()
    err = c.Send(request, response)
    return
}

func NewFaceFusionRequest() (request *FaceFusionRequest) {
    request = &FaceFusionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("facefusion", APIVersion, "FaceFusion")
    return
}

func NewFaceFusionResponse() (response *FaceFusionResponse) {
    response = &FaceFusionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口用于人脸融合，用户上传人脸图片，获取与模板融合后的人脸图片。未发布的活动请求频率限制为1次/秒，已发布的活动请求频率限制50次/秒。如有需要提高活动的请求频率限制，请在控制台中申请。
// >     
// - 公共参数中的签名方式必须指定为V3版本，即配置SignatureMethod参数为TC3-HMAC-SHA256。
func (c *Client) FaceFusion(request *FaceFusionRequest) (response *FaceFusionResponse, err error) {
    if request == nil {
        request = NewFaceFusionRequest()
    }
    response = NewFaceFusionResponse()
    err = c.Send(request, response)
    return
}

func NewFaceFusionLiteRequest() (request *FaceFusionLiteRequest) {
    request = &FaceFusionLiteRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("facefusion", APIVersion, "FaceFusionLite")
    return
}

func NewFaceFusionLiteResponse() (response *FaceFusionLiteResponse) {
    response = &FaceFusionLiteResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 人脸融合活动专用版，不推荐使用。人脸融合接口建议使用[人脸融合](https://cloud.tencent.com/document/product/670/31061)或[选脸融合](https://cloud.tencent.com/document/product/670/37736)接口
func (c *Client) FaceFusionLite(request *FaceFusionLiteRequest) (response *FaceFusionLiteResponse, err error) {
    if request == nil {
        request = NewFaceFusionLiteRequest()
    }
    response = NewFaceFusionLiteResponse()
    err = c.Send(request, response)
    return
}

func NewFuseFaceRequest() (request *FuseFaceRequest) {
    request = &FuseFaceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("facefusion", APIVersion, "FuseFace")
    return
}

func NewFuseFaceResponse() (response *FuseFaceResponse) {
    response = &FuseFaceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口用于单脸、多脸融合，用户上传人脸图片，获取与模板融合后的人脸图片。查看 <a href="https://cloud.tencent.com/document/product/670/38247" target="_blank">选脸融合接入指引</a>。
// 
// 未发布的活动请求频率限制为1次/秒，已发布的活动请求频率限制50次/秒。如有需要提高活动的请求频率限制，请在控制台中申请。
// >
// - 公共参数中的签名方式必须指定为V3版本，即配置SignatureMethod参数为TC3-HMAC-SHA256。
func (c *Client) FuseFace(request *FuseFaceRequest) (response *FuseFaceResponse, err error) {
    if request == nil {
        request = NewFuseFaceRequest()
    }
    response = NewFuseFaceResponse()
    err = c.Send(request, response)
    return
}
