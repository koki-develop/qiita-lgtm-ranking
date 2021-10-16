variable "stage" {
  type    = string
  default = "dev"
}

terraform {
  backend "s3" {
    profile = "default"
    region  = "us-east-1"
    bucket  = "qiita-lgtm-ranking-tfstates"
    key     = "dev/terraform.tfstate"
  }
}
