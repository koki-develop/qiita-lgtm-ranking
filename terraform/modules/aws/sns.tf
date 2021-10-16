resource "aws_sns_topic" "default" {
  name = local.prefix
}

resource "aws_sns_topic_subscription" "email" {
  topic_arn = aws_sns_topic.default.arn
  protocol  = "email"
  endpoint  = "kou.pg.0131@gmail.com"
}
