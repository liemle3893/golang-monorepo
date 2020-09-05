|GitHub Actions|
|:-----|
|![Main](https://github.com/liemle3893/golang-monorepo/workflows/Main/badge.svg)|

## Overview

This is an example of a golang-based monorepo. It has the following features:

- Only build the services or cmds that are modified in a commit;
- Build all services and/or cmds that are affected by changes in common codes (i.e. `pkg`);
- Build all services and/or cmds that are affected by changes in `vendor` codes.

For now, [CircleCI 2.1](./.circleci/config.yml), [GitHub Actions](https://github.com/liemle3893/golang-monorepo/actions), [Gitlab](https://gitlab.com) are supported. But since it uses bash scripts and Makefiles, it should be fairly straightforward to port to [TravisCI](https://travis-ci.org/) or [AppVeyor](https://www.appveyor.com/), etc.

At the moment, CI is setup to use Go 1.14[.x] with `GO111MODULE=on` and `GOFLAGS=-mod=vendor` environment variables enabled during build. See sample [dockerfile](./services/samplesvc/dockerfile.samplesvc) for more details.

## How does it work

During CI builds, [build.sh](./build.sh) iterates the updated files within the commit range (`CIRCLE_COMPARE_URL` environment variable in CircleCI) or the modified files within a single commit (when the value is not a valid range), excluding hidden files, `pkg`, and `vendor` folders. It will then try to walk up the directory path until it can find a Makefile (excluding root Makefile). Once found, the [root Makefile](./Makefile) will include that Makefile and call the `custom` rule as target, thus, initiating the build.

When the changes belong to either `pkg` or `vendor`, the script will then try to determine the services (and cmds) that have dependencies using the `go list` command. All dependent services will then be built using the same process described above.

You can override the `COMMIT_RANGE` environment variable for your own CI. If this is set, `build.sh` will use its value. You also want to set `CIRCLE_SHA1` to your commit SHA (`CIRCLE_SHA1` is CircleCI-specific). Example for GitHub Actions is [here](https://liemlhd.com/golang/monorepo-example/blob/master/.github/workflows/main.yml). Something like:
```bash
# If your commit range is correct:
COMMIT_RANGE: aaaaa..bbbbb
CIRCLE_SHA1: aaaaa

# If no valid commit range:
COMMIT_RANGE: <your_commit_sha>
CIRCLE_SHA1: <your_commit_sha>
```

## Directory structure

- `services/` - Basically, long running services.
- `cmd/` - CLI-based tools that are not long running.
- `pkg/` - Shared codes, or libraries common across the repo.
- `vendor/` - Third party codes from different vendors.

Although we have this structure, there is no limitation into where should you put your services/cmds. Any subdirectory structure is fine as long as a Makefile is provided.

## How to add a service/cmd

```shell script
$ ./new-service <service-name>
```

## Need help
PR's are welcome!
- [x] Support for GitHub Actions
- [x] Support for GitLab
- [ ] Run `go test ...` for `pkg/` when Makefile is root
- [ ] Make it work without the `vendor` folder as well

## Misc
https://tech.mobingi.com/2018/09/25/ouchan-monorepo.html
https://github.com/flowerinthenight/golang-monorepo
