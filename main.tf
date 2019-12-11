provider "alicloud" {
  version              = ">=1.56.0"
  region               = var.region != "" ? var.region : null
  configuration_source = "terraform-will-modules/kvstore"
}

#################
# ons instance
#################
module "mq_instance" {
    source = "./modules/ons_instance"
    name   = var.name
    remark = var.remark
}