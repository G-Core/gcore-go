# Changelog

## 0.9.0 (2025-07-31)

Full Changelog: [v0.8.0...v0.9.0](https://github.com/G-Core/gcore-go/compare/v0.8.0...v0.9.0)

### Features

* **api:** aggregated API specs update ([3599a2a](https://github.com/G-Core/gcore-go/commit/3599a2ab610a6aedb498b0427b877410caf35132))
* **api:** aggregated API specs update ([763343d](https://github.com/G-Core/gcore-go/commit/763343d73c3cc14cf923c5c4a58a530442c31fa3))
* **fastedge:** add binaries create method ([7b4ca39](https://github.com/G-Core/gcore-go/commit/7b4ca39e49487aca3dc3a222a9309309c1602f47))
* **security:** add security api ([f5d1461](https://github.com/G-Core/gcore-go/commit/f5d146115fb04f3d0a4a20c41ecf91be5f4c18a4))

## 0.8.0 (2025-07-29)

Full Changelog: [v0.7.0...v0.8.0](https://github.com/G-Core/gcore-go/compare/v0.7.0...v0.8.0)

### Features

* **api:** aggregated API specs update ([4f640ce](https://github.com/G-Core/gcore-go/commit/4f640ce257ee8738bfbaf3a9c92da307842b4d28))
* **api:** aggregated API specs update ([0255eb0](https://github.com/G-Core/gcore-go/commit/0255eb0a2bf68d657f98a914e46898c813198bb6))


### Bug Fixes

* **iam:** remove obsolete pagination scheme ([8a20964](https://github.com/G-Core/gcore-go/commit/8a209643372a49fa77cba95b823ea407ae9e9faf))
* **iam:** user model path ([e11de38](https://github.com/G-Core/gcore-go/commit/e11de38530c01f43186b6313ec2f567aae8505e0))

## 0.7.0 (2025-07-24)

Full Changelog: [v0.6.0...v0.7.0](https://github.com/G-Core/gcore-go/compare/v0.6.0...v0.7.0)

### Features

* **api:** aggregated API specs update ([d09df2a](https://github.com/G-Core/gcore-go/commit/d09df2a7a3d988a489f9131e416c16cbaa65231e))
* **api:** aggregated API specs update ([623ce0a](https://github.com/G-Core/gcore-go/commit/623ce0aff79f0cb17857cac57d3f587ebdfdb7dd))
* **cloud:** add cost and usage reports ([49f7536](https://github.com/G-Core/gcore-go/commit/49f75366c2af49d806a08262bc4198abc124960f))
* **streaming:** add streaming api ([eec84c8](https://github.com/G-Core/gcore-go/commit/eec84c865d4c958f61c776ce380348906cf79abd))


### Bug Fixes

* **client:** process custom base url ahead of time ([26874df](https://github.com/G-Core/gcore-go/commit/26874df5ef23f2f1d045ac2fbd5c344b5fb3481a))
* **cloud:** update example of replacing inference registry credential ([943b201](https://github.com/G-Core/gcore-go/commit/943b2016aa713d315696b1663c47876d018040fa))

## 0.6.0 (2025-07-21)

Full Changelog: [v0.5.0...v0.6.0](https://github.com/G-Core/gcore-go/compare/v0.5.0...v0.6.0)

### Features

* **api:** aggregated API specs update ([ec9c7b9](https://github.com/G-Core/gcore-go/commit/ec9c7b948b6fb3936525be03bf895274a5ed24fc))
* **api:** aggregated API specs update ([56dd867](https://github.com/G-Core/gcore-go/commit/56dd867a75bb81a68715556bb7bb1dfc5c703b9f))
* **cloud:** add audit logs ([31eb179](https://github.com/G-Core/gcore-go/commit/31eb17931efa44e0920a8eed7cac78a304438d03))
* **cloud:** add baremetal examples ([e876469](https://github.com/G-Core/gcore-go/commit/e876469cdfe033042ca684d694d4adaf72f08ea2))
* **cloud:** add inference api_keys subresource ([d025d80](https://github.com/G-Core/gcore-go/commit/d025d8027028069e55d91a552719bc17f9caa0f6))


### Bug Fixes

* **cloud:** make name optional in file share update example ([c8f03a8](https://github.com/G-Core/gcore-go/commit/c8f03a8552f417d6dda76f2ade6cb22b22887bad))

## 0.5.0 (2025-07-14)

Full Changelog: [v0.4.0...v0.5.0](https://github.com/G-Core/gcore-go/compare/v0.4.0...v0.5.0)

### ⚠ BREAKING CHANGES

* **cloud:** remove deprecated secrets NewAndPoll method
* **cloud:** refactor cloud inference models

### Features

* **api:** aggregated API specs update ([8e885ea](https://github.com/G-Core/gcore-go/commit/8e885ea585de4a91f4d5057d34d9e7852db233c4))
* **api:** aggregated API specs update ([93b112d](https://github.com/G-Core/gcore-go/commit/93b112d32a6b39130d3b091e12c4874ee72a1767))
* **api:** manual updates ([4ca75cc](https://github.com/G-Core/gcore-go/commit/4ca75cc27e04576fc4e9e369895c66a3122e64dd))
* **api:** manual upload of aggregated API specs ([7193379](https://github.com/G-Core/gcore-go/commit/7193379c170faf959cd54afc8528f6b1befad16c))
* **cloud:** add inference examples ([cdb2685](https://github.com/G-Core/gcore-go/commit/cdb268583ddfab449742b1adaf5cd56faa2ae721))
* **cloud:** add UploadTlsCertificateAndPoll method for secrets ([4c768f7](https://github.com/G-Core/gcore-go/commit/4c768f78f93ec6fce51031d91869a79bb457fa13))
* **fastedge:** add api ([b5c6ad4](https://github.com/G-Core/gcore-go/commit/b5c6ad4fd0a3499c2df99f2500e523290ee8ba78))


### Bug Fixes

* **cloud:** remove deprecated secrets NewAndPoll method ([e01efb7](https://github.com/G-Core/gcore-go/commit/e01efb726742a2bd99bd988d351de4b9f5860613))


### Chores

* **internal:** fix lint script for tests ([558c5da](https://github.com/G-Core/gcore-go/commit/558c5da71b5a1491814fc96a4999bc11d77a8d2a))
* lint tests ([a496baf](https://github.com/G-Core/gcore-go/commit/a496baf9678f5075048f319aa6736bbfeff0ee00))
* lint tests in subpackages ([1f71f43](https://github.com/G-Core/gcore-go/commit/1f71f434a3e048e47204d7548d745ddfe6363c45))


### Refactors

* **cloud:** refactor cloud inference models ([ab17e97](https://github.com/G-Core/gcore-go/commit/ab17e97c7e6fc3d72f487ddc83bcd52cf3c8a3d4))

## 0.4.0 (2025-07-04)

Full Changelog: [v0.3.0...v0.4.0](https://github.com/G-Core/gcore-go/compare/v0.3.0...v0.4.0)

### ⚠ BREAKING CHANGES

* **cloud:** remove list suitable from bm flavors
* remove list suitable and list for resize from instance flavors

### Features

* **api:** aggregated API specs update ([1e48d08](https://github.com/G-Core/gcore-go/commit/1e48d0817df211333522b45adba4a7db08ef5a14))
* **api:** aggregated API specs update ([6d3c18b](https://github.com/G-Core/gcore-go/commit/6d3c18b6b6d3d8bac27d9d2820233ef5bcfdad4c))
* **api:** aggregated API specs update ([f3a4e7b](https://github.com/G-Core/gcore-go/commit/f3a4e7bf1229e5924a6f4b0d31c8e8a7f7685a12))
* **api:** aggregated API specs update ([f91a4ef](https://github.com/G-Core/gcore-go/commit/f91a4ef15c3ff5143c5e1d8bc251f8491c826bb8))
* **api:** aggregated API specs update ([2546f6a](https://github.com/G-Core/gcore-go/commit/2546f6a5c3526783421aa65ddf5755411413f700))
* **api:** aggregated API specs update ([4b00db0](https://github.com/G-Core/gcore-go/commit/4b00db038beb3b1458912dcb8fb1c71054b3a80e))
* **client:** add escape hatch for null slice & maps ([a7eddc2](https://github.com/G-Core/gcore-go/commit/a7eddc28b2a70062147e288cd401e1a898856aab))
* **iam:** add IAM ([6257c1a](https://github.com/G-Core/gcore-go/commit/6257c1aa14437102e1662767c2ff2ce11a7a9475))


### Bug Fixes

* **cloud:** fix type in volume example ([c1a6818](https://github.com/G-Core/gcore-go/commit/c1a6818d94672ec1fa782131735f32da2862ed16))
* **cloud:** name is optional in update network ([a421627](https://github.com/G-Core/gcore-go/commit/a421627ad863d4abb176a8789dd6e83cb15db7d0))
* don't try to deserialize as json when ResponseBodyInto is []byte ([4183a08](https://github.com/G-Core/gcore-go/commit/4183a0838d99fbac53ad045a84c67dad832c9d5b))
* **pagination:** check if page data is empty in GetNextPage ([832e6f7](https://github.com/G-Core/gcore-go/commit/832e6f7c6173737adf43bc14877fcc3d014bc169))
* **waap:** remove duplicate method for acct overview ([ed1a8fe](https://github.com/G-Core/gcore-go/commit/ed1a8fe6fca7fdb2bf8225f0ca94ea896fbc0a1c))


### Chores

* **ci:** only run for pushes and fork pull requests ([4780623](https://github.com/G-Core/gcore-go/commit/478062347a5b54288f75082d31d065fea17f0673))
* **cloud:** use port id env in floating ip example ([#82](https://github.com/G-Core/gcore-go/issues/82)) ([2a3b963](https://github.com/G-Core/gcore-go/commit/2a3b96346d65905d4762ac61fb61c01aab524cf6))
* fix documentation of null map ([2911457](https://github.com/G-Core/gcore-go/commit/2911457a6b6b42cacc74b7dc096f0f8aa36fa784))
* **internal:** updates ([ba988ac](https://github.com/G-Core/gcore-go/commit/ba988ac5c0e5b7b4621c0e6a18926664a98ee52e))


### Refactors

* **cloud:** remove list suitable from bm flavors ([3869a14](https://github.com/G-Core/gcore-go/commit/3869a143a6ffd36745832116dbf22a7db3e5955a))
* remove list suitable and list for resize from instance flavors ([ab849fa](https://github.com/G-Core/gcore-go/commit/ab849fa81e81af3d2ee844cfbfd08a28ccd50a15))
* remove list suitable flavors from examples ([c6e0213](https://github.com/G-Core/gcore-go/commit/c6e021361a919d163480688efb63961994b2d789))

## 0.3.0 (2025-06-17)

Full Changelog: [v0.2.0...v0.3.0](https://github.com/G-Core/gcore-go/compare/v0.2.0...v0.3.0)

### Features

* **api:** aggregated API specs update ([c6e1502](https://github.com/G-Core/gcore-go/commit/c6e1502b417c5903d9e8c3e919267501eab56a8a))
* **api:** aggregated API specs update ([bc5dd64](https://github.com/G-Core/gcore-go/commit/bc5dd6460a07248d5661fc41aa8b5e7cb6161aaa))
* **api:** manual upload of aggregated API specs ([2ee8aeb](https://github.com/G-Core/gcore-go/commit/2ee8aebe900d1481d0a0feba573e803a532b8355))
* **api:** manual upload of aggregated API specs ([5ba0d9c](https://github.com/G-Core/gcore-go/commit/5ba0d9c30393d03c3b7e5ff4ae5861a77b257064))
* **client:** add debug log helper ([a0fb055](https://github.com/G-Core/gcore-go/commit/a0fb05502ff9ee51a2afb19d1f09622de361cb33))
* **client:** allow overriding unions ([3f5ba85](https://github.com/G-Core/gcore-go/commit/3f5ba8548708290885441fe4cd211055692ba340))
* **cloud,ssh_keys:** pass envs for SSH key examples ([#69](https://github.com/G-Core/gcore-go/issues/69)) ([04f9c92](https://github.com/G-Core/gcore-go/commit/04f9c921bac748ddd0955c88c7e22a57715c47ba))
* **cloud:** add file share access rules examples ([2c6c7c1](https://github.com/G-Core/gcore-go/commit/2c6c7c1aa66583a28f380a1a4c124c940a8f0d42))
* **cloud:** add file shares examples ([#60](https://github.com/G-Core/gcore-go/issues/60)) ([7c0534b](https://github.com/G-Core/gcore-go/commit/7c0534b2843926fcf5661d01c04cc40871aef780))
* **cloud:** add floating IPs example ([#67](https://github.com/G-Core/gcore-go/issues/67)) ([dda87a6](https://github.com/G-Core/gcore-go/commit/dda87a6dce05ed91dd882473ca8084a671efb76c))
* **cloud:** add images examples ([#61](https://github.com/G-Core/gcore-go/issues/61)) ([adc3dce](https://github.com/G-Core/gcore-go/commit/adc3dcef99d59c575246bde2ad3cf94a34c71174))
* **cloud:** add instances examples ([99a2e89](https://github.com/G-Core/gcore-go/commit/99a2e89c6eaa2b81d6b5a42dfc24432d89cc24cd))
* **cloud:** add loadbalancers examples ([9e55398](https://github.com/G-Core/gcore-go/commit/9e553989c28f3a137b4b13bff717926535b19edf))
* **cloud:** add networks and subnets examples ([7db4151](https://github.com/G-Core/gcore-go/commit/7db4151c459231048e09e9fcb08e26a4e914921e))
* **cloud:** add quotas examples ([#57](https://github.com/G-Core/gcore-go/issues/57)) ([d038e72](https://github.com/G-Core/gcore-go/commit/d038e72a2642d58c73322ef0fce7ffef27d606a4))
* **cloud:** add volumes examples ([#51](https://github.com/G-Core/gcore-go/issues/51)) ([fa4f0b0](https://github.com/G-Core/gcore-go/commit/fa4f0b04d3313c89f9a0a754ef9883123d24a8e1))
* **cloud:** implement routers examples ([#56](https://github.com/G-Core/gcore-go/issues/56)) ([74bec88](https://github.com/G-Core/gcore-go/commit/74bec8827e7af08bdbaa0de3669f13aa10263203))
* **cloud:** move all network examples into 1 dir ([af8160b](https://github.com/G-Core/gcore-go/commit/af8160bddab7fa7b3ff36aa66e0058e4957cfabb))
* **cloud:** rename network examples dir ([52940c6](https://github.com/G-Core/gcore-go/commit/52940c665ea0ef1d4663d4357ed70b9980663e12))
* **cloud:** rename tag in floating ips example ([#73](https://github.com/G-Core/gcore-go/issues/73)) ([c448846](https://github.com/G-Core/gcore-go/commit/c448846691fa4037e61a0bba9d6b2df0974e1e40))
* **cloud:** unify naming in examples ([#75](https://github.com/G-Core/gcore-go/issues/75)) ([0b7eb27](https://github.com/G-Core/gcore-go/commit/0b7eb271c2f5338352c59ddc25397a3342e1a391))
* **cloud:** use *AndPoll methods in images examples ([#74](https://github.com/G-Core/gcore-go/issues/74)) ([5a863ce](https://github.com/G-Core/gcore-go/commit/5a863ce75d42b38b006e5099562252828b920e61))
* **reserved_ips:** adapt examples for reserved fixed IPs ([#66](https://github.com/G-Core/gcore-go/issues/66)) ([239e47c](https://github.com/G-Core/gcore-go/commit/239e47cb4cdf5e9c172bd874f3da29fbca84edee))
* **security_groups:** implement security groups example ([#68](https://github.com/G-Core/gcore-go/issues/68)) ([245fad4](https://github.com/G-Core/gcore-go/commit/245fad46bdfb7aa9ecd92d12acb81e780b45f464))
* **waap:** add domain analytics, api_paths, insights and insight_silences; and ip_info ([2d8985f](https://github.com/G-Core/gcore-go/commit/2d8985fe47e02e8a6384f4064b78ea79d04b425a))
* **waap:** add domain custom, firewall and advanced rules; custom page sets, advanced rules and tags ([94727ea](https://github.com/G-Core/gcore-go/commit/94727ea439330ba09de47b4c226e15a692dc3113))


### Bug Fixes

* **client:** cast to raw message when converting to params ([e4b7ca1](https://github.com/G-Core/gcore-go/commit/e4b7ca1dc57ca9f24b8628a773e9588120897a6a))
* **cloud:** fix GCORE_BASE_URL env name in examples ([7ffed44](https://github.com/G-Core/gcore-go/commit/7ffed440d1afb7aaf50975e4b6616292f3046504))


### Chores

* **change-detection:** filter newly generated files ([f5afaca](https://github.com/G-Core/gcore-go/commit/f5afaca1ad7c39001c5f45e9bf2c7a86f7c13f5e))
* **ci:** enable for pull requests ([6b54e43](https://github.com/G-Core/gcore-go/commit/6b54e43c57e2660dbc6c42193f09f7f38ef62985))
* **cloud:** rename example files ([a4bd5bc](https://github.com/G-Core/gcore-go/commit/a4bd5bc009560b05d76a8e9fbe59a17f7ffed668))
* **cloud:** split examples into multiple files ([824c5ba](https://github.com/G-Core/gcore-go/commit/824c5ba10b20e2ee9fc5b9d179c9f9ce73036b86))

## 0.2.0 (2025-05-31)

Full Changelog: [v0.1.0...v0.2.0](https://github.com/G-Core/gcore-go/compare/v0.1.0...v0.2.0)

### Features

* **api:** aggregated API specs update ([305ca08](https://github.com/G-Core/gcore-go/commit/305ca085feb74b165189478e3db2b698c7a9a7a2))
* **api:** aggregated API specs update ([2773920](https://github.com/G-Core/gcore-go/commit/277392043efa9455f073245893bdafb63a5cd85a))
* **api:** aggregated API specs update ([baf23f3](https://github.com/G-Core/gcore-go/commit/baf23f364f20ea963c9c0bc27114c19418142f88))
* **api:** aggregated API specs update ([6edcefd](https://github.com/G-Core/gcore-go/commit/6edcefd749a376c763b6539ae0ced377de002d9f))
* **api:** aggregated API specs update ([b0ef96e](https://github.com/G-Core/gcore-go/commit/b0ef96e07096fb42689ec38775f29e9fb530c4c6))
* **api:** aggregated API specs update ([08a51f0](https://github.com/G-Core/gcore-go/commit/08a51f0af8f09cc2edad98519724f2063489e1df))
* **baremetal:** add polling methods ([#44](https://github.com/G-Core/gcore-go/issues/44)) ([8ecfae6](https://github.com/G-Core/gcore-go/commit/8ecfae61c990cef04009f147fb189577fdfa68eb))
* **client:** add support for endpoint-specific base URLs in python ([c4544bf](https://github.com/G-Core/gcore-go/commit/c4544bf5b6de1a6130122c82b86474cabb250ced))
* **gpu_cloud:** add polling methods ([#48](https://github.com/G-Core/gcore-go/issues/48)) ([29120d3](https://github.com/G-Core/gcore-go/commit/29120d343b3f5aebabeffef18a3af56d1d5bc29f))
* **inference:** add polling methods ([#47](https://github.com/G-Core/gcore-go/issues/47)) ([701ac2d](https://github.com/G-Core/gcore-go/commit/701ac2dec8b816a472e0557f4081fcafc1c9e107))
* **loadbalancers:** add polling methods ([#46](https://github.com/G-Core/gcore-go/issues/46)) ([06b83ed](https://github.com/G-Core/gcore-go/commit/06b83ed4db6b0f89bab7e7ba9402969d1884a1d2))


### Bug Fixes

* **ci:** do not always skip breaking change detection ([09df252](https://github.com/G-Core/gcore-go/commit/09df252f58c27a46a23edca36e53600976967aba))
* **client:** correctly set stream key for multipart ([12358b0](https://github.com/G-Core/gcore-go/commit/12358b02ba69ea667aabe16444bf9aba202790cf))
* **client:** don't panic on marshal with extra null field ([8d58bb5](https://github.com/G-Core/gcore-go/commit/8d58bb5b48c78f21fe33b7eae7d33d2bfabc1d0e))
* correct unmarshalling of root body params ([14f5785](https://github.com/G-Core/gcore-go/commit/14f57856f77acf5edde9453cdc1e0b572138475a))
* fix error ([3283861](https://github.com/G-Core/gcore-go/commit/32838616080bacdadd68cb58f38c8f366cefa6cd))
* **instances,baremetal,loadbalancers,inference,gpu_cloud:** don't fail if nr tasks gt 1 ([ab92204](https://github.com/G-Core/gcore-go/commit/ab9220404ac62b430b603adc7689356f75a8df4e))


### Chores

* **api:** mark some methods as deprecated ([44e481e](https://github.com/G-Core/gcore-go/commit/44e481e0661d563a2c184393af02a9099c4a6ca4))
* **docs:** grammar improvements ([d620989](https://github.com/G-Core/gcore-go/commit/d62098926b2b66925667a38e24527307d0d700a2))
* improve devcontainer setup ([327fbee](https://github.com/G-Core/gcore-go/commit/327fbee5d54ceee8cc2a0c7482d397193af95ca6))
* **internal:** codegen related update ([7eb57b9](https://github.com/G-Core/gcore-go/commit/7eb57b9cb996447db74f75f716f0486b3be75af5))
* make go mod tidy continue on error ([6312a23](https://github.com/G-Core/gcore-go/commit/6312a23250cdb8ecd2f6e949b880bbfb950adb76))


### Refactors

* **instances,secrets,reserved_ips:** harmonize polling methods ([#50](https://github.com/G-Core/gcore-go/issues/50)) ([918660c](https://github.com/G-Core/gcore-go/commit/918660c91fda57cfd3742c7c498aba223a4db2fe))
* **loadbalancers:** change oas schema names ([a3595c5](https://github.com/G-Core/gcore-go/commit/a3595c50e185b15d3cdda8bc5433b602ec670188))
* **loadbalancers:** use correct schema for loadbalancer pool ([7bedefc](https://github.com/G-Core/gcore-go/commit/7bedefc237d163e153ad2f0e1246316cff53b4f2))

## 0.1.0 (2025-05-08)

Full Changelog: [v0.1.0-alpha.2...v0.1.0](https://github.com/G-Core/gcore-go/compare/v0.1.0-alpha.2...v0.1.0)

### Features

* **api:** aggregated API specs update ([25ef71e](https://github.com/G-Core/gcore-go/commit/25ef71edc7712eb81350f2bf71d313d4192dcb33))
* **client:** experimental support for unmarshalling into param structs ([efcbafb](https://github.com/G-Core/gcore-go/commit/efcbafb002002b2b29ac5a8ba4f7f1481893420d))


### Bug Fixes

* **client:** clean up reader resources ([6a70d50](https://github.com/G-Core/gcore-go/commit/6a70d5044390fc7e459bda20118657ac5b94d5d6))
* **client:** correctly update body in WithJSONSet ([f9c6ace](https://github.com/G-Core/gcore-go/commit/f9c6ace9a743446071759c92731e07cc3c4ebbb5))
* **client:** unmarshal responses properly ([4be7a0c](https://github.com/G-Core/gcore-go/commit/4be7a0c66414b11a645f6f1660a58090bed5c0d3))

## 0.1.0-alpha.2 (2025-05-06)

Full Changelog: [v0.1.0-alpha.1...v0.1.0-alpha.2](https://github.com/G-Core/gcore-go/compare/v0.1.0-alpha.1...v0.1.0-alpha.2)

### Documentation

* update links ([f0389fc](https://github.com/G-Core/gcore-go/commit/f0389fcb22dde17bc38373d51cd0a01c618f0814))

## 0.1.0-alpha.1 (2025-05-06)

Full Changelog: [v0.0.1-alpha.0...v0.1.0-alpha.1](https://github.com/G-Core/gcore-go/compare/v0.0.1-alpha.0...v0.1.0-alpha.1)

### ⚠ BREAKING CHANGES

* **client:** rename resp package
* **client:** improve core function names
* **client:** better method root unions
* **client:** improve param subunions & deduplicate types

### Features

* **api:** add nested_params readme example ([878507f](https://github.com/G-Core/gcore-go/commit/878507f35f70fbc247d2fd7529ed1925b2f403d4))
* **api:** aggregated API specs update ([8a3c701](https://github.com/G-Core/gcore-go/commit/8a3c701e184a544db2408f3656e3e3bfe80449ab))
* **api:** aggregated API specs update ([016497c](https://github.com/G-Core/gcore-go/commit/016497c30b981d9e084cd8e0cc6d4f3fcb571133))
* **api:** aggregated API specs update ([36425c2](https://github.com/G-Core/gcore-go/commit/36425c23f26b71cf768b0d7abe0636cd7d2fab20))
* **api:** aggregated API specs update ([adf5ad9](https://github.com/G-Core/gcore-go/commit/adf5ad940c01bf7fc92ce70754eeea68145d83b3))
* **api:** aggregated API specs update ([e5c21e7](https://github.com/G-Core/gcore-go/commit/e5c21e7e341f04a07b3486ae8db7bed77c6deac4))
* **api:** aggregated API specs update ([30868dc](https://github.com/G-Core/gcore-go/commit/30868dc26d267ed5f8efda527dee35e2539aea14))
* **api:** aggregated API specs update ([947e8dd](https://github.com/G-Core/gcore-go/commit/947e8dd7597a763e396f44e31d8f47ccfbbad394))
* **api:** aggregated API specs update ([e1e05fc](https://github.com/G-Core/gcore-go/commit/e1e05fc01c88288b3e42b507d87bf8bce8368602))
* **api:** aggregated API specs update ([e6aeb64](https://github.com/G-Core/gcore-go/commit/e6aeb644d4a854a3dfb693f00052758715113390))
* **api:** aggregated API specs update ([703d896](https://github.com/G-Core/gcore-go/commit/703d896562734edf40afdb512acfbfa10dd0ead2))
* **api:** aggregated API specs update ([92672fe](https://github.com/G-Core/gcore-go/commit/92672fe282e8c6056e9653787be578927d6d686f))
* **api:** aggregated API specs update ([ffd390e](https://github.com/G-Core/gcore-go/commit/ffd390e842968b159faa524ab15835f11638046b))
* **api:** aggregated API specs update ([fd27c71](https://github.com/G-Core/gcore-go/commit/fd27c71e29acf3f07f0a59e9bdf777b568f8f8dc))
* **api:** aggregated API specs update ([3b44923](https://github.com/G-Core/gcore-go/commit/3b44923c0ecb13c9b95e407a6375db50026d8791))
* **api:** cloud and projects as standalone apis ([1c0809e](https://github.com/G-Core/gcore-go/commit/1c0809e940bdb3a6bb00f99b4dada861da1b61e4))
* **api:** Config update for algis-dumbris/cloud-quotas ([3930e04](https://github.com/G-Core/gcore-go/commit/3930e04b8b80cea1ead5458a6e36b9b77fd11167))
* **api:** manual updates ([5796ff2](https://github.com/G-Core/gcore-go/commit/5796ff2a6bf30e3a7325f3a3af045fa8ebdfb43b))
* **api:** manual updates ([7e746ef](https://github.com/G-Core/gcore-go/commit/7e746ef733032ede1b1846da7997c87f49ccbeed))
* **api:** manual updates ([f9e5ca1](https://github.com/G-Core/gcore-go/commit/f9e5ca1202f3c8e8178b472668a7bc5ffd8277a5))
* **api:** manual updates ([f48fb2f](https://github.com/G-Core/gcore-go/commit/f48fb2ffedb86a7da740e92fde370c545bc3eccc))
* **api:** manual updates ([1d93144](https://github.com/G-Core/gcore-go/commit/1d93144c59fd492f7c992c6ca03749b24c559567))
* **api:** manual updates ([90a0cbc](https://github.com/G-Core/gcore-go/commit/90a0cbc484eb7a5bf91fd63f6b09cfd914217cf5))
* **api:** manual updates ([8a916ba](https://github.com/G-Core/gcore-go/commit/8a916ba238944feade9ad4a815ed9f4904e58519))
* **api:** manual upload of aggregated API specs ([c7badf3](https://github.com/G-Core/gcore-go/commit/c7badf30801c408727d2bdc72e2fd8b26460997e))
* **api:** remove duplicates ([63f3f37](https://github.com/G-Core/gcore-go/commit/63f3f376ca81ab3831ec6322e98258e36e000ec4))
* **api:** remove quotas ([b1458a3](https://github.com/G-Core/gcore-go/commit/b1458a34db79da01062df79685baec52d061fd32))
* **api:** rename regions.retrieve() to rregions.get() ([a89247f](https://github.com/G-Core/gcore-go/commit/a89247f9ae0cb76e9e745bdd382be6b48ce8e945))
* **api:** simplify env vars ([dde81ad](https://github.com/G-Core/gcore-go/commit/dde81ad4d73c33259de4fc4900250a57b0db2a22))
* **api:** trigger codegen ([0ec3a6a](https://github.com/G-Core/gcore-go/commit/0ec3a6aa503e597bd567ebfe265374eaa322aab2))
* changes for loadbalancers-renaming ([623bee0](https://github.com/G-Core/gcore-go/commit/623bee0f1f0b799141ca55cb4a53a1e2ddaa454c))
* **client:** add escape hatch to omit required param fields ([6fca194](https://github.com/G-Core/gcore-go/commit/6fca194f5a18becad17443cc53f529b428fa4baf))
* **client:** add helper method to generate constant structs ([c035982](https://github.com/G-Core/gcore-go/commit/c0359824ee549004889d9b979306f3c51258c5e9))
* **client:** add support for reading base URL from environment variable ([59f386d](https://github.com/G-Core/gcore-go/commit/59f386d07dbebae8fa887a9cba9336d520a01c05))
* **client:** better method root unions ([77086ad](https://github.com/G-Core/gcore-go/commit/77086ad2de51a6442678df6990b859800f0dc05d))
* **client:** improve param subunions & deduplicate types ([c7b77f5](https://github.com/G-Core/gcore-go/commit/c7b77f5d457ccdf151ba22a7b33d1c3c0d4ddd28))
* **client:** rename resp package ([4a60dcf](https://github.com/G-Core/gcore-go/commit/4a60dcf2255f3ff85fefdaefc69366aa5319dd3d))
* **client:** support custom http clients ([368158e](https://github.com/G-Core/gcore-go/commit/368158ee2d4009b519f0208329eb342ced9f2b01))
* **client:** support more time formats ([2b0fb7d](https://github.com/G-Core/gcore-go/commit/2b0fb7d5992908e528868e563939080156bb96d7))
* **client:** support unions in query and forms ([8de2d1f](https://github.com/G-Core/gcore-go/commit/8de2d1f87ab5c39a01545e4148769bd1cae85378))
* **instances:** add polling methods ([#40](https://github.com/G-Core/gcore-go/issues/40)) ([a7f5912](https://github.com/G-Core/gcore-go/commit/a7f5912b09349784a3ab610862d9fa809befd3ae))
* **ipranges:** add examples ([#13](https://github.com/G-Core/gcore-go/issues/13)) ([8d75e9e](https://github.com/G-Core/gcore-go/commit/8d75e9e9e4a836958caa661e39a9c146fb17a8c3))
* **oas:** update to v14.150.0 ([a17fdda](https://github.com/G-Core/gcore-go/commit/a17fdda4b0544f5f12068f79f6edf0091c613ef9))
* **opts:** polling interval in secs ([2ee7b70](https://github.com/G-Core/gcore-go/commit/2ee7b7025009b8247b83f979608f6956ea22ebe6))
* **projects:** add examples ([#1](https://github.com/G-Core/gcore-go/issues/1)) ([2caa4e1](https://github.com/G-Core/gcore-go/commit/2caa4e17f0f173f2a7ccdbffd02f9c1d8ef7d377))
* **regions:** move regions examples ([#14](https://github.com/G-Core/gcore-go/issues/14)) ([e1e153d](https://github.com/G-Core/gcore-go/commit/e1e153da2561d56011597a0ef6b424b753787db5))
* **reserved_fixed_ips:** add NewAndPoll and examples ([#20](https://github.com/G-Core/gcore-go/issues/20)) ([68d35cb](https://github.com/G-Core/gcore-go/commit/68d35cbf93083eda1e032786ac25ffb5bdb67928))
* **secrets:** add NewAndPoll method ([#10](https://github.com/G-Core/gcore-go/issues/10)) ([b2fd175](https://github.com/G-Core/gcore-go/commit/b2fd175922fc67fcf0265858cbaba94c1bd1c271))
* **ssh_keys:** add examples ([#11](https://github.com/G-Core/gcore-go/issues/11)) ([188ba48](https://github.com/G-Core/gcore-go/commit/188ba482fbf15a3c85115a34ac3709c25a19653d))
* **tasks:** add Poll method ([#7](https://github.com/G-Core/gcore-go/issues/7)) ([a340e49](https://github.com/G-Core/gcore-go/commit/a340e49122606a0f7f87cf93dc630a56eb4d4fa8))
* **tasks:** make polling interval in secs ([#39](https://github.com/G-Core/gcore-go/issues/39)) ([2af355c](https://github.com/G-Core/gcore-go/commit/2af355c5b6da276bfe0becf1ed6df22efb608781))


### Bug Fixes

* **client:** custom code after opt rename ([#42](https://github.com/G-Core/gcore-go/issues/42)) ([91653b9](https://github.com/G-Core/gcore-go/commit/91653b94020111a7794a91dd5699024998e237b9))
* **client:** improve core function names ([894a36a](https://github.com/G-Core/gcore-go/commit/894a36a83fa39d15fda4079467967e52108dea68))
* **client:** improve docs ([276e21f](https://github.com/G-Core/gcore-go/commit/276e21f311ce902cadee07185aca28963c943ad4))
* **client:** renamed client opts ([88b66f9](https://github.com/G-Core/gcore-go/commit/88b66f9d35e9843ead00c0cdc22b723fc9d83793))
* **client:** resolve issue with optional multipart files ([627f832](https://github.com/G-Core/gcore-go/commit/627f83288e44c072289eb612d1db230d9bea2efc))
* **client:** time format encoding fix ([bf6ba85](https://github.com/G-Core/gcore-go/commit/bf6ba852dfb91a86d01fabe56240bc5f5b78f808))
* **cloud:** move and/or rename models ([2c247bf](https://github.com/G-Core/gcore-go/commit/2c247bfdc61cd4ee247f549c3da06507a9f3bf22))
* **cloud:** remove workaround ([bac0f24](https://github.com/G-Core/gcore-go/commit/bac0f24fd117cfadf3dd35b0c9e4e0803339f389))
* **examples:** adapt to breaking changes on root unions ([#34](https://github.com/G-Core/gcore-go/issues/34)) ([ace8d9e](https://github.com/G-Core/gcore-go/commit/ace8d9ece37e4202f7271832970bc63de8f6f9ab))
* **examples:** correct environment variable name for region ID and add task management example ([#36](https://github.com/G-Core/gcore-go/issues/36)) ([b4d5fb2](https://github.com/G-Core/gcore-go/commit/b4d5fb2fd12d9eb0654ff6542b4379a4f6548308))
* **examples:** update environment variable references for region ID and add task management example ([0c8820c](https://github.com/G-Core/gcore-go/commit/0c8820c453eed42876f5f736e3c1a85d3b21aee0))
* **floating_ips:** workaround ([c10823c](https://github.com/G-Core/gcore-go/commit/c10823c7a2114e1020003f6e1458f7030a375699))
* handle empty bodies in WithJSONSet ([9d4b1ed](https://github.com/G-Core/gcore-go/commit/9d4b1eda800f559c4fce4fffaaaaceb82a319d5d))
* **opts:** fix custom code after opts rename ([#31](https://github.com/G-Core/gcore-go/issues/31)) ([87edd17](https://github.com/G-Core/gcore-go/commit/87edd176bec1eb7dad1f2ad59bca4d008061a45d))
* **pagination:** handle errors when applying options ([335f127](https://github.com/G-Core/gcore-go/commit/335f12732a478c08a5ab0e8c3d267f780868cca2))
* **secrets:** validation on created resources ([#19](https://github.com/G-Core/gcore-go/issues/19)) ([fda9520](https://github.com/G-Core/gcore-go/commit/fda9520a869d505b8798d55a3080323c611b9bd7))


### Chores

* **ci:** add timeout thresholds for CI jobs ([b44aaf3](https://github.com/G-Core/gcore-go/commit/b44aaf373269ef194777330af36088183af340e5))
* **ci:** fix formatting for debug mode ([d155849](https://github.com/G-Core/gcore-go/commit/d155849935cf4c54c3e1ca68a5c97dc857443afd))
* **ci:** only use depot for staging repos ([99ad139](https://github.com/G-Core/gcore-go/commit/99ad13975e96e5eec072d189fd58e5ea104524d2))
* **docs:** doc improvements ([0d2424f](https://github.com/G-Core/gcore-go/commit/0d2424ff07e91f69aab325218bf97a29bba914be))
* **docs:** document pre-request options ([b8f54dc](https://github.com/G-Core/gcore-go/commit/b8f54dc89b331d134acde224eaac5b88eee3af62))
* **docs:** readme improvements ([eb5de32](https://github.com/G-Core/gcore-go/commit/eb5de326ca12ce76331e364b8781c593efb6d804))
* **docs:** update respjson package name ([9e3c117](https://github.com/G-Core/gcore-go/commit/9e3c117bbfee29be2f27e8b968b2efb4fd60d450))
* **internal:** codegen related update ([cf8cfac](https://github.com/G-Core/gcore-go/commit/cf8cfac3a4e653556ef272f02d2379309dc47a84))
* **internal:** expand CI branch coverage ([68f8aca](https://github.com/G-Core/gcore-go/commit/68f8aca3e44ecd041cbd0ec60966a3fc1df2b6c9))
* **internal:** reduce CI branch coverage ([c0fa359](https://github.com/G-Core/gcore-go/commit/c0fa35997a920d2c3f6b0a063d921fc367a2a793))
* **readme:** improve formatting ([4a0cb95](https://github.com/G-Core/gcore-go/commit/4a0cb956b84eefd5deef28e5a9a7706333fafd55))
* **tests:** improve enum examples ([bbfb3f3](https://github.com/G-Core/gcore-go/commit/bbfb3f3fba04026af90463a274f80eeca3a74997))
* update SDK settings ([311a915](https://github.com/G-Core/gcore-go/commit/311a915c45d60825490959237a9555c93975f036))
* update SDK settings ([e6f71ec](https://github.com/G-Core/gcore-go/commit/e6f71ecdc24f13b80087ed60275aef92466bc202))
* **utils:** add internal resp to param utility ([a2f0b87](https://github.com/G-Core/gcore-go/commit/a2f0b873458092537728b91e4d9bb1da9dad6011))


### Documentation

* update documentation links to be more uniform ([a1b4e08](https://github.com/G-Core/gcore-go/commit/a1b4e08a46d2b38684bb0e19470b468df379304e))
