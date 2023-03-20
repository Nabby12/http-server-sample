# 設計概要

- クリーンアーキテクチャを採用

## アーキテクチャ詳細

### ディレクトリ構成

| ディレクトリ          | レイヤー                  | 概要                                                                           |
| --------------------- | ------------------------- | ------------------------------------------------------------------------------ |
| infrastructure        | Frameworks & Drivers      | 外部に依存するロジック（DBとのやり取りや、外部との通信）                       |
| cmd, internal/adapter | Interface Adapters        | Application Business Rule間でのやりとりを抽象化                                |
| internal/usecase      | Application Business Rule | ユースケースをカプセル化、実装し、エンティティを利用してユースケースを実現する |
| internal/domain       | Enterprise Business Rule  | ビジネスロジックをエンティティとして表現（objectやrepositoryの仕様を定義）     |

### レイヤーごとの考え方

#### infrastructure層

- ディレクトリ

```plain
infrastructure
├── domain_impl
│   ├── model
│   ├── repository
│   └── service
└── persistence
    └── database
```

- 役割
  - データベースや外部APIなど、アプリケーションの外部とのやり取りを担当
  - ドメインロジックに依存

#### adapter層

- ディレクトリ

```plain
# エントリーポイント
cmd
├── consumer
│   └── main.go
├── job
│   └── main.go
└── server
    └── main.go

# internal/adapter
internal
└── adapter
    ├── configuration
    ├── handler
    └── router
```

- 役割
  - `cmd`
    - アプリケーションのエントリーポイントとしてふるまう（main.go）
      - 外部からの各種リクエストを受けて処理を開始する
      - 環境変数読み込み、dbとのトランザクション確立、フレームワークの初期化、consumer起動などを行う
  - `internal/adapter`
    - 外部とのやり取りを内部のデータ構造に変換する役割を持つ

#### usecase層

- ディレクトリ

```plain
internal
└── usecase
```

- 役割
  - ビジネスロジックを実装する層で、システムのコアとなる
  - 外部システムからの入力や出力に依存することなく、ドメインモデルを操作する

#### domain

- ディレクトリ

```plain
internal
└── domain
    ├── model
    ├── repository
    └── service
```

- 役割
  - ビジネスロジックを定義
  - アプリケーションのビジネスルールを表すドメインモデルを実装
