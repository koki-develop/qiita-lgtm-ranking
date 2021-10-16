resource "aws_cloudwatch_dashboard" "default" {
  dashboard_name = local.prefix
  dashboard_body = jsonencode({
    widgets = [
      {
        properties = {
          metrics = [
            [
              "AWS/Lambda",
              "Invocations",
              "FunctionName",
              data.aws_lambda_function.update_weekly.function_name,
            ],
          ]
          period  = 300
          region  = "us-east-1"
          stacked = false
          stat    = "Sum"
          title   = "週間ランキング - 実行回数"
          view    = "timeSeries"
        }
        type   = "metric"
        width  = 12
        height = 6
        x      = 0
        y      = 0
      },
      {
        properties = {
          annotations = {
            alarms = [
              aws_cloudwatch_metric_alarm.weekly_errors.arn
            ]
          }
          period  = 300
          region  = "us-east-1"
          stacked = false
          stat    = "Sum"
          title   = "週間ランキング - エラー回数"
          view    = "timeSeries"
        }
        type   = "metric"
        width  = 12
        height = 6
        x      = 12
        y      = 0
      },
      {
        properties = {
          metrics = [
            [
              "AWS/Lambda",
              "Invocations",
              "FunctionName",
              data.aws_lambda_function.update_weekly_by_tag.function_name,
            ],
          ]
          period  = 300
          region  = "us-east-1"
          stacked = false
          stat    = "Sum"
          title   = "週間ランキング ( タグ別 ) - 実行回数"
          view    = "timeSeries"
        }
        type   = "metric"
        width  = 12
        height = 6
        x      = 0
        y      = 6
      },
      {
        properties = {
          annotations = {
            alarms = [
              aws_cloudwatch_metric_alarm.weekly_by_tag_errors.arn
            ]
          }
          period  = 300
          region  = "us-east-1"
          stacked = false
          stat    = "Sum"
          title   = "週間ランキング ( タグ別 ) - エラー回数"
          view    = "timeSeries"
        }
        type   = "metric"
        width  = 12
        height = 6
        x      = 12
        y      = 6
      },
    ]
  })
}

resource "aws_cloudwatch_metric_alarm" "weekly_errors" {
  alarm_name          = "${local.prefix}-weekly-errors"
  actions_enabled     = true
  comparison_operator = "GreaterThanOrEqualToThreshold"
  datapoints_to_alarm = 1
  dimensions = {
    "FunctionName" = data.aws_lambda_function.update_weekly.function_name
    "Resource"     = data.aws_lambda_function.update_weekly.function_name
  }
  evaluation_periods = 1
  metric_name        = "Errors"
  namespace          = "AWS/Lambda"
  period             = 300
  statistic          = "Sum"
  threshold          = 1
  treat_missing_data = "notBreaching"
  ok_actions         = [aws_sns_topic.default.arn]
  alarm_actions      = [aws_sns_topic.default.arn]
}

resource "aws_cloudwatch_metric_alarm" "weekly_by_tag_errors" {
  alarm_name          = "${local.prefix}-weekly-by-tag-errors"
  actions_enabled     = true
  comparison_operator = "GreaterThanOrEqualToThreshold"
  datapoints_to_alarm = 1
  dimensions = {
    "FunctionName" = data.aws_lambda_function.update_weekly_by_tag.function_name
    "Resource"     = data.aws_lambda_function.update_weekly_by_tag.function_name
  }
  evaluation_periods = 1
  metric_name        = "Errors"
  namespace          = "AWS/Lambda"
  period             = 300
  statistic          = "Sum"
  threshold          = 1
  treat_missing_data = "notBreaching"
  ok_actions         = [aws_sns_topic.default.arn]
  alarm_actions      = [aws_sns_topic.default.arn]
}
