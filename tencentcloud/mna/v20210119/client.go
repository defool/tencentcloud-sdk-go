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

package v20210119

import (
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2021-01-19"

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


func NewCreateQosRequest() (request *CreateQosRequest) {
    request = &CreateQosRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mna", APIVersion, "CreateQos")
    return
}

func NewCreateQosResponse() (response *CreateQosResponse) {
    response = &CreateQosResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 移动网络发起Qos加速过程
func (c *Client) CreateQos(request *CreateQosRequest) (response *CreateQosResponse, err error) {
    if request == nil {
        request = NewCreateQosRequest()
    }
    response = NewCreateQosResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteQosRequest() (request *DeleteQosRequest) {
    request = &DeleteQosRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mna", APIVersion, "DeleteQos")
    return
}

func NewDeleteQosResponse() (response *DeleteQosResponse) {
    response = &DeleteQosResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 移动网络停止Qos加速过程
func (c *Client) DeleteQos(request *DeleteQosRequest) (response *DeleteQosResponse, err error) {
    if request == nil {
        request = NewDeleteQosRequest()
    }
    response = NewDeleteQosResponse()
    err = c.Send(request, response)
    return
}
