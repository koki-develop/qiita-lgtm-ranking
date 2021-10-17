resource "aws_cloudwatch_dashboard" "default" {
  count = var.stage == "prod" ? 1 : 0

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
              data.aws_lambda_function.update_daily.function_name,
            ],
          ]
          period  = 300
          region  = "us-east-1"
          stacked = false
          stat    = "Sum"
          title   = "デイリーランキング - 実行回数"
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
              aws_cloudwatch_metric_alarm.daily_errors[0].arn
            ]
          }
          period  = 300
          region  = "us-east-1"
          stacked = false
          stat    = "Sum"
          title   = "デイリーランキング - エラー回数"
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
              data.aws_lambda_function.update_daily_by_tag.function_name,
            ],
          ]
          period  = 300
          region  = "us-east-1"
          stacked = false
          stat    = "Sum"
          title   = "デイリーランキング ( タグ別 ) - 実行回数"
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
              aws_cloudwatch_metric_alarm.daily_by_tag_errors[0].arn
            ]
          }
          period  = 300
          region  = "us-east-1"
          stacked = false
          stat    = "Sum"
          title   = "デイリーランキング ( タグ別 ) - エラー回数"
          view    = "timeSeries"
        }
        type   = "metric"
        width  = 12
        height = 6
        x      = 12
        y      = 6
      },

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
        y      = 12
      },
      {
        properties = {
          annotations = {
            alarms = [
              aws_cloudwatch_metric_alarm.weekly_errors[0].arn
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
        y      = 12
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
        y      = 18
      },
      {
        properties = {
          annotations = {
            alarms = [
              aws_cloudwatch_metric_alarm.weekly_by_tag_errors[0].arn
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
        y      = 18
      },
    ]
  })
}

resource "aws_cloudwatch_metric_alarm" "daily_errors" {
  count = var.stage == "prod" ? 1 : 0

  alarm_name          = "${local.prefix}-daily-errors"
  actions_enabled     = true
  comparison_operator = "GreaterThanOrEqualToThreshold"
  datapoints_to_alarm = 1
  dimensions = {
    "FunctionName" = data.aws_lambda_function.update_daily.function_name
    "Resource"     = data.aws_lambda_function.update_daily.function_name
  }
  evaluation_periods = 1
  metric_name        = "Errors"
  namespace          = "AWS/Lambda"
  period             = 3600
  statistic          = "Sum"
  threshold          = 1
  treat_missing_data = "notBreaching"
  ok_actions         = [aws_sns_topic.default[0].arn]
  alarm_actions      = [aws_sns_topic.default[0].arn]
}

resource "aws_cloudwatch_metric_alarm" "daily_by_tag_errors" {
  count = var.stage == "prod" ? 1 : 0

  alarm_name          = "${local.prefix}-daily-by-tag-errors"
  actions_enabled     = true
  comparison_operator = "GreaterThanOrEqualToThreshold"
  datapoints_to_alarm = 1
  dimensions = {
    "FunctionName" = data.aws_lambda_function.update_daily_by_tag.function_name
    "Resource"     = data.aws_lambda_function.update_daily_by_tag.function_name
  }
  evaluation_periods = 1
  metric_name        = "Errors"
  namespace          = "AWS/Lambda"
  period             = 3600
  statistic          = "Sum"
  threshold          = 1
  treat_missing_data = "notBreaching"
  ok_actions         = [aws_sns_topic.default[0].arn]
  alarm_actions      = [aws_sns_topic.default[0].arn]
}

resource "aws_cloudwatch_metric_alarm" "weekly_errors" {
  count = var.stage == "prod" ? 1 : 0

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
  period             = 3600
  statistic          = "Sum"
  threshold          = 1
  treat_missing_data = "notBreaching"
  ok_actions         = [aws_sns_topic.default[0].arn]
  alarm_actions      = [aws_sns_topic.default[0].arn]
}

resource "aws_cloudwatch_metric_alarm" "weekly_by_tag_errors" {
  count = var.stage == "prod" ? 1 : 0

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
  period             = 3600
  statistic          = "Sum"
  threshold          = 1
  treat_missing_data = "notBreaching"
  ok_actions         = [aws_sns_topic.default[0].arn]
  alarm_actions      = [aws_sns_topic.default[0].arn]
}
