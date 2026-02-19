# Iam

Params Types:

- <a href="https://pkg.go.dev/github.com/G-Core/gcore-go/iam">iam</a>.<a href="https://pkg.go.dev/github.com/G-Core/gcore-go/iam#AuthType">AuthType</a>
- <a href="https://pkg.go.dev/github.com/G-Core/gcore-go/iam">iam</a>.<a href="https://pkg.go.dev/github.com/G-Core/gcore-go/iam#UserGroupParam">UserGroupParam</a>
- <a href="https://pkg.go.dev/github.com/G-Core/gcore-go/iam">iam</a>.<a href="https://pkg.go.dev/github.com/G-Core/gcore-go/iam#UserLanguage">UserLanguage</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/G-Core/gcore-go/iam">iam</a>.<a href="https://pkg.go.dev/github.com/G-Core/gcore-go/iam#AccountOverview">AccountOverview</a>
- <a href="https://pkg.go.dev/github.com/G-Core/gcore-go/iam">iam</a>.<a href="https://pkg.go.dev/github.com/G-Core/gcore-go/iam#AuthType">AuthType</a>
- <a href="https://pkg.go.dev/github.com/G-Core/gcore-go/iam">iam</a>.<a href="https://pkg.go.dev/github.com/G-Core/gcore-go/iam#UserGroup">UserGroup</a>
- <a href="https://pkg.go.dev/github.com/G-Core/gcore-go/iam">iam</a>.<a href="https://pkg.go.dev/github.com/G-Core/gcore-go/iam#UserLanguage">UserLanguage</a>

Methods:

- <code title="get /iam/clients/me">client.Iam.<a href="https://pkg.go.dev/github.com/G-Core/gcore-go/iam#IamService.GetAccountOverview">GetAccountOverview</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/G-Core/gcore-go/iam">iam</a>.<a href="https://pkg.go.dev/github.com/G-Core/gcore-go/iam#AccountOverview">AccountOverview</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## APITokens

Response Types:

- <a href="https://pkg.go.dev/github.com/G-Core/gcore-go/iam">iam</a>.<a href="https://pkg.go.dev/github.com/G-Core/gcore-go/iam#APIToken">APIToken</a>
- <a href="https://pkg.go.dev/github.com/G-Core/gcore-go/iam">iam</a>.<a href="https://pkg.go.dev/github.com/G-Core/gcore-go/iam#APITokenCreated">APITokenCreated</a>
- <a href="https://pkg.go.dev/github.com/G-Core/gcore-go/iam">iam</a>.<a href="https://pkg.go.dev/github.com/G-Core/gcore-go/iam#APITokenList">APITokenList</a>

Methods:

- <code title="post /iam/clients/{clientId}/tokens">client.Iam.APITokens.<a href="https://pkg.go.dev/github.com/G-Core/gcore-go/iam#APITokenService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, clientID <a href="https://pkg.go.dev/builtin#int64">int64</a>, body <a href="https://pkg.go.dev/github.com/G-Core/gcore-go/iam">iam</a>.<a href="https://pkg.go.dev/github.com/G-Core/gcore-go/iam#APITokenNewParams">APITokenNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/G-Core/gcore-go/iam">iam</a>.<a href="https://pkg.go.dev/github.com/G-Core/gcore-go/iam#APITokenCreated">APITokenCreated</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /iam/clients/{clientId}/tokens">client.Iam.APITokens.<a href="https://pkg.go.dev/github.com/G-Core/gcore-go/iam#APITokenService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, clientID <a href="https://pkg.go.dev/builtin#int64">int64</a>, query <a href="https://pkg.go.dev/github.com/G-Core/gcore-go/iam">iam</a>.<a href="https://pkg.go.dev/github.com/G-Core/gcore-go/iam#APITokenListParams">APITokenListParams</a>) (\*<a href="https://pkg.go.dev/github.com/G-Core/gcore-go/iam">iam</a>.<a href="https://pkg.go.dev/github.com/G-Core/gcore-go/iam#APITokenList">APITokenList</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /iam/clients/{clientId}/tokens/{tokenId}">client.Iam.APITokens.<a href="https://pkg.go.dev/github.com/G-Core/gcore-go/iam#APITokenService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, tokenID <a href="https://pkg.go.dev/builtin#int64">int64</a>, body <a href="https://pkg.go.dev/github.com/G-Core/gcore-go/iam">iam</a>.<a href="https://pkg.go.dev/github.com/G-Core/gcore-go/iam#APITokenDeleteParams">APITokenDeleteParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="get /iam/clients/{clientId}/tokens/{tokenId}">client.Iam.APITokens.<a href="https://pkg.go.dev/github.com/G-Core/gcore-go/iam#APITokenService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, tokenID <a href="https://pkg.go.dev/builtin#int64">int64</a>, query <a href="https://pkg.go.dev/github.com/G-Core/gcore-go/iam">iam</a>.<a href="https://pkg.go.dev/github.com/G-Core/gcore-go/iam#APITokenGetParams">APITokenGetParams</a>) (\*<a href="https://pkg.go.dev/github.com/G-Core/gcore-go/iam">iam</a>.<a href="https://pkg.go.dev/github.com/G-Core/gcore-go/iam#APIToken">APIToken</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Users

Response Types:

- <a href="https://pkg.go.dev/github.com/G-Core/gcore-go/iam">iam</a>.<a href="https://pkg.go.dev/github.com/G-Core/gcore-go/iam#User">User</a>
- <a href="https://pkg.go.dev/github.com/G-Core/gcore-go/iam">iam</a>.<a href="https://pkg.go.dev/github.com/G-Core/gcore-go/iam#UserInvited">UserInvited</a>
- <a href="https://pkg.go.dev/github.com/G-Core/gcore-go/iam">iam</a>.<a href="https://pkg.go.dev/github.com/G-Core/gcore-go/iam#UserType">UserType</a>

Methods:

- <code title="patch /iam/users/{userId}">client.Iam.Users.<a href="https://pkg.go.dev/github.com/G-Core/gcore-go/iam#UserService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, userID <a href="https://pkg.go.dev/builtin#int64">int64</a>, body <a href="https://pkg.go.dev/github.com/G-Core/gcore-go/iam">iam</a>.<a href="https://pkg.go.dev/github.com/G-Core/gcore-go/iam#UserUpdateParams">UserUpdateParams</a>) (\*<a href="https://pkg.go.dev/github.com/G-Core/gcore-go/iam">iam</a>.<a href="https://pkg.go.dev/github.com/G-Core/gcore-go/iam#User">User</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /iam/users">client.Iam.Users.<a href="https://pkg.go.dev/github.com/G-Core/gcore-go/iam#UserService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/G-Core/gcore-go/iam">iam</a>.<a href="https://pkg.go.dev/github.com/G-Core/gcore-go/iam#UserListParams">UserListParams</a>) (\*<a href="https://pkg.go.dev/github.com/G-Core/gcore-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/G-Core/gcore-go/packages/pagination#OffsetPage">OffsetPage</a>[<a href="https://pkg.go.dev/github.com/G-Core/gcore-go/iam">iam</a>.<a href="https://pkg.go.dev/github.com/G-Core/gcore-go/iam#User">User</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /iam/clients/{clientId}/client-users/{userId}">client.Iam.Users.<a href="https://pkg.go.dev/github.com/G-Core/gcore-go/iam#UserService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, userID <a href="https://pkg.go.dev/builtin#int64">int64</a>, body <a href="https://pkg.go.dev/github.com/G-Core/gcore-go/iam">iam</a>.<a href="https://pkg.go.dev/github.com/G-Core/gcore-go/iam#UserDeleteParams">UserDeleteParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="get /iam/users/{userId}">client.Iam.Users.<a href="https://pkg.go.dev/github.com/G-Core/gcore-go/iam#UserService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, userID <a href="https://pkg.go.dev/builtin#int64">int64</a>) (\*<a href="https://pkg.go.dev/github.com/G-Core/gcore-go/iam">iam</a>.<a href="https://pkg.go.dev/github.com/G-Core/gcore-go/iam#User">User</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /iam/clients/invite_user">client.Iam.Users.<a href="https://pkg.go.dev/github.com/G-Core/gcore-go/iam#UserService.Invite">Invite</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/G-Core/gcore-go/iam">iam</a>.<a href="https://pkg.go.dev/github.com/G-Core/gcore-go/iam#UserInviteParams">UserInviteParams</a>) (\*<a href="https://pkg.go.dev/github.com/G-Core/gcore-go/iam">iam</a>.<a href="https://pkg.go.dev/github.com/G-Core/gcore-go/iam#UserInvited">UserInvited</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
