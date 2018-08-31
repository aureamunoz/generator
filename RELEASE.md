# Release the project and generate go release

## The easy way

Simply create a release using the script `./script/tag_release_manually.sh` where yoou pass as parameter your `GITHUB_API_TOKEN` and
`ID` of the release to be created :

```bash
./script/tag_release_manually.sh GITHUB_API_TOKEN VERSION
```

where `VERSION` corresponds to a string starting with the prefix `release-` and next by the number of the release (e.g release-0.0.1, release-0.0.2, .... release-0.0.n)

Example

```bash
./script/tag_release_manually.sh  aaaabbbbcccccdddddeeeeeffff release-0.0.1
```

This will cause `CircleCI` to perform a build/release and will create a GitHub release named `0.0.x`
that will contain the built binaries for both MacOS and Linux.

**Note**: This assumes that the `GITHUB_API_TOKEN` has been set in the CircleCI UI for the job !

When performing a CircleCI release, a new docker image will also be created under this name `snowdrop/spring-boot-generator` 
which is published on `quay.io`.

**Note**: This assumes that the `QUAY_ROBOT_USER` and `QUAY_ROBOT_TOKEN` has been set in the CircleCI UI for the job

When the process is terminated, you can delete the temporary release and tag created under GitHub using the following bash script

```bash
./script/delete_release_manually.sh GITHUB_API_TOKEN VERSION
```

Example

```bash
./script/delete_release_manually.sh aaaabbbbcccccdddddeeeeeffff release-0.0.1
```

   