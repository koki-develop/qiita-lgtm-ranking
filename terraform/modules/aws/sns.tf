resource "aws_sns_topic" "errors" {
  name = "${local.prefix}-errors"
}
