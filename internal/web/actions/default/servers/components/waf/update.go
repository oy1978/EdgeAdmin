package waf

import (
	"encoding/json"
	"net/http"

	"github.com/iwind/TeaGo/actions"
	"github.com/iwind/TeaGo/maps"
	"github.com/iwind/TeaGo/types"
	"github.com/oy1978/EdgeAdmin/internal/web/actions/actionutils"
	"github.com/oy1978/EdgeCommon/pkg/langs/codes"
	"github.com/oy1978/EdgeCommon/pkg/rpc/dao"
	"github.com/oy1978/EdgeCommon/pkg/rpc/pb"
	"github.com/oy1978/EdgeCommon/pkg/serverconfigs/firewallconfigs"
)

type UpdateAction struct {
	actionutils.ParentAction
}

func (this *UpdateAction) Init() {
	this.Nav("", "", "update")
}

func (this *UpdateAction) RunGet(params struct {
	FirewallPolicyId int64
}) {
	firewallPolicy, err := dao.SharedHTTPFirewallPolicyDAO.FindEnabledHTTPFirewallPolicyConfig(this.AdminContext(), params.FirewallPolicyId)
	if err != nil {
		this.ErrorPage(err)
		return
	}
	if firewallPolicy == nil {
		this.NotFound("firewallPolicy", params.FirewallPolicyId)
		return
	}

	if firewallPolicy.BlockOptions == nil {
		firewallPolicy.BlockOptions = &firewallconfigs.HTTPFirewallBlockAction{
			StatusCode: http.StatusForbidden,
			Body:       "Blocked By WAF",
			URL:        "",
			Timeout:    60,
		}
	}

	// mode
	if len(firewallPolicy.Mode) == 0 {
		firewallPolicy.Mode = firewallconfigs.FirewallModeDefend
	}
	this.Data["modes"] = firewallconfigs.FindAllFirewallModes()

	// syn flood
	if firewallPolicy.SYNFlood == nil {
		firewallPolicy.SYNFlood = &firewallconfigs.SYNFloodConfig{
			IsOn:           false,
			MinAttempts:    10,
			TimeoutSeconds: 600,
			IgnoreLocal:    true,
		}
	}

	// log
	if firewallPolicy.Log == nil {
		firewallPolicy.Log = firewallconfigs.DefaultHTTPFirewallPolicyLogConfig
	}

	this.Data["firewallPolicy"] = maps.Map{
		"id":                 firewallPolicy.Id,
		"name":               firewallPolicy.Name,
		"description":        firewallPolicy.Description,
		"isOn":               firewallPolicy.IsOn,
		"mode":               firewallPolicy.Mode,
		"blockOptions":       firewallPolicy.BlockOptions,
		"captchaOptions":     firewallPolicy.CaptchaOptions,
		"silOptions":         firewallPolicy.SilOptions,
		"useLocalFirewall":   firewallPolicy.UseLocalFirewall,
		"synFloodConfig":     firewallPolicy.SYNFlood,
		"log":                firewallPolicy.Log,
		"maxRequestBodySize": types.String(firewallPolicy.MaxRequestBodySize),
		"denyCountryHTML":    firewallPolicy.DenyCountryHTML,
		"denyProvinceHTML":   firewallPolicy.DenyProvinceHTML,
	}

	// 预置分组
	var groups = []maps.Map{}
	templatePolicy := firewallconfigs.HTTPFirewallTemplate()
	for _, group := range templatePolicy.AllRuleGroups() {
		if len(group.Code) > 0 {
			usedGroup := firewallPolicy.FindRuleGroupWithCode(group.Code)
			if usedGroup != nil {
				group.IsOn = usedGroup.IsOn
			}
		}

		groups = append(groups, maps.Map{
			"code": group.Code,
			"name": group.Name,
			"isOn": group.IsOn,
		})
	}
	this.Data["groups"] = groups

	this.Show()
}

func (this *UpdateAction) RunPost(params struct {
	FirewallPolicyId   int64
	Name               string
	GroupCodes         []string
	BlockOptionsJSON   []byte
	CaptchaOptionsJSON []byte
	SilOptionsJSON     []byte
	Description        string
	IsOn               bool
	Mode               string
	UseLocalFirewall   bool
	SynFloodJSON       []byte
	LogJSON            []byte
	MaxRequestBodySize int64
	DenyCountryHTML    string
	DenyProvinceHTML   string

	Must *actions.Must
}) {
	// 日志
	defer this.CreateLogInfo(codes.WAFPolicy_LogUpdateWAFPolicy, params.FirewallPolicyId)

	params.Must.
		Field("name", params.Name).
		Require("请输入策略名称")

	// 校验拦截选项JSON
	var blockOptions = &firewallconfigs.HTTPFirewallBlockAction{}
	err := json.Unmarshal(params.BlockOptionsJSON, blockOptions)
	if err != nil {
		this.Fail("拦截动作参数校验失败：" + err.Error())
	}

	// 校验验证码选项JSON
	var captchaOptions = &firewallconfigs.HTTPFirewallCaptchaAction{}
	err = json.Unmarshal(params.CaptchaOptionsJSON, captchaOptions)
	if err != nil {
		this.Fail("验证码动作参数校验失败：" + err.Error())
	}

	var silOptions = &firewallconfigs.HTTPFirewallSilAction{}
	err = json.Unmarshal(params.SilOptionsJSON, silOptions)
	if err != nil {
		this.Fail("验证码动作参数校验失败：" + err.Error())
	}

	// 最大内容尺寸
	if params.MaxRequestBodySize < 0 {
		params.MaxRequestBodySize = 0
	}

	_, err = this.RPC().HTTPFirewallPolicyRPC().UpdateHTTPFirewallPolicy(this.AdminContext(), &pb.UpdateHTTPFirewallPolicyRequest{
		HttpFirewallPolicyId: params.FirewallPolicyId,
		IsOn:                 params.IsOn,
		Name:                 params.Name,
		Description:          params.Description,
		FirewallGroupCodes:   params.GroupCodes,
		BlockOptionsJSON:     params.BlockOptionsJSON,
		CaptchaOptionsJSON:   params.CaptchaOptionsJSON,
		SilOptionsJSON:       params.SilOptionsJSON,
		Mode:                 params.Mode,
		UseLocalFirewall:     params.UseLocalFirewall,
		SynFloodJSON:         params.SynFloodJSON,
		LogJSON:              params.LogJSON,
		MaxRequestBodySize:   params.MaxRequestBodySize,
		DenyCountryHTML:      params.DenyCountryHTML,
		DenyProvinceHTML:     params.DenyProvinceHTML,
	})
	if err != nil {
		this.ErrorPage(err)
		return
	}

	this.Success()
}
