# Making a new release

* Bump the version in version.go
* Commit the version bump and push it to git
* Run release.sh. This will:
  * Tag the current version
  * Push the tag to git
  * Run goreleaser
