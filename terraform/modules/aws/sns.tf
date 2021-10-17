resource "aws_sns_topic" "default" {
  count = var.stage == "prod" ? 1 : 0

  name = local.prefix
}

resource "aws_sns_topic_subscription" "email" {
  count = var.stage == "prod" ? 1 : 0

  topic_arn = aws_sns_topic.default[0].arn
  protocol  = "email"
  endpoint  = "kou.pg.0131@gmail.com"
}

resource "aws_sns_topic_subscription" "chatbot" {
  count = var.stage == "prod" ? 1 : 0

  topic_arn = aws_sns_topic.default[0].arn
  protocol  = "https"
  endpoint  = "https://global.sns-api.chatbot.amazonaws.com"
}
