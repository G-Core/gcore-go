# Changelog

## 0.19.0 (2025-11-07)

Full Changelog: [v0.18.0...v0.19.0](https://github.com/G-Core/gcore-go/compare/v0.18.0...v0.19.0)

### Features

* **api:** aggregated API specs update ([843bd35](https://github.com/G-Core/gcore-go/commit/843bd35bdccd4d12990e0e11a1e946e5db0880a3))
* **api:** aggregated API specs update ([9dc8473](https://github.com/G-Core/gcore-go/commit/9dc84730118a792f55b46d05cffe8202e3cc20c6))
* **api:** aggregated API specs update ([003cd96](https://github.com/G-Core/gcore-go/commit/003cd966ae4af17a478366171561898f17c4faeb))
* **api:** aggregated API specs update ([d73045f](https://github.com/G-Core/gcore-go/commit/d73045fda8dd183f066933290c7801088af3ba4c))


### Bug Fixes

* **examples:** update storage location to string ([e9ea78e](https://github.com/G-Core/gcore-go/commit/e9ea78e94a6876fe8b0295647db1eb6eb90a0363))
* remove readonly parameters from request params ([825a766](https://github.com/G-Core/gcore-go/commit/825a766a2cbfeff545359b277d2e6d1025d28ebc))

## 0.18.0 (2025-11-04)

Full Changelog: [v0.17.0...v0.18.0](https://github.com/G-Core/gcore-go/compare/v0.17.0...v0.18.0)

### Features

* **api:** aggregated API specs update ([410ee99](https://github.com/G-Core/gcore-go/commit/410ee99998501fae43838121e4ae6cef7efaf379))
* **api:** aggregated API specs update ([bbecad5](https://github.com/G-Core/gcore-go/commit/bbecad533162983e18ef394c1225c5623aefb35d))
* **api:** aggregated API specs update ([16853b1](https://github.com/G-Core/gcore-go/commit/16853b10dfb9fd3dc6b922296eaa821fd2cd3e68))
* **api:** aggregated API specs update ([d6ea9ad](https://github.com/G-Core/gcore-go/commit/d6ea9ad1c3c42caf8aa946479524af7b1e31b826))
* **api:** aggregated API specs update ([4c37826](https://github.com/G-Core/gcore-go/commit/4c3782693de7e65edb98fd743b0794b5f1e4253f))
* **api:** aggregated API specs update ([417769e](https://github.com/G-Core/gcore-go/commit/417769e6cff42f234a7a8ee053970d23a0078ec3))
* **api:** aggregated API specs update ([2202759](https://github.com/G-Core/gcore-go/commit/220275985846733c431453ef8d9b539a9aeb0839))
* **api:** aggregated API specs update ([c10e1b8](https://github.com/G-Core/gcore-go/commit/c10e1b8e5ee2b5a00192aadf638c1ee11d3a27d9))
* **api:** aggregated API specs update ([993f01a](https://github.com/G-Core/gcore-go/commit/993f01a2a86fb6e0f35d626c4a345d81022c4bbf))
* **cloud:** add polling methods to postgres clusters ([3d06488](https://github.com/G-Core/gcore-go/commit/3d06488064032dc77f1c589376c5fa28ff12c03b))
* **cloud:** add support for postgres ([da38400](https://github.com/G-Core/gcore-go/commit/da384003a4237eec8cec624f279d6cbf5be5f383))


### Bug Fixes

* **client:** make sure to import param package when used ([a689cfd](https://github.com/G-Core/gcore-go/commit/a689cfd3b21e99e0e8841eeb68c9c97eafbd650a))
* **cloud:** internal task service initialization in instance ([8dc5dd0](https://github.com/G-Core/gcore-go/commit/8dc5dd0ff71119dadbab850df4acd8b3c6071bd8))


### Chores

* **internal:** grammar fix (it's -&gt; its) ([53e9730](https://github.com/G-Core/gcore-go/commit/53e9730eb72a5a621734ade0c73d52fc9c742df0))

## 0.17.0 (2025-10-21)

Full Changelog: [v0.16.0...v0.17.0](https://github.com/G-Core/gcore-go/compare/v0.16.0...v0.17.0)

### ⚠ BREAKING CHANGES

* **cloud:** rename to projects update
* **cloud:** use new PATCH files shares endpoint

### Features

* **api:** aggregated API specs update ([558fa44](https://github.com/G-Core/gcore-go/commit/558fa447f33919c7ac06e66cc2d21584db51c9b5))
* **cdn:** add methods to list aws and alibaba regions ([4ee3849](https://github.com/G-Core/gcore-go/commit/4ee3849d2ee55a8b620bea84b2bd2fcd235eb856))
* **client:** add client opt for cloud polling timeout ([95edc63](https://github.com/G-Core/gcore-go/commit/95edc639b26fa1f8786aff68514f6ba9bff779a3))
* **cloud:** add DeleteAndPoll to placement groups ([a287293](https://github.com/G-Core/gcore-go/commit/a287293020adad577b035a23661d3144fa78d7cc))
* **cloud:** add DeleteAndPoll to projects ([16f7b78](https://github.com/G-Core/gcore-go/commit/16f7b78eaa7b44c5d99af98d214949514d53190a))
* **cloud:** add Secret.DeleteAndPoll method ([#146](https://github.com/G-Core/gcore-go/issues/146)) ([4267ffd](https://github.com/G-Core/gcore-go/commit/4267ffd4fa26e734e02c8eec45b78fe77ed3f0ef))
* **cloud:** enable TF for placement groups ([8c81678](https://github.com/G-Core/gcore-go/commit/8c81678e22ea49f8fdfefd94d1440cd6d97c6af8))


### Chores

* **cloud:** rename to projects update ([d17ea13](https://github.com/G-Core/gcore-go/commit/d17ea13bcf4f4511e562d0236d0f9352c583d418))
* **cloud:** use new PATCH files shares endpoint ([1c63726](https://github.com/G-Core/gcore-go/commit/1c637263aff81c5251731fede31901dfa5d3030c))


### Refactors

* **cloud:** improve opts concatenation in poll methods ([dc3204c](https://github.com/G-Core/gcore-go/commit/dc3204cbf684712ddd754e62b3aff87aa401a73e))
* **spec:** remove CDN deprecated endpoints ([2608b61](https://github.com/G-Core/gcore-go/commit/2608b6171c4d782bdc41ff65531c2f4ae1c35146))

## 0.16.0 (2025-10-16)

Full Changelog: [v0.15.0...v0.16.0](https://github.com/G-Core/gcore-go/compare/v0.15.0...v0.16.0)

### ⚠ BREAKING CHANGES

* **cloud:** remove get and update list method for billing reservations
* **cloud:** rename to load_balancer_id path param
* **cloud:** rename inference applications deployments update method

### Features

* **api:** aggregated API specs update ([888df9e](https://github.com/G-Core/gcore-go/commit/888df9ec0d4f282d6a70cd95fe718a9f019015e5))
* **api:** aggregated API specs update ([9b66b72](https://github.com/G-Core/gcore-go/commit/9b66b72da0c1ad78b63d4ca6729105a153b4e347))
* **api:** aggregated API specs update ([554de1e](https://github.com/G-Core/gcore-go/commit/554de1e99f089f0910eddf54f99462fea753346f))
* **api:** aggregated API specs update ([0a0de0b](https://github.com/G-Core/gcore-go/commit/0a0de0befb37e9254afb70b9a5edbca65e6684c8))
* **api:** aggregated API specs update ([e663d7a](https://github.com/G-Core/gcore-go/commit/e663d7ac6bc5c338e3251fb01da45c59702e0e19))
* **api:** aggregated API specs update ([37bd6ab](https://github.com/G-Core/gcore-go/commit/37bd6ab35fba956455bd60d4ec92ac7b6e6fe800))
* **api:** aggregated API specs update ([e6fc584](https://github.com/G-Core/gcore-go/commit/e6fc584ade819d92e17497d8b64cd3f3534b4522))
* **api:** aggregated API specs update ([fb41f9f](https://github.com/G-Core/gcore-go/commit/fb41f9f73d2846e276219d49e3b3923622d1376e))
* **api:** aggregated API specs update ([9197fbf](https://github.com/G-Core/gcore-go/commit/9197fbf615634b2c165e82e3344c3fc828f43441))
* **cloud:** add NewAndPoll and DeleteAndPoll methods to NetworkRouterService ([#136](https://github.com/G-Core/gcore-go/issues/136)) ([53b38be](https://github.com/G-Core/gcore-go/commit/53b38beaed90e595819d28c191f3fd4fb76b563e))
* **cloude:** remove cloud_lbmember name ([00d3139](https://github.com/G-Core/gcore-go/commit/00d313988440ba6782325932fe74f07b65962d47))
* **cloud:** implement AddAndPoll and RemoveAndPoll methods for LoadBalancerPoolMemberService ([#137](https://github.com/G-Core/gcore-go/issues/137)) ([3314f5b](https://github.com/G-Core/gcore-go/commit/3314f5bd889ec0cb8ce5dfccefe8b6c09db2614c))
* **cloud:** remove get and update list method for billing reservations ([58d2841](https://github.com/G-Core/gcore-go/commit/58d2841510f0e147da56e42e0ccc9d5c1e3a6ccd))


### Bug Fixes

* **cloud:** rename to load_balancer_id path param ([ec29fce](https://github.com/G-Core/gcore-go/commit/ec29fce204368ac71d36d4c23a6a751bb3c40729))
* **examples:** make name optional in cloud instance update ([9072242](https://github.com/G-Core/gcore-go/commit/9072242e9f8b84c3f606108a55a2b91ac3187eeb))


### Chores

* add pull request template ([1375c8c](https://github.com/G-Core/gcore-go/commit/1375c8c18b2315a79a79b961262751399d635e26))
* **ci:** add fossa ([252fef9](https://github.com/G-Core/gcore-go/commit/252fef94dc187569fb70fef1e50e3623ffa26be7))
* **cloud:** rename inference applications deployments update method ([ce36a20](https://github.com/G-Core/gcore-go/commit/ce36a20667c7c9bb4b80ec00f5c012b208e664bf))
* **cloud:** rename loadBalancerID parameter ([a516645](https://github.com/G-Core/gcore-go/commit/a51664588e91947b12250c5f214c6ab10f57f39a))

## 0.15.0 (2025-10-02)

Full Changelog: [v0.14.0...v0.15.0](https://github.com/G-Core/gcore-go/compare/v0.14.0...v0.15.0)

### Features

* **api:** Add missing reserved_fixed_ips update method ([67057ab](https://github.com/G-Core/gcore-go/commit/67057abac2c98dbe1ac6a53e7fa3b988b4c80b19))
* **api:** aggregated API specs update ([f635f85](https://github.com/G-Core/gcore-go/commit/f635f8570d532431857d505fd46e2e0d8259720f))
* **api:** aggregated API specs update ([ab636ba](https://github.com/G-Core/gcore-go/commit/ab636bac399f5da60ec0e50163539148fc8cdb01))
* **api:** aggregated API specs update ([974cae5](https://github.com/G-Core/gcore-go/commit/974cae5fa07bd88c3b5cfa98d005dfdb0fc8bbd6))
* **cloud:** add DeleteAndPoll for subnets ([578d222](https://github.com/G-Core/gcore-go/commit/578d2222ebf5299930965c1b7023d93afc1c568c))


### Bug Fixes

* **examples:** update examples to match updated SDK signatures ([#123](https://github.com/G-Core/gcore-go/issues/123)) ([4b536bd](https://github.com/G-Core/gcore-go/commit/4b536bde1bc6677d49023c2ae45bfc57d5cf59a4))

## 0.14.0 (2025-09-30)

Full Changelog: [v0.13.0...v0.14.0](https://github.com/G-Core/gcore-go/compare/v0.13.0...v0.14.0)

### Features

* **api:** aggregated API specs update ([9563541](https://github.com/G-Core/gcore-go/commit/95635416cf5835eca96e2741a54ccf008d41d39b))
* **api:** aggregated API specs update ([ff36e57](https://github.com/G-Core/gcore-go/commit/ff36e573b89d11a5a4fa3291dc1ea9eff1bd4cd6))
* **api:** aggregated API specs update ([27a0d29](https://github.com/G-Core/gcore-go/commit/27a0d29637f4413e1bff44ed971ac92cbc6e2521))
* **api:** aggregated API specs update ([5722309](https://github.com/G-Core/gcore-go/commit/5722309174710cf4567fa7f76c7fe456208fd517))
* **api:** aggregated API specs update ([225f9e2](https://github.com/G-Core/gcore-go/commit/225f9e2616eb940261fa5e979babbfd97c6d00f3))
* **cdn:** add API support ([e33a359](https://github.com/G-Core/gcore-go/commit/e33a3593b7bf7d54a8d72b5e079d7f3371766415))
* **cloud:** enable TF for floating IPs ([77c64f0](https://github.com/G-Core/gcore-go/commit/77c64f00c492ffd84a8d8168783dadae20bde17a))
* **storage:** add examples ([855c53b](https://github.com/G-Core/gcore-go/commit/855c53b7b408f0004877581b7628019e9182288c))


### Bug Fixes

* bugfix for setting JSON keys with special characters ([2de8afe](https://github.com/G-Core/gcore-go/commit/2de8afe98cab189fa4426d56ecb4452c6bc16d09))
* **client:** correctly generate K8sClusterSlurmAddonV2Serializers ([562e94d](https://github.com/G-Core/gcore-go/commit/562e94dc6b700dbc82e4286ed7b2884940b644e7))
* use slices.Concat instead of sometimes modifying r.Options ([0ba8706](https://github.com/G-Core/gcore-go/commit/0ba87060842833644f5574037ac4018b53de9ddf))


### Chores

* bump minimum go version to 1.22 ([e233cc3](https://github.com/G-Core/gcore-go/commit/e233cc3089abcb83b1e0b58d4abfe16dea0e6dfa))
* do not install brew dependencies in ./scripts/bootstrap by default ([0bdb83b](https://github.com/G-Core/gcore-go/commit/0bdb83b6211f83f7b5b93d6955391f252342a1ce))
* improve example values ([7b94270](https://github.com/G-Core/gcore-go/commit/7b94270ac1443b6ae7d38953993b93bbf965ad68))
* update more docs for 1.22 ([983e697](https://github.com/G-Core/gcore-go/commit/983e697c23318ff9a1fbfc69d9ac156c32675ca1))

## 0.13.0 (2025-09-16)

Full Changelog: [v0.12.0...v0.13.0](https://github.com/G-Core/gcore-go/compare/v0.12.0...v0.13.0)

### ⚠ BREAKING CHANGES

* **waap:** model references

### Features

* **api:** aggregated API specs update ([3a6f913](https://github.com/G-Core/gcore-go/commit/3a6f9130b631a95cc358d4dd28db56219eb59e7f))
* **api:** aggregated API specs update ([1ef659c](https://github.com/G-Core/gcore-go/commit/1ef659c9187ddc2c0567952cbb67bf439681c073))
* **api:** aggregated API specs update ([2fd148e](https://github.com/G-Core/gcore-go/commit/2fd148ee9349d222f6fd02701289109d656295bb))
* **api:** aggregated API specs update ([b5b74dd](https://github.com/G-Core/gcore-go/commit/b5b74ddf140edd99ada4922bedb113ef40dfb28e))
* **cloud:** support floating IPs update ([f6665c0](https://github.com/G-Core/gcore-go/commit/f6665c0b0ecc1c0a5c6a85cca4326f5e937e525d))
* **dns:** replace post with get in check delegation status ([d45f9a5](https://github.com/G-Core/gcore-go/commit/d45f9a5261df3e8e399f5e48cc47e13bec9b7804))


### Bug Fixes

* **waap:** model references ([ec1e84e](https://github.com/G-Core/gcore-go/commit/ec1e84e1c862511899b7eb0b0a86da5af334d0a7))

## 0.12.0 (2025-09-11)

Full Changelog: [v0.11.0...v0.12.0](https://github.com/G-Core/gcore-go/compare/v0.11.0...v0.12.0)

### Features

* **api:** aggregated API specs update ([b7792f7](https://github.com/G-Core/gcore-go/commit/b7792f7fd405f1ef222724ea181a49cb74606ff7))
* **cloud:** add DeleteAndPoll() for reserved fixed ip ([84c4297](https://github.com/G-Core/gcore-go/commit/84c42971e672e257dc5eda7468ac12afc1718cce))
* **cloud:** add polling methods to volumes ([1d9acc3](https://github.com/G-Core/gcore-go/commit/1d9acc307040f44bf63b75df3c9a24d727693646))


### Refactors

* **storage:** use v2 endpoint ([9637ab9](https://github.com/G-Core/gcore-go/commit/9637ab9a494966011a0bb1b8749e5d339907a312))

## 0.11.0 (2025-09-09)

Full Changelog: [v0.10.0...v0.11.0](https://github.com/G-Core/gcore-go/compare/v0.10.0...v0.11.0)

### ⚠ BREAKING CHANGES

* **cloud:** migrate baremetal gpu cluster from v1 to v3
* **cloud:** support inference applications

### Features

* **api:** aggregated API specs update ([99c1e45](https://github.com/G-Core/gcore-go/commit/99c1e459efb53a699f637ac035104f2ff5843d28))
* **api:** aggregated API specs update ([3377fd7](https://github.com/G-Core/gcore-go/commit/3377fd76fa6207e368f193428b022c95dd3a1d58))
* **api:** aggregated API specs update ([b66b088](https://github.com/G-Core/gcore-go/commit/b66b0887cb34a4139bff60a78c8dfe234fe94eda))
* **api:** aggregated API specs update ([6994f5d](https://github.com/G-Core/gcore-go/commit/6994f5d1ed046b38c533c2a07f0fd860510291dd))
* **api:** aggregated API specs update ([12fc447](https://github.com/G-Core/gcore-go/commit/12fc447a91e20325bfd98db446e3ab4b2cbe6930))
* **api:** aggregated API specs update ([66442b4](https://github.com/G-Core/gcore-go/commit/66442b4582d4e1307502faa11f7d6537d7f8cd41))
* **api:** aggregated API specs update ([b437753](https://github.com/G-Core/gcore-go/commit/b43775377b9afe0d7fdf511c81b885ad5c43e913))
* **api:** aggregated API specs update ([e3e2989](https://github.com/G-Core/gcore-go/commit/e3e29891210e597e91d1affe31800e8e19c1ae18))
* **api:** aggregated API specs update ([8f064ef](https://github.com/G-Core/gcore-go/commit/8f064ef675929a32520d62b60392752e7961bf1c))
* **api:** aggregated API specs update ([277b318](https://github.com/G-Core/gcore-go/commit/277b318dc2012500722626754ab31b2d37ca8645))
* **api:** aggregated API specs update ([3c33290](https://github.com/G-Core/gcore-go/commit/3c332904c5249d4af49cb7b7f8c4fc9653a541a6))
* **api:** aggregated API specs update ([6e56534](https://github.com/G-Core/gcore-go/commit/6e56534fed327c6aae6b87263494c92cf14e90a2))
* **api:** aggregated API specs update ([2770f2c](https://github.com/G-Core/gcore-go/commit/2770f2c2dd5546d2217810b21b3421ed7761ad75))
* **api:** aggregated API specs update ([d87d6b1](https://github.com/G-Core/gcore-go/commit/d87d6b1866719abbb9d544906ebba515bdd3d491))
* **api:** aggregated API specs update ([6c44f8a](https://github.com/G-Core/gcore-go/commit/6c44f8a6cdc738320f3d7d516f087f758df95f3f))
* **api:** aggregated API specs update ([8643fe8](https://github.com/G-Core/gcore-go/commit/8643fe8df688758015a04f97b810ab2a8b9be886))
* **api:** aggregated API specs update ([7043dc0](https://github.com/G-Core/gcore-go/commit/7043dc00f239f352b19f5d19d37502fde6bd18f7))
* **api:** api update ([d3d64b3](https://github.com/G-Core/gcore-go/commit/d3d64b3d3215551bb123ad77010776151b8766bc))
* **api:** manual updates ([8a55e76](https://github.com/G-Core/gcore-go/commit/8a55e766e0075c12abde3b506b1f2b9fb6f96206))
* **api:** manual upload of aggregated API specs ([fc2cc17](https://github.com/G-Core/gcore-go/commit/fc2cc17f98a62222de4fd1dac2115103cc392609))
* **api:** manual upload of aggregated API specs ([b0adde7](https://github.com/G-Core/gcore-go/commit/b0adde7eca2ed186c6121eb3ed2e72070c78c07b))
* **api:** update field_value type ([591847e](https://github.com/G-Core/gcore-go/commit/591847e4582d6402bd31306858c91686d722ef46))
* **cloud:** add managed k8s ([8a25bd4](https://github.com/G-Core/gcore-go/commit/8a25bd49ea615e2300746786cc4e676ab6e645b1))
* **cloud:** add NewAndPoll() and DeleteAndPoll() for floating ips ([2798254](https://github.com/G-Core/gcore-go/commit/27982547cc49f1d7e2b8fd8aab35625763d52665))
* **cloud:** add NewAndPoll() and DeleteAndPoll() for networks ([b787f73](https://github.com/G-Core/gcore-go/commit/b787f7374b9c7e28e313733fe3caddf641e99800))
* **cloud:** add NewAndPoll() for subnets ([3a64336](https://github.com/G-Core/gcore-go/commit/3a6433661e5173bb5f2357ac67908e6b0845b246))
* **cloud:** fetch client_id from iam in cloud quotas examples ([0aab78e](https://github.com/G-Core/gcore-go/commit/0aab78e76fb9833c7e43e3c5a7ed33105aa29045))
* **cloud:** migrate baremetal gpu cluster from v1 to v3 ([b9f3997](https://github.com/G-Core/gcore-go/commit/b9f3997fcd8d049c713cba316fcb4cc1ddc03bed))
* **cloud:** remove inference model examples ([58221cb](https://github.com/G-Core/gcore-go/commit/58221cbb31ae86ec49273eb3a0b9ae348eae57d4))
* **cloud:** support inference applications ([31c00d6](https://github.com/G-Core/gcore-go/commit/31c00d643e7fcebb084b480d5dd88c57a9156b3b))
* **cloud:** use PATCH /v2/lbpools ([4ed5b55](https://github.com/G-Core/gcore-go/commit/4ed5b552f58708bb38c308970656ba095085975f))
* **s3:** add object storage ([b258b9a](https://github.com/G-Core/gcore-go/commit/b258b9afc8c6ef37b3e79c36ef8610cd2dca124a))
* **storage:** make list storage locations paginated ([53f76bf](https://github.com/G-Core/gcore-go/commit/53f76bf96287ba5a153aafcc2ced681b5f0dd23f))


### Bug Fixes

* close body before retrying ([4aa95c1](https://github.com/G-Core/gcore-go/commit/4aa95c145147ef4e3c3d1373b5310be103755b52))
* **dns:** fix dns methods ([24c71bd](https://github.com/G-Core/gcore-go/commit/24c71bd87377d6f5829654f52849ac1048037fb3))
* **internal:** unmarshal correctly when there are multiple discriminators ([db7b996](https://github.com/G-Core/gcore-go/commit/db7b9967959dad6ad23c8755af7f3fec398881f6))
* remove null from release please manifest ([9de166b](https://github.com/G-Core/gcore-go/commit/9de166b82fc37cf7e239047b7d3ba2704f6c0941))
* use release please annotations on more places ([1549cda](https://github.com/G-Core/gcore-go/commit/1549cda8e66331bfbfbfdaf9c5bd8faf4feffc8a))
* **waap:** fix component name ([41f5fca](https://github.com/G-Core/gcore-go/commit/41f5fcaf2093b92d8dd26ae836c84d80c8c02a88))


### Chores

* **internal:** codegen related update ([ed63ff2](https://github.com/G-Core/gcore-go/commit/ed63ff2ed044bbdeeaa0578694b88b2bf31d730e))
* **internal:** detect breaking changes when removing endpoints ([1edc306](https://github.com/G-Core/gcore-go/commit/1edc306a797456fb7876952b707389fc64dea91b))
* **internal:** update comment in script ([19e7f4a](https://github.com/G-Core/gcore-go/commit/19e7f4ae96857b9c7cd08f82386fa548533ad6f6))
* update @stainless-api/prism-cli to v5.15.0 ([ec528d1](https://github.com/G-Core/gcore-go/commit/ec528d179d42becbf3f0f576a7f0d6f0e265fc6a))

## 0.10.0 (2025-08-07)

Full Changelog: [v0.9.0...v0.10.0](https://github.com/G-Core/gcore-go/compare/v0.9.0...v0.10.0)

### ⚠ BREAKING CHANGES

* **security:** rename bgp_announces change() to toggle()
* **waap:** refactor WAAP models

### Features

* add example snippet to invite user and assign cloud role ([9edcaf7](https://github.com/G-Core/gcore-go/commit/9edcaf786326f751b73a63dee517a0b9e3d467c8))
* **api:** aggregated API specs update ([1c847a5](https://github.com/G-Core/gcore-go/commit/1c847a5c4ea9b9b2748e7e466e0a2f5db96dc6ad))
* **api:** aggregated API specs update ([659c741](https://github.com/G-Core/gcore-go/commit/659c741a78d9c94714e87142edbacedc814280cb))
* **client:** support optional json html escaping ([7917576](https://github.com/G-Core/gcore-go/commit/7917576809148638a6d82eaae5c7ecf1f0d087a5))


### Bug Fixes

* **security:** rename bgp_announces change() to toggle() ([561648c](https://github.com/G-Core/gcore-go/commit/561648c57050f4136d19dee618e13f5c8fa2f14c))


### Refactors

* **waap:** refactor WAAP models ([b07cb4f](https://github.com/G-Core/gcore-go/commit/b07cb4f0daa4e01306fa09d0df34ef97a8bca798))

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
