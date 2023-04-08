# Todo-Sample-Backend

Todo サンプルのバックエンド側。  
この Readme ではバックエンド側に閉じた実装イメージや工夫点について説明します。

## 基本構成

### 概要

| 検討項目       | 採用技術                                               | 採用理由                                                                                      | 備考                                                                                                                          |
| -------------- | ------------------------------------------------------ | --------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------- |
| クラウド基盤   | AWS</br>(API-Gateyway ＆ Cognito & Lambda ＆ DynamoDB) | AWS が一番慣れているため。</br>また、サーバレスにすると、利用されていない時間は課金されない。 | [Amplify](https://aws.amazon.com/jp/amplify/)で基盤レイヤーは作成。                                                           |
| 言語           | [GO](https://go.dev/)                                  | Lambda においてはコールドスタート、</br>ウォームスタート共にハイパフォーマンスであったため。  | [参考サイト](filia-aleks.medium.com/aws-lambda-battle-2021-performance-comparison-for-all-languages-c1b441005fd1)             |
| フレームワーク | [Fiber](https://gofiber.io/)                           | メモリ効率が高く、Fast-HTTP を使用していることによるハイパフォーマンス                        | [参考 URL](https://github.com/gofiber/fiber#-benchmarks)</br>ただ、[Gin](https://gin-gonic.com/ja/)のほうが有名なのがネック。 |
