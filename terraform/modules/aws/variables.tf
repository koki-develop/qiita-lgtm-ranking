variable "stage" {
  type = string
}

locals {
  prefix = "qiita-lgtm-ranking-${var.stage}"
}
