package driver

import (
	"context"
	"github.com/container-storage-interface/spec/lib/go/csi"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"k8s.io/klog"
)

var (
	controllerCaps = []csi.ControllerServiceCapability_RPC_Type{
		csi.ControllerServiceCapability_RPC_CREATE_DELETE_VOLUME,     // 创建、删除 volume
		csi.ControllerServiceCapability_RPC_PUBLISH_UNPUBLISH_VOLUME, // 挂载、卸载 volume
	}
)

type ControllerServer struct {
}

func NewControllerServer() *ControllerServer {
	return &ControllerServer{}
}

func (s *ControllerServer) DeleteVolume(ctx context.Context, request *csi.DeleteVolumeRequest) (*csi.DeleteVolumeResponse, error) {
	klog.V(4).Infof("DeleteVolume: called with args: %+v", *request)
	return &csi.DeleteVolumeResponse{}, nil
}

func (s *ControllerServer) ControllerPublishVolume(ctx context.Context, request *csi.ControllerPublishVolumeRequest) (*csi.ControllerPublishVolumeResponse, error) {
	klog.Info("发布PublishVolume")
	//TODO implement me
	return &csi.ControllerPublishVolumeResponse{}, nil
}

func (s *ControllerServer) ControllerUnpublishVolume(ctx context.Context, request *csi.ControllerUnpublishVolumeRequest) (*csi.ControllerUnpublishVolumeResponse, error) {
	//TODO implement me
	klog.Info("执行UnpublishVolume")
	return &csi.ControllerUnpublishVolumeResponse{}, nil
}

func (s *ControllerServer) ValidateVolumeCapabilities(ctx context.Context, request *csi.ValidateVolumeCapabilitiesRequest) (*csi.ValidateVolumeCapabilitiesResponse, error) {
	//TODO implement me
	return nil, status.Errorf(codes.Unimplemented, "method ValidateVolumeCapabilities not implemented")
}

func (s *ControllerServer) ListVolumes(ctx context.Context, request *csi.ListVolumesRequest) (*csi.ListVolumesResponse, error) {
	klog.Info("列出volume")
	//TODO implement me
	return &csi.ListVolumesResponse{}, nil
}

func (s *ControllerServer) GetCapacity(ctx context.Context, request *csi.GetCapacityRequest) (*csi.GetCapacityResponse, error) {
	return &csi.GetCapacityResponse{
		AvailableCapacity: 100 * 1024 * 1024,
	}, nil
}

func (s *ControllerServer) ControllerGetCapabilities(ctx context.Context, request *csi.ControllerGetCapabilitiesRequest) (*csi.ControllerGetCapabilitiesResponse, error) {
	klog.V(4).Infof("ControllerGetCapabilities: called with args %+v", *request)

	var caps []*csi.ControllerServiceCapability
	for _, ca := range controllerCaps {
		c := &csi.ControllerServiceCapability{
			Type: &csi.ControllerServiceCapability_Rpc{
				Rpc: &csi.ControllerServiceCapability_RPC{
					Type: ca,
				},
			},
		}
		caps = append(caps, c)
	}

	return &csi.ControllerGetCapabilitiesResponse{Capabilities: caps}, nil
}

func (s *ControllerServer) CreateSnapshot(ctx context.Context, request *csi.CreateSnapshotRequest) (*csi.CreateSnapshotResponse, error) {
	//TODO implement me
	return nil, status.Errorf(codes.Unimplemented, "method CreateSnapshot not implemented")
}

func (s *ControllerServer) DeleteSnapshot(ctx context.Context, request *csi.DeleteSnapshotRequest) (*csi.DeleteSnapshotResponse, error) {
	//TODO implement me
	return nil, status.Errorf(codes.Unimplemented, "method DeleteSnapshot not implemented")
}

func (s *ControllerServer) ListSnapshots(ctx context.Context, request *csi.ListSnapshotsRequest) (*csi.ListSnapshotsResponse, error) {
	//TODO implement me
	return nil, status.Errorf(codes.Unimplemented, "method ListSnapshots not implemented")
}

func (s *ControllerServer) ControllerExpandVolume(ctx context.Context, request *csi.ControllerExpandVolumeRequest) (*csi.ControllerExpandVolumeResponse, error) {
	//TODO implement me
	return nil, status.Errorf(codes.Unimplemented, "method ControllerExpandVolume not implemented")
}

func (s *ControllerServer) ControllerGetVolume(ctx context.Context, request *csi.ControllerGetVolumeRequest) (*csi.ControllerGetVolumeResponse, error) {
	klog.V(4).Infof("ControllerGetVolume: called with args: %+v", *request)
	return &csi.ControllerGetVolumeResponse{}, nil
}

func (*ControllerServer) CreateVolume(ctx context.Context, req *csi.CreateVolumeRequest) (*csi.CreateVolumeResponse, error) {
	klog.V(4).Infof("CreateVolume: called with args %+v", *req)

	// 这里先返回一个假数据，模拟我们创建出了一块id为"jhzhang-123456"容量为20G的云盘
	return &csi.CreateVolumeResponse{
		Volume: &csi.Volume{
			CapacityBytes: 100 * 1024 * 1024,
			VolumeId:      "jhzhang-123456",
			VolumeContext: req.GetParameters(),
		},
	}, nil
}

