// Copyright 2016-2020 terraform-provider-sakuracloud authors
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

package sakuracloud

import (
	"context"
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/libsacloud/v2/utils/power"
)

func resourceSakuraCloudDatabaseReadReplica() *schema.Resource {
	resourceName := "read-replica database"
	return &schema.Resource{
		Create: resourceSakuraCloudDatabaseReadReplicaCreate,
		Read:   resourceSakuraCloudDatabaseReadReplicaRead,
		Update: resourceSakuraCloudDatabaseReadReplicaUpdate,
		Delete: resourceSakuraCloudDatabaseReadReplicaDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(60 * time.Minute),
			Update: schema.DefaultTimeout(60 * time.Minute),
			Delete: schema.DefaultTimeout(20 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"master_id": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validateSakuracloudIDType,
				Description:  "The id of the replication master database",
			},
			"name": schemaResourceName(resourceName),
			"network_interface": {
				Type:     schema.TypeList,
				Required: true,
				MinItems: 1,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"switch_id": {
							Type:         schema.TypeString,
							ForceNew:     true,
							Optional:     true,
							Computed:     true,
							ValidateFunc: validateSakuracloudIDType,
							Description:  descf("The id of the switch to which the %s connects. If `switch_id` isn't specified, it will be set to the same value of the master database", resourceName),
						},
						"ip_address": {
							Type:        schema.TypeString,
							ForceNew:    true,
							Required:    true,
							Description: descf("The IP address to assign to the %s", resourceName),
						},
						"netmask": {
							Type:         schema.TypeInt,
							ForceNew:     true,
							Optional:     true,
							Computed:     true,
							ValidateFunc: validation.IntBetween(8, 29),
							Description: descf(
								"The bit length of the subnet to assign to the %s. %s. If `netmask` isn't specified, it will be set to the same value of the master database",
								resourceName,
								descRange(8, 29),
							),
						},
						"gateway": {
							Type:        schema.TypeString,
							ForceNew:    true,
							Optional:    true,
							Computed:    true,
							Description: descf("The IP address of the gateway used by %s. If `gateway` isn't specified, it will be set to the same value of the master database", resourceName),
						},
						"source_ranges": {
							Type:     schema.TypeList,
							Optional: true,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Description: descf(
								"The range of source IP addresses that allow to access to the %s via network",
								resourceName,
							),
						},
					},
				},
			},
			"icon_id":     schemaResourceIconID(resourceName),
			"description": schemaResourceDescription(resourceName),
			"tags":        schemaResourceTags(resourceName),
			"zone":        schemaResourceZone(resourceName),
		},
	}
}

func resourceSakuraCloudDatabaseReadReplicaCreate(d *schema.ResourceData, meta interface{}) error {
	client, zone, err := sakuraCloudClient(d, meta)
	if err != nil {
		return err
	}
	ctx, cancel := operationContext(d, schema.TimeoutCreate)
	defer cancel()

	// validate master instance
	builder, err := expandDatabaseReadReplicaBuilder(ctx, d, client, zone)
	if err != nil {
		return nil
	}

	db, err := builder.Build(ctx, zone)
	if db != nil {
		d.SetId(db.ID.String())
	}
	if err != nil {
		return fmt.Errorf("creating SakuraCloud Database ReadReplica is failed: %s", err)
	}

	// HACK データベースアプライアンスの電源投入後すぐに他の操作(Updateなど)を行うと202(Accepted)が返ってくるものの無視される。
	// この挙動はテストなどで問題となる。このためここで少しsleepすることで対応する。
	time.Sleep(client.databaseWaitAfterCreateDuration)

	return setDatabaseReadReplicaResourceData(ctx, d, client, db)
}

func resourceSakuraCloudDatabaseReadReplicaRead(d *schema.ResourceData, meta interface{}) error {
	client, zone, err := sakuraCloudClient(d, meta)
	if err != nil {
		return err
	}
	ctx, cancel := operationContext(d, schema.TimeoutRead)
	defer cancel()

	dbOp := sacloud.NewDatabaseOp(client)

	data, err := dbOp.Read(ctx, zone, sakuraCloudID(d.Id()))
	if err != nil {
		if sacloud.IsNotFoundError(err) {
			d.SetId("")
			return nil
		}
		return fmt.Errorf("could not find SakuraCloud Database ReadReplica[%s] : %s", d.Id(), err)
	}
	return setDatabaseReadReplicaResourceData(ctx, d, client, data)
}

func resourceSakuraCloudDatabaseReadReplicaUpdate(d *schema.ResourceData, meta interface{}) error {
	client, zone, err := sakuraCloudClient(d, meta)
	if err != nil {
		return err
	}
	ctx, cancel := operationContext(d, schema.TimeoutUpdate)
	defer cancel()

	dbOp := sacloud.NewDatabaseOp(client)

	db, err := dbOp.Read(ctx, zone, sakuraCloudID(d.Id()))
	if err != nil {
		return fmt.Errorf("could not read SakuraCloud Database[%s]: %s", d.Id(), err)
	}

	builder, err := expandDatabaseReadReplicaBuilder(ctx, d, client, zone)
	if err != nil {
		return nil
	}

	db, err = builder.Update(ctx, zone, db.ID)
	if err != nil {
		return fmt.Errorf("updating SakuraCloud Database ReadReplica[%s] is failed: %s", db.ID, err)
	}

	return setDatabaseReadReplicaResourceData(ctx, d, client, db)
}

func resourceSakuraCloudDatabaseReadReplicaDelete(d *schema.ResourceData, meta interface{}) error {
	client, zone, err := sakuraCloudClient(d, meta)
	if err != nil {
		return err
	}
	ctx, cancel := operationContext(d, schema.TimeoutDelete)
	defer cancel()

	dbOp := sacloud.NewDatabaseOp(client)

	data, err := dbOp.Read(ctx, zone, sakuraCloudID(d.Id()))
	if err != nil {
		if sacloud.IsNotFoundError(err) {
			d.SetId("")
			return nil
		}
		return fmt.Errorf("could not read SakuraCloud Database[%s]: %s", d.Id(), err)
	}

	// shutdown(force) if running
	if data.InstanceStatus.IsUp() {
		if err := power.ShutdownDatabase(ctx, dbOp, zone, data.ID, true); err != nil {
			return err
		}
	}

	// delete
	if err = dbOp.Delete(ctx, zone, data.ID); err != nil {
		return fmt.Errorf("deleting SakuraCloud Database[%s] is failed: %s", data.ID, err)
	}

	d.SetId("")
	return nil
}

func setDatabaseReadReplicaResourceData(ctx context.Context, d *schema.ResourceData, client *APIClient, data *sacloud.Database) error {
	if data.Availability.IsFailed() {
		d.SetId("")
		return fmt.Errorf("got unexpected state: Database[%d].Availability is failed", data.ID)
	}

	if err := d.Set("tags", flattenDatabaseTags(data)); err != nil {
		return err
	}

	d.Set("master_id", data.ReplicationSetting.ApplianceID.String()) // nolint
	d.Set("name", data.Name)                                         // nolint
	if err := d.Set("network_interface", flattenDatabaseReadReplicaNetworkInterface(data)); err != nil {
		return err
	}
	d.Set("icon_id", data.IconID.String()) // nolint
	d.Set("description", data.Description) // nolint
	d.Set("zone", getZone(d, client))      // nolint
	return nil
}
