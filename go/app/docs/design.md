# 設計概要

- クリーンアーキテクチャを採用

## アーキテクチャ詳細

### ディレクトリ構成

| レイヤー - （外部 -> 内部） | ディレクトリ    |
| ------------------------- | -------------- |
| Frameworks & Drivers      | infrastructure |
| Interface Adapters        | interface      |
| Use Cases                 | usecase        |
| Entities                  | domain         |

### ディレクトリごとの機能概要

| ディレクトリ    | 機能概要                                                                                                       |
| -------------- | ------------------------------------------------------------------------------------------------------------- |
| infrastructure | フレームワーク等、外部ツールの管理（本プロジェクトは置換を想定し、標準パッケージのみを管理）                          |
| interface      | 外部ツールとのやり取りを行うアダプタ（外部ツールを意識しない基本的な機能を管理）<br />repositoryとして機能の実体も管理 |
| usecase        | 各種入力をどう出力するかを定義（ControllerからどのRepositoryの処理実体を呼び出すかを管理）                          |
| domain         | 各種モデルを管理（各ページがどんな情報を持っているかを管理）                                                       |
