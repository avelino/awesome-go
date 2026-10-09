#!/bin/bash

set -eu -o pipefail

declare -r EXPECTED_SHA="${1:?missing EXPECTED_SHA}"
declare -r GH_PR_NUMBER="${2:?missing GH_PR_NUMBER}"
declare -r GH_REPOSITORY="${3:?missing GH_REPOSITORY}"
declare -r GH_CURRENT_RUN_ID="${4:?missing GH_CURRENT_RUN_ID}"

# 1. Fetch current PR HEAD SHA to ensure no new commit was pushed while this ran
current_sha=$(gh pr view "${GH_PR_NUMBER}" \
	--repo "${GH_REPOSITORY}" \
	--json headRefOid --jq '.headRefOid')

if [ "${current_sha}" != "${EXPECTED_SHA}" ]; then
	echo "A new commit (${current_sha}) was pushed after this run started (${EXPECTED_SHA}). Aborting auto-merge."
	exit 1
fi

# 2. Allowlist filter: select any check whose state is NOT in [SUCCESS, SKIPPED, NEUTRAL],
#   filtering out the current workflow run by checking if its link contains GH_CURRENT_RUN_ID.
pr_checks=$(gh pr checks "${GH_PR_NUMBER}" \
	--repo "${GH_REPOSITORY}" \
	--json workflow,name,state,link
)
non_passing_checks=$(echo "${pr_checks}" |
	jq -c --arg current_run_id "${GH_CURRENT_RUN_ID}" \
		'.[] | select(
       (.link | contains($current_run_id) | not)
       and
       (.state | test("^(SUCCESS|SKIPPED|NEUTRAL)$") | not)
     )')

# helpful for debugging
printf >&2 "# GH_CURRENT_RUN_ID: %s
# pr_checks: %s
# non_passing_checks: %s" "${GH_CURRENT_RUN_ID}" "${pr_checks}" "${non_passing_checks}"

if [ -n "${non_passing_checks}" ]; then
	echo "Cannot auto-merge. The following checks are not passing:"
	echo "${non_passing_checks}" | jq -c '{workflow,name,state}'
	exit 1
fi
