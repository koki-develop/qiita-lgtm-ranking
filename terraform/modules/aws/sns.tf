resource "aws_sns_topic" "errors" {
  name = "${local.prefix}-errors"
}

resource "aws_sns_topic_subscription" "errors_email" {
  topic_arn = aws_sns_topic.errors.arn
  protocol  = "email"
  endpoint  = "kou.pg.0131@gmail.com"
}
