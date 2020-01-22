// Copyright 2016-2020 The Libsacloud Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// generated by 'github.com/sacloud/libsacloud/internal/tools/gen-fake-store'; DO NOT EDIT

package fake

import (
	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/libsacloud/v2/sacloud/accessor"
	"github.com/sacloud/libsacloud/v2/sacloud/types"
)

func getArchive(zone string) []*sacloud.Archive {
	values := ds().List(ResourceArchive, zone)
	var ret []*sacloud.Archive
	for _, v := range values {
		if v, ok := v.(*sacloud.Archive); ok {
			ret = append(ret, v)
		}
	}
	return ret
}

func getArchiveByID(zone string, id types.ID) *sacloud.Archive {
	v := ds().Get(ResourceArchive, zone, id)
	if v, ok := v.(*sacloud.Archive); ok {
		return v
	}
	return nil
}

func putArchive(zone string, value *sacloud.Archive) {
	var v interface{} = value
	if id, ok := v.(accessor.ID); ok {
		ds().Put(ResourceArchive, zone, id.GetID(), value)
		return
	}
	ds().Put(ResourceArchive, zone, 0, value)
}

func getAuthStatus(zone string) []*sacloud.AuthStatus {
	values := ds().List(ResourceAuthStatus, zone)
	var ret []*sacloud.AuthStatus
	for _, v := range values {
		if v, ok := v.(*sacloud.AuthStatus); ok {
			ret = append(ret, v)
		}
	}
	return ret
}

func getAuthStatusByID(zone string, id types.ID) *sacloud.AuthStatus {
	v := ds().Get(ResourceAuthStatus, zone, id)
	if v, ok := v.(*sacloud.AuthStatus); ok {
		return v
	}
	return nil
}

func putAuthStatus(zone string, value *sacloud.AuthStatus) {
	var v interface{} = value
	if id, ok := v.(accessor.ID); ok {
		ds().Put(ResourceAuthStatus, zone, id.GetID(), value)
		return
	}
	ds().Put(ResourceAuthStatus, zone, 0, value)
}

func getAutoBackup(zone string) []*sacloud.AutoBackup {
	values := ds().List(ResourceAutoBackup, zone)
	var ret []*sacloud.AutoBackup
	for _, v := range values {
		if v, ok := v.(*sacloud.AutoBackup); ok {
			ret = append(ret, v)
		}
	}
	return ret
}

func getAutoBackupByID(zone string, id types.ID) *sacloud.AutoBackup {
	v := ds().Get(ResourceAutoBackup, zone, id)
	if v, ok := v.(*sacloud.AutoBackup); ok {
		return v
	}
	return nil
}

func putAutoBackup(zone string, value *sacloud.AutoBackup) {
	var v interface{} = value
	if id, ok := v.(accessor.ID); ok {
		ds().Put(ResourceAutoBackup, zone, id.GetID(), value)
		return
	}
	ds().Put(ResourceAutoBackup, zone, 0, value)
}

func getBill(zone string) []*sacloud.Bill {
	values := ds().List(ResourceBill, zone)
	var ret []*sacloud.Bill
	for _, v := range values {
		if v, ok := v.(*sacloud.Bill); ok {
			ret = append(ret, v)
		}
	}
	return ret
}

func getBillByID(zone string, id types.ID) *sacloud.Bill {
	v := ds().Get(ResourceBill, zone, id)
	if v, ok := v.(*sacloud.Bill); ok {
		return v
	}
	return nil
}

func putBill(zone string, value *sacloud.Bill) {
	var v interface{} = value
	if id, ok := v.(accessor.ID); ok {
		ds().Put(ResourceBill, zone, id.GetID(), value)
		return
	}
	ds().Put(ResourceBill, zone, 0, value)
}

func getBridge(zone string) []*sacloud.Bridge {
	values := ds().List(ResourceBridge, zone)
	var ret []*sacloud.Bridge
	for _, v := range values {
		if v, ok := v.(*sacloud.Bridge); ok {
			ret = append(ret, v)
		}
	}
	return ret
}

func getBridgeByID(zone string, id types.ID) *sacloud.Bridge {
	v := ds().Get(ResourceBridge, zone, id)
	if v, ok := v.(*sacloud.Bridge); ok {
		return v
	}
	return nil
}

func putBridge(zone string, value *sacloud.Bridge) {
	var v interface{} = value
	if id, ok := v.(accessor.ID); ok {
		ds().Put(ResourceBridge, zone, id.GetID(), value)
		return
	}
	ds().Put(ResourceBridge, zone, 0, value)
}

func getCDROM(zone string) []*sacloud.CDROM {
	values := ds().List(ResourceCDROM, zone)
	var ret []*sacloud.CDROM
	for _, v := range values {
		if v, ok := v.(*sacloud.CDROM); ok {
			ret = append(ret, v)
		}
	}
	return ret
}

func getCDROMByID(zone string, id types.ID) *sacloud.CDROM {
	v := ds().Get(ResourceCDROM, zone, id)
	if v, ok := v.(*sacloud.CDROM); ok {
		return v
	}
	return nil
}

func putCDROM(zone string, value *sacloud.CDROM) {
	var v interface{} = value
	if id, ok := v.(accessor.ID); ok {
		ds().Put(ResourceCDROM, zone, id.GetID(), value)
		return
	}
	ds().Put(ResourceCDROM, zone, 0, value)
}

func getContainerRegistry(zone string) []*sacloud.ContainerRegistry {
	values := ds().List(ResourceContainerRegistry, zone)
	var ret []*sacloud.ContainerRegistry
	for _, v := range values {
		if v, ok := v.(*sacloud.ContainerRegistry); ok {
			ret = append(ret, v)
		}
	}
	return ret
}

func getContainerRegistryByID(zone string, id types.ID) *sacloud.ContainerRegistry {
	v := ds().Get(ResourceContainerRegistry, zone, id)
	if v, ok := v.(*sacloud.ContainerRegistry); ok {
		return v
	}
	return nil
}

func putContainerRegistry(zone string, value *sacloud.ContainerRegistry) {
	var v interface{} = value
	if id, ok := v.(accessor.ID); ok {
		ds().Put(ResourceContainerRegistry, zone, id.GetID(), value)
		return
	}
	ds().Put(ResourceContainerRegistry, zone, 0, value)
}

func getCoupon(zone string) []*sacloud.Coupon {
	values := ds().List(ResourceCoupon, zone)
	var ret []*sacloud.Coupon
	for _, v := range values {
		if v, ok := v.(*sacloud.Coupon); ok {
			ret = append(ret, v)
		}
	}
	return ret
}

func getCouponByID(zone string, id types.ID) *sacloud.Coupon {
	v := ds().Get(ResourceCoupon, zone, id)
	if v, ok := v.(*sacloud.Coupon); ok {
		return v
	}
	return nil
}

func putCoupon(zone string, value *sacloud.Coupon) {
	var v interface{} = value
	if id, ok := v.(accessor.ID); ok {
		ds().Put(ResourceCoupon, zone, id.GetID(), value)
		return
	}
	ds().Put(ResourceCoupon, zone, 0, value)
}

func getDatabase(zone string) []*sacloud.Database {
	values := ds().List(ResourceDatabase, zone)
	var ret []*sacloud.Database
	for _, v := range values {
		if v, ok := v.(*sacloud.Database); ok {
			ret = append(ret, v)
		}
	}
	return ret
}

func getDatabaseByID(zone string, id types.ID) *sacloud.Database {
	v := ds().Get(ResourceDatabase, zone, id)
	if v, ok := v.(*sacloud.Database); ok {
		return v
	}
	return nil
}

func putDatabase(zone string, value *sacloud.Database) {
	var v interface{} = value
	if id, ok := v.(accessor.ID); ok {
		ds().Put(ResourceDatabase, zone, id.GetID(), value)
		return
	}
	ds().Put(ResourceDatabase, zone, 0, value)
}

func getDisk(zone string) []*sacloud.Disk {
	values := ds().List(ResourceDisk, zone)
	var ret []*sacloud.Disk
	for _, v := range values {
		if v, ok := v.(*sacloud.Disk); ok {
			ret = append(ret, v)
		}
	}
	return ret
}

func getDiskByID(zone string, id types.ID) *sacloud.Disk {
	v := ds().Get(ResourceDisk, zone, id)
	if v, ok := v.(*sacloud.Disk); ok {
		return v
	}
	return nil
}

func putDisk(zone string, value *sacloud.Disk) {
	var v interface{} = value
	if id, ok := v.(accessor.ID); ok {
		ds().Put(ResourceDisk, zone, id.GetID(), value)
		return
	}
	ds().Put(ResourceDisk, zone, 0, value)
}

func getDiskPlan(zone string) []*sacloud.DiskPlan {
	values := ds().List(ResourceDiskPlan, zone)
	var ret []*sacloud.DiskPlan
	for _, v := range values {
		if v, ok := v.(*sacloud.DiskPlan); ok {
			ret = append(ret, v)
		}
	}
	return ret
}

func getDiskPlanByID(zone string, id types.ID) *sacloud.DiskPlan {
	v := ds().Get(ResourceDiskPlan, zone, id)
	if v, ok := v.(*sacloud.DiskPlan); ok {
		return v
	}
	return nil
}

func putDiskPlan(zone string, value *sacloud.DiskPlan) {
	var v interface{} = value
	if id, ok := v.(accessor.ID); ok {
		ds().Put(ResourceDiskPlan, zone, id.GetID(), value)
		return
	}
	ds().Put(ResourceDiskPlan, zone, 0, value)
}

func getDNS(zone string) []*sacloud.DNS {
	values := ds().List(ResourceDNS, zone)
	var ret []*sacloud.DNS
	for _, v := range values {
		if v, ok := v.(*sacloud.DNS); ok {
			ret = append(ret, v)
		}
	}
	return ret
}

func getDNSByID(zone string, id types.ID) *sacloud.DNS {
	v := ds().Get(ResourceDNS, zone, id)
	if v, ok := v.(*sacloud.DNS); ok {
		return v
	}
	return nil
}

func putDNS(zone string, value *sacloud.DNS) {
	var v interface{} = value
	if id, ok := v.(accessor.ID); ok {
		ds().Put(ResourceDNS, zone, id.GetID(), value)
		return
	}
	ds().Put(ResourceDNS, zone, 0, value)
}

func getGSLB(zone string) []*sacloud.GSLB {
	values := ds().List(ResourceGSLB, zone)
	var ret []*sacloud.GSLB
	for _, v := range values {
		if v, ok := v.(*sacloud.GSLB); ok {
			ret = append(ret, v)
		}
	}
	return ret
}

func getGSLBByID(zone string, id types.ID) *sacloud.GSLB {
	v := ds().Get(ResourceGSLB, zone, id)
	if v, ok := v.(*sacloud.GSLB); ok {
		return v
	}
	return nil
}

func putGSLB(zone string, value *sacloud.GSLB) {
	var v interface{} = value
	if id, ok := v.(accessor.ID); ok {
		ds().Put(ResourceGSLB, zone, id.GetID(), value)
		return
	}
	ds().Put(ResourceGSLB, zone, 0, value)
}

func getIcon(zone string) []*sacloud.Icon {
	values := ds().List(ResourceIcon, zone)
	var ret []*sacloud.Icon
	for _, v := range values {
		if v, ok := v.(*sacloud.Icon); ok {
			ret = append(ret, v)
		}
	}
	return ret
}

func getIconByID(zone string, id types.ID) *sacloud.Icon {
	v := ds().Get(ResourceIcon, zone, id)
	if v, ok := v.(*sacloud.Icon); ok {
		return v
	}
	return nil
}

func putIcon(zone string, value *sacloud.Icon) {
	var v interface{} = value
	if id, ok := v.(accessor.ID); ok {
		ds().Put(ResourceIcon, zone, id.GetID(), value)
		return
	}
	ds().Put(ResourceIcon, zone, 0, value)
}

func getInterface(zone string) []*sacloud.Interface {
	values := ds().List(ResourceInterface, zone)
	var ret []*sacloud.Interface
	for _, v := range values {
		if v, ok := v.(*sacloud.Interface); ok {
			ret = append(ret, v)
		}
	}
	return ret
}

func getInterfaceByID(zone string, id types.ID) *sacloud.Interface {
	v := ds().Get(ResourceInterface, zone, id)
	if v, ok := v.(*sacloud.Interface); ok {
		return v
	}
	return nil
}

func putInterface(zone string, value *sacloud.Interface) {
	var v interface{} = value
	if id, ok := v.(accessor.ID); ok {
		ds().Put(ResourceInterface, zone, id.GetID(), value)
		return
	}
	ds().Put(ResourceInterface, zone, 0, value)
}

func getInternet(zone string) []*sacloud.Internet {
	values := ds().List(ResourceInternet, zone)
	var ret []*sacloud.Internet
	for _, v := range values {
		if v, ok := v.(*sacloud.Internet); ok {
			ret = append(ret, v)
		}
	}
	return ret
}

func getInternetByID(zone string, id types.ID) *sacloud.Internet {
	v := ds().Get(ResourceInternet, zone, id)
	if v, ok := v.(*sacloud.Internet); ok {
		return v
	}
	return nil
}

func putInternet(zone string, value *sacloud.Internet) {
	var v interface{} = value
	if id, ok := v.(accessor.ID); ok {
		ds().Put(ResourceInternet, zone, id.GetID(), value)
		return
	}
	ds().Put(ResourceInternet, zone, 0, value)
}

func getInternetPlan(zone string) []*sacloud.InternetPlan {
	values := ds().List(ResourceInternetPlan, zone)
	var ret []*sacloud.InternetPlan
	for _, v := range values {
		if v, ok := v.(*sacloud.InternetPlan); ok {
			ret = append(ret, v)
		}
	}
	return ret
}

func getInternetPlanByID(zone string, id types.ID) *sacloud.InternetPlan {
	v := ds().Get(ResourceInternetPlan, zone, id)
	if v, ok := v.(*sacloud.InternetPlan); ok {
		return v
	}
	return nil
}

func putInternetPlan(zone string, value *sacloud.InternetPlan) {
	var v interface{} = value
	if id, ok := v.(accessor.ID); ok {
		ds().Put(ResourceInternetPlan, zone, id.GetID(), value)
		return
	}
	ds().Put(ResourceInternetPlan, zone, 0, value)
}

func getIPAddress(zone string) []*sacloud.IPAddress {
	values := ds().List(ResourceIPAddress, zone)
	var ret []*sacloud.IPAddress
	for _, v := range values {
		if v, ok := v.(*sacloud.IPAddress); ok {
			ret = append(ret, v)
		}
	}
	return ret
}

func getIPAddressByID(zone string, id types.ID) *sacloud.IPAddress {
	v := ds().Get(ResourceIPAddress, zone, id)
	if v, ok := v.(*sacloud.IPAddress); ok {
		return v
	}
	return nil
}

func putIPAddress(zone string, value *sacloud.IPAddress) {
	var v interface{} = value
	if id, ok := v.(accessor.ID); ok {
		ds().Put(ResourceIPAddress, zone, id.GetID(), value)
		return
	}
	ds().Put(ResourceIPAddress, zone, 0, value)
}

func getIPv6Net(zone string) []*sacloud.IPv6Net {
	values := ds().List(ResourceIPv6Net, zone)
	var ret []*sacloud.IPv6Net
	for _, v := range values {
		if v, ok := v.(*sacloud.IPv6Net); ok {
			ret = append(ret, v)
		}
	}
	return ret
}

func getIPv6NetByID(zone string, id types.ID) *sacloud.IPv6Net {
	v := ds().Get(ResourceIPv6Net, zone, id)
	if v, ok := v.(*sacloud.IPv6Net); ok {
		return v
	}
	return nil
}

func putIPv6Net(zone string, value *sacloud.IPv6Net) {
	var v interface{} = value
	if id, ok := v.(accessor.ID); ok {
		ds().Put(ResourceIPv6Net, zone, id.GetID(), value)
		return
	}
	ds().Put(ResourceIPv6Net, zone, 0, value)
}

func getIPv6Addr(zone string) []*sacloud.IPv6Addr {
	values := ds().List(ResourceIPv6Addr, zone)
	var ret []*sacloud.IPv6Addr
	for _, v := range values {
		if v, ok := v.(*sacloud.IPv6Addr); ok {
			ret = append(ret, v)
		}
	}
	return ret
}

func getIPv6AddrByID(zone string, id types.ID) *sacloud.IPv6Addr {
	v := ds().Get(ResourceIPv6Addr, zone, id)
	if v, ok := v.(*sacloud.IPv6Addr); ok {
		return v
	}
	return nil
}

func putIPv6Addr(zone string, value *sacloud.IPv6Addr) {
	var v interface{} = value
	if id, ok := v.(accessor.ID); ok {
		ds().Put(ResourceIPv6Addr, zone, id.GetID(), value)
		return
	}
	ds().Put(ResourceIPv6Addr, zone, 0, value)
}

func getLicense(zone string) []*sacloud.License {
	values := ds().List(ResourceLicense, zone)
	var ret []*sacloud.License
	for _, v := range values {
		if v, ok := v.(*sacloud.License); ok {
			ret = append(ret, v)
		}
	}
	return ret
}

func getLicenseByID(zone string, id types.ID) *sacloud.License {
	v := ds().Get(ResourceLicense, zone, id)
	if v, ok := v.(*sacloud.License); ok {
		return v
	}
	return nil
}

func putLicense(zone string, value *sacloud.License) {
	var v interface{} = value
	if id, ok := v.(accessor.ID); ok {
		ds().Put(ResourceLicense, zone, id.GetID(), value)
		return
	}
	ds().Put(ResourceLicense, zone, 0, value)
}

func getLicenseInfo(zone string) []*sacloud.LicenseInfo {
	values := ds().List(ResourceLicenseInfo, zone)
	var ret []*sacloud.LicenseInfo
	for _, v := range values {
		if v, ok := v.(*sacloud.LicenseInfo); ok {
			ret = append(ret, v)
		}
	}
	return ret
}

func getLicenseInfoByID(zone string, id types.ID) *sacloud.LicenseInfo {
	v := ds().Get(ResourceLicenseInfo, zone, id)
	if v, ok := v.(*sacloud.LicenseInfo); ok {
		return v
	}
	return nil
}

func putLicenseInfo(zone string, value *sacloud.LicenseInfo) {
	var v interface{} = value
	if id, ok := v.(accessor.ID); ok {
		ds().Put(ResourceLicenseInfo, zone, id.GetID(), value)
		return
	}
	ds().Put(ResourceLicenseInfo, zone, 0, value)
}

func getLoadBalancer(zone string) []*sacloud.LoadBalancer {
	values := ds().List(ResourceLoadBalancer, zone)
	var ret []*sacloud.LoadBalancer
	for _, v := range values {
		if v, ok := v.(*sacloud.LoadBalancer); ok {
			ret = append(ret, v)
		}
	}
	return ret
}

func getLoadBalancerByID(zone string, id types.ID) *sacloud.LoadBalancer {
	v := ds().Get(ResourceLoadBalancer, zone, id)
	if v, ok := v.(*sacloud.LoadBalancer); ok {
		return v
	}
	return nil
}

func putLoadBalancer(zone string, value *sacloud.LoadBalancer) {
	var v interface{} = value
	if id, ok := v.(accessor.ID); ok {
		ds().Put(ResourceLoadBalancer, zone, id.GetID(), value)
		return
	}
	ds().Put(ResourceLoadBalancer, zone, 0, value)
}

func getLocalRouter(zone string) []*sacloud.LocalRouter {
	values := ds().List(ResourceLocalRouter, zone)
	var ret []*sacloud.LocalRouter
	for _, v := range values {
		if v, ok := v.(*sacloud.LocalRouter); ok {
			ret = append(ret, v)
		}
	}
	return ret
}

func getLocalRouterByID(zone string, id types.ID) *sacloud.LocalRouter {
	v := ds().Get(ResourceLocalRouter, zone, id)
	if v, ok := v.(*sacloud.LocalRouter); ok {
		return v
	}
	return nil
}

func putLocalRouter(zone string, value *sacloud.LocalRouter) {
	var v interface{} = value
	if id, ok := v.(accessor.ID); ok {
		ds().Put(ResourceLocalRouter, zone, id.GetID(), value)
		return
	}
	ds().Put(ResourceLocalRouter, zone, 0, value)
}

func getMobileGateway(zone string) []*sacloud.MobileGateway {
	values := ds().List(ResourceMobileGateway, zone)
	var ret []*sacloud.MobileGateway
	for _, v := range values {
		if v, ok := v.(*sacloud.MobileGateway); ok {
			ret = append(ret, v)
		}
	}
	return ret
}

func getMobileGatewayByID(zone string, id types.ID) *sacloud.MobileGateway {
	v := ds().Get(ResourceMobileGateway, zone, id)
	if v, ok := v.(*sacloud.MobileGateway); ok {
		return v
	}
	return nil
}

func putMobileGateway(zone string, value *sacloud.MobileGateway) {
	var v interface{} = value
	if id, ok := v.(accessor.ID); ok {
		ds().Put(ResourceMobileGateway, zone, id.GetID(), value)
		return
	}
	ds().Put(ResourceMobileGateway, zone, 0, value)
}

func getNFS(zone string) []*sacloud.NFS {
	values := ds().List(ResourceNFS, zone)
	var ret []*sacloud.NFS
	for _, v := range values {
		if v, ok := v.(*sacloud.NFS); ok {
			ret = append(ret, v)
		}
	}
	return ret
}

func getNFSByID(zone string, id types.ID) *sacloud.NFS {
	v := ds().Get(ResourceNFS, zone, id)
	if v, ok := v.(*sacloud.NFS); ok {
		return v
	}
	return nil
}

func putNFS(zone string, value *sacloud.NFS) {
	var v interface{} = value
	if id, ok := v.(accessor.ID); ok {
		ds().Put(ResourceNFS, zone, id.GetID(), value)
		return
	}
	ds().Put(ResourceNFS, zone, 0, value)
}

func getNote(zone string) []*sacloud.Note {
	values := ds().List(ResourceNote, zone)
	var ret []*sacloud.Note
	for _, v := range values {
		if v, ok := v.(*sacloud.Note); ok {
			ret = append(ret, v)
		}
	}
	return ret
}

func getNoteByID(zone string, id types.ID) *sacloud.Note {
	v := ds().Get(ResourceNote, zone, id)
	if v, ok := v.(*sacloud.Note); ok {
		return v
	}
	return nil
}

func putNote(zone string, value *sacloud.Note) {
	var v interface{} = value
	if id, ok := v.(accessor.ID); ok {
		ds().Put(ResourceNote, zone, id.GetID(), value)
		return
	}
	ds().Put(ResourceNote, zone, 0, value)
}

func getPacketFilter(zone string) []*sacloud.PacketFilter {
	values := ds().List(ResourcePacketFilter, zone)
	var ret []*sacloud.PacketFilter
	for _, v := range values {
		if v, ok := v.(*sacloud.PacketFilter); ok {
			ret = append(ret, v)
		}
	}
	return ret
}

func getPacketFilterByID(zone string, id types.ID) *sacloud.PacketFilter {
	v := ds().Get(ResourcePacketFilter, zone, id)
	if v, ok := v.(*sacloud.PacketFilter); ok {
		return v
	}
	return nil
}

func putPacketFilter(zone string, value *sacloud.PacketFilter) {
	var v interface{} = value
	if id, ok := v.(accessor.ID); ok {
		ds().Put(ResourcePacketFilter, zone, id.GetID(), value)
		return
	}
	ds().Put(ResourcePacketFilter, zone, 0, value)
}

func getPrivateHost(zone string) []*sacloud.PrivateHost {
	values := ds().List(ResourcePrivateHost, zone)
	var ret []*sacloud.PrivateHost
	for _, v := range values {
		if v, ok := v.(*sacloud.PrivateHost); ok {
			ret = append(ret, v)
		}
	}
	return ret
}

func getPrivateHostByID(zone string, id types.ID) *sacloud.PrivateHost {
	v := ds().Get(ResourcePrivateHost, zone, id)
	if v, ok := v.(*sacloud.PrivateHost); ok {
		return v
	}
	return nil
}

func putPrivateHost(zone string, value *sacloud.PrivateHost) {
	var v interface{} = value
	if id, ok := v.(accessor.ID); ok {
		ds().Put(ResourcePrivateHost, zone, id.GetID(), value)
		return
	}
	ds().Put(ResourcePrivateHost, zone, 0, value)
}

func getPrivateHostPlan(zone string) []*sacloud.PrivateHostPlan {
	values := ds().List(ResourcePrivateHostPlan, zone)
	var ret []*sacloud.PrivateHostPlan
	for _, v := range values {
		if v, ok := v.(*sacloud.PrivateHostPlan); ok {
			ret = append(ret, v)
		}
	}
	return ret
}

func getPrivateHostPlanByID(zone string, id types.ID) *sacloud.PrivateHostPlan {
	v := ds().Get(ResourcePrivateHostPlan, zone, id)
	if v, ok := v.(*sacloud.PrivateHostPlan); ok {
		return v
	}
	return nil
}

func putPrivateHostPlan(zone string, value *sacloud.PrivateHostPlan) {
	var v interface{} = value
	if id, ok := v.(accessor.ID); ok {
		ds().Put(ResourcePrivateHostPlan, zone, id.GetID(), value)
		return
	}
	ds().Put(ResourcePrivateHostPlan, zone, 0, value)
}

func getProxyLB(zone string) []*sacloud.ProxyLB {
	values := ds().List(ResourceProxyLB, zone)
	var ret []*sacloud.ProxyLB
	for _, v := range values {
		if v, ok := v.(*sacloud.ProxyLB); ok {
			ret = append(ret, v)
		}
	}
	return ret
}

func getProxyLBByID(zone string, id types.ID) *sacloud.ProxyLB {
	v := ds().Get(ResourceProxyLB, zone, id)
	if v, ok := v.(*sacloud.ProxyLB); ok {
		return v
	}
	return nil
}

func putProxyLB(zone string, value *sacloud.ProxyLB) {
	var v interface{} = value
	if id, ok := v.(accessor.ID); ok {
		ds().Put(ResourceProxyLB, zone, id.GetID(), value)
		return
	}
	ds().Put(ResourceProxyLB, zone, 0, value)
}

func getRegion(zone string) []*sacloud.Region {
	values := ds().List(ResourceRegion, zone)
	var ret []*sacloud.Region
	for _, v := range values {
		if v, ok := v.(*sacloud.Region); ok {
			ret = append(ret, v)
		}
	}
	return ret
}

func getRegionByID(zone string, id types.ID) *sacloud.Region {
	v := ds().Get(ResourceRegion, zone, id)
	if v, ok := v.(*sacloud.Region); ok {
		return v
	}
	return nil
}

func putRegion(zone string, value *sacloud.Region) {
	var v interface{} = value
	if id, ok := v.(accessor.ID); ok {
		ds().Put(ResourceRegion, zone, id.GetID(), value)
		return
	}
	ds().Put(ResourceRegion, zone, 0, value)
}

func getServer(zone string) []*sacloud.Server {
	values := ds().List(ResourceServer, zone)
	var ret []*sacloud.Server
	for _, v := range values {
		if v, ok := v.(*sacloud.Server); ok {
			ret = append(ret, v)
		}
	}
	return ret
}

func getServerByID(zone string, id types.ID) *sacloud.Server {
	v := ds().Get(ResourceServer, zone, id)
	if v, ok := v.(*sacloud.Server); ok {
		return v
	}
	return nil
}

func putServer(zone string, value *sacloud.Server) {
	var v interface{} = value
	if id, ok := v.(accessor.ID); ok {
		ds().Put(ResourceServer, zone, id.GetID(), value)
		return
	}
	ds().Put(ResourceServer, zone, 0, value)
}

func getServerPlan(zone string) []*sacloud.ServerPlan {
	values := ds().List(ResourceServerPlan, zone)
	var ret []*sacloud.ServerPlan
	for _, v := range values {
		if v, ok := v.(*sacloud.ServerPlan); ok {
			ret = append(ret, v)
		}
	}
	return ret
}

func getServerPlanByID(zone string, id types.ID) *sacloud.ServerPlan {
	v := ds().Get(ResourceServerPlan, zone, id)
	if v, ok := v.(*sacloud.ServerPlan); ok {
		return v
	}
	return nil
}

func putServerPlan(zone string, value *sacloud.ServerPlan) {
	var v interface{} = value
	if id, ok := v.(accessor.ID); ok {
		ds().Put(ResourceServerPlan, zone, id.GetID(), value)
		return
	}
	ds().Put(ResourceServerPlan, zone, 0, value)
}

func getServiceClass(zone string) []*sacloud.ServiceClass {
	values := ds().List(ResourceServiceClass, zone)
	var ret []*sacloud.ServiceClass
	for _, v := range values {
		if v, ok := v.(*sacloud.ServiceClass); ok {
			ret = append(ret, v)
		}
	}
	return ret
}

func getServiceClassByID(zone string, id types.ID) *sacloud.ServiceClass {
	v := ds().Get(ResourceServiceClass, zone, id)
	if v, ok := v.(*sacloud.ServiceClass); ok {
		return v
	}
	return nil
}

func putServiceClass(zone string, value *sacloud.ServiceClass) {
	var v interface{} = value
	if id, ok := v.(accessor.ID); ok {
		ds().Put(ResourceServiceClass, zone, id.GetID(), value)
		return
	}
	ds().Put(ResourceServiceClass, zone, 0, value)
}

func getSIM(zone string) []*sacloud.SIM {
	values := ds().List(ResourceSIM, zone)
	var ret []*sacloud.SIM
	for _, v := range values {
		if v, ok := v.(*sacloud.SIM); ok {
			ret = append(ret, v)
		}
	}
	return ret
}

func getSIMByID(zone string, id types.ID) *sacloud.SIM {
	v := ds().Get(ResourceSIM, zone, id)
	if v, ok := v.(*sacloud.SIM); ok {
		return v
	}
	return nil
}

func putSIM(zone string, value *sacloud.SIM) {
	var v interface{} = value
	if id, ok := v.(accessor.ID); ok {
		ds().Put(ResourceSIM, zone, id.GetID(), value)
		return
	}
	ds().Put(ResourceSIM, zone, 0, value)
}

func getSimpleMonitor(zone string) []*sacloud.SimpleMonitor {
	values := ds().List(ResourceSimpleMonitor, zone)
	var ret []*sacloud.SimpleMonitor
	for _, v := range values {
		if v, ok := v.(*sacloud.SimpleMonitor); ok {
			ret = append(ret, v)
		}
	}
	return ret
}

func getSimpleMonitorByID(zone string, id types.ID) *sacloud.SimpleMonitor {
	v := ds().Get(ResourceSimpleMonitor, zone, id)
	if v, ok := v.(*sacloud.SimpleMonitor); ok {
		return v
	}
	return nil
}

func putSimpleMonitor(zone string, value *sacloud.SimpleMonitor) {
	var v interface{} = value
	if id, ok := v.(accessor.ID); ok {
		ds().Put(ResourceSimpleMonitor, zone, id.GetID(), value)
		return
	}
	ds().Put(ResourceSimpleMonitor, zone, 0, value)
}

func getSSHKey(zone string) []*sacloud.SSHKey {
	values := ds().List(ResourceSSHKey, zone)
	var ret []*sacloud.SSHKey
	for _, v := range values {
		if v, ok := v.(*sacloud.SSHKey); ok {
			ret = append(ret, v)
		}
	}
	return ret
}

func getSSHKeyByID(zone string, id types.ID) *sacloud.SSHKey {
	v := ds().Get(ResourceSSHKey, zone, id)
	if v, ok := v.(*sacloud.SSHKey); ok {
		return v
	}
	return nil
}

func putSSHKey(zone string, value *sacloud.SSHKey) {
	var v interface{} = value
	if id, ok := v.(accessor.ID); ok {
		ds().Put(ResourceSSHKey, zone, id.GetID(), value)
		return
	}
	ds().Put(ResourceSSHKey, zone, 0, value)
}

func getSubnet(zone string) []*sacloud.Subnet {
	values := ds().List(ResourceSubnet, zone)
	var ret []*sacloud.Subnet
	for _, v := range values {
		if v, ok := v.(*sacloud.Subnet); ok {
			ret = append(ret, v)
		}
	}
	return ret
}

func getSubnetByID(zone string, id types.ID) *sacloud.Subnet {
	v := ds().Get(ResourceSubnet, zone, id)
	if v, ok := v.(*sacloud.Subnet); ok {
		return v
	}
	return nil
}

func putSubnet(zone string, value *sacloud.Subnet) {
	var v interface{} = value
	if id, ok := v.(accessor.ID); ok {
		ds().Put(ResourceSubnet, zone, id.GetID(), value)
		return
	}
	ds().Put(ResourceSubnet, zone, 0, value)
}

func getSwitch(zone string) []*sacloud.Switch {
	values := ds().List(ResourceSwitch, zone)
	var ret []*sacloud.Switch
	for _, v := range values {
		if v, ok := v.(*sacloud.Switch); ok {
			ret = append(ret, v)
		}
	}
	return ret
}

func getSwitchByID(zone string, id types.ID) *sacloud.Switch {
	v := ds().Get(ResourceSwitch, zone, id)
	if v, ok := v.(*sacloud.Switch); ok {
		return v
	}
	return nil
}

func putSwitch(zone string, value *sacloud.Switch) {
	var v interface{} = value
	if id, ok := v.(accessor.ID); ok {
		ds().Put(ResourceSwitch, zone, id.GetID(), value)
		return
	}
	ds().Put(ResourceSwitch, zone, 0, value)
}

func getVPCRouter(zone string) []*sacloud.VPCRouter {
	values := ds().List(ResourceVPCRouter, zone)
	var ret []*sacloud.VPCRouter
	for _, v := range values {
		if v, ok := v.(*sacloud.VPCRouter); ok {
			ret = append(ret, v)
		}
	}
	return ret
}

func getVPCRouterByID(zone string, id types.ID) *sacloud.VPCRouter {
	v := ds().Get(ResourceVPCRouter, zone, id)
	if v, ok := v.(*sacloud.VPCRouter); ok {
		return v
	}
	return nil
}

func putVPCRouter(zone string, value *sacloud.VPCRouter) {
	var v interface{} = value
	if id, ok := v.(accessor.ID); ok {
		ds().Put(ResourceVPCRouter, zone, id.GetID(), value)
		return
	}
	ds().Put(ResourceVPCRouter, zone, 0, value)
}

func getWebAccel(zone string) []*sacloud.WebAccel {
	values := ds().List(ResourceWebAccel, zone)
	var ret []*sacloud.WebAccel
	for _, v := range values {
		if v, ok := v.(*sacloud.WebAccel); ok {
			ret = append(ret, v)
		}
	}
	return ret
}

func getWebAccelByID(zone string, id types.ID) *sacloud.WebAccel {
	v := ds().Get(ResourceWebAccel, zone, id)
	if v, ok := v.(*sacloud.WebAccel); ok {
		return v
	}
	return nil
}

func putWebAccel(zone string, value *sacloud.WebAccel) {
	var v interface{} = value
	if id, ok := v.(accessor.ID); ok {
		ds().Put(ResourceWebAccel, zone, id.GetID(), value)
		return
	}
	ds().Put(ResourceWebAccel, zone, 0, value)
}

func getZone(zone string) []*sacloud.Zone {
	values := ds().List(ResourceZone, zone)
	var ret []*sacloud.Zone
	for _, v := range values {
		if v, ok := v.(*sacloud.Zone); ok {
			ret = append(ret, v)
		}
	}
	return ret
}

func getZoneByID(zone string, id types.ID) *sacloud.Zone {
	v := ds().Get(ResourceZone, zone, id)
	if v, ok := v.(*sacloud.Zone); ok {
		return v
	}
	return nil
}

func putZone(zone string, value *sacloud.Zone) {
	var v interface{} = value
	if id, ok := v.(accessor.ID); ok {
		ds().Put(ResourceZone, zone, id.GetID(), value)
		return
	}
	ds().Put(ResourceZone, zone, 0, value)
}
