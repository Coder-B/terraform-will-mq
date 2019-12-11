# alicloud_ons_instance

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|:----:|:-----:|:-----:|
|name|The name of mq instance| string| "" | yes|
|remark|The remark of mq instance| string| "" | no|


## Outputs

| Name | Description |
|------|-------------|
|id|The key of the resource supplied above|
|instance_type | The edition of instance. 1 represents the postPaid edition, and 2 represents the platinum edition|
|instance_status |The status of instance. 1 represents the platinum edition instance is in deployment. 2 represents the postpaid edition instance are overdue. 5 represents the postpaid or platinum edition instance is in service. 7 represents the platinum version instance is in upgrade and the service is available|
|release_time|Platinum edition instance expiration time|