// Copyright 2016-2019 terraform-provider-sakuracloud authors
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

package main

import (
	"fmt"
	"log"
	"os"

	"github.com/sacloud/terraform-provider-sakuracloud/sakuracloud"
	"github.com/sacloud/terraform-provider-sakuracloud/tools/tfdocgen"
)

const (
	templateDir = "templates"
	exampleDir  = "examples"
)

func main() {
	if len(os.Args) < 2 || 2 < len(os.Args) {
		fmt.Println("Usage: gen-sakuracloud-docs <destination-directory>")
		os.Exit(1)
	}

	destination := os.Args[1]
	provider := tfdocgen.Provider{
		Name:              "sakuracloud",
		TerraformProvider: sakuracloud.Provider(),
		DisplayNameFunc: func(name string) string {
			d, ok := definitions[name]
			if !ok {
				return name
			}
			return d.displayName
		},
		CategoryNameFunc: func(name string) string {
			d, ok := definitions[name]
			if !ok {
				return ""
			}
			return d.category
		},
		CategoriesFunc: func() []string {
			return categories
		},
	}
	if err := provider.GenerateDocs(templateDir, exampleDir, destination); err != nil {
		log.Fatal(err)
	}
}

type definition struct {
	displayName string
	category    string
}

const (
	CategoryProvider      = "Provider Data Sources"
	CategoryCompute       = "Compute"
	CategoryStorage       = "Storage"
	CategoryNetworking    = "Networking"
	CategoryAppliance     = "Appliance"
	CategoryGlobal        = "Global"
	CategorySecureMobile  = "SecureMobile"
	CategoryObjectStorage = "ObjectStorage"
	CategoryLab           = "Lab"
	CategoryMisc          = "Misc"
)

var categories = []string{
	CategoryProvider,
	CategoryCompute,
	CategoryStorage,
	CategoryNetworking,
	CategoryAppliance,
	CategoryGlobal,
	CategorySecureMobile,
	CategoryLab,
	CategoryMisc,
	CategoryObjectStorage,
}

var definitions = map[string]definition{
	"sakuracloud": {
		displayName: "SakuraCloud",
	},
	"sakuracloud_archive": {
		displayName: "Archive",
		category:    CategoryStorage,
	},
	"sakuracloud_auto_backup": {
		displayName: "AutoBackup",
		category:    CategoryAppliance,
	},
	"sakuracloud_bridge": {
		displayName: "Bridge",
		category:    CategoryNetworking,
	},
	"sakuracloud_bucket_object": {
		displayName: "Bucket Object",
		category:    CategoryObjectStorage,
	},
	"sakuracloud_cdrom": {
		displayName: "CD-ROM",
		category:    CategoryStorage,
	},
	"sakuracloud_container_registry": {
		displayName: "Container Registry",
		category:    CategoryLab,
	},
	"sakuracloud_database": {
		displayName: "Database",
		category:    CategoryAppliance,
	},
	"sakuracloud_database_read_replica": {
		displayName: "Database Read Replica",
		category:    CategoryAppliance,
	},
	"sakuracloud_disk": {
		displayName: "Disk",
		category:    CategoryStorage,
	},
	"sakuracloud_dns": {
		displayName: "DNS",
		category:    CategoryGlobal,
	},
	"sakuracloud_dns_record": {
		displayName: "DNS Record",
		category:    CategoryGlobal,
	},
	"sakuracloud_gslb": {
		displayName: "GSLB",
		category:    CategoryGlobal,
	},
	"sakuracloud_icon": {
		displayName: "Icon",
		category:    CategoryMisc,
	},
	"sakuracloud_internet": {
		displayName: "Internet",
		category:    CategoryNetworking,
	},
	"sakuracloud_ipv4_ptr": {
		displayName: "IPv4 PTR",
		category:    CategoryNetworking,
	},
	"sakuracloud_load_balancer": {
		displayName: "LoadBalancer",
		category:    CategoryAppliance,
	},
	"sakuracloud_mobile_gateway": {
		displayName: "MobileGateway",
		category:    CategorySecureMobile,
	},
	"sakuracloud_nfs": {
		displayName: "NFS",
		category:    CategoryAppliance,
	},
	"sakuracloud_note": {
		displayName: "Note",
		category:    CategoryMisc,
	},
	"sakuracloud_packet_filter": {
		displayName: "PacketFilter",
		category:    CategoryNetworking,
	},
	"sakuracloud_packet_filter_rules": {
		displayName: "PacketFilter Rules",
		category:    CategoryNetworking,
	},
	"sakuracloud_private_host": {
		displayName: "PrivateHost",
		category:    CategoryCompute,
	},
	"sakuracloud_proxylb": {
		displayName: "ProxyLB",
		category:    CategoryGlobal,
	},
	"sakuracloud_proxylb_acme": {
		displayName: "ProxyLB ACME Setting",
		category:    CategoryGlobal,
	},
	"sakuracloud_server": {
		displayName: "Server",
		category:    CategoryCompute,
	},
	"sakuracloud_sim": {
		displayName: "SIM",
		category:    CategorySecureMobile,
	},
	"sakuracloud_simple_monitor": {
		displayName: "SimpleMonitor",
		category:    CategoryGlobal,
	},
	"sakuracloud_ssh_key": {
		displayName: "SSHKey",
		category:    CategoryMisc,
	},
	"sakuracloud_ssh_key_gen": {
		displayName: "SSHKey Gen",
		category:    CategoryMisc,
	},
	"sakuracloud_subnet": {
		displayName: "Subnet",
		category:    CategoryNetworking,
	},
	"sakuracloud_switch": {
		displayName: "Switch",
		category:    CategoryNetworking,
	},
	"sakuracloud_vpc_router": {
		displayName: "VPCRouter",
		category:    CategoryAppliance,
	},
	"sakuracloud_zone": {
		displayName: "Zone",
		category:    CategoryProvider,
	},
}