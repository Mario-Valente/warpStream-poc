resource "aws_s3_bucket" "warpstream_bucket" {
  bucket = "my-warpstream-bucket-data-store-evolution"

  tags = {
    Name        = "my-warpstream-bucket-data-store-evolution"
    Environment = "staging"
  }
}

resource "aws_s3_bucket_metric" "warpstream_bucket_metrics" {
 bucket = aws_s3_bucket.warpstream_bucket.id
 name   = "EntireBucket"
}

resource "aws_s3_bucket_lifecycle_configuration" "warpstream_bucket_lifecycle" {
  bucket = aws_s3_bucket.warpstream_bucket.id


  rule {
    id     = "7d multi-part"
    status = "Enabled"
    abort_incomplete_multipart_upload {
      days_after_initiation = 7
    }
  }
}