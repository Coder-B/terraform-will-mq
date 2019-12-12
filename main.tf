provider "alicloud" {
  version              = ">=1.56.0"
  access_key           = var.access_key != "" ? var.access_key : null
  secret_key           = var.secret_key != "" ? var.secret_key : null
  region               = var.region != "" ? var.region : null
  configuration_source = "terraform-will-modules/mq"
}

#################
# ons instance
#################
module "mq_instance" {
    source = "./modules/ons_instance"
    name   = var.name
    remark = var.remark
}