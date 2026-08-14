#!/bin/bash
set -eu -o pipefail

declare -r expected_sha="${1:?missing expected_sha}"
declare -r gh_pr_number="${2:?missing gh_pr_number}"
declare -r gh_repository="${3:?missing gh_repository}"
declare -r gh_current_run_id="${4:?missing gh_current_run_id}"

# 1. Fetch current PR HEAD SHA to ensure no new commit was pushed while this ran
CURRENT_SHA=$(gh pr view "${gh_pr_number}" \
	--repo "${gh_repository}" \
	--json headRefOid --jq '.headRefOid')

if [ "${CURRENT_SHA}" != "${expected_sha}" ]; then
	echo "A new commit (${CURRENT_SHA}) was pushed after this run started (${expected_sha}). Aborting auto-merge."
	exit 1
fi

# 2. Allowlist filter: select any check whose state is NOT in [SUCCESS, SKIPPED, NEUTRAL],
#   filtering out the current workflow run by checking if its link contains gh_current_run_id.
NON_PASSING_CHECKS=$(gh pr checks "${gh_pr_number}" \
	--repo "${gh_repository}" \
	--json workflow,name,state,link |
	jq -c --arg current_run_id "${gh_current_run_id}" \
		'.[] | select(
       (.link | contains($current_run_id) | not)
       and
       (.state | test("^(SUCCESS|SKIPPED|NEUTRAL)$") | not)
     )')

if [ -n "${NON_PASSING_CHECKS}" ]; then
	echo "Cannot auto-merge. The following checks are not passing:"
	echo "${NON_PASSING_CHECKS}" | jq -r '"workflow: \(.workflow), name: \(.name), state: \(.state)"'
	exit 1
fi
