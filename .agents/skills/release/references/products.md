# Gcore Products — Canonical Names & File Path Heuristics

Authoritative list of product and sub-product names for release notes.
Order products alphabetically; place "Other" last. Order sub-products
alphabetically within each product.

## CDN

SDK package: `cdn/`

| Sub-Product | File patterns |
|---|---|
| Activity Logs | `cdn/auditlog*` |
| Advanced Analytics | `cdn/advancedanalytic*` |
| CA Certificates | `cdn/cacertificate*` |
| CDN Resources | `cdn/cdnresource*` (not rules, not shields) |
| CDN Service | `cdn/cdn.go` (service-level config) |
| IP Ranges | `cdn/iprange*` |
| Let's Encrypt Certificates | `cdn/letsencrypt*` |
| Log Viewer | `cdn/log.go`, `cdn/log_test.go` |
| Logs Uploader | `cdn/logsuploader*` |
| Origin Shielding | `cdn/cdnresourceshield*` |
| Origins | `cdn/origin*` |
| Purge History | `cdn/purgehistory*` |
| Rule Templates | `cdn/ruletemplate*` |
| Rules | `cdn/cdnresourcerule*` |
| SSL Certificates | `cdn/certificate*` |
| Statistics | `cdn/statistic*` |
| Tools | `cdn/tool*`, `cdn/prefetch*`, `cdn/purge*` |

## Cloud

SDK package: `cloud/`

| Sub-Product | File patterns |
|---|---|
| Bare Metal | `cloud/baremetal*` |
| Container as a Service | `cloud/caas*`, `cloud/container*` |
| Cost Reports | `cloud/costreport*`, `cloud/usagereport*` |
| DDoS Protection | `cloud/ddos*` |
| Everywhere Inference | `cloud/inference*` (not `inferenceapplication*`) |
| Everywhere Inference Apps | `cloud/inferenceapplication*` |
| File Shares | `cloud/fileshare*` |
| Floating IPs | `cloud/floatingip*` |
| Function as a Service | `cloud/faas*`, `cloud/function*` |
| GPU Bare Metal | `cloud/gpubaremetal*` |
| GPU Virtual | `cloud/gpuvirtual*` |
| Images | `cloud/instanceimage*`, `cloud/baremetalimage*` |
| Instances | `cloud/instance*` (not `instanceimage*`) |
| IP Ranges | `cloud/iprange*` |
| Load Balancers | `cloud/loadbalancer*` |
| Logging | `cloud/logging*` |
| Managed Kubernetes | `cloud/k8s*` |
| Managed PostgreSQL | `cloud/database*` |
| Networks | `cloud/network*` (not `networkrouter*`) |
| Placement Groups | `cloud/placementgroup*` |
| Projects | `cloud/project*` |
| Quotas | `cloud/quota*` |
| Regions | `cloud/region*` |
| Registry | `cloud/registry*` |
| Reservations | `cloud/billingreservation*` |
| Reserved IPs | `cloud/reservedfixedip*` |
| Routers | `cloud/networkrouter*` |
| Secrets | `cloud/secret*` |
| Security Groups | `cloud/securitygroup*` |
| Snapshot Schedules | `cloud/snapshotschedule*` |
| Snapshots | `cloud/volumesnapshot*` |
| SSH Keys | `cloud/sshkey*` |
| Tasks | `cloud/task*` |
| User Actions | `cloud/auditlog*` |
| User Role Assignments | `cloud/userroleassignment*` |
| Volumes | `cloud/volume*` (not `volumesnapshot*`) |

## DDoS Protection

SDK package: `security/`

| Sub-Product | File patterns |
|---|---|
| BGP Announces | `security/bgpannounce*` |
| Event Logs | `security/event*` |
| Profiles | `security/profile*` (not `profiletemplate*`) |
| Security Templates | `security/profiletemplate*` |

## DNS

SDK package: `dns/`

| Sub-Product | File patterns |
|---|---|
| DNSSEC | `dns/zonednssec*` |
| Locations | `dns/location*` |
| Metrics | `dns/metric*` |
| Network Mappings | `dns/networkmapping*` |
| Pickers | `dns/picker*` |
| RRsets | `dns/zonerrset*` |
| Zones | `dns/zone.go`, `dns/zone_test.go` |

## FastEdge

SDK package: `fastedge/`

| Sub-Product | File patterns |
|---|---|
| App Logs | `fastedge/applog*` |
| Apps | `fastedge/app*` (not `applog*`) |
| Binaries | `fastedge/binary*` |
| Edge Storage | `fastedge/kvstore*` |
| Secrets | `fastedge/secret*` |
| Statistics | `fastedge/statistic*` |
| Templates | `fastedge/template*` |

## IAM

SDK package: `iam/`

| Sub-Product | File patterns |
|---|---|
| Account | `iam/iam.go` (service-level auth) |
| API Tokens | `iam/apitoken*` |
| Users | `iam/user*` |

## Object Storage

SDK package: `storage/`

| Sub-Product | File patterns |
|---|---|
| Buckets | `storage/bucket*` (not `bucketcor*`, `bucketlifecycle*`, `bucketpolicy*`) |
| CORS | `storage/bucketcor*` |
| Credentials | `storage/credential*` |
| Lifecycle | `storage/bucketlifecycle*` |
| Locations | `storage/location*` |
| Policies | `storage/bucketpolicy*` |
| Statistics | `storage/statistic*` |

## Streaming

SDK package: `streaming/`

| Sub-Product | File patterns |
|---|---|
| AI | `streaming/aitask*` |
| Broadcasts | `streaming/broadcast*` |
| Directories | `streaming/directory*` |
| Overlays | `streaming/streamoverlay*` |
| Players | `streaming/player*` |
| Playlists | `streaming/playlist*` |
| Quality Sets | `streaming/qualityset*` |
| Restreams | `streaming/restream*` |
| Statistics | `streaming/statistic*` |
| Streams | `streaming/stream.go`, `streaming/stream_test.go` |
| Subtitles | `streaming/videosubtitle*` |
| Videos | `streaming/video.go`, `streaming/video_test.go` |

## WAAP

SDK package: `waap/`

| Sub-Product | File patterns |
|---|---|
| Advanced Rules | `waap/advancedrule*`, `waap/domainadvancedrule*` |
| Analytics | `waap/domainstatistic*`, `waap/statistic*` |
| API Discovery | `waap/domainapidiscovery*`, `waap/domainapipath*` |
| Custom Page Sets | `waap/custompageset*` |
| Custom Rules | `waap/domaincustomrule*` |
| Domains | `waap/domain.go`, `waap/domain_test.go`, `waap/domainsetting*` |
| Firewall Rules | `waap/domainfirewallrule*` |
| IP Spotlight | `waap/ipinfo*` |
| Network Organizations | `waap/organization*` |
| Security Insights | `waap/domaininsight*`, `waap/insight*` |
| Tags | `waap/tag*` |
| WAAP Service | `waap/waap.go` |

## Other

Changes not specific to any product (SDK internals, utilities, shared code).

| File patterns |
|---|
| `.github/*`, `.stats.yml`, `go.mod`, `go.sum`, `README.md` |
| `client.go`, `field.go`, `aliases.go` (root level) |
| `packages/*`, `internal/*`, `option/*`, `shared/*`, `lib/*` |
