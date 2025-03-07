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

package v20190820

import (
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2019-08-20"

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


func NewApplyApplicationMaterialRequest() (request *ApplyApplicationMaterialRequest) {
    request = &ApplyApplicationMaterialRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cpdp", APIVersion, "ApplyApplicationMaterial")
    return
}

func NewApplyApplicationMaterialResponse() (response *ApplyApplicationMaterialResponse) {
    response = &ApplyApplicationMaterialResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 跨境-提交申报材料。申报材料的主体是付款人，需要提前调用【跨境-付款人申请】接口提交付款人信息且审核通过后调用。
func (c *Client) ApplyApplicationMaterial(request *ApplyApplicationMaterialRequest) (response *ApplyApplicationMaterialResponse, err error) {
    if request == nil {
        request = NewApplyApplicationMaterialRequest()
    }
    response = NewApplyApplicationMaterialResponse()
    err = c.Send(request, response)
    return
}

func NewApplyOutwardOrderRequest() (request *ApplyOutwardOrderRequest) {
    request = &ApplyOutwardOrderRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cpdp", APIVersion, "ApplyOutwardOrder")
    return
}

func NewApplyOutwardOrderResponse() (response *ApplyOutwardOrderResponse) {
    response = &ApplyOutwardOrderResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 跨境-汇出指令申请。通过该接口可将对接方账户中的人民币余额汇兑成外币，再汇出至指定银行账户。
func (c *Client) ApplyOutwardOrder(request *ApplyOutwardOrderRequest) (response *ApplyOutwardOrderResponse, err error) {
    if request == nil {
        request = NewApplyOutwardOrderRequest()
    }
    response = NewApplyOutwardOrderResponse()
    err = c.Send(request, response)
    return
}

func NewApplyPayerInfoRequest() (request *ApplyPayerInfoRequest) {
    request = &ApplyPayerInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cpdp", APIVersion, "ApplyPayerInfo")
    return
}

func NewApplyPayerInfoResponse() (response *ApplyPayerInfoResponse) {
    response = &ApplyPayerInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 跨境-付款人申请。通过该接口提交付款人信息并进行 kyc 审核。
func (c *Client) ApplyPayerInfo(request *ApplyPayerInfoRequest) (response *ApplyPayerInfoResponse, err error) {
    if request == nil {
        request = NewApplyPayerInfoRequest()
    }
    response = NewApplyPayerInfoResponse()
    err = c.Send(request, response)
    return
}

func NewApplyReWithdrawalRequest() (request *ApplyReWithdrawalRequest) {
    request = &ApplyReWithdrawalRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cpdp", APIVersion, "ApplyReWithdrawal")
    return
}

func NewApplyReWithdrawalResponse() (response *ApplyReWithdrawalResponse) {
    response = &ApplyReWithdrawalResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 正常结算提现失败情况下，发起重新提现的请求接口
func (c *Client) ApplyReWithdrawal(request *ApplyReWithdrawalRequest) (response *ApplyReWithdrawalResponse, err error) {
    if request == nil {
        request = NewApplyReWithdrawalRequest()
    }
    response = NewApplyReWithdrawalResponse()
    err = c.Send(request, response)
    return
}

func NewApplyTradeRequest() (request *ApplyTradeRequest) {
    request = &ApplyTradeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cpdp", APIVersion, "ApplyTrade")
    return
}

func NewApplyTradeResponse() (response *ApplyTradeResponse) {
    response = &ApplyTradeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 跨境-提交贸易材料。通过提交贸易材料接口可为对接方累计贸易额度，在额度范围内可发起汇兑汇出交易。
func (c *Client) ApplyTrade(request *ApplyTradeRequest) (response *ApplyTradeResponse, err error) {
    if request == nil {
        request = NewApplyTradeRequest()
    }
    response = NewApplyTradeResponse()
    err = c.Send(request, response)
    return
}

func NewApplyWithdrawalRequest() (request *ApplyWithdrawalRequest) {
    request = &ApplyWithdrawalRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cpdp", APIVersion, "ApplyWithdrawal")
    return
}

func NewApplyWithdrawalResponse() (response *ApplyWithdrawalResponse) {
    response = &ApplyWithdrawalResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 商户提现
func (c *Client) ApplyWithdrawal(request *ApplyWithdrawalRequest) (response *ApplyWithdrawalResponse, err error) {
    if request == nil {
        request = NewApplyWithdrawalRequest()
    }
    response = NewApplyWithdrawalResponse()
    err = c.Send(request, response)
    return
}

func NewBindAcctRequest() (request *BindAcctRequest) {
    request = &BindAcctRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cpdp", APIVersion, "BindAcct")
    return
}

func NewBindAcctResponse() (response *BindAcctResponse) {
    response = &BindAcctResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 商户绑定提现银行卡，每个商户只能绑定一张提现银行卡
func (c *Client) BindAcct(request *BindAcctRequest) (response *BindAcctResponse, err error) {
    if request == nil {
        request = NewBindAcctRequest()
    }
    response = NewBindAcctResponse()
    err = c.Send(request, response)
    return
}

func NewBindRelateAccReUnionPayRequest() (request *BindRelateAccReUnionPayRequest) {
    request = &BindRelateAccReUnionPayRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cpdp", APIVersion, "BindRelateAccReUnionPay")
    return
}

func NewBindRelateAccReUnionPayResponse() (response *BindRelateAccReUnionPayResponse) {
    response = &BindRelateAccReUnionPayResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 会员绑定提现账户-回填银联鉴权短信码。用于会员填写动态验证码后，发往银行进行验证，验证成功则完成绑定。
func (c *Client) BindRelateAccReUnionPay(request *BindRelateAccReUnionPayRequest) (response *BindRelateAccReUnionPayResponse, err error) {
    if request == nil {
        request = NewBindRelateAccReUnionPayRequest()
    }
    response = NewBindRelateAccReUnionPayResponse()
    err = c.Send(request, response)
    return
}

func NewBindRelateAcctSmallAmountRequest() (request *BindRelateAcctSmallAmountRequest) {
    request = &BindRelateAcctSmallAmountRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cpdp", APIVersion, "BindRelateAcctSmallAmount")
    return
}

func NewBindRelateAcctSmallAmountResponse() (response *BindRelateAcctSmallAmountResponse) {
    response = &BindRelateAcctSmallAmountResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 会员绑定提现账户-小额鉴权。会员申请绑定提现账户，绑定后从会员子账户中提现到绑定账户。
// 转账鉴权有两种形式：往账鉴权和来账鉴权。
// 往账鉴权：该接口发起成功后，银行会向提现账户转入小于等于0.5元的随机金额，并短信通知客户查看，客户查看后，需将收到的金额大小，在电商平台页面上回填，并通知银行。银行验证通过后，完成提现账户绑定。
// 来账鉴权：该接口发起成功后，银行以短信通知客户查看，客户查看后，需通过待绑定的账户往市场的监管账户转入短信上指定的金额。银行检索到该笔指定金额的来账是源自待绑定账户，则绑定成功。平安银行的账户，即BankType送1时，大小额行号和超级网银号都不用送。
func (c *Client) BindRelateAcctSmallAmount(request *BindRelateAcctSmallAmountRequest) (response *BindRelateAcctSmallAmountResponse, err error) {
    if request == nil {
        request = NewBindRelateAcctSmallAmountRequest()
    }
    response = NewBindRelateAcctSmallAmountResponse()
    err = c.Send(request, response)
    return
}

func NewBindRelateAcctUnionPayRequest() (request *BindRelateAcctUnionPayRequest) {
    request = &BindRelateAcctUnionPayRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cpdp", APIVersion, "BindRelateAcctUnionPay")
    return
}

func NewBindRelateAcctUnionPayResponse() (response *BindRelateAcctUnionPayResponse) {
    response = &BindRelateAcctUnionPayResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 会员绑定提现账户-银联鉴权。用于会员申请绑定提现账户，申请后银行前往银联验证卡信息：姓名、证件、卡号、银行预留手机是否相符，相符则发送给会员手机动态验证码并返回成功，不相符则返回失败。
// 平台接收到银行返回成功后，进入输入动态验证码的页面，有效期120秒，若120秒未输入，客户可点击重新发送动态验证码，这个步骤重新调用该接口即可。
// 平安银行的账户，大小额行号和超级网银号都不用送。
// 超级网银号：单笔转账金额不超过5万，不限制笔数，只用选XX银行，不用具体到支行，可实时知道对方是否收款成功。
// 大小额联行号：单笔转账可超过5万，需具体到支行，不能实时知道对方是否收款成功。金额超过5万的，在工作日的8点30-17点间才会成功。
func (c *Client) BindRelateAcctUnionPay(request *BindRelateAcctUnionPayRequest) (response *BindRelateAcctUnionPayResponse, err error) {
    if request == nil {
        request = NewBindRelateAcctUnionPayRequest()
    }
    response = NewBindRelateAcctUnionPayResponse()
    err = c.Send(request, response)
    return
}

func NewCheckAcctRequest() (request *CheckAcctRequest) {
    request = &CheckAcctRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cpdp", APIVersion, "CheckAcct")
    return
}

func NewCheckAcctResponse() (response *CheckAcctResponse) {
    response = &CheckAcctResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 商户绑定提现银行卡的验证接口
func (c *Client) CheckAcct(request *CheckAcctRequest) (response *CheckAcctResponse, err error) {
    if request == nil {
        request = NewCheckAcctRequest()
    }
    response = NewCheckAcctResponse()
    err = c.Send(request, response)
    return
}

func NewCheckAmountRequest() (request *CheckAmountRequest) {
    request = &CheckAmountRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cpdp", APIVersion, "CheckAmount")
    return
}

func NewCheckAmountResponse() (response *CheckAmountResponse) {
    response = &CheckAmountResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 验证鉴权金额。此接口可受理BindRelateAcctSmallAmount接口发起的转账金额（往账鉴权方式）的验证处理。若所回填的验证金额验证通过，则会绑定原申请中的银行账户作为提现账户。通过此接口也可以查得BindRelateAcctSmallAmount接口发起的来账鉴权方式的申请的当前状态。
func (c *Client) CheckAmount(request *CheckAmountRequest) (response *CheckAmountResponse, err error) {
    if request == nil {
        request = NewCheckAmountRequest()
    }
    response = NewCheckAmountResponse()
    err = c.Send(request, response)
    return
}

func NewCloseOrderRequest() (request *CloseOrderRequest) {
    request = &CloseOrderRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cpdp", APIVersion, "CloseOrder")
    return
}

func NewCloseOrderResponse() (response *CloseOrderResponse) {
    response = &CloseOrderResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 通过此接口关闭此前已创建的订单，关闭后，用户将无法继续付款。仅能关闭创建后未支付的订单
func (c *Client) CloseOrder(request *CloseOrderRequest) (response *CloseOrderResponse, err error) {
    if request == nil {
        request = NewCloseOrderRequest()
    }
    response = NewCloseOrderResponse()
    err = c.Send(request, response)
    return
}

func NewContractOrderRequest() (request *ContractOrderRequest) {
    request = &ContractOrderRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cpdp", APIVersion, "ContractOrder")
    return
}

func NewContractOrderResponse() (response *ContractOrderResponse) {
    response = &ContractOrderResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 应用需要先带上签约信息调用本接口生成支付订单号，并将应答的PayInfo透传给聚鑫SDK，拉起客户端（包括微信公众号/微信小程序/客户端App）支付。
func (c *Client) ContractOrder(request *ContractOrderRequest) (response *ContractOrderResponse, err error) {
    if request == nil {
        request = NewContractOrderRequest()
    }
    response = NewContractOrderResponse()
    err = c.Send(request, response)
    return
}

func NewCreateAcctRequest() (request *CreateAcctRequest) {
    request = &CreateAcctRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cpdp", APIVersion, "CreateAcct")
    return
}

func NewCreateAcctResponse() (response *CreateAcctResponse) {
    response = &CreateAcctResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 子商户入驻聚鑫平台
func (c *Client) CreateAcct(request *CreateAcctRequest) (response *CreateAcctResponse, err error) {
    if request == nil {
        request = NewCreateAcctRequest()
    }
    response = NewCreateAcctResponse()
    err = c.Send(request, response)
    return
}

func NewCreateAgentTaxPaymentInfosRequest() (request *CreateAgentTaxPaymentInfosRequest) {
    request = &CreateAgentTaxPaymentInfosRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cpdp", APIVersion, "CreateAgentTaxPaymentInfos")
    return
}

func NewCreateAgentTaxPaymentInfosResponse() (response *CreateAgentTaxPaymentInfosResponse) {
    response = &CreateAgentTaxPaymentInfosResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 直播平台-代理商完税信息录入
func (c *Client) CreateAgentTaxPaymentInfos(request *CreateAgentTaxPaymentInfosRequest) (response *CreateAgentTaxPaymentInfosResponse, err error) {
    if request == nil {
        request = NewCreateAgentTaxPaymentInfosRequest()
    }
    response = NewCreateAgentTaxPaymentInfosResponse()
    err = c.Send(request, response)
    return
}

func NewCreateCustAcctIdRequest() (request *CreateCustAcctIdRequest) {
    request = &CreateCustAcctIdRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cpdp", APIVersion, "CreateCustAcctId")
    return
}

func NewCreateCustAcctIdResponse() (response *CreateCustAcctIdResponse) {
    response = &CreateCustAcctIdResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 会员子账户开立。会员在银行注册，并开立会员子账户，交易网会员代码即会员在平台端系统的会员编号。
// 平台需保存银行返回的子账户账号，后续交易接口都会用到。会员属性字段为预留扩展字段，当前必须送默认值。
func (c *Client) CreateCustAcctId(request *CreateCustAcctIdRequest) (response *CreateCustAcctIdResponse, err error) {
    if request == nil {
        request = NewCreateCustAcctIdRequest()
    }
    response = NewCreateCustAcctIdResponse()
    err = c.Send(request, response)
    return
}

func NewCreateInvoiceRequest() (request *CreateInvoiceRequest) {
    request = &CreateInvoiceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cpdp", APIVersion, "CreateInvoice")
    return
}

func NewCreateInvoiceResponse() (response *CreateInvoiceResponse) {
    response = &CreateInvoiceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 智慧零售-发票开具
func (c *Client) CreateInvoice(request *CreateInvoiceRequest) (response *CreateInvoiceResponse, err error) {
    if request == nil {
        request = NewCreateInvoiceRequest()
    }
    response = NewCreateInvoiceResponse()
    err = c.Send(request, response)
    return
}

func NewCreateInvoiceV2Request() (request *CreateInvoiceV2Request) {
    request = &CreateInvoiceV2Request{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cpdp", APIVersion, "CreateInvoiceV2")
    return
}

func NewCreateInvoiceV2Response() (response *CreateInvoiceV2Response) {
    response = &CreateInvoiceV2Response{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 智慧零售-发票开具V2
func (c *Client) CreateInvoiceV2(request *CreateInvoiceV2Request) (response *CreateInvoiceV2Response, err error) {
    if request == nil {
        request = NewCreateInvoiceV2Request()
    }
    response = NewCreateInvoiceV2Response()
    err = c.Send(request, response)
    return
}

func NewCreateMerchantRequest() (request *CreateMerchantRequest) {
    request = &CreateMerchantRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cpdp", APIVersion, "CreateMerchant")
    return
}

func NewCreateMerchantResponse() (response *CreateMerchantResponse) {
    response = &CreateMerchantResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 智慧零售-商户注册
func (c *Client) CreateMerchant(request *CreateMerchantRequest) (response *CreateMerchantResponse, err error) {
    if request == nil {
        request = NewCreateMerchantRequest()
    }
    response = NewCreateMerchantResponse()
    err = c.Send(request, response)
    return
}

func NewCreateRedInvoiceRequest() (request *CreateRedInvoiceRequest) {
    request = &CreateRedInvoiceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cpdp", APIVersion, "CreateRedInvoice")
    return
}

func NewCreateRedInvoiceResponse() (response *CreateRedInvoiceResponse) {
    response = &CreateRedInvoiceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 智慧零售-发票红冲
func (c *Client) CreateRedInvoice(request *CreateRedInvoiceRequest) (response *CreateRedInvoiceResponse, err error) {
    if request == nil {
        request = NewCreateRedInvoiceRequest()
    }
    response = NewCreateRedInvoiceResponse()
    err = c.Send(request, response)
    return
}

func NewCreateRedInvoiceV2Request() (request *CreateRedInvoiceV2Request) {
    request = &CreateRedInvoiceV2Request{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cpdp", APIVersion, "CreateRedInvoiceV2")
    return
}

func NewCreateRedInvoiceV2Response() (response *CreateRedInvoiceV2Response) {
    response = &CreateRedInvoiceV2Response{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 智慧零售-发票红冲V2
func (c *Client) CreateRedInvoiceV2(request *CreateRedInvoiceV2Request) (response *CreateRedInvoiceV2Response, err error) {
    if request == nil {
        request = NewCreateRedInvoiceV2Request()
    }
    response = NewCreateRedInvoiceV2Response()
    err = c.Send(request, response)
    return
}

func NewCreateSinglePayRequest() (request *CreateSinglePayRequest) {
    request = &CreateSinglePayRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cpdp", APIVersion, "CreateSinglePay")
    return
}

func NewCreateSinglePayResponse() (response *CreateSinglePayResponse) {
    response = &CreateSinglePayResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 银企直连-单笔支付接口
func (c *Client) CreateSinglePay(request *CreateSinglePayRequest) (response *CreateSinglePayResponse, err error) {
    if request == nil {
        request = NewCreateSinglePayRequest()
    }
    response = NewCreateSinglePayResponse()
    err = c.Send(request, response)
    return
}

func NewCreateTransferBatchRequest() (request *CreateTransferBatchRequest) {
    request = &CreateTransferBatchRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cpdp", APIVersion, "CreateTransferBatch")
    return
}

func NewCreateTransferBatchResponse() (response *CreateTransferBatchResponse) {
    response = &CreateTransferBatchResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 微信商户发起批量转账
func (c *Client) CreateTransferBatch(request *CreateTransferBatchRequest) (response *CreateTransferBatchResponse, err error) {
    if request == nil {
        request = NewCreateTransferBatchRequest()
    }
    response = NewCreateTransferBatchResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteAgentTaxPaymentInfoRequest() (request *DeleteAgentTaxPaymentInfoRequest) {
    request = &DeleteAgentTaxPaymentInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cpdp", APIVersion, "DeleteAgentTaxPaymentInfo")
    return
}

func NewDeleteAgentTaxPaymentInfoResponse() (response *DeleteAgentTaxPaymentInfoResponse) {
    response = &DeleteAgentTaxPaymentInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 直播平台-删除代理商完税信息
func (c *Client) DeleteAgentTaxPaymentInfo(request *DeleteAgentTaxPaymentInfoRequest) (response *DeleteAgentTaxPaymentInfoResponse, err error) {
    if request == nil {
        request = NewDeleteAgentTaxPaymentInfoRequest()
    }
    response = NewDeleteAgentTaxPaymentInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteAgentTaxPaymentInfosRequest() (request *DeleteAgentTaxPaymentInfosRequest) {
    request = &DeleteAgentTaxPaymentInfosRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cpdp", APIVersion, "DeleteAgentTaxPaymentInfos")
    return
}

func NewDeleteAgentTaxPaymentInfosResponse() (response *DeleteAgentTaxPaymentInfosResponse) {
    response = &DeleteAgentTaxPaymentInfosResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 直播平台-删除代理商完税信息
func (c *Client) DeleteAgentTaxPaymentInfos(request *DeleteAgentTaxPaymentInfosRequest) (response *DeleteAgentTaxPaymentInfosResponse, err error) {
    if request == nil {
        request = NewDeleteAgentTaxPaymentInfosRequest()
    }
    response = NewDeleteAgentTaxPaymentInfosResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeChargeDetailRequest() (request *DescribeChargeDetailRequest) {
    request = &DescribeChargeDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cpdp", APIVersion, "DescribeChargeDetail")
    return
}

func NewDescribeChargeDetailResponse() (response *DescribeChargeDetailResponse) {
    response = &DescribeChargeDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询充值明细接口
func (c *Client) DescribeChargeDetail(request *DescribeChargeDetailRequest) (response *DescribeChargeDetailResponse, err error) {
    if request == nil {
        request = NewDescribeChargeDetailRequest()
    }
    response = NewDescribeChargeDetailResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeOrderStatusRequest() (request *DescribeOrderStatusRequest) {
    request = &DescribeOrderStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cpdp", APIVersion, "DescribeOrderStatus")
    return
}

func NewDescribeOrderStatusResponse() (response *DescribeOrderStatusResponse) {
    response = &DescribeOrderStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询单笔订单交易状态
func (c *Client) DescribeOrderStatus(request *DescribeOrderStatusRequest) (response *DescribeOrderStatusResponse, err error) {
    if request == nil {
        request = NewDescribeOrderStatusRequest()
    }
    response = NewDescribeOrderStatusResponse()
    err = c.Send(request, response)
    return
}

func NewDownloadBillRequest() (request *DownloadBillRequest) {
    request = &DownloadBillRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cpdp", APIVersion, "DownloadBill")
    return
}

func NewDownloadBillResponse() (response *DownloadBillResponse) {
    response = &DownloadBillResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 账单下载接口，根据本接口返回的URL地址，在D+1日下载对账单。注意，本接口返回的URL地址有时效，请尽快下载。URL超时时效后，请重新调用本接口再次获取。
func (c *Client) DownloadBill(request *DownloadBillRequest) (response *DownloadBillResponse, err error) {
    if request == nil {
        request = NewDownloadBillRequest()
    }
    response = NewDownloadBillResponse()
    err = c.Send(request, response)
    return
}

func NewExecuteMemberTransactionRequest() (request *ExecuteMemberTransactionRequest) {
    request = &ExecuteMemberTransactionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cpdp", APIVersion, "ExecuteMemberTransaction")
    return
}

func NewExecuteMemberTransactionResponse() (response *ExecuteMemberTransactionResponse) {
    response = &ExecuteMemberTransactionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 会员间交易接口
func (c *Client) ExecuteMemberTransaction(request *ExecuteMemberTransactionRequest) (response *ExecuteMemberTransactionResponse, err error) {
    if request == nil {
        request = NewExecuteMemberTransactionRequest()
    }
    response = NewExecuteMemberTransactionResponse()
    err = c.Send(request, response)
    return
}

func NewMigrateOrderRefundRequest() (request *MigrateOrderRefundRequest) {
    request = &MigrateOrderRefundRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cpdp", APIVersion, "MigrateOrderRefund")
    return
}

func NewMigrateOrderRefundResponse() (response *MigrateOrderRefundResponse) {
    response = &MigrateOrderRefundResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 山姆聚合支付项目-存量订单退款接口。可以通过本接口将支付款全部或部分退还给付款方，在收到用户退款请求并且验证成功之后，按照退款规则将支付款按原路退回到支付帐号。
func (c *Client) MigrateOrderRefund(request *MigrateOrderRefundRequest) (response *MigrateOrderRefundResponse, err error) {
    if request == nil {
        request = NewMigrateOrderRefundRequest()
    }
    response = NewMigrateOrderRefundResponse()
    err = c.Send(request, response)
    return
}

func NewMigrateOrderRefundQueryRequest() (request *MigrateOrderRefundQueryRequest) {
    request = &MigrateOrderRefundQueryRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cpdp", APIVersion, "MigrateOrderRefundQuery")
    return
}

func NewMigrateOrderRefundQueryResponse() (response *MigrateOrderRefundQueryResponse) {
    response = &MigrateOrderRefundQueryResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 提交退款申请后，通过调用该接口查询退款状态。退款可能有一定延时。
func (c *Client) MigrateOrderRefundQuery(request *MigrateOrderRefundQueryRequest) (response *MigrateOrderRefundQueryResponse, err error) {
    if request == nil {
        request = NewMigrateOrderRefundQueryRequest()
    }
    response = NewMigrateOrderRefundQueryResponse()
    err = c.Send(request, response)
    return
}

func NewModifyAgentTaxPaymentInfoRequest() (request *ModifyAgentTaxPaymentInfoRequest) {
    request = &ModifyAgentTaxPaymentInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cpdp", APIVersion, "ModifyAgentTaxPaymentInfo")
    return
}

func NewModifyAgentTaxPaymentInfoResponse() (response *ModifyAgentTaxPaymentInfoResponse) {
    response = &ModifyAgentTaxPaymentInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 直播平台-修改代理商完税信息
func (c *Client) ModifyAgentTaxPaymentInfo(request *ModifyAgentTaxPaymentInfoRequest) (response *ModifyAgentTaxPaymentInfoResponse, err error) {
    if request == nil {
        request = NewModifyAgentTaxPaymentInfoRequest()
    }
    response = NewModifyAgentTaxPaymentInfoResponse()
    err = c.Send(request, response)
    return
}

func NewModifyMntMbrBindRelateAcctBankCodeRequest() (request *ModifyMntMbrBindRelateAcctBankCodeRequest) {
    request = &ModifyMntMbrBindRelateAcctBankCodeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cpdp", APIVersion, "ModifyMntMbrBindRelateAcctBankCode")
    return
}

func NewModifyMntMbrBindRelateAcctBankCodeResponse() (response *ModifyMntMbrBindRelateAcctBankCodeResponse) {
    response = &ModifyMntMbrBindRelateAcctBankCodeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 维护会员绑定提现账户联行号。此接口可以支持市场修改会员的提现账户的开户行信息，具体包括开户行行名、开户行的银行联行号（大小额联行号）和超级网银行号。
func (c *Client) ModifyMntMbrBindRelateAcctBankCode(request *ModifyMntMbrBindRelateAcctBankCodeRequest) (response *ModifyMntMbrBindRelateAcctBankCodeResponse, err error) {
    if request == nil {
        request = NewModifyMntMbrBindRelateAcctBankCodeRequest()
    }
    response = NewModifyMntMbrBindRelateAcctBankCodeResponse()
    err = c.Send(request, response)
    return
}

func NewQueryAcctBindingRequest() (request *QueryAcctBindingRequest) {
    request = &QueryAcctBindingRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cpdp", APIVersion, "QueryAcctBinding")
    return
}

func NewQueryAcctBindingResponse() (response *QueryAcctBindingResponse) {
    response = &QueryAcctBindingResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 聚鑫-查询子账户绑定银行卡
func (c *Client) QueryAcctBinding(request *QueryAcctBindingRequest) (response *QueryAcctBindingResponse, err error) {
    if request == nil {
        request = NewQueryAcctBindingRequest()
    }
    response = NewQueryAcctBindingResponse()
    err = c.Send(request, response)
    return
}

func NewQueryAcctInfoRequest() (request *QueryAcctInfoRequest) {
    request = &QueryAcctInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cpdp", APIVersion, "QueryAcctInfo")
    return
}

func NewQueryAcctInfoResponse() (response *QueryAcctInfoResponse) {
    response = &QueryAcctInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 聚鑫-开户信息查询
func (c *Client) QueryAcctInfo(request *QueryAcctInfoRequest) (response *QueryAcctInfoResponse, err error) {
    if request == nil {
        request = NewQueryAcctInfoRequest()
    }
    response = NewQueryAcctInfoResponse()
    err = c.Send(request, response)
    return
}

func NewQueryAcctInfoListRequest() (request *QueryAcctInfoListRequest) {
    request = &QueryAcctInfoListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cpdp", APIVersion, "QueryAcctInfoList")
    return
}

func NewQueryAcctInfoListResponse() (response *QueryAcctInfoListResponse) {
    response = &QueryAcctInfoListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 聚鑫-开户信息列表查询, 查询某一段时间的开户信息
func (c *Client) QueryAcctInfoList(request *QueryAcctInfoListRequest) (response *QueryAcctInfoListResponse, err error) {
    if request == nil {
        request = NewQueryAcctInfoListRequest()
    }
    response = NewQueryAcctInfoListResponse()
    err = c.Send(request, response)
    return
}

func NewQueryAgentStatementsRequest() (request *QueryAgentStatementsRequest) {
    request = &QueryAgentStatementsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cpdp", APIVersion, "QueryAgentStatements")
    return
}

func NewQueryAgentStatementsResponse() (response *QueryAgentStatementsResponse) {
    response = &QueryAgentStatementsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 直播平台-查询代理商结算单链接
func (c *Client) QueryAgentStatements(request *QueryAgentStatementsRequest) (response *QueryAgentStatementsResponse, err error) {
    if request == nil {
        request = NewQueryAgentStatementsRequest()
    }
    response = NewQueryAgentStatementsResponse()
    err = c.Send(request, response)
    return
}

func NewQueryAgentTaxPaymentBatchRequest() (request *QueryAgentTaxPaymentBatchRequest) {
    request = &QueryAgentTaxPaymentBatchRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cpdp", APIVersion, "QueryAgentTaxPaymentBatch")
    return
}

func NewQueryAgentTaxPaymentBatchResponse() (response *QueryAgentTaxPaymentBatchResponse) {
    response = &QueryAgentTaxPaymentBatchResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 直播平台-查询批次信息
func (c *Client) QueryAgentTaxPaymentBatch(request *QueryAgentTaxPaymentBatchRequest) (response *QueryAgentTaxPaymentBatchResponse, err error) {
    if request == nil {
        request = NewQueryAgentTaxPaymentBatchRequest()
    }
    response = NewQueryAgentTaxPaymentBatchResponse()
    err = c.Send(request, response)
    return
}

func NewQueryAnchorContractInfoRequest() (request *QueryAnchorContractInfoRequest) {
    request = &QueryAnchorContractInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cpdp", APIVersion, "QueryAnchorContractInfo")
    return
}

func NewQueryAnchorContractInfoResponse() (response *QueryAnchorContractInfoResponse) {
    response = &QueryAnchorContractInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 直播平台-查询主播签约信息
func (c *Client) QueryAnchorContractInfo(request *QueryAnchorContractInfoRequest) (response *QueryAnchorContractInfoResponse, err error) {
    if request == nil {
        request = NewQueryAnchorContractInfoRequest()
    }
    response = NewQueryAnchorContractInfoResponse()
    err = c.Send(request, response)
    return
}

func NewQueryApplicationMaterialRequest() (request *QueryApplicationMaterialRequest) {
    request = &QueryApplicationMaterialRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cpdp", APIVersion, "QueryApplicationMaterial")
    return
}

func NewQueryApplicationMaterialResponse() (response *QueryApplicationMaterialResponse) {
    response = &QueryApplicationMaterialResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 跨境-成功申报材料查询。查询成功入库的申报材料。
func (c *Client) QueryApplicationMaterial(request *QueryApplicationMaterialRequest) (response *QueryApplicationMaterialResponse, err error) {
    if request == nil {
        request = NewQueryApplicationMaterialRequest()
    }
    response = NewQueryApplicationMaterialResponse()
    err = c.Send(request, response)
    return
}

func NewQueryBalanceRequest() (request *QueryBalanceRequest) {
    request = &QueryBalanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cpdp", APIVersion, "QueryBalance")
    return
}

func NewQueryBalanceResponse() (response *QueryBalanceResponse) {
    response = &QueryBalanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 子商户余额查询
func (c *Client) QueryBalance(request *QueryBalanceRequest) (response *QueryBalanceResponse, err error) {
    if request == nil {
        request = NewQueryBalanceRequest()
    }
    response = NewQueryBalanceResponse()
    err = c.Send(request, response)
    return
}

func NewQueryBankClearRequest() (request *QueryBankClearRequest) {
    request = &QueryBankClearRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cpdp", APIVersion, "QueryBankClear")
    return
}

func NewQueryBankClearResponse() (response *QueryBankClearResponse) {
    response = &QueryBankClearResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询银行在途清算结果。查询时间段内交易网的在途清算结果。
func (c *Client) QueryBankClear(request *QueryBankClearRequest) (response *QueryBankClearResponse, err error) {
    if request == nil {
        request = NewQueryBankClearRequest()
    }
    response = NewQueryBankClearResponse()
    err = c.Send(request, response)
    return
}

func NewQueryBankTransactionDetailsRequest() (request *QueryBankTransactionDetailsRequest) {
    request = &QueryBankTransactionDetailsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cpdp", APIVersion, "QueryBankTransactionDetails")
    return
}

func NewQueryBankTransactionDetailsResponse() (response *QueryBankTransactionDetailsResponse) {
    response = &QueryBankTransactionDetailsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询银行时间段内交易明细。查询时间段的会员成功交易。
func (c *Client) QueryBankTransactionDetails(request *QueryBankTransactionDetailsRequest) (response *QueryBankTransactionDetailsResponse, err error) {
    if request == nil {
        request = NewQueryBankTransactionDetailsRequest()
    }
    response = NewQueryBankTransactionDetailsResponse()
    err = c.Send(request, response)
    return
}

func NewQueryBankWithdrawCashDetailsRequest() (request *QueryBankWithdrawCashDetailsRequest) {
    request = &QueryBankWithdrawCashDetailsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cpdp", APIVersion, "QueryBankWithdrawCashDetails")
    return
}

func NewQueryBankWithdrawCashDetailsResponse() (response *QueryBankWithdrawCashDetailsResponse) {
    response = &QueryBankWithdrawCashDetailsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询银行时间段内清分提现明细。查询银行时间段内清分提现明细接口：若为“见证+收单退款”“见证+收单充值”记录时备注Note为“见证+收单充值,订单号”“见证+收单退款,订单号”，此接口可以查到T0/T1的充值明细和退款记录。查询标志：充值记录仍用3清分选项查询，退款记录同提现用2选项查询。
func (c *Client) QueryBankWithdrawCashDetails(request *QueryBankWithdrawCashDetailsRequest) (response *QueryBankWithdrawCashDetailsResponse, err error) {
    if request == nil {
        request = NewQueryBankWithdrawCashDetailsRequest()
    }
    response = NewQueryBankWithdrawCashDetailsResponse()
    err = c.Send(request, response)
    return
}

func NewQueryBillDownloadURLRequest() (request *QueryBillDownloadURLRequest) {
    request = &QueryBillDownloadURLRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cpdp", APIVersion, "QueryBillDownloadURL")
    return
}

func NewQueryBillDownloadURLResponse() (response *QueryBillDownloadURLResponse) {
    response = &QueryBillDownloadURLResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取单笔代发转账对账单下载URL
func (c *Client) QueryBillDownloadURL(request *QueryBillDownloadURLRequest) (response *QueryBillDownloadURLResponse, err error) {
    if request == nil {
        request = NewQueryBillDownloadURLRequest()
    }
    response = NewQueryBillDownloadURLResponse()
    err = c.Send(request, response)
    return
}

func NewQueryCommonTransferRechargeRequest() (request *QueryCommonTransferRechargeRequest) {
    request = &QueryCommonTransferRechargeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cpdp", APIVersion, "QueryCommonTransferRecharge")
    return
}

func NewQueryCommonTransferRechargeResponse() (response *QueryCommonTransferRechargeResponse) {
    response = &QueryCommonTransferRechargeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询普通转账充值明细。接口用于查询会员主动转账进资金汇总账户的明细情况。若会员使用绑定账号转入，则直接入账到会员子账户。若未使用绑定账号转入，则系统无法自动清分到对应子账户，则转入挂账子账户由平台自行清分。若是 “见证+收单充值”T0充值记录时备注Note为“见证+收单充值,订单号” 此接口可以查到T0到账的“见证+收单充值”充值记录。
func (c *Client) QueryCommonTransferRecharge(request *QueryCommonTransferRechargeRequest) (response *QueryCommonTransferRechargeResponse, err error) {
    if request == nil {
        request = NewQueryCommonTransferRechargeRequest()
    }
    response = NewQueryCommonTransferRechargeResponse()
    err = c.Send(request, response)
    return
}

func NewQueryContractRequest() (request *QueryContractRequest) {
    request = &QueryContractRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cpdp", APIVersion, "QueryContract")
    return
}

func NewQueryContractResponse() (response *QueryContractResponse) {
    response = &QueryContractResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 通过此接口查询签约数据
func (c *Client) QueryContract(request *QueryContractRequest) (response *QueryContractResponse, err error) {
    if request == nil {
        request = NewQueryContractRequest()
    }
    response = NewQueryContractResponse()
    err = c.Send(request, response)
    return
}

func NewQueryCustAcctIdBalanceRequest() (request *QueryCustAcctIdBalanceRequest) {
    request = &QueryCustAcctIdBalanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cpdp", APIVersion, "QueryCustAcctIdBalance")
    return
}

func NewQueryCustAcctIdBalanceResponse() (response *QueryCustAcctIdBalanceResponse) {
    response = &QueryCustAcctIdBalanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询银行子账户余额。查询会员子账户以及平台的功能子账户的余额。
func (c *Client) QueryCustAcctIdBalance(request *QueryCustAcctIdBalanceRequest) (response *QueryCustAcctIdBalanceResponse, err error) {
    if request == nil {
        request = NewQueryCustAcctIdBalanceRequest()
    }
    response = NewQueryCustAcctIdBalanceResponse()
    err = c.Send(request, response)
    return
}

func NewQueryExchangeRateRequest() (request *QueryExchangeRateRequest) {
    request = &QueryExchangeRateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cpdp", APIVersion, "QueryExchangeRate")
    return
}

func NewQueryExchangeRateResponse() (response *QueryExchangeRateResponse) {
    response = &QueryExchangeRateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 跨境-查询汇率
func (c *Client) QueryExchangeRate(request *QueryExchangeRateRequest) (response *QueryExchangeRateResponse, err error) {
    if request == nil {
        request = NewQueryExchangeRateRequest()
    }
    response = NewQueryExchangeRateResponse()
    err = c.Send(request, response)
    return
}

func NewQueryInvoiceRequest() (request *QueryInvoiceRequest) {
    request = &QueryInvoiceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cpdp", APIVersion, "QueryInvoice")
    return
}

func NewQueryInvoiceResponse() (response *QueryInvoiceResponse) {
    response = &QueryInvoiceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 智慧零售-发票查询
func (c *Client) QueryInvoice(request *QueryInvoiceRequest) (response *QueryInvoiceResponse, err error) {
    if request == nil {
        request = NewQueryInvoiceRequest()
    }
    response = NewQueryInvoiceResponse()
    err = c.Send(request, response)
    return
}

func NewQueryInvoiceV2Request() (request *QueryInvoiceV2Request) {
    request = &QueryInvoiceV2Request{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cpdp", APIVersion, "QueryInvoiceV2")
    return
}

func NewQueryInvoiceV2Response() (response *QueryInvoiceV2Response) {
    response = &QueryInvoiceV2Response{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 智慧零售-发票查询V2
func (c *Client) QueryInvoiceV2(request *QueryInvoiceV2Request) (response *QueryInvoiceV2Response, err error) {
    if request == nil {
        request = NewQueryInvoiceV2Request()
    }
    response = NewQueryInvoiceV2Response()
    err = c.Send(request, response)
    return
}

func NewQueryMemberBindRequest() (request *QueryMemberBindRequest) {
    request = &QueryMemberBindRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cpdp", APIVersion, "QueryMemberBind")
    return
}

func NewQueryMemberBindResponse() (response *QueryMemberBindResponse) {
    response = &QueryMemberBindResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 会员绑定信息查询。查询标志为“单个会员”的情况下，返回该会员的有效的绑定账户信息。
// 查询标志为“全部会员”的情况下，返回市场下的全部的有效的绑定账户信息。查询标志为“单个会员的证件信息”的情况下，返回市场下的指定的会员的留存在电商见证宝系统的证件信息。
func (c *Client) QueryMemberBind(request *QueryMemberBindRequest) (response *QueryMemberBindResponse, err error) {
    if request == nil {
        request = NewQueryMemberBindRequest()
    }
    response = NewQueryMemberBindResponse()
    err = c.Send(request, response)
    return
}

func NewQueryMemberTransactionRequest() (request *QueryMemberTransactionRequest) {
    request = &QueryMemberTransactionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cpdp", APIVersion, "QueryMemberTransaction")
    return
}

func NewQueryMemberTransactionResponse() (response *QueryMemberTransactionResponse) {
    response = &QueryMemberTransactionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 会员间交易-不验证。此接口可以实现会员间的余额的交易，实现资金在会员之间流动。
func (c *Client) QueryMemberTransaction(request *QueryMemberTransactionRequest) (response *QueryMemberTransactionResponse, err error) {
    if request == nil {
        request = NewQueryMemberTransactionRequest()
    }
    response = NewQueryMemberTransactionResponse()
    err = c.Send(request, response)
    return
}

func NewQueryMerchantBalanceRequest() (request *QueryMerchantBalanceRequest) {
    request = &QueryMerchantBalanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cpdp", APIVersion, "QueryMerchantBalance")
    return
}

func NewQueryMerchantBalanceResponse() (response *QueryMerchantBalanceResponse) {
    response = &QueryMerchantBalanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 跨境-对接方账户余额查询
func (c *Client) QueryMerchantBalance(request *QueryMerchantBalanceRequest) (response *QueryMerchantBalanceResponse, err error) {
    if request == nil {
        request = NewQueryMerchantBalanceRequest()
    }
    response = NewQueryMerchantBalanceResponse()
    err = c.Send(request, response)
    return
}

func NewQueryMerchantInfoForManagementRequest() (request *QueryMerchantInfoForManagementRequest) {
    request = &QueryMerchantInfoForManagementRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cpdp", APIVersion, "QueryMerchantInfoForManagement")
    return
}

func NewQueryMerchantInfoForManagementResponse() (response *QueryMerchantInfoForManagementResponse) {
    response = &QueryMerchantInfoForManagementResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 智慧零售-查询管理端商户
func (c *Client) QueryMerchantInfoForManagement(request *QueryMerchantInfoForManagementRequest) (response *QueryMerchantInfoForManagementResponse, err error) {
    if request == nil {
        request = NewQueryMerchantInfoForManagementRequest()
    }
    response = NewQueryMerchantInfoForManagementResponse()
    err = c.Send(request, response)
    return
}

func NewQueryOrderRequest() (request *QueryOrderRequest) {
    request = &QueryOrderRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cpdp", APIVersion, "QueryOrder")
    return
}

func NewQueryOrderResponse() (response *QueryOrderResponse) {
    response = &QueryOrderResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 根据订单号，或者用户Id，查询支付订单状态 
func (c *Client) QueryOrder(request *QueryOrderRequest) (response *QueryOrderResponse, err error) {
    if request == nil {
        request = NewQueryOrderRequest()
    }
    response = NewQueryOrderResponse()
    err = c.Send(request, response)
    return
}

func NewQueryOutwardOrderRequest() (request *QueryOutwardOrderRequest) {
    request = &QueryOutwardOrderRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cpdp", APIVersion, "QueryOutwardOrder")
    return
}

func NewQueryOutwardOrderResponse() (response *QueryOutwardOrderResponse) {
    response = &QueryOutwardOrderResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 跨境-查询汇出结果
func (c *Client) QueryOutwardOrder(request *QueryOutwardOrderRequest) (response *QueryOutwardOrderResponse, err error) {
    if request == nil {
        request = NewQueryOutwardOrderRequest()
    }
    response = NewQueryOutwardOrderResponse()
    err = c.Send(request, response)
    return
}

func NewQueryPayerInfoRequest() (request *QueryPayerInfoRequest) {
    request = &QueryPayerInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cpdp", APIVersion, "QueryPayerInfo")
    return
}

func NewQueryPayerInfoResponse() (response *QueryPayerInfoResponse) {
    response = &QueryPayerInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 跨境-付款人查询
func (c *Client) QueryPayerInfo(request *QueryPayerInfoRequest) (response *QueryPayerInfoResponse, err error) {
    if request == nil {
        request = NewQueryPayerInfoRequest()
    }
    response = NewQueryPayerInfoResponse()
    err = c.Send(request, response)
    return
}

func NewQueryReconciliationDocumentRequest() (request *QueryReconciliationDocumentRequest) {
    request = &QueryReconciliationDocumentRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cpdp", APIVersion, "QueryReconciliationDocument")
    return
}

func NewQueryReconciliationDocumentResponse() (response *QueryReconciliationDocumentResponse) {
    response = &QueryReconciliationDocumentResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询对账文件信息。平台调用该接口获取需下载对账文件的文件名称以及密钥。 平台获取到信息后， 可以再调用OPENAPI的文件下载功能。
func (c *Client) QueryReconciliationDocument(request *QueryReconciliationDocumentRequest) (response *QueryReconciliationDocumentResponse, err error) {
    if request == nil {
        request = NewQueryReconciliationDocumentRequest()
    }
    response = NewQueryReconciliationDocumentResponse()
    err = c.Send(request, response)
    return
}

func NewQueryRefundRequest() (request *QueryRefundRequest) {
    request = &QueryRefundRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cpdp", APIVersion, "QueryRefund")
    return
}

func NewQueryRefundResponse() (response *QueryRefundResponse) {
    response = &QueryRefundResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 提交退款申请后，通过调用该接口查询退款状态。退款可能有一定延时，用微信零钱支付的退款约20分钟内到账，银行卡支付的退款约3个工作日后到账。
func (c *Client) QueryRefund(request *QueryRefundRequest) (response *QueryRefundResponse, err error) {
    if request == nil {
        request = NewQueryRefundRequest()
    }
    response = NewQueryRefundResponse()
    err = c.Send(request, response)
    return
}

func NewQuerySinglePayRequest() (request *QuerySinglePayRequest) {
    request = &QuerySinglePayRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cpdp", APIVersion, "QuerySinglePay")
    return
}

func NewQuerySinglePayResponse() (response *QuerySinglePayResponse) {
    response = &QuerySinglePayResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 银企直连-单笔支付状态查询接口
func (c *Client) QuerySinglePay(request *QuerySinglePayRequest) (response *QuerySinglePayResponse, err error) {
    if request == nil {
        request = NewQuerySinglePayRequest()
    }
    response = NewQuerySinglePayResponse()
    err = c.Send(request, response)
    return
}

func NewQuerySingleTransactionStatusRequest() (request *QuerySingleTransactionStatusRequest) {
    request = &QuerySingleTransactionStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cpdp", APIVersion, "QuerySingleTransactionStatus")
    return
}

func NewQuerySingleTransactionStatusResponse() (response *QuerySingleTransactionStatusResponse) {
    response = &QuerySingleTransactionStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询银行单笔交易状态。查询单笔交易的状态。
func (c *Client) QuerySingleTransactionStatus(request *QuerySingleTransactionStatusRequest) (response *QuerySingleTransactionStatusResponse, err error) {
    if request == nil {
        request = NewQuerySingleTransactionStatusRequest()
    }
    response = NewQuerySingleTransactionStatusResponse()
    err = c.Send(request, response)
    return
}

func NewQuerySmallAmountTransferRequest() (request *QuerySmallAmountTransferRequest) {
    request = &QuerySmallAmountTransferRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cpdp", APIVersion, "QuerySmallAmountTransfer")
    return
}

func NewQuerySmallAmountTransferResponse() (response *QuerySmallAmountTransferResponse) {
    response = &QuerySmallAmountTransferResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询小额鉴权转账结果。查询小额往账鉴权的转账状态。
func (c *Client) QuerySmallAmountTransfer(request *QuerySmallAmountTransferRequest) (response *QuerySmallAmountTransferResponse, err error) {
    if request == nil {
        request = NewQuerySmallAmountTransferRequest()
    }
    response = NewQuerySmallAmountTransferResponse()
    err = c.Send(request, response)
    return
}

func NewQueryTradeRequest() (request *QueryTradeRequest) {
    request = &QueryTradeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cpdp", APIVersion, "QueryTrade")
    return
}

func NewQueryTradeResponse() (response *QueryTradeResponse) {
    response = &QueryTradeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 跨境-贸易材料明细查询。
func (c *Client) QueryTrade(request *QueryTradeRequest) (response *QueryTradeResponse, err error) {
    if request == nil {
        request = NewQueryTradeRequest()
    }
    response = NewQueryTradeResponse()
    err = c.Send(request, response)
    return
}

func NewQueryTransferBatchRequest() (request *QueryTransferBatchRequest) {
    request = &QueryTransferBatchRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cpdp", APIVersion, "QueryTransferBatch")
    return
}

func NewQueryTransferBatchResponse() (response *QueryTransferBatchResponse) {
    response = &QueryTransferBatchResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 通过商家批次单号或者微信批次号查询批次单
func (c *Client) QueryTransferBatch(request *QueryTransferBatchRequest) (response *QueryTransferBatchResponse, err error) {
    if request == nil {
        request = NewQueryTransferBatchRequest()
    }
    response = NewQueryTransferBatchResponse()
    err = c.Send(request, response)
    return
}

func NewQueryTransferDetailRequest() (request *QueryTransferDetailRequest) {
    request = &QueryTransferDetailRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cpdp", APIVersion, "QueryTransferDetail")
    return
}

func NewQueryTransferDetailResponse() (response *QueryTransferDetailResponse) {
    response = &QueryTransferDetailResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 通过商家或者微信批次明细单号查询明细单
func (c *Client) QueryTransferDetail(request *QueryTransferDetailRequest) (response *QueryTransferDetailResponse, err error) {
    if request == nil {
        request = NewQueryTransferDetailRequest()
    }
    response = NewQueryTransferDetailResponse()
    err = c.Send(request, response)
    return
}

func NewQueryTransferResultRequest() (request *QueryTransferResultRequest) {
    request = &QueryTransferResultRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cpdp", APIVersion, "QueryTransferResult")
    return
}

func NewQueryTransferResultResponse() (response *QueryTransferResultResponse) {
    response = &QueryTransferResultResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 智能代发-单笔代发转账查询接口
func (c *Client) QueryTransferResult(request *QueryTransferResultRequest) (response *QueryTransferResultResponse, err error) {
    if request == nil {
        request = NewQueryTransferResultRequest()
    }
    response = NewQueryTransferResultResponse()
    err = c.Send(request, response)
    return
}

func NewRechargeByThirdPayRequest() (request *RechargeByThirdPayRequest) {
    request = &RechargeByThirdPayRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cpdp", APIVersion, "RechargeByThirdPay")
    return
}

func NewRechargeByThirdPayResponse() (response *RechargeByThirdPayResponse) {
    response = &RechargeByThirdPayResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 会员在途充值(经第三方支付渠道)接口
func (c *Client) RechargeByThirdPay(request *RechargeByThirdPayRequest) (response *RechargeByThirdPayResponse, err error) {
    if request == nil {
        request = NewRechargeByThirdPayRequest()
    }
    response = NewRechargeByThirdPayResponse()
    err = c.Send(request, response)
    return
}

func NewRechargeMemberThirdPayRequest() (request *RechargeMemberThirdPayRequest) {
    request = &RechargeMemberThirdPayRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cpdp", APIVersion, "RechargeMemberThirdPay")
    return
}

func NewRechargeMemberThirdPayResponse() (response *RechargeMemberThirdPayResponse) {
    response = &RechargeMemberThirdPayResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 见证宝-会员在途充值(经第三方支付渠道)
func (c *Client) RechargeMemberThirdPay(request *RechargeMemberThirdPayRequest) (response *RechargeMemberThirdPayResponse, err error) {
    if request == nil {
        request = NewRechargeMemberThirdPayRequest()
    }
    response = NewRechargeMemberThirdPayResponse()
    err = c.Send(request, response)
    return
}

func NewRefundRequest() (request *RefundRequest) {
    request = &RefundRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cpdp", APIVersion, "Refund")
    return
}

func NewRefundResponse() (response *RefundResponse) {
    response = &RefundResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 如交易订单需退款，可以通过本接口将支付款全部或部分退还给付款方，聚鑫将在收到退款请求并且验证成功之后，按照退款规则将支付款按原路退回到支付帐号。最长支持1年的订单退款。在订单包含多个子订单的情况下，如果使用本接口传入OutTradeNo或TransactionId退款，则只支持全单退款；如果需要部分退款，请通过传入子订单的方式来指定部分金额退款。 
func (c *Client) Refund(request *RefundRequest) (response *RefundResponse, err error) {
    if request == nil {
        request = NewRefundRequest()
    }
    response = NewRefundResponse()
    err = c.Send(request, response)
    return
}

func NewRefundMemberTransactionRequest() (request *RefundMemberTransactionRequest) {
    request = &RefundMemberTransactionRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cpdp", APIVersion, "RefundMemberTransaction")
    return
}

func NewRefundMemberTransactionResponse() (response *RefundMemberTransactionResponse) {
    response = &RefundMemberTransactionResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 会员间交易退款
func (c *Client) RefundMemberTransaction(request *RefundMemberTransactionRequest) (response *RefundMemberTransactionResponse, err error) {
    if request == nil {
        request = NewRefundMemberTransactionRequest()
    }
    response = NewRefundMemberTransactionResponse()
    err = c.Send(request, response)
    return
}

func NewRegisterBillRequest() (request *RegisterBillRequest) {
    request = &RegisterBillRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cpdp", APIVersion, "RegisterBill")
    return
}

func NewRegisterBillResponse() (response *RegisterBillResponse) {
    response = &RegisterBillResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 登记挂账(支持撤销)
func (c *Client) RegisterBill(request *RegisterBillRequest) (response *RegisterBillResponse, err error) {
    if request == nil {
        request = NewRegisterBillRequest()
    }
    response = NewRegisterBillResponse()
    err = c.Send(request, response)
    return
}

func NewRegisterBillSupportWithdrawRequest() (request *RegisterBillSupportWithdrawRequest) {
    request = &RegisterBillSupportWithdrawRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cpdp", APIVersion, "RegisterBillSupportWithdraw")
    return
}

func NewRegisterBillSupportWithdrawResponse() (response *RegisterBillSupportWithdrawResponse) {
    response = &RegisterBillSupportWithdrawResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 登记挂账(支持撤销)。此接口可实现把不明来账或自有资金等已登记在挂账子账户下的资金调整到普通会员子账户。即通过申请调用此接口，将会减少挂账子账户的资金，调增指定的普通会员子账户的可提现余额及可用余额。此接口不支持把挂账子账户资金清分到功能子账户。
func (c *Client) RegisterBillSupportWithdraw(request *RegisterBillSupportWithdrawRequest) (response *RegisterBillSupportWithdrawResponse, err error) {
    if request == nil {
        request = NewRegisterBillSupportWithdrawRequest()
    }
    response = NewRegisterBillSupportWithdrawResponse()
    err = c.Send(request, response)
    return
}

func NewRevResigterBillSupportWithdrawRequest() (request *RevResigterBillSupportWithdrawRequest) {
    request = &RevResigterBillSupportWithdrawRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cpdp", APIVersion, "RevResigterBillSupportWithdraw")
    return
}

func NewRevResigterBillSupportWithdrawResponse() (response *RevResigterBillSupportWithdrawResponse) {
    response = &RevResigterBillSupportWithdrawResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 登记挂账撤销。此接口可以实现把RegisterBillSupportWithdraw接口完成的登记挂账进行撤销，即调减普通会员子账户的可提现和可用余额，调增挂账子账户的可用余额。
func (c *Client) RevResigterBillSupportWithdraw(request *RevResigterBillSupportWithdrawRequest) (response *RevResigterBillSupportWithdrawResponse, err error) {
    if request == nil {
        request = NewRevResigterBillSupportWithdrawRequest()
    }
    response = NewRevResigterBillSupportWithdrawResponse()
    err = c.Send(request, response)
    return
}

func NewReviseMbrPropertyRequest() (request *ReviseMbrPropertyRequest) {
    request = &ReviseMbrPropertyRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cpdp", APIVersion, "ReviseMbrProperty")
    return
}

func NewReviseMbrPropertyResponse() (response *ReviseMbrPropertyResponse) {
    response = &ReviseMbrPropertyResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 修改会员属性-普通商户子账户。修改会员的会员属性。
func (c *Client) ReviseMbrProperty(request *ReviseMbrPropertyRequest) (response *ReviseMbrPropertyResponse, err error) {
    if request == nil {
        request = NewReviseMbrPropertyRequest()
    }
    response = NewReviseMbrPropertyResponse()
    err = c.Send(request, response)
    return
}

func NewRevokeMemberRechargeThirdPayRequest() (request *RevokeMemberRechargeThirdPayRequest) {
    request = &RevokeMemberRechargeThirdPayRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cpdp", APIVersion, "RevokeMemberRechargeThirdPay")
    return
}

func NewRevokeMemberRechargeThirdPayResponse() (response *RevokeMemberRechargeThirdPayResponse) {
    response = &RevokeMemberRechargeThirdPayResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 撤销会员在途充值(经第三方支付渠道)
func (c *Client) RevokeMemberRechargeThirdPay(request *RevokeMemberRechargeThirdPayRequest) (response *RevokeMemberRechargeThirdPayResponse, err error) {
    if request == nil {
        request = NewRevokeMemberRechargeThirdPayRequest()
    }
    response = NewRevokeMemberRechargeThirdPayResponse()
    err = c.Send(request, response)
    return
}

func NewRevokeRechargeByThirdPayRequest() (request *RevokeRechargeByThirdPayRequest) {
    request = &RevokeRechargeByThirdPayRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cpdp", APIVersion, "RevokeRechargeByThirdPay")
    return
}

func NewRevokeRechargeByThirdPayResponse() (response *RevokeRechargeByThirdPayResponse) {
    response = &RevokeRechargeByThirdPayResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 撤销会员在途充值(经第三方支付渠道)接口
func (c *Client) RevokeRechargeByThirdPay(request *RevokeRechargeByThirdPayRequest) (response *RevokeRechargeByThirdPayResponse, err error) {
    if request == nil {
        request = NewRevokeRechargeByThirdPayRequest()
    }
    response = NewRevokeRechargeByThirdPayResponse()
    err = c.Send(request, response)
    return
}

func NewSyncContractDataRequest() (request *SyncContractDataRequest) {
    request = &SyncContractDataRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cpdp", APIVersion, "SyncContractData")
    return
}

func NewSyncContractDataResponse() (response *SyncContractDataResponse) {
    response = &SyncContractDataResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 对于存量的签约关系导入或者部分场景下米大师无法收到签约通知的场景，需要由调用方主动将签约状态同步至米大师
func (c *Client) SyncContractData(request *SyncContractDataRequest) (response *SyncContractDataResponse, err error) {
    if request == nil {
        request = NewSyncContractDataRequest()
    }
    response = NewSyncContractDataResponse()
    err = c.Send(request, response)
    return
}

func NewTerminateContractRequest() (request *TerminateContractRequest) {
    request = &TerminateContractRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cpdp", APIVersion, "TerminateContract")
    return
}

func NewTerminateContractResponse() (response *TerminateContractResponse) {
    response = &TerminateContractResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 通过此接口进行解约
func (c *Client) TerminateContract(request *TerminateContractRequest) (response *TerminateContractResponse, err error) {
    if request == nil {
        request = NewTerminateContractRequest()
    }
    response = NewTerminateContractResponse()
    err = c.Send(request, response)
    return
}

func NewTransferSinglePayRequest() (request *TransferSinglePayRequest) {
    request = &TransferSinglePayRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cpdp", APIVersion, "TransferSinglePay")
    return
}

func NewTransferSinglePayResponse() (response *TransferSinglePayResponse) {
    response = &TransferSinglePayResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 智能代发-单笔代发转账接口
func (c *Client) TransferSinglePay(request *TransferSinglePayRequest) (response *TransferSinglePayResponse, err error) {
    if request == nil {
        request = NewTransferSinglePayRequest()
    }
    response = NewTransferSinglePayResponse()
    err = c.Send(request, response)
    return
}

func NewUnBindAcctRequest() (request *UnBindAcctRequest) {
    request = &UnBindAcctRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cpdp", APIVersion, "UnBindAcct")
    return
}

func NewUnBindAcctResponse() (response *UnBindAcctResponse) {
    response = &UnBindAcctResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 商户解除绑定的提现银行卡
func (c *Client) UnBindAcct(request *UnBindAcctRequest) (response *UnBindAcctResponse, err error) {
    if request == nil {
        request = NewUnBindAcctRequest()
    }
    response = NewUnBindAcctResponse()
    err = c.Send(request, response)
    return
}

func NewUnbindRelateAcctRequest() (request *UnbindRelateAcctRequest) {
    request = &UnbindRelateAcctRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cpdp", APIVersion, "UnbindRelateAcct")
    return
}

func NewUnbindRelateAcctResponse() (response *UnbindRelateAcctResponse) {
    response = &UnbindRelateAcctResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 会员解绑提现账户。此接口可以支持会员解除名下的绑定账户关系。
func (c *Client) UnbindRelateAcct(request *UnbindRelateAcctRequest) (response *UnbindRelateAcctResponse, err error) {
    if request == nil {
        request = NewUnbindRelateAcctRequest()
    }
    response = NewUnbindRelateAcctResponse()
    err = c.Send(request, response)
    return
}

func NewUnifiedOrderRequest() (request *UnifiedOrderRequest) {
    request = &UnifiedOrderRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cpdp", APIVersion, "UnifiedOrder")
    return
}

func NewUnifiedOrderResponse() (response *UnifiedOrderResponse) {
    response = &UnifiedOrderResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 应用需要先调用本接口生成支付订单号，并将应答的PayInfo透传给聚鑫SDK，拉起客户端（包括微信公众号/微信小程序/客户端App）支付。
func (c *Client) UnifiedOrder(request *UnifiedOrderRequest) (response *UnifiedOrderResponse, err error) {
    if request == nil {
        request = NewUnifiedOrderRequest()
    }
    response = NewUnifiedOrderResponse()
    err = c.Send(request, response)
    return
}

func NewUploadTaxListRequest() (request *UploadTaxListRequest) {
    request = &UploadTaxListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cpdp", APIVersion, "UploadTaxList")
    return
}

func NewUploadTaxListResponse() (response *UploadTaxListResponse) {
    response = &UploadTaxListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 直播平台-上传代理商完税列表
func (c *Client) UploadTaxList(request *UploadTaxListRequest) (response *UploadTaxListResponse, err error) {
    if request == nil {
        request = NewUploadTaxListRequest()
    }
    response = NewUploadTaxListResponse()
    err = c.Send(request, response)
    return
}

func NewUploadTaxPaymentRequest() (request *UploadTaxPaymentRequest) {
    request = &UploadTaxPaymentRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cpdp", APIVersion, "UploadTaxPayment")
    return
}

func NewUploadTaxPaymentResponse() (response *UploadTaxPaymentResponse) {
    response = &UploadTaxPaymentResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 直播平台-上传代理商完税证明
func (c *Client) UploadTaxPayment(request *UploadTaxPaymentRequest) (response *UploadTaxPaymentResponse, err error) {
    if request == nil {
        request = NewUploadTaxPaymentRequest()
    }
    response = NewUploadTaxPaymentResponse()
    err = c.Send(request, response)
    return
}

func NewWithdrawCashMembershipRequest() (request *WithdrawCashMembershipRequest) {
    request = &WithdrawCashMembershipRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("cpdp", APIVersion, "WithdrawCashMembership")
    return
}

func NewWithdrawCashMembershipResponse() (response *WithdrawCashMembershipResponse) {
    response = &WithdrawCashMembershipResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 会员提现-不验证。此接口受理会员发起的提现申请。会员子账户的可提现余额、可用余额会减少，市场的资金汇总账户(监管账户)会减少相应的发生金额，提现到会员申请的收款账户。		
func (c *Client) WithdrawCashMembership(request *WithdrawCashMembershipRequest) (response *WithdrawCashMembershipResponse, err error) {
    if request == nil {
        request = NewWithdrawCashMembershipRequest()
    }
    response = NewWithdrawCashMembershipResponse()
    err = c.Send(request, response)
    return
}
