Генерация таблиц

```shell
ent new --target internal/domain/schema [TableName]
```

Генерация орм

```shell
ent generate --target ./pkg/ent ./internal/domain/schema
```

Перед использованием надо поставить ent локально, либо запускать через go run