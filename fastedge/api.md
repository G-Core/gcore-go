# Fastedge

Response Types:

- <a href="https://pkg.go.dev/github.com/G-Core/gcore-go/fastedge">fastedge</a>.<a href="https://pkg.go.dev/github.com/G-Core/gcore-go/fastedge#Client">Client</a>

Methods:

- <code title="get /fastedge/v1/me">client.Fastedge.<a href="https://pkg.go.dev/github.com/G-Core/gcore-go/fastedge#FastedgeService.GetAccountOverview">GetAccountOverview</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/G-Core/gcore-go/fastedge">fastedge</a>.<a href="https://pkg.go.dev/github.com/G-Core/gcore-go/fastedge#Client">Client</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Templates

Params Types:

- <a href="https://pkg.go.dev/github.com/G-Core/gcore-go/fastedge">fastedge</a>.<a href="https://pkg.go.dev/github.com/G-Core/gcore-go/fastedge#TemplateParam">TemplateParam</a>
- <a href="https://pkg.go.dev/github.com/G-Core/gcore-go/fastedge">fastedge</a>.<a href="https://pkg.go.dev/github.com/G-Core/gcore-go/fastedge#TemplateParameter">TemplateParameter</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/G-Core/gcore-go/fastedge">fastedge</a>.<a href="https://pkg.go.dev/github.com/G-Core/gcore-go/fastedge#TemplateShort">TemplateShort</a>

Methods:

- <code title="post /fastedge/v1/template">client.Fastedge.Templates.<a href="https://pkg.go.dev/github.com/G-Core/gcore-go/fastedge#TemplateService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/G-Core/gcore-go/fastedge">fastedge</a>.<a href="https://pkg.go.dev/github.com/G-Core/gcore-go/fastedge#TemplateNewParams">TemplateNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/G-Core/gcore-go/fastedge">fastedge</a>.<a href="https://pkg.go.dev/github.com/G-Core/gcore-go/fastedge#TemplateShort">TemplateShort</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /fastedge/v1/template">client.Fastedge.Templates.<a href="https://pkg.go.dev/github.com/G-Core/gcore-go/fastedge#TemplateService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/G-Core/gcore-go/fastedge">fastedge</a>.<a href="https://pkg.go.dev/github.com/G-Core/gcore-go/fastedge#TemplateListParams">TemplateListParams</a>) (\*<a href="https://pkg.go.dev/github.com/G-Core/gcore-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/G-Core/gcore-go/packages/pagination#OffsetPageFastedgeTemplates">OffsetPageFastedgeTemplates</a>[<a href="https://pkg.go.dev/github.com/G-Core/gcore-go/fastedge">fastedge</a>.<a href="https://pkg.go.dev/github.com/G-Core/gcore-go/fastedge#TemplateShort">TemplateShort</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Secrets

Params Types:

- <a href="https://pkg.go.dev/github.com/G-Core/gcore-go/fastedge">fastedge</a>.<a href="https://pkg.go.dev/github.com/G-Core/gcore-go/fastedge#SecretParam">SecretParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/G-Core/gcore-go/fastedge">fastedge</a>.<a href="https://pkg.go.dev/github.com/G-Core/gcore-go/fastedge#Secret">Secret</a>
- <a href="https://pkg.go.dev/github.com/G-Core/gcore-go/fastedge">fastedge</a>.<a href="https://pkg.go.dev/github.com/G-Core/gcore-go/fastedge#SecretShort">SecretShort</a>
- <a href="https://pkg.go.dev/github.com/G-Core/gcore-go/fastedge">fastedge</a>.<a href="https://pkg.go.dev/github.com/G-Core/gcore-go/fastedge#SecretNewResponse">SecretNewResponse</a>
- <a href="https://pkg.go.dev/github.com/G-Core/gcore-go/fastedge">fastedge</a>.<a href="https://pkg.go.dev/github.com/G-Core/gcore-go/fastedge#SecretListResponse">SecretListResponse</a>

Methods:

- <code title="post /fastedge/v1/secrets">client.Fastedge.Secrets.<a href="https://pkg.go.dev/github.com/G-Core/gcore-go/fastedge#SecretService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/G-Core/gcore-go/fastedge">fastedge</a>.<a href="https://pkg.go.dev/github.com/G-Core/gcore-go/fastedge#SecretNewParams">SecretNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/G-Core/gcore-go/fastedge">fastedge</a>.<a href="https://pkg.go.dev/github.com/G-Core/gcore-go/fastedge#SecretNewResponse">SecretNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /fastedge/v1/secrets">client.Fastedge.Secrets.<a href="https://pkg.go.dev/github.com/G-Core/gcore-go/fastedge#SecretService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/G-Core/gcore-go/fastedge">fastedge</a>.<a href="https://pkg.go.dev/github.com/G-Core/gcore-go/fastedge#SecretListParams">SecretListParams</a>) (\*<a href="https://pkg.go.dev/github.com/G-Core/gcore-go/fastedge">fastedge</a>.<a href="https://pkg.go.dev/github.com/G-Core/gcore-go/fastedge#SecretListResponse">SecretListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Binaries

Response Types:

- <a href="https://pkg.go.dev/github.com/G-Core/gcore-go/fastedge">fastedge</a>.<a href="https://pkg.go.dev/github.com/G-Core/gcore-go/fastedge#BinaryShort">BinaryShort</a>
- <a href="https://pkg.go.dev/github.com/G-Core/gcore-go/fastedge">fastedge</a>.<a href="https://pkg.go.dev/github.com/G-Core/gcore-go/fastedge#BinaryListResponse">BinaryListResponse</a>

Methods:

- <code title="post /fastedge/v1/binaries/raw">client.Fastedge.Binaries.<a href="https://pkg.go.dev/github.com/G-Core/gcore-go/fastedge#BinaryService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/builtin#io.Reader">io.Reader</a>) (\*<a href="https://pkg.go.dev/github.com/G-Core/gcore-go/fastedge">fastedge</a>.<a href="https://pkg.go.dev/github.com/G-Core/gcore-go/fastedge#BinaryShort">BinaryShort</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /fastedge/v1/binaries">client.Fastedge.Binaries.<a href="https://pkg.go.dev/github.com/G-Core/gcore-go/fastedge#BinaryService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/G-Core/gcore-go/fastedge">fastedge</a>.<a href="https://pkg.go.dev/github.com/G-Core/gcore-go/fastedge#BinaryListResponse">BinaryListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Statistics

Response Types:

- <a href="https://pkg.go.dev/github.com/G-Core/gcore-go/fastedge">fastedge</a>.<a href="https://pkg.go.dev/github.com/G-Core/gcore-go/fastedge#CallStatus">CallStatus</a>
- <a href="https://pkg.go.dev/github.com/G-Core/gcore-go/fastedge">fastedge</a>.<a href="https://pkg.go.dev/github.com/G-Core/gcore-go/fastedge#DurationStats">DurationStats</a>
- <a href="https://pkg.go.dev/github.com/G-Core/gcore-go/fastedge">fastedge</a>.<a href="https://pkg.go.dev/github.com/G-Core/gcore-go/fastedge#StatisticGetCallSeriesResponse">StatisticGetCallSeriesResponse</a>
- <a href="https://pkg.go.dev/github.com/G-Core/gcore-go/fastedge">fastedge</a>.<a href="https://pkg.go.dev/github.com/G-Core/gcore-go/fastedge#StatisticGetDurationSeriesResponse">StatisticGetDurationSeriesResponse</a>

Methods:

- <code title="get /fastedge/v1/stats/calls">client.Fastedge.Statistics.<a href="https://pkg.go.dev/github.com/G-Core/gcore-go/fastedge#StatisticService.GetCallSeries">GetCallSeries</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/G-Core/gcore-go/fastedge">fastedge</a>.<a href="https://pkg.go.dev/github.com/G-Core/gcore-go/fastedge#StatisticGetCallSeriesParams">StatisticGetCallSeriesParams</a>) (\*<a href="https://pkg.go.dev/github.com/G-Core/gcore-go/fastedge">fastedge</a>.<a href="https://pkg.go.dev/github.com/G-Core/gcore-go/fastedge#StatisticGetCallSeriesResponse">StatisticGetCallSeriesResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /fastedge/v1/stats/app_duration">client.Fastedge.Statistics.<a href="https://pkg.go.dev/github.com/G-Core/gcore-go/fastedge#StatisticService.GetDurationSeries">GetDurationSeries</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/G-Core/gcore-go/fastedge">fastedge</a>.<a href="https://pkg.go.dev/github.com/G-Core/gcore-go/fastedge#StatisticGetDurationSeriesParams">StatisticGetDurationSeriesParams</a>) (\*<a href="https://pkg.go.dev/github.com/G-Core/gcore-go/fastedge">fastedge</a>.<a href="https://pkg.go.dev/github.com/G-Core/gcore-go/fastedge#StatisticGetDurationSeriesResponse">StatisticGetDurationSeriesResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Apps

Params Types:

- <a href="https://pkg.go.dev/github.com/G-Core/gcore-go/fastedge">fastedge</a>.<a href="https://pkg.go.dev/github.com/G-Core/gcore-go/fastedge#AppParam">AppParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/G-Core/gcore-go/fastedge">fastedge</a>.<a href="https://pkg.go.dev/github.com/G-Core/gcore-go/fastedge#AppShort">AppShort</a>

Methods:

- <code title="post /fastedge/v1/apps">client.Fastedge.Apps.<a href="https://pkg.go.dev/github.com/G-Core/gcore-go/fastedge#AppService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/G-Core/gcore-go/fastedge">fastedge</a>.<a href="https://pkg.go.dev/github.com/G-Core/gcore-go/fastedge#AppNewParams">AppNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/G-Core/gcore-go/fastedge">fastedge</a>.<a href="https://pkg.go.dev/github.com/G-Core/gcore-go/fastedge#AppShort">AppShort</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /fastedge/v1/apps">client.Fastedge.Apps.<a href="https://pkg.go.dev/github.com/G-Core/gcore-go/fastedge#AppService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/G-Core/gcore-go/fastedge">fastedge</a>.<a href="https://pkg.go.dev/github.com/G-Core/gcore-go/fastedge#AppListParams">AppListParams</a>) (\*<a href="https://pkg.go.dev/github.com/G-Core/gcore-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/G-Core/gcore-go/packages/pagination#OffsetPageFastedgeApps">OffsetPageFastedgeApps</a>[<a href="https://pkg.go.dev/github.com/G-Core/gcore-go/fastedge">fastedge</a>.<a href="https://pkg.go.dev/github.com/G-Core/gcore-go/fastedge#AppShort">AppShort</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Logs

## KvStores

Params Types:

- <a href="https://pkg.go.dev/github.com/G-Core/gcore-go/fastedge">fastedge</a>.<a href="https://pkg.go.dev/github.com/G-Core/gcore-go/fastedge#KvStoreParam">KvStoreParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/G-Core/gcore-go/fastedge">fastedge</a>.<a href="https://pkg.go.dev/github.com/G-Core/gcore-go/fastedge#KvStore">KvStore</a>
- <a href="https://pkg.go.dev/github.com/G-Core/gcore-go/fastedge">fastedge</a>.<a href="https://pkg.go.dev/github.com/G-Core/gcore-go/fastedge#KvStoreShort">KvStoreShort</a>
- <a href="https://pkg.go.dev/github.com/G-Core/gcore-go/fastedge">fastedge</a>.<a href="https://pkg.go.dev/github.com/G-Core/gcore-go/fastedge#KvStoreNewResponse">KvStoreNewResponse</a>
- <a href="https://pkg.go.dev/github.com/G-Core/gcore-go/fastedge">fastedge</a>.<a href="https://pkg.go.dev/github.com/G-Core/gcore-go/fastedge#KvStoreListResponse">KvStoreListResponse</a>

Methods:

- <code title="post /fastedge/v1/kv">client.Fastedge.KvStores.<a href="https://pkg.go.dev/github.com/G-Core/gcore-go/fastedge#KvStoreService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/G-Core/gcore-go/fastedge">fastedge</a>.<a href="https://pkg.go.dev/github.com/G-Core/gcore-go/fastedge#KvStoreNewParams">KvStoreNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/G-Core/gcore-go/fastedge">fastedge</a>.<a href="https://pkg.go.dev/github.com/G-Core/gcore-go/fastedge#KvStoreNewResponse">KvStoreNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /fastedge/v1/kv">client.Fastedge.KvStores.<a href="https://pkg.go.dev/github.com/G-Core/gcore-go/fastedge#KvStoreService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/G-Core/gcore-go/fastedge">fastedge</a>.<a href="https://pkg.go.dev/github.com/G-Core/gcore-go/fastedge#KvStoreListParams">KvStoreListParams</a>) (\*<a href="https://pkg.go.dev/github.com/G-Core/gcore-go/fastedge">fastedge</a>.<a href="https://pkg.go.dev/github.com/G-Core/gcore-go/fastedge#KvStoreListResponse">KvStoreListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
