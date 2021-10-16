data "aws_lambda_function" "update_daily" {
  function_name = "${local.prefix}-updateDaily"
}

data "aws_lambda_function" "update_weekly" {
  function_name = "${local.prefix}-updateWeekly"
}

data "aws_lambda_function" "update_weekly_by_tag" {
  function_name = "${local.prefix}-updateWeeklyByTag"
}
