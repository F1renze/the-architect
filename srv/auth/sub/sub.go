package sub

import (
	"context"
	"encoding/json"
	"os"
	"strconv"
	"time"

	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	sms "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/sms/v20190711" //引入sms

	"github.com/f1renze/the-architect/common/constant"
	"github.com/f1renze/the-architect/common/errno"
	"github.com/f1renze/the-architect/common/utils/log"
	"github.com/f1renze/the-architect/srv/auth/model"
	pb "github.com/f1renze/the-architect/srv/auth/proto"
)

func NewSubscriber() *Sub {
	credential := common.NewCredential(
		os.Getenv("TX_CLOUD_SECRET_ID"),
		os.Getenv("TX_CLOUD_SECRET_KEY"),
	)
	client, _ := sms.NewClient(credential, "ap-guangzhou", profile.NewClientProfile())

	return &Sub{
		smsCli: client,
		otp:    model.NewOtpSrv(),
	}
}

type Sub struct {
	smsCli *sms.Client
	otp    model.OtpSrv
}

func (s *Sub) SendRegistrationEmail() func(context.Context, *pb.ConfirmEmail) error {
	return func(ctx context.Context, event *pb.ConfirmEmail) error {
		log.InfoF("[srv.auth.sub::SendRegistrationEmail] Received event %+v", event)

		// todo send email
		return nil
	}
}

// https://github.com/TencentCloud/tencentcloud-sdk-go/blob/master/examples/sms/v20190711/SendSms.go
func (s *Sub) SendRegistrationSms() func(context.Context, *pb.ConfirmMobile) error {
	return func(ctx context.Context, event *pb.ConfirmMobile) error {
		log.InfoF("[srv.auth.sub::SendRegistrationSms] received event %+v", event)

		code, err := s.otp.GenerateCode(event.Mobile)
		if err != nil {
			log.ErrorF("[srv.auth.sub::SendRegistrationSms] generate code error: %s", err)
			return err
		}

		req := sms.NewSendSmsRequest()
		// 短信应用ID: 短信SdkAppid在 [短信控制台] 添加应用后生成的实际SdkAppid
		req.SmsSdkAppid = common.StringPtr(os.Getenv("TX_CLOUD_SMS_SDK_APP_ID"))
		// 短信签名内容: 使用 UTF-8 编码，必须填写已审核通过的签名，签名信息可登录 [短信控制台] 查看
		req.Sign = common.StringPtr(os.Getenv("TX_CLOUD_SIGN"))
		// 模板 ID: 必须填写已审核通过的模板 ID。模板ID可登录 [短信控制台] 查看
		req.TemplateID = common.StringPtr(os.Getenv("TX_CLOUD_TEMPLATE_ID"))
		// 模板参数: 若无模板参数，则设置为空
		exp := strconv.Itoa(int(constant.SmsCodeExpiredTime / time.Minute))
		req.TemplateParamSet = common.StringPtrs([]string{code, exp})
		/* 下发手机号码，采用 e.164 标准，+[国家或地区码][手机号]
		 * 示例如：+8613711112222， 其中前面有一个+号 ，86为国家码，13711112222为手机号，最多不要超过200个手机号*/
		req.PhoneNumberSet = common.StringPtrs([]string{"+86" + event.Mobile})

		resp, err := s.smsCli.SendSms(req)
		// 处理异常
		if _, ok := err.(*errors.TencentCloudSDKError); ok {
			log.ErrorF("[srv.auth.sub::SendRegistrationSms] sms api return a error", err)
			return err
		}
		// 非SDK异常，直接失败。实际代码中可以加入其他的处理。
		if err != nil {
			log.ErrorF("[srv.auth.sub::SendRegistrationSms] error occur", err)
			return errno.SystemErr.With(err)
		}
		b, _ := json.Marshal(resp.Response)
		// 打印返回的json字符串
		log.InfoF("[srv.auth.sub::SendRegistrationSms] received from sms api: %s", b)
		//{"SendStatusSet":[{"SerialNo":"2019:2692126464494930333","PhoneNumber":"+8613826886196","Fee":1,"SessionContext":"","Code":"Ok","Message":"send success"}],"RequestId":"36df623a-da55-4efd-a1b6-a075b7955ae7"}

		return nil
	}
}
