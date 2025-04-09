# Cloud

## Projects

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/gcore-go/cloud">cloud</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/gcore-go/cloud#Project">Project</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/gcore-go/cloud">cloud</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/gcore-go/cloud#ProjectListResponse">ProjectListResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/gcore-go/cloud">cloud</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/gcore-go/cloud#ProjectDeleteResponse">ProjectDeleteResponse</a>

Methods:

- <code title="post /cloud/v1/projects">client.Cloud.Projects.<a href="https://pkg.go.dev/github.com/stainless-sdks/gcore-go/cloud#ProjectService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/gcore-go/cloud">cloud</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/gcore-go/cloud#ProjectNewParams">ProjectNewParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/gcore-go/cloud">cloud</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/gcore-go/cloud#Project">Project</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /cloud/v1/projects/{project_id}">client.Cloud.Projects.<a href="https://pkg.go.dev/github.com/stainless-sdks/gcore-go/cloud#ProjectService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/gcore-go/cloud">cloud</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/gcore-go/cloud#ProjectGetParams">ProjectGetParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/gcore-go/cloud">cloud</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/gcore-go/cloud#Project">Project</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /cloud/v1/projects/{project_id}">client.Cloud.Projects.<a href="https://pkg.go.dev/github.com/stainless-sdks/gcore-go/cloud#ProjectService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/stainless-sdks/gcore-go/cloud">cloud</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/gcore-go/cloud#ProjectUpdateParams">ProjectUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/gcore-go/cloud">cloud</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/gcore-go/cloud#Project">Project</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /cloud/v1/projects">client.Cloud.Projects.<a href="https://pkg.go.dev/github.com/stainless-sdks/gcore-go/cloud#ProjectService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/gcore-go/cloud">cloud</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/gcore-go/cloud#ProjectListParams">ProjectListParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/gcore-go/cloud">cloud</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/gcore-go/cloud#ProjectListResponse">ProjectListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /cloud/v1/projects/{project_id}">client.Cloud.Projects.<a href="https://pkg.go.dev/github.com/stainless-sdks/gcore-go/cloud#ProjectService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/gcore-go/cloud">cloud</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/gcore-go/cloud#ProjectDeleteParams">ProjectDeleteParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/gcore-go/cloud">cloud</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/gcore-go/cloud#ProjectDeleteResponse">ProjectDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Quotas

### ClientQuotas

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/gcore-go/cloud">cloud</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/gcore-go/cloud#AllClientQuotas">AllClientQuotas</a>

Methods:

- <code title="get /cloud/v2/client_quotas">client.Cloud.Quotas.ClientQuotas.<a href="https://pkg.go.dev/github.com/stainless-sdks/gcore-go/cloud#QuotaClientQuotaService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/gcore-go/cloud">cloud</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/gcore-go/cloud#AllClientQuotas">AllClientQuotas</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Global

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/gcore-go/cloud">cloud</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/gcore-go/cloud#GlobalQuotas">GlobalQuotas</a>

Methods:

- <code title="get /cloud/v2/global_quotas/{client_id}">client.Cloud.Quotas.Global.<a href="https://pkg.go.dev/github.com/stainless-sdks/gcore-go/cloud#QuotaGlobalService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, clientID <a href="https://pkg.go.dev/builtin#int64">int64</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/gcore-go/cloud">cloud</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/gcore-go/cloud#GlobalQuotas">GlobalQuotas</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Regional

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/gcore-go/cloud">cloud</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/gcore-go/cloud#RegionalQuotas">RegionalQuotas</a>

Methods:

- <code title="get /cloud/v2/regional_quotas/{client_id}/{region_id}">client.Cloud.Quotas.Regional.<a href="https://pkg.go.dev/github.com/stainless-sdks/gcore-go/cloud#QuotaRegionalService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, clientID <a href="https://pkg.go.dev/builtin#int64">int64</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/gcore-go/cloud">cloud</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/gcore-go/cloud#QuotaRegionalGetParams">QuotaRegionalGetParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/gcore-go/cloud">cloud</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/gcore-go/cloud#RegionalQuotas">RegionalQuotas</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### LimitsRequest

Params Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/gcore-go/cloud">cloud</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/gcore-go/cloud#LimitsRequestCreateParam">LimitsRequestCreateParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/gcore-go/cloud">cloud</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/gcore-go/cloud#LimitsRequest">LimitsRequest</a>

Methods:

- <code title="post /cloud/v2/limits_request">client.Cloud.Quotas.LimitsRequest.<a href="https://pkg.go.dev/github.com/stainless-sdks/gcore-go/cloud#QuotaLimitsRequestService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/gcore-go/cloud">cloud</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/gcore-go/cloud#QuotaLimitsRequestNewParams">QuotaLimitsRequestNewParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="get /cloud/v2/limits_request/{request_id}">client.Cloud.Quotas.LimitsRequest.<a href="https://pkg.go.dev/github.com/stainless-sdks/gcore-go/cloud#QuotaLimitsRequestService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, requestID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/gcore-go/cloud">cloud</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/gcore-go/cloud#LimitsRequest">LimitsRequest</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /cloud/v2/limits_request">client.Cloud.Quotas.LimitsRequest.<a href="https://pkg.go.dev/github.com/stainless-sdks/gcore-go/cloud#QuotaLimitsRequestService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/gcore-go/cloud">cloud</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/gcore-go/cloud#QuotaLimitsRequestListParams">QuotaLimitsRequestListParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="delete /cloud/v2/limits_request/{request_id}">client.Cloud.Quotas.LimitsRequest.<a href="https://pkg.go.dev/github.com/stainless-sdks/gcore-go/cloud#QuotaLimitsRequestService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, requestID <a href="https://pkg.go.dev/builtin#string">string</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>

## Regions

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/gcore-go/cloud">cloud</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/gcore-go/cloud#Region">Region</a>

Methods:

- <code title="get /cloud/v1/regions/{region_id}">client.Cloud.Regions.<a href="https://pkg.go.dev/github.com/stainless-sdks/gcore-go/cloud#RegionService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/stainless-sdks/gcore-go/cloud">cloud</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/gcore-go/cloud#RegionGetParams">RegionGetParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/gcore-go/cloud">cloud</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/gcore-go/cloud#Region">Region</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /cloud/v1/regions">client.Cloud.Regions.<a href="https://pkg.go.dev/github.com/stainless-sdks/gcore-go/cloud#RegionService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/gcore-go/cloud">cloud</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/gcore-go/cloud#RegionListParams">RegionListParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/gcore-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/gcore-go/packages/pagination#OffsetPage">OffsetPage</a>[<a href="https://pkg.go.dev/github.com/stainless-sdks/gcore-go/cloud">cloud</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/gcore-go/cloud#Region">Region</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
