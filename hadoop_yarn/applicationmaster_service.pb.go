// Code generated by protoc-gen-go.
// source: applicationmaster_protocol.proto
// DO NOT EDIT!

package hadoop_yarn

import proto "github.com/golang/protobuf/proto"
import json "encoding/json"
import math "math"

import "github.com/hortonworks/gohadoop"
import hadoop_ipc_client "github.com/hortonworks/gohadoop/hadoop_common/ipc/client"
import yarn_conf "github.com/hortonworks/gohadoop/hadoop_yarn/conf"
import "github.com/nu7hatch/gouuid"

// Reference proto, json, and math imports to suppress error if they are not otherwise used.
var _ = proto.Marshal
var _ = &json.SyntaxError{}
var _ = math.Inf

var APPLICATION_MASTER_PROTOCOL = "org.apache.hadoop.yarn.api.ApplicationMasterProtocolPB"

func init() {
}

type ApplicationMasterProtocolService interface {
	RegisterApplicationMaster(in *RegisterApplicationMasterRequestProto, out *RegisterApplicationMasterResponseProto) error
	FinishApplicationMaster(in *FinishApplicationMasterRequestProto, out *FinishApplicationMasterResponseProto) error
	Allocate(in *AllocateRequestProto, out *AllocateResponseProto) error
}

type ApplicationMasterProtocolServiceClient struct {
	*hadoop_ipc_client.Client
}

func (c *ApplicationMasterProtocolServiceClient) RegisterApplicationMaster(in *RegisterApplicationMasterRequestProto, out *RegisterApplicationMasterResponseProto) error {
	return c.Call(gohadoop.GetCalleeRPCRequestHeaderProto(&APPLICATION_MASTER_PROTOCOL), in, out)
}
func (c *ApplicationMasterProtocolServiceClient) FinishApplicationMaster(in *FinishApplicationMasterRequestProto, out *FinishApplicationMasterResponseProto) error {
	return c.Call(gohadoop.GetCalleeRPCRequestHeaderProto(&APPLICATION_MASTER_PROTOCOL), in, out)
}
func (c *ApplicationMasterProtocolServiceClient) Allocate(in *AllocateRequestProto, out *AllocateResponseProto) error {
	return c.Call(gohadoop.GetCalleeRPCRequestHeaderProto(&APPLICATION_MASTER_PROTOCOL), in, out)
}

func DialApplicationMasterProtocolService(conf yarn_conf.YarnConfiguration) (ApplicationMasterProtocolService, error) {
  clientId, _ := uuid.NewV4()
  ugi, _ := gohadoop.CreateSimpleUGIProto()
  serverAddress, _ := conf.GetRMSchedulerAddress()
  c := &hadoop_ipc_client.Client{ClientId: clientId, Ugi: ugi, ServerAddress: serverAddress}
	return &ApplicationMasterProtocolServiceClient{c}, nil
}

/*
// DialApplicationMasterProtocolService connects to an ApplicationMasterProtocolService at the specified network address.
// DialApplicationMasterProtocolServiceTimeout connects to an ApplicationMasterProtocolService at the specified network address.
func DialApplicationMasterProtocolServiceTimeout(network, addr string,
	timeout time.Duration) (*ApplicationMasterProtocolServiceClient, *rpc.Client, error) {
	c, err := protorpc.DialTimeout(network, addr, timeout)
	if err != nil {
		return nil, nil, err
	}
	return &ApplicationMasterProtocolServiceClient{c}, c, nil
}
*/
