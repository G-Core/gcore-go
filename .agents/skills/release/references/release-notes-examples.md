# Release Notes Examples

Real examples from this repository showing the target format. Use these as
style references when generating human-readable release notes (Part 1).

## Example 1: v0.32.0 (multiple product areas, no breaking changes section)

```markdown
We're excited to announce version 0.32.0!

### **CDN**

* **CDN Resources**
  * Added `DeactivateAndDelete()` convenience method — deactivates a CDN resource then deletes it
  * ⚠ BREAKING CHANGE: `CDNResourceNewParams.Origin` changed from `string` (required) to `param.Opt[string]` (optional)
  * ⚠ BREAKING CHANGE: `CDNResourceNewParams.OriginGroup` changed from `int64` (required) to `param.Opt[int64]` (optional) — exactly one of `Origin` or `OriginGroup` must be provided

* **Logs Uploader**
  * Added `Endpoint` field to S3 OSS config — custom S3 endpoint now supported in `LogsUploaderTarget` and related params

### **Cloud**

* **Load Balancers**
  * Added `AdminStateUp` field across all load balancer, listener, pool, and health monitor response and param types
  * ⚠ BREAKING CHANGE: `LoadBalancerNewParamsListener.SecretID` changed from `param.Opt[string]` to `string` with empty-string enum validation
  * ⚠ BREAKING CHANGE: `LoadBalancerListenerNewParams.SecretID` changed from `param.Opt[string]` to `LoadBalancerListenerNewParamsSecretID` string type
  * Deprecated `LoadBalancers.Update()` — use PATCH `/v2/loadbalancers/...` instead
  * Deprecated `NetworkRouters.Update()` — use PATCH `/v2/routers/...` instead

* **Managed Kubernetes**
  * Added `IncludeCapacity` param to `K8SFlavorListParams` for flavor capacity info
  * Added `ClusterRebuild` and `ClusterServerRebuild` task states
  * ⚠ BREAKING CHANGE: `K8SClusterKubeconfig.CreatedAt` and `ExpiresAt` changed from `nullable` to `required`

### **FastEdge**

* ⚠ BREAKING CHANGE: `App.Stores` type changed from `map[string]int64` to `map[string]AppStore` — stores now carry `ID`, `Name`, and `Comment`
* ⚠ BREAKING CHANGE: `KvStore` response restructured — `ID` field removed, `Name` is now primary identifier, `Updated` renamed to `UpdatedAt`, added `Revision` and `Size` fields
* ⚠ BREAKING CHANGE: `KvStoreService.New` return type changed from `*KvStore` to `*KvStoreNewResponse`
* ⚠ BREAKING CHANGE: `KvStoreService.Get` return type changed from `*KvStoreGetResponse` to `*KvStore`
* ⚠ BREAKING CHANGE: Removed `KvStoreStats` and `KvStoreGetResponse` types
* ⚠ BREAKING CHANGE: Removed `AppLimit`, `DailyLimit`, `HourlyLimit` fields from `Client` type
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
  * ⚠ BREAKING CHANGE: Renamed all `Cdn*` types to `CDN*` for consistent Go naming conventions - `CdnService` is now `CDNService`, `CdnAccount` is now `CDNAccount`, `CdnAccountLimits` is now `CDNAccountLimits`, etc.
  * ⚠ BREAKING CHANGE: Renamed `Resources` service to `CDNResources` - access CDN resources via `client.CDN.CDNResources` instead of `client.CDN.Resources`

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
5. **Breaking changes**: Inline with `⚠ BREAKING CHANGE:` prefix, include old type/value -> new type/value
6. **Deprecations**: Use format `Deprecated \`Method()\` — use \`Alternative()\` instead`
7. **New fields/methods**: Use format `Added \`FieldName\` field to \`TypeName\` — description`
8. **Fixes**: Use format `Fixed {description} — {detail}`
9. **Type references**: Always use backtick-wrapped Go identifiers (e.g., `param.Opt[string]`)
10. **Descriptions**: Short, specific, user-actionable. No commit hashes in Part 1.
11. **Product area order**: Alphabetical. Place "Other" last.
12. **Sub-area order**: Alphabetical within each product area.
