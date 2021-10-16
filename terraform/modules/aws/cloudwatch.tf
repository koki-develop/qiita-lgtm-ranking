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
          metrics = [
            [
              "AWS/Lambda",
              "Errors",
              "FunctionName",
              data.aws_lambda_function.update_weekly.function_name,
              "Resource",
              data.aws_lambda_function.update_weekly.function_name,
              {
                color = "#d62728"
              },
            ],
          ]
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
          metrics = [
            [
              "AWS/Lambda",
              "Errors",
              "FunctionName",
              data.aws_lambda_function.update_weekly_by_tag.function_name,
              "Resource",
              data.aws_lambda_function.update_weekly_by_tag.function_name,
              {
                color = "#d62728"
              },
            ],
          ]
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
