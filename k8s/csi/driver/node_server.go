package driver

import (
	"context"
	"github.com/container-storage-interface/spec/lib/go/csi"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"k8s.io/klog"
	"k8s.io/utils/mount"
)

type NodeServer struct {
	nodeId  string
	mounter mount.Interface
}

func NewNodeServer(nodeId string) *NodeServer {
	return &NodeServer{
		nodeId:  nodeId,
		mounter: mount.New(""),
	}
}

// 如果使用云盘，
// 就会将云硬盘格式化成对应文件系统 将volume mount到一个全局的目录
func (n *NodeServer) NodeStageVolume(ctx context.Context, req *csi.NodeStageVolumeRequest) (*csi.NodeStageVolumeResponse, error) {
	klog.V(4).Infof("NodeStageVolume: called with args %+v", *req)

	return &csi.NodeStageVolumeResponse{}, nil
}

func (n *NodeServer) NodeUnstageVolume(ctx context.Context, req *csi.NodeUnstageVolumeRequest) (*csi.NodeUnstageVolumeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NodeUnstageVolume not implemented")
}
func (n *NodeServer) NodePublishVolume(ctx context.Context, req *csi.NodePublishVolumeRequest) (*csi.NodePublishVolumeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NodePublishVolume not implemented")
}
func (n *NodeServer) NodeUnpublishVolume(ctx context.Context, req *csi.NodeUnpublishVolumeRequest) (*csi.NodeUnpublishVolumeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NodeUnpublishVolume not implemented")
}
func (n *NodeServer) NodeGetVolumeStats(ctx context.Context, req *csi.NodeGetVolumeStatsRequest) (*csi.NodeGetVolumeStatsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NodeGetVolumeStats not implemented")
}
func (n *NodeServer) NodeExpandVolume(ctx context.Context, req *csi.NodeExpandVolumeRequest) (*csi.NodeExpandVolumeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NodeExpandVolume not implemented")
}
func (n *NodeServer) NodeGetCapabilities(ctx context.Context, req *csi.NodeGetCapabilitiesRequest) (*csi.NodeGetCapabilitiesResponse, error) {
	klog.V(4).Infof("NodeGetCapabilities: called with args %+v", *req)
	
	return &csi.NodeGetCapabilitiesResponse{
		Capabilities: []*csi.NodeServiceCapability{
			{
				Type: &csi.NodeServiceCapability_Rpc{
					Rpc: &csi.NodeServiceCapability_RPC{
						Type: csi.NodeServiceCapability_RPC_STAGE_UNSTAGE_VOLUME,
					},
				},
			},
			{
				Type: &csi.NodeServiceCapability_Rpc{
					Rpc: &csi.NodeServiceCapability_RPC{
						Type: csi.NodeServiceCapability_RPC_UNKNOWN,
					},
				},
			},
			{
				Type: &csi.NodeServiceCapability_Rpc{
					Rpc: &csi.NodeServiceCapability_RPC{
						Type: csi.NodeServiceCapability_RPC_GET_VOLUME_STATS,
					},
				},
			},
			{
				Type: &csi.NodeServiceCapability_Rpc{
					Rpc: &csi.NodeServiceCapability_RPC{
						Type: csi.NodeServiceCapability_RPC_SINGLE_NODE_MULTI_WRITER,
					},
				},
			},
		},
	},nil

}
func (n *NodeServer) NodeGetInfo(ctx context.Context, req *csi.NodeGetInfoRequest) (*csi.NodeGetInfoResponse, error) {
	klog.Infoln("NodeGetInfo")
	return &csi.NodeGetInfoResponse{
		NodeId: n.nodeId,
	}, nil
}
