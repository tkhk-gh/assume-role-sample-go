# aws-sdk-go で assume role 使ったサンプル

## 前提条件

- $HOME/.aws/config に ```source_profile``` と ```role_arn``` を設定

```
[profile sample]
role_arn = arn:aws:iam::<account>:role/<role_name>
source_profile = default
```

## Sample

``` bash
AWS_PROFILE=default AWS_REGION=ap-northeast-1 go run main.go
AWS_PROFILE=sample AWS_REGION=ap-northeast-1 go run main.go
```

## メモ

- assume role の設定はデフォルトだと ```$HOME/.aws/config``` に記述する
- ```NewSession()``` では config を読んでくれない
  - SharedConfigEnable にした ```session.Options``` を渡す必要がある
