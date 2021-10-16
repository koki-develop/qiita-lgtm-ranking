variable "stage" {
  type    = string
  default = "prod"
}

terraform {
  backend "s3" {
    profile = "default"
    region  = "us-east-1"
    bucket  = "qiita-lgtm-ranking-tfstates"
    key     = "prod/terraform.tfstate"
  }
}
