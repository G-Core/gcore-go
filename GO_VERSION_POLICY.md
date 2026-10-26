# Go version support policy

**Status: proposed.** This draft is for review. It does not raise the SDK's minimum Go version.
The support and testing commitments below take effect after the policy and its CI coverage are adopted.

## Supported Go versions

The SDK supports the current stable Go release family and the previous stable family,
following the [Go release policy][1]. We test the latest patch release of each family.
Release candidates are outside this support guarantee.

When a new stable Go family is released, it enters the support window and the oldest family leaves it.
Users should install the latest patch release of a supported family to receive Go security fixes.

## Minimum required Go version

The `go` directive in [`go.mod`](go.mod) defines the minimum Go version required to build the SDK.
This is separate from the supported release window and the toolchain used in CI.
See the [Go toolchain documentation][2] for how this requirement affects downstream modules.

We may retain a lower minimum to avoid forcing unnecessary upgrades on SDK users.
A new Go release or an upstream end-of-support date does not automatically raise this minimum.

We raise the minimum when there is a concrete reason, such as:

- A useful SDK feature needs newer Go language features or standard-library APIs.
- A necessary dependency update requires a newer Go version.
- Removing compatibility code provides a clear maintenance benefit.
- A Go fix is needed for correct or secure SDK behavior.

Minimum-version changes must remain compatible with the two supported Go families.
Any exception that drops a supported family requires an explicit policy change and explanation.

Older Go versions may still build the SDK, but receive no ongoing support or security guarantee.
We check the declared minimum in a separate compatibility job, with `GOTOOLCHAIN=local` to prevent automatic upgrades.
Passing this job confirms build and test compatibility; it does not make an unsupported Go release secure.

## Announcing minimum-version changes

We announce planned increases at least 30 days before release, in a public repository issue and the changelog.
Minimum-version increases ship in SDK minor releases. Routine patch releases preserve the existing minimum.
Under this policy, a Go minimum increase alone does not require a new SDK major version.

Each announcement and release note identifies:

- The new minimum Go version and the reason for the change.
- The first SDK release that requires it.
- The last SDK release compatible with the previous minimum.

An urgent security fix may require shorter notice or an increase in a patch release.
We explain the exception and the required upgrade steps in the release notes.

Previously published SDK versions remain available, but receive no guaranteed maintenance or security backports.

## CI and security updates

We update CI toolchains promptly when security patches become available.
The CI toolchain is configured independently of the `go.mod` minimum.

The SDK is distributed as source code. Applications use the Go standard library from the toolchain that builds them.
Users must rebuild and redeploy affected applications with a patched Go toolchain.
Updating SDK CI or the SDK dependency alone does not patch an existing application binary.

## Current state and adoption

As of October 26, 2026:

| Item | Current state |
| --- | --- |
| Declared SDK minimum | Go 1.22 |
| Upstream-supported Go families | Go 1.26 and Go 1.27 |
| Main CI lint and test toolchain | Go 1.26.10 |

Before adopting this policy, add tests for both supported Go families and the declared minimum.
Preserve the policy and README link through the `stlc-go` generation workflow.
Review the support window after each stable Go release and keep CI patch versions current.

This proposal does not schedule a minimum-version increase.
Any later increase follows the criteria and notice process above.

[1]: https://go.dev/doc/devel/release#policy
[2]: https://go.dev/doc/toolchain
