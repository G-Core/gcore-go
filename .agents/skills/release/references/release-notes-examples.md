# Release Notes Examples

Real examples from this repository showing the target format. Use these as
style references when generating human-readable release notes (Part 1).

## Example 1: v0.32.0 (multiple product areas, no breaking changes section)

```markdown
We're excited to announce version 0.32.0!

### **CDN**

* **CDN Resources**
  * ⚠ BREAKING CHANGE: `CDNResourceNewParams` fields `Origin` and `OriginGroup` are now optional — exactly one must be provided
  * Added `DeactivateAndDelete()` convenience method — deactivates a CDN resource then deletes it

* **Logs Uploader**
  * Added `Endpoint` field to S3 OSS config — custom S3 endpoint now supported in `LogsUploaderTarget` and related params

### **Cloud**

* **Load Balancers**
  * ⚠ BREAKING CHANGE: `SecretID` in load balancer and listener create params changed from `param.Opt[string]` to a `string` type with empty-string enum validation
  * Added `AdminStateUp` field across all load balancer, listener, pool, and health monitor response and param types
  * Deprecated `LoadBalancers.Update()` — use PATCH `/v2/loadbalancers/...` instead
  * Deprecated `NetworkRouters.Update()` — use PATCH `/v2/routers/...` instead

* **Managed Kubernetes**
  * ⚠ BREAKING CHANGE: `K8SClusterKubeconfig.CreatedAt` and `ExpiresAt` changed from nullable to required
  * Added `IncludeCapacity` param to `K8SFlavorListParams` for flavor capacity info
  * Added `ClusterRebuild` and `ClusterServerRebuild` task states

### **FastEdge**

* ⚠ BREAKING CHANGE: Edge Storage API changes
  * `App.Stores` changed from `map[string]int64` to `map[string]AppStore` — stores now carry `ID`, `Name`, and `Comment`
  * `KvStore` restructured — `ID` removed, `Name` is now the primary identifier, `Updated` renamed to `UpdatedAt`, `Revision` and `Size` added
  * `KvStoreService.New` and `Get` now return `*KvStoreNewResponse` and `*KvStore`
  * Removed `KvStoreStats`, `KvStoreGetResponse`, and the `AppLimit`, `DailyLimit`, `HourlyLimit` fields from `Client`
* Renamed "KV Storage" to "Edge Storage" across all endpoints
* Added `store` as new `TemplateParameterDataType` enum value

### **WAAP**

* **Analytics**
  * Added `Decision` and `OptionalAction` fields to `WaapRequestDetails` and `WaapRequestSummary`
  * Added `DomainID` field to `WaapRequestSummary`
  * Deprecated `DomainStatistics.GetTrafficSeries()` — use `GET /v1/analytics/traffic` instead
```

## Example 2: v0.31.0 (smaller release, refactors + features)

```markdown
We're excited to announce version 0.31.0!

### **CDN**

* **API Naming**
  * ⚠ BREAKING CHANGE: Renamed types and services for consistent Go naming
    * All `Cdn*` types are now `CDN*` (for example, `CdnService` is now `CDNService`)
    * The `Resources` service is now `CDNResources` — use `client.CDN.CDNResources` instead of `client.CDN.Resources`

* **User-Agent ACL**
  * Added regex pattern support - you can now use regular expressions in User-Agent ACL rules with `~` (case-sensitive) or `~*` (case-insensitive) prefix

### **Cloud**

* **GPU Bare Metal**
  * Fixed cluster ID extraction in `NewAndPoll()` for GPU baremetal clusters - properly extracts cluster ID from response
  * Added GPU Cloud examples demonstrating usage patterns

### **WAAP**

* **Analytics**
  * Deprecated `GetRequestsSeries()` and `GetRequestsSeriesAutoPaging()` methods - use the new `/v1/analytics/requests` endpoint instead
  * Updated action filter values - changed from "block", "captcha", "handshake", "monitor" to "allow", "block", "captcha", "handshake"

### **Other**

* Added `param.SetJSON()` helper — convenient function for setting JSON values in request parameters
```

## Example 3: v0.33.0 (no breaking changes, fine-grained sub-areas)

```markdown
We're excited to announce version 0.33.0!

### **CDN**

* **Logs Uploader**
  * Added `LogSampleRate` field to `LogsUploaderPolicy`, `LogsUploaderPolicyNewParams`, `LogsUploaderPolicyUpdateParams`, and `LogsUploaderPolicyReplaceParams` — controls the fraction of log entries to collect (0 to 1)

### **Cloud**

* **Bare Metal**
  * Added `ReservedCapacity` field to `BaremetalFlavor` — shows available instances from reservations

* **Everywhere Inference**
  * Fixed option filtering in `InferenceDeploymentService.NewAndPoll()`, `.DeleteAndPoll()`, and `.UpdateAndPoll()` — properly excludes `WithResponseBodyInto` and clears request body for intermediate Poll/Get calls

* **GPU Bare Metal**
  * Added `ReservedCapacity` field to `GPUBaremetalFlavorBareMetalGPUFlavorsChemaWithoutPrice` and `GPUBaremetalFlavorBareMetalGPUFlavorsChemaWithPrice` — shows available instances from reservations

* **Security Groups**
  * Added `SecurityGroupRules` field to `TaskCreatedResources` — tracks created security group rule IDs
  * Added `create_security_group_rule` and `delete_security_group_rule` task types

### **WAAP**

* **Analytics**
  * Added `SessionID` field to `WaapRequestDetails` and `WaapRequestSummary` — the session ID associated with the request

* **Rules**
  * `Allow`, `Captcha`, `Handshake`, and `Monitor` action fields changed from `any` to `map[string]any` in `WaapAdvancedRuleAction`, `WaapCustomRuleAction`, and `WaapFirewallRuleAction` (and corresponding param types)
```

## Style Rules (inferred from examples)

1. **Opening line**: Always `We're excited to announce version {VERSION}!`
2. **Product area headers**: `### **{Area}**` — bold inside h3
3. **Sub-area items**: `* **{Sub-area}**` — bold bullet, followed by indented child bullets
4. **When a product area has no sub-areas** (e.g., FastEdge above): list items directly under the product header without a sub-area bullet
5. **Breaking changes**: One `⚠ BREAKING CHANGE:` bullet per sub-area with a short summary; details go in nested bullets without the prefix. Include old type/value -> new type/value. Never repeat the prefix on each bullet
6. **Deprecations**: Use format `Deprecated \`Method()\` — use \`Alternative()\` instead`
7. **New fields/methods**: Use format `Added \`FieldName\` field to \`TypeName\` — description`
8. **Fixes**: Use format `Fixed {description} — {detail}`
9. **Type references**: Always use backtick-wrapped Go identifiers (e.g., `param.Opt[string]`)
10. **Descriptions**: Short, specific, user-actionable. No commit hashes in Part 1.
11. **Product area order**: Alphabetical. Place "Other" last.
12. **Sub-area order**: Alphabetical within each product area.

## Breaking Changes Example

Preferred (one marker, nested details, few identifiers):

```markdown
* **Users**
  * ⚠ BREAKING CHANGE: User methods and response models now follow the IAM v2 API
    * `Get()` now returns `*User` instead of `*UserGetResponse`
    * `Update()` now returns `*UserUpdated` instead of `*User`
  * Added `GetMe()` — returns the authenticated user's profile, clients, and roles
```

Avoid (prefix on every bullet, too many types):

```markdown
* **Users**
  * ⚠ BREAKING CHANGE: `Get()` now returns `*User` instead of `*UserGetResponse`
  * ⚠ BREAKING CHANGE: `Update()` now returns `*UserUpdated` instead of `*User`
  * ⚠ BREAKING CHANGE: Removed `UserGetResponse`, `UserStats`, and `UserLimits` types
```
