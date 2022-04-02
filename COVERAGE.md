# Code Coverage

While we recommend using one of the free websites available for monitoring code coverage
during your continuous integration process, below is an example of how you can incorporate
code coverage during the continuous integration process provided by github actions and 
generate a code coverage report without one of those services.

This yaml file will run tests on multiple system configurations, but will produce
a code coverage report on only one of those. It will then create a code coverage badge
and add it to the README file.

This file should be put in the `.github/workflows` directory of your repo. 

```yaml
name: Go  # The name of the workflow that will appear on Github

on:
  push:
    branches: [ main ]
  pull_request:
    branches: [ main ]
  # Allows you to run this workflow manually from the Actions tab
  workflow_dispatch:

jobs:

  build:
    runs-on: ${{ matrix.os }}
    strategy:
      matrix:
        os: [ubuntu-latest, windows-latest]
        go: [1.16, 1.17]
    steps:
    - uses: actions/checkout@v2

    - name: Set up Go
      uses: actions/setup-go@v2
      with:
        go-version: ${{ matrix.go }}

    - name: Build
      run: go install

    - name: Test
      run: |
        go test -v -cover ./... -coverprofile coverage.out -coverpkg ./...
        go tool cover -func coverage.out -o coverage.out  # Replaces coverage.out with the analysis of coverage.out

    - name: Go Coverage Badge
      uses: tj-actions/coverage-badge-go@v1
      if: ${{ runner.os == 'Linux' && matrix.go == '1.17' }} # Runs this on only one of the ci builds.
      with:
        green: 80
        filename: coverage.out

    - uses: stefanzweifel/git-auto-commit-action@v4
      id: auto-commit-action
      with:
        commit_message: Apply Code Coverage Badge
        skip_fetch: true
        skip_checkout: true
        file_pattern: ./README.md

    - name: Push Changes
      if: steps.auto-commit-action.outputs.changes_detected == 'true'
      uses: ad-m/github-push-action@master
      with:
        github_token: ${{ github.token }}
        branch: ${{ github.ref }}

```
