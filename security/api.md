# Security

## Events

Response Types:

- <a href="https://pkg.go.dev/github.com/G-Core/gcore-go/security">security</a>.<a href="https://pkg.go.dev/github.com/G-Core/gcore-go/security#ClientView">ClientView</a>

Methods:

- <code title="get /security/notifier/v1/event_logs">client.Security.Events.<a href="https://pkg.go.dev/github.com/G-Core/gcore-go/security#EventService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/G-Core/gcore-go/security">security</a>.<a href="https://pkg.go.dev/github.com/G-Core/gcore-go/security#EventListParams">EventListParams</a>) (\*<a href="https://pkg.go.dev/github.com/G-Core/gcore-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/G-Core/gcore-go/packages/pagination#OffsetPage">OffsetPage</a>[<a href="https://pkg.go.dev/github.com/G-Core/gcore-go/security">security</a>.<a href="https://pkg.go.dev/github.com/G-Core/gcore-go/security#ClientView">ClientView</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## BgpAnnounces

Response Types:

- <a href="https://pkg.go.dev/github.com/G-Core/gcore-go/security">security</a>.<a href="https://pkg.go.dev/github.com/G-Core/gcore-go/security#ClientAnnounce">ClientAnnounce</a>
- <a href="https://pkg.go.dev/github.com/G-Core/gcore-go/security">security</a>.<a href="https://pkg.go.dev/github.com/G-Core/gcore-go/security#BgpAnnounceToggleResponse">BgpAnnounceToggleResponse</a>

Methods:

- <code title="get /security/sifter/v2/protected_addresses/announces">client.Security.BgpAnnounces.<a href="https://pkg.go.dev/github.com/G-Core/gcore-go/security#BgpAnnounceService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/G-Core/gcore-go/security">security</a>.<a href="https://pkg.go.dev/github.com/G-Core/gcore-go/security#BgpAnnounceListParams">BgpAnnounceListParams</a>) (\*[]<a href="https://pkg.go.dev/github.com/G-Core/gcore-go/security">security</a>.<a href="https://pkg.go.dev/github.com/G-Core/gcore-go/security#ClientAnnounce">ClientAnnounce</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /security/sifter/v2/protected_addresses/announces">client.Security.BgpAnnounces.<a href="https://pkg.go.dev/github.com/G-Core/gcore-go/security#BgpAnnounceService.Toggle">Toggle</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/G-Core/gcore-go/security">security</a>.<a href="https://pkg.go.dev/github.com/G-Core/gcore-go/security#BgpAnnounceToggleParams">BgpAnnounceToggleParams</a>) (\*<a href="https://pkg.go.dev/github.com/G-Core/gcore-go/security">security</a>.<a href="https://pkg.go.dev/github.com/G-Core/gcore-go/security#BgpAnnounceToggleResponse">BgpAnnounceToggleResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## ProfileTemplates

Response Types:

- <a href="https://pkg.go.dev/github.com/G-Core/gcore-go/security">security</a>.<a href="https://pkg.go.dev/github.com/G-Core/gcore-go/security#ClientProfileTemplate">ClientProfileTemplate</a>

Methods:

- <code title="get /security/iaas/profile-templates">client.Security.ProfileTemplates.<a href="https://pkg.go.dev/github.com/G-Core/gcore-go/security#ProfileTemplateService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*[]<a href="https://pkg.go.dev/github.com/G-Core/gcore-go/security">security</a>.<a href="https://pkg.go.dev/github.com/G-Core/gcore-go/security#ClientProfileTemplate">ClientProfileTemplate</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Profiles

Response Types:

- <a href="https://pkg.go.dev/github.com/G-Core/gcore-go/security">security</a>.<a href="https://pkg.go.dev/github.com/G-Core/gcore-go/security#ClientProfile">ClientProfile</a>

Methods:

- <code title="post /security/iaas/v2/profiles">client.Security.Profiles.<a href="https://pkg.go.dev/github.com/G-Core/gcore-go/security#ProfileService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/G-Core/gcore-go/security">security</a>.<a href="https://pkg.go.dev/github.com/G-Core/gcore-go/security#ProfileNewParams">ProfileNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/G-Core/gcore-go/security">security</a>.<a href="https://pkg.go.dev/github.com/G-Core/gcore-go/security#ClientProfile">ClientProfile</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /security/iaas/v2/profiles">client.Security.Profiles.<a href="https://pkg.go.dev/github.com/G-Core/gcore-go/security#ProfileService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/G-Core/gcore-go/security">security</a>.<a href="https://pkg.go.dev/github.com/G-Core/gcore-go/security#ProfileListParams">ProfileListParams</a>) (\*[]<a href="https://pkg.go.dev/github.com/G-Core/gcore-go/security">security</a>.<a href="https://pkg.go.dev/github.com/G-Core/gcore-go/security#ClientProfile">ClientProfile</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /security/iaas/v2/profiles/{id}">client.Security.Profiles.<a href="https://pkg.go.dev/github.com/G-Core/gcore-go/security#ProfileService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#int64">int64</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="get /security/iaas/v2/profiles/{id}">client.Security.Profiles.<a href="https://pkg.go.dev/github.com/G-Core/gcore-go/security#ProfileService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#int64">int64</a>) (\*<a href="https://pkg.go.dev/github.com/G-Core/gcore-go/security">security</a>.<a href="https://pkg.go.dev/github.com/G-Core/gcore-go/security#ClientProfile">ClientProfile</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /security/iaas/v2/profiles/{id}/recreate">client.Security.Profiles.<a href="https://pkg.go.dev/github.com/G-Core/gcore-go/security#ProfileService.Recreate">Recreate</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#int64">int64</a>, body <a href="https://pkg.go.dev/github.com/G-Core/gcore-go/security">security</a>.<a href="https://pkg.go.dev/github.com/G-Core/gcore-go/security#ProfileRecreateParams">ProfileRecreateParams</a>) (\*<a href="https://pkg.go.dev/github.com/G-Core/gcore-go/security">security</a>.<a href="https://pkg.go.dev/github.com/G-Core/gcore-go/security#ClientProfile">ClientProfile</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /security/iaas/v2/profiles/{id}">client.Security.Profiles.<a href="https://pkg.go.dev/github.com/G-Core/gcore-go/security#ProfileService.Replace">Replace</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#int64">int64</a>, body <a href="https://pkg.go.dev/github.com/G-Core/gcore-go/security">security</a>.<a href="https://pkg.go.dev/github.com/G-Core/gcore-go/security#ProfileReplaceParams">ProfileReplaceParams</a>) (\*<a href="https://pkg.go.dev/github.com/G-Core/gcore-go/security">security</a>.<a href="https://pkg.go.dev/github.com/G-Core/gcore-go/security#ClientProfile">ClientProfile</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
