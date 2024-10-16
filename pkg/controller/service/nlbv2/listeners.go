package nlbv2

import (
	"context"
	"fmt"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/mohae/deepcopy"
	v1 "k8s.io/api/core/v1"
	"k8s.io/cloud-provider-alibaba-cloud/pkg/controller/service/reconcile/annotation"
	svcCtx "k8s.io/cloud-provider-alibaba-cloud/pkg/controller/service/reconcile/context"
	"k8s.io/cloud-provider-alibaba-cloud/pkg/model"
	nlbmodel "k8s.io/cloud-provider-alibaba-cloud/pkg/model/nlb"
	prvd "k8s.io/cloud-provider-alibaba-cloud/pkg/provider"
	"k8s.io/cloud-provider-alibaba-cloud/pkg/provider/alibaba/base"
	"k8s.io/cloud-provider-alibaba-cloud/pkg/provider/dryrun"
	"k8s.io/cloud-provider-alibaba-cloud/pkg/util"
	"strconv"
	"strings"
)

func NewListenerManager(cloud prvd.Provider) *ListenerManager {
	return &ListenerManager{
		cloud: cloud,
	}
}

type ListenerManager struct {
	cloud prvd.Provider
}

// serverGroup find the vGroup id associated with the specific ServicePort
func serverGroup(annotation string, port v1.ServicePort) (string, error) {
	for _, v := range strings.Split(annotation, ",") {
		pp := strings.Split(v, ":")
		if len(pp) < 2 {
			return "", fmt.Errorf("server group id and "+
				"protocol format must be like 'sg-xxx:443' with colon separated. got=[%+v]", pp)
		}

		if pp[1] == fmt.Sprintf("%d", port.Port) {
			return pp[0], nil
		}
	}
	return "", nil
}

func (mgr *ListenerManager) BuildLocalModel(reqCtx *svcCtx.RequestContext, mdl *nlbmodel.NetworkLoadBalancer) error {
	for _, port := range reqCtx.Service.Spec.Ports {
		listener, err := mgr.buildListenerFromServicePort(reqCtx, port, mdl.LoadBalancerAttribute.IsUserManaged)
		if err != nil {
			return fmt.Errorf("build listener from servicePort %d error: %s", port.Port, err.Error())
		}
		mdl.Listeners = append(mdl.Listeners, listener)
	}
	return nil
}

func (mgr *ListenerManager) BuildRemoteModel(reqCtx *svcCtx.RequestContext, mdl *nlbmodel.NetworkLoadBalancer) error {
	listeners, err := mgr.ListListeners(reqCtx, mdl.LoadBalancerAttribute.LoadBalancerId)
	if err != nil {
		return fmt.Errorf("DescribeNLBListeners error:%s", err.Error())
	}
	mdl.Listeners = listeners
	return nil
}

func (mgr *ListenerManager) buildListenerFromServicePort(reqCtx *svcCtx.RequestContext, port v1.ServicePort,
	isUserManagedLB bool) (*nlbmodel.ListenerAttribute, error) {
	listener := &nlbmodel.ListenerAttribute{
		NamedKey: &nlbmodel.ListenerNamedKey{
			NamedKey: nlbmodel.NamedKey{
				Prefix:      model.DEFAULT_PREFIX,
				CID:         base.CLUSTER_ID,
				Namespace:   reqCtx.Service.Namespace,
				ServiceName: reqCtx.Service.Name,
			},
			Port: port.Port,
		},
		ServicePort:  &port,
		ListenerPort: port.Port,
	}

	proto, err := nlbListenerProtocol(reqCtx.Anno.Get(annotation.ProtocolPort), port)
	if err != nil {
		return listener, err
	}
	listener.ListenerProtocol = proto
	listener.NamedKey.Protocol = proto

	listener.ListenerDescription = listener.NamedKey.Key()
	listener.ServerGroupName = getServerGroupNamedKey(reqCtx.Service, proto, &port).Key()

	if isUserManagedLB && reqCtx.Anno.Get(annotation.VGroupPort) != "" {
		serverGroupId, err := serverGroup(reqCtx.Anno.Get(annotation.VGroupPort), port)
		if err != nil {
			return listener, err
		}
		listener.ServerGroupId = serverGroupId
	}

	if reqCtx.Anno.Get(annotation.IdleTimeout) != "" {
		idleTimeout, err := strconv.Atoi(reqCtx.Anno.Get(annotation.IdleTimeout))
		if err != nil {
			return listener, fmt.Errorf("parse IdleTimeout error: %s", err.Error())
		}
		listener.IdleTimeout = int32(idleTimeout)
	}
	if reqCtx.Anno.Get(annotation.TLSCipherPolicy) != "" {
		listener.SecurityPolicyId = reqCtx.Anno.Get(annotation.TLSCipherPolicy)
	}

	if reqCtx.Anno.Get(annotation.ProxyProtocol) != "" {
		listener.ProxyProtocolEnabled = tea.Bool(strings.EqualFold(reqCtx.Anno.Get(annotation.ProxyProtocol), string(model.OnFlag)))
	}
	if reqCtx.Anno.Get(annotation.CertID) != "" {
		listener.CertificateIds = strings.Split(reqCtx.Anno.Get(annotation.CertID), ",")
	}
	if reqCtx.Anno.Get(annotation.CaCertID) != "" {
		listener.CaCertificateIds = strings.Split(reqCtx.Anno.Get(annotation.CaCertID), ",")
	}
	if reqCtx.Anno.Get(annotation.CaCert) != "" {
		listener.CaEnabled = tea.Bool(strings.EqualFold(reqCtx.Anno.Get(annotation.CaCert), string(model.OnFlag)))
	}
	if reqCtx.Anno.Get(annotation.Cps) != "" {
		cps, err := strconv.Atoi(reqCtx.Anno.Get(annotation.Cps))
		if err != nil {
			return listener, fmt.Errorf("parse Mss error: %s", err.Error())
		}
		listener.Cps = tea.Int32(int32(cps))
	}

	if reqCtx.Anno.Get(annotation.Ppv2PrivateLinkEpIdEnabled) != "" {
		listener.ProxyProtocolV2Config.PrivateLinkEpIdEnabled = tea.Bool(strings.EqualFold(reqCtx.Anno.Get(annotation.Ppv2PrivateLinkEpIdEnabled), string(model.OnFlag)))
	}
	if reqCtx.Anno.Get(annotation.Ppv2PrivateLinkEpsIdEnabled) != "" {
		listener.ProxyProtocolV2Config.PrivateLinkEpsIdEnabled = tea.Bool(strings.EqualFold(reqCtx.Anno.Get(annotation.Ppv2PrivateLinkEpsIdEnabled), string(model.OnFlag)))
	}
	if reqCtx.Anno.Get(annotation.Ppv2VpcIdEnabled) != "" {
		listener.ProxyProtocolV2Config.VpcIdEnabled = tea.Bool(strings.EqualFold(reqCtx.Anno.Get(annotation.Ppv2VpcIdEnabled), string(model.OnFlag)))
	}

	if listener.ListenerProtocol == nlbmodel.TCPSSL {
		if reqCtx.Anno.Get(annotation.AlpnEnabled) != "" {
			listener.AlpnEnabled = tea.Bool(strings.EqualFold(reqCtx.Anno.Get(annotation.AlpnEnabled), string(model.OnFlag)))
		}

		if reqCtx.Anno.Get(annotation.AlpnPolicy) != "" {
			listener.AlpnPolicy = reqCtx.Anno.Get(annotation.AlpnPolicy)
		}
	}

	return listener, nil
}

func (mgr *ListenerManager) ListListeners(reqCtx *svcCtx.RequestContext, lbId string,
) ([]*nlbmodel.ListenerAttribute, error) {
	return mgr.cloud.ListNLBListeners(reqCtx.Ctx, lbId)
}

func (mgr *ListenerManager) CreateListener(reqCtx *svcCtx.RequestContext, lbId string, local *nlbmodel.ListenerAttribute) error {
	return mgr.cloud.CreateNLBListener(reqCtx.Ctx, lbId, local)
}

func (mgr *ListenerManager) UpdateNLBListener(reqCtx *svcCtx.RequestContext, local, remote *nlbmodel.ListenerAttribute,
) error {
	if remote.ListenerStatus == nlbmodel.StoppedListenerStatus {
		if err := mgr.cloud.StartNLBListener(reqCtx.Ctx, remote.ListenerId); err != nil {
			return fmt.Errorf("start listener %s error: %s", remote.ListenerId, err.Error())
		}
	}

	update := deepcopy.Copy(remote).(*nlbmodel.ListenerAttribute)
	needUpdate := false
	updateDetail := ""

	if remote.ListenerDescription != local.ListenerDescription {
		needUpdate = true
		update.ListenerDescription = local.ListenerDescription
		updateDetail += fmt.Sprintf("ListenerDescription %v should be changed to %v;",
			remote.ListenerDescription, local.ListenerDescription)
	}
	if remote.ServerGroupId != local.ServerGroupId {
		needUpdate = true
		update.ServerGroupId = local.ServerGroupId
		updateDetail += fmt.Sprintf("ServerGroupId %v should be changed to %v;",
			remote.ServerGroupId, local.ServerGroupId)
	}
	if local.ProxyProtocolEnabled != nil &&
		tea.BoolValue(remote.ProxyProtocolEnabled) != tea.BoolValue(local.ProxyProtocolEnabled) {
		needUpdate = true
		update.ProxyProtocolEnabled = local.ProxyProtocolEnabled
		updateDetail += fmt.Sprintf("ProxyProtocolEnabled %v should be changed to %v;",
			tea.BoolValue(remote.ProxyProtocolEnabled), tea.BoolValue(local.ProxyProtocolEnabled))
	}
	if local.ProxyProtocolV2Config.PrivateLinkEpIdEnabled != nil &&
		(remote.ProxyProtocolV2Config.PrivateLinkEpIdEnabled == nil || tea.BoolValue(remote.ProxyProtocolV2Config.PrivateLinkEpIdEnabled) != tea.BoolValue(local.ProxyProtocolV2Config.PrivateLinkEpIdEnabled)) {
		needUpdate = true
		update.ProxyProtocolV2Config.PrivateLinkEpIdEnabled = local.ProxyProtocolV2Config.PrivateLinkEpIdEnabled
		updateDetail += fmt.Sprintf("PrivateLinkEpIdEnabled %v should be changed to %v;",
			tea.BoolValue(remote.ProxyProtocolV2Config.PrivateLinkEpIdEnabled), tea.BoolValue(local.ProxyProtocolV2Config.PrivateLinkEpIdEnabled))
	}
	if local.ProxyProtocolV2Config.PrivateLinkEpsIdEnabled != nil &&
		(remote.ProxyProtocolV2Config.PrivateLinkEpsIdEnabled == nil || tea.BoolValue(remote.ProxyProtocolV2Config.PrivateLinkEpsIdEnabled) != tea.BoolValue(local.ProxyProtocolV2Config.PrivateLinkEpsIdEnabled)) {
		needUpdate = true
		update.ProxyProtocolV2Config.PrivateLinkEpsIdEnabled = local.ProxyProtocolV2Config.PrivateLinkEpsIdEnabled
		updateDetail += fmt.Sprintf("PrivateLinkEpsIdEnabled %v should be changed to %v;",
			tea.BoolValue(remote.ProxyProtocolV2Config.PrivateLinkEpsIdEnabled), tea.BoolValue(local.ProxyProtocolV2Config.PrivateLinkEpsIdEnabled))
	}
	if local.ProxyProtocolV2Config.VpcIdEnabled != nil &&
		(remote.ProxyProtocolV2Config.VpcIdEnabled == nil || tea.BoolValue(remote.ProxyProtocolV2Config.VpcIdEnabled) != tea.BoolValue(local.ProxyProtocolV2Config.VpcIdEnabled)) {
		needUpdate = true
		update.ProxyProtocolV2Config.VpcIdEnabled = local.ProxyProtocolV2Config.VpcIdEnabled
		updateDetail += fmt.Sprintf("VpcIdEnabled %v should be changed to %v;",
			tea.BoolValue(remote.ProxyProtocolV2Config.VpcIdEnabled), tea.BoolValue(local.ProxyProtocolV2Config.VpcIdEnabled))
	}
	// idle timeout
	if local.IdleTimeout != 0 && remote.IdleTimeout != local.IdleTimeout {
		needUpdate = true
		update.IdleTimeout = local.IdleTimeout
		updateDetail += fmt.Sprintf("IdleTimeout %v should be changed to %v;",
			remote.IdleTimeout, local.IdleTimeout)
	}
	if local.Cps != nil && tea.Int32Value(local.Cps) != tea.Int32Value(remote.Cps) {
		needUpdate = true
		update.Cps = local.Cps
		updateDetail += fmt.Sprintf("Cps %v should be changed to %v;", tea.Int32Value(remote.Cps), tea.Int32Value(local.Cps))
	}

	// only for TCPSSL protocol
	if isTCPSSL(local.ListenerProtocol) {
		// certId
		if len(local.CertificateIds) != 0 &&
			!util.IsStringSliceEqual(local.CertificateIds, remote.CertificateIds) {
			needUpdate = true
			update.CertificateIds = local.CertificateIds
			updateDetail += fmt.Sprintf("CertificateIds %v should be changed to %v;",
				remote.CertificateIds, local.CertificateIds)
		}
		// cacertId
		if len(local.CaCertificateIds) != 0 &&
			!util.IsStringSliceEqual(local.CaCertificateIds, remote.CaCertificateIds) {
			needUpdate = true
			update.CaCertificateIds = local.CaCertificateIds
			updateDetail += fmt.Sprintf("CaCertificateIds %v should be changed to %v;",
				remote.CaCertificateIds, local.CaCertificateIds)
		}
		if local.CaEnabled != nil &&
			tea.BoolValue(local.CaEnabled) != tea.BoolValue(remote.CaEnabled) {
			needUpdate = true
			update.CaEnabled = local.CaEnabled
			updateDetail += fmt.Sprintf("CaEnabled %v should be changed to %v;", tea.BoolValue(remote.CaEnabled),
				tea.BoolValue(local.CaEnabled))
		}
		if local.SecurityPolicyId != "" &&
			local.SecurityPolicyId != remote.SecurityPolicyId {
			needUpdate = true
			update.SecurityPolicyId = local.SecurityPolicyId
			updateDetail += fmt.Sprintf("SecurityPolicyId %v should be changed to %v;",
				remote.SecurityPolicyId, local.SecurityPolicyId)
		}
		if local.AlpnEnabled != nil &&
			*local.AlpnEnabled != tea.BoolValue(remote.AlpnEnabled) {
			needUpdate = true
			update.AlpnEnabled = local.AlpnEnabled
			updateDetail += fmt.Sprintf("AlpnEnabled %v should be changed to %v;", tea.BoolValue(remote.AlpnEnabled),
				tea.BoolValue(local.AlpnEnabled))
		}
		if tea.BoolValue(local.AlpnEnabled) && local.AlpnPolicy != "" &&
			local.AlpnPolicy != remote.AlpnPolicy {
			needUpdate = true
			update.AlpnPolicy = local.AlpnPolicy
			updateDetail += fmt.Sprintf("AlpnPolicy %v should be changed to %v;",
				remote.AlpnPolicy, local.AlpnPolicy)
		}
	}

	if needUpdate {
		reqCtx.Ctx = context.WithValue(reqCtx.Ctx, dryrun.ContextMessage, updateDetail)
		reqCtx.Log.Info(fmt.Sprintf("update listener: %s [%d] changed, detail %s", local.ListenerProtocol, local.ListenerPort, updateDetail))

		return mgr.cloud.UpdateNLBListener(reqCtx.Ctx, update)
	}

	reqCtx.Log.Info(fmt.Sprintf("update listener: %s [%d] not changed, skip", local.ListenerProtocol, local.ListenerPort))
	return nil
}

func (mgr *ListenerManager) DeleteListener(reqCtx *svcCtx.RequestContext, lisId string) error {
	return mgr.cloud.DeleteNLBListener(reqCtx.Ctx, lisId)
}

func nlbListenerProtocol(annotation string, port v1.ServicePort) (string, error) {

	if annotation == "" {
		return strings.ToUpper(string(port.Protocol)), nil
	}
	for _, v := range strings.Split(annotation, ",") {
		pp := strings.Split(v, ":")
		if len(pp) < 2 {
			return "", fmt.Errorf("port and "+
				"protocol format must be like 'https:443' with colon separated. got=[%+v]", pp)
		}

		if strings.ToUpper(pp[0]) != nlbmodel.TCP &&
			strings.ToUpper(pp[0]) != nlbmodel.UDP &&
			strings.ToUpper(pp[0]) != nlbmodel.TCPSSL {
			return "", fmt.Errorf("port protocol"+
				" format must be either [TCP|UDP|TCPSSL], protocol not supported wit [%s]\n", pp[0])
		}

		if pp[1] == fmt.Sprintf("%d", port.Port) {
			util.NLBLog.Info(fmt.Sprintf("port [%d] transform protocol from %s to %s", port.Port, port.Protocol, strings.ToUpper(pp[0])))
			return strings.ToUpper(pp[0]), nil
		}
	}
	return strings.ToUpper(string(port.Protocol)), nil
}

func isTCPSSL(proto string) bool {
	return proto == nlbmodel.TCPSSL
}
