This resource was made by the Go community and wouldn't be possible without you! 
We appreciate and recognize [all contributors](https://github.com/avelino/awesome-go/graphs/contributors).

# Contribution Guidelines

> Please be aware that we want to accept your contribution, but we have **some rules to keep the minimum quality** of the packages listed here. All reviews are **not personal feedback**, even if you are a _developer reviewing your contribution_. **Sorry, if we can't meet your expectations; we do our best**.

- **To add, remove, or change things on the list:** Submit a pull request

To set this list apart from and complement the excellent [Go wiki Projects page](https://golang.org/wiki/Projects), 
and other lists, awesome-go is a specially curated list of high-quality, actively maintained Go packages and resources.

Please contribute links to packages/projects you have used or are familiar with. This will help ensure high-quality entries.

## Quality standards

To be on the list, project repositories should adhere to the following quality standards. 
(https://goreportcard.com/report/github.com/ **github_user** / **github_repo**):

- have an **open source license**, [see list of allowed licenses](https://opensource.org/licenses/alphabetical);
- function as documented and expected;
- be generally useful to the wider community of Go programmers;
- be actively maintained with:
  - regular, recent commits;
  - or, for finished projects, issues and pull requests are responded to generally within 2 weeks;
- be stable or progressing toward stable;
- be thoroughly documented (README, pkg.go.dev doc comments, etc.) in the English language, so everyone is able to understand the project's intention and how it works. All public functions and types should have a Go-style documentation header;
- if the library/program is testable, then coverage should be >= 80% for non-data-related packages and >=90% for data-related packages. (**Note**: the tests will be reviewed too. We will check your coverage manually if your package's coverage is just a benchmark result);
- have at least one official version-numbered release that allows go.mod files to list the file by version number of the form vX.X.X.

Categories must have at least 3 items.

## Preparing for review

Projects listed must have the following in their documentation. When submitting, you will be asked
to provide them.

- A link to the project's pkg.go.dev page
- A link to the project's Go Report Card report
- A link to a code coverage report

One way to accomplish the above is to add badges to your project's README file.
- Use https://pkg.go.dev/badge/ to create the pkg.go.dev link.
- Go to https://goreportcard.com/ to generate a Go Report Card report, then click on the report badge in the upper-right corner to see details on how to add the badge to your README.
- Codecov, coveralls, and gocover all offer ways to create badges for code coverage reports. Another option is to generate a badge as part of a continuous integration process. See [Code Coverage](COVERAGE.md) for an example.

## How to add an item to the list

Open a pull request against the README.md document that adds the repository to the list.

- The pull request should add one and only one item to the list.
- The added item should be in alphabetical order within its category.
- The link should be the name of the package or project.
- Descriptions should be clear, concise, and non-promotional.
- Descriptions should follow the link on the same line and end with a punctuation mark.
- Remember to put a period `.` at the end of the project description.

If you are creating a new category, move the projects that apply to the new category, ensuring
that the resulting list has at least 3 projects in every category, and that the categories are alphabetized.

Fill out the template in your PR with the links asked for. If you accidentally remove the PR template from the submission, you can find it [here](https://github.com/avelino/awesome-go/blob/main/.github/PULL_REQUEST_TEMPLATE.md).


## Congrats, your project got accepted - what now?
You are an outstanding project now! Feel encouraged to tell others about it by adding one of these badges:  
[![Mentioned in Awesome Go](https://awesome.re/mentioned-badge.svg)](https://github.com/avelino/awesome-go)  
[![Mentioned in Awesome Go](https://awesome.re/mentioned-badge-flat.svg)](https://github.com/avelino/awesome-go)

```md
[![Mentioned in Awesome Go](https://awesome.re/mentioned-badge.svg)](https://github.com/avelino/awesome-go)  
[![Mentioned in Awesome Go](https://awesome.re/mentioned-badge-flat.svg)](https://github.com/avelino/awesome-go)
```
## Maintenance expectations for projects listed here

To prevent removal from awesome-go, your project must maintain the following quality standards.
- Development should be ongoing and maintain code quality. Official releases should be at least once a year if the project is ongoing.
- Or, if development has halted because the project is mature and stable, that can be demonstrated by having no bug reports in the Issues list that are older than 6 months.
- All links to quality reports should be to the most recent official release or current ongoing development.

Highly recommended but not required:
- A continuous integration process to be part of the ongoing development process
- That the project uses a pull-request process, and the owners do not commit directly to the repository
- That the pull-request process requires the continuous-integration tests to pass before a pull request can be merged

## How to remove an item from the list

- Open a pull request that deletes the line of the project in question.
- Delete the submission template and substitute a description of which criteria the project is not meeting. It should be a combination of the following.
  - The project has not made an official release within the last year and has open issues.
  - The project is not responding to bug reports issued within 6 months of submission.
  - The project is not meeting quality standards as indicated by the Go Report Card or Code Coverage tests.
  - The quality standard links have been removed from the documentation.
  - The project is no longer open-sourced.
  - The project is incompatible with any Go version issued within the last year (there is hopefully an open PR about this at the project).

If the project is hosted on GitHub, include a link to the project's submitter and/or author so
that they will be notified of the desire to remove the project and have an opportunity to respond.
The link should be of the form @githubID.

If the project is not hosted on GitHub, open an issue at the project in question's repository linking to the PR 
and stating the following:

>This project is currently listed at awesome-go at https://github.com/avelino/awesome-go. 
However, it appears that the project is not maintaining the quality standards required to continue to be listed at the awesome-go project.
This project is scheduled to be removed within 2 weeks of this posting. To continue to be listed at awesome-go, please respond at:
  -- link to above PR --
  
Then, comment on your PR at awesome-go with a link to the removal issue at the project.

## Maintainers

To make sure every PR is checked, we have [team maintainers](MAINTAINERS). Every PR MUST be reviewed by at least one maintainer before it can get merged.

The maintainers will review your PR and notify you and tag it in case any
information is still missing. They will wait 15 days for your interaction, after
that the PR will be closed.


## Reporting issues

Please open an issue if you would like to discuss anything that could be improved or have suggestions for making the list a more valuable resource. We realize sometimes packages fall into abandonment or have breaking builds for extended periods of time, so if you see that, feel free to change its listing, or please let us know. We also realize that sometimes projects are just going through transitions or are more experimental in nature. These can still be cool, but we can indicate them as transitory or experimental.

Removal changes will not be applied until they have been pending for a minimum of 1 week (7 days). This grace window benefits projects that may be going through a temporary transition, but are otherwise worthy of being on the list.

Thanks, everyone!

## How decisions are made

The official group of maintainers has the final decision on what PRs are accepted. Discussions are made openly in issues. Decisions are made by consensus.
