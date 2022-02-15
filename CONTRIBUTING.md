This resource was made by the Go community and wouldn't be possible without you! 
We appreciate and recognize [all contributors](https://github.com/avelino/awesome-go/graphs/contributors).

# Contribution Guidelines

> Please be aware that we want to accept your contribution, but we have **some rules to keep the minimum quality** 
of the packages listed here. All reviews are **not personal feedback**, 
even if you are a _developer reviewing your contribution_. **Sorry if we can't meet your expectations, we do our best**.

- **To add, remove, or change things on the list:** Submit a pull request

To set this list apart from and complement the excellent [Go wiki Projects page](https://golang.org/wiki/Projects), 
and other lists, awesome-go is a specially curated list for high-quality, actively maintained Go packages and resources.

- At least 3 items are needed to create a new category;
- The package or project had to be maintained under **open source license** ( *we make a brief review of the code before the link enters the list* ), [see list of allowed licenses](https://opensource.org/licenses/alphabetical);

Please contribute links to packages/projects you have used or are familiar with. This will help ensure high-quality entries.

If you removed our PR template you can find it [here](https://github.com/avelino/awesome-go/blob/main/.github/PULL_REQUEST_TEMPLATE.md).


## General quality standards

To be on the list, project repositories should adhere to these quality standards 
(https://goreportcard.com/report/github.com/ **github_user** / **github_repo**):

- Code functions as documented and expected;
- Generally useful to the wider community of Go programmer
- Actively maintained
  - Regular, recent commits
  - Or, for finished projects, issues and pull requests are responded to generally within 2 weeks
- Stable or progressing toward stable
- Thoroughly documented (README, pkg.go.dev doc comments, etc.) in english language, so everyone is able to understand the project's intention and how it works
- Tests, where practical. If the library/program is testable, then coverage should be >= 80% for non-data-related packages and >=90% for data related packages. **Notice**: the tests will be reviewed too. We will check your coverage manually if your package's coverage is just a benchmark results.
- The project has a Go Report Card value of A.
- The project must have at least one official version-numbered release that allows go.mod files
  to list the file by version number. It should be of the form vX.X.X.


## Preparing for review

To make it easy to review, we ask that you do the following to the README file in your proect:
- Add a badge that points to the matching pkg.go.dev link. You can use https://pkg.go.dev/badge/ to create the badge;
- Add a badge that points to the matching Go Report Card report. To generate the report, go to https://goreportcard.com/. Click on the report badge in the upper right corner to see details on how to add the badge to your Readme. The report card must correspond to the latest official release of the project or the current development tip;
- Add a badge that points to a matching coverage report. There are many free ones you can choose from, including codecov, coveralls, and gocover. Also acceptable is generating your own badge as part of a continuous integration process. The report card must correspond to the latest official release of the project or the current development tip;

You will be asked to provide these links when submitting.

## How to add an item to the list

Open a pull request against the README.md document that adds the repository to the list.

- The pull request should add one and only one item to the list;
- The added item should be in alphabetical order within its category;
- The link should be the name of the package or project;
- Descriptions should be clear, concise, and non-promotional;
- Descriptions should follow the link, on the same line and end with a punctuation mark;
- Remember to put a period `.` at end of the project description;
- Links to projects that are version 2 or higher should end in /vX/, where X is the major version number;

Then fill out the template in your PR with the links asked for in your README file.

## Congrats, your project got accepted - what now?
You are an awesome project now! Feel encouraged to tell others about it by adding one of these badges:  
[![Mentioned in Awesome Go](https://awesome.re/mentioned-badge.svg)](https://github.com/avelino/awesome-go)  
[![Mentioned in Awesome Go](https://awesome.re/mentioned-badge-flat.svg)](https://github.com/avelino/awesome-go)

```md
[![Mentioned in Awesome Go](https://awesome.re/mentioned-badge.svg)](https://github.com/avelino/awesome-go)  
[![Mentioned in Awesome Go](https://awesome.re/mentioned-badge-flat.svg)](https://github.com/avelino/awesome-go)
```
## Maintenance expectations for projects listed here

To prevent removal from go-awesome, your project must maintain the following quality standards:
- On-going development while maintaining code quality. Official release should be at least once a year if the project is ongoing. OR, 
- If development has halted because the project is mature and stable, that can be demonstrated by having no bug reports in the Issues list that are older than 6 months.
- All badges should link to reports showing the most recent official release or current ongoing development. Badges should continue to show the same quality standards as the acceptance criteria above.

Highly recommended but not required:
- A continuous integration process be part of the ongoing development process
- That the project uses a pull-request process and the owners do not commit directly to the repository
- That the pull-request process requires the continuous-integration tests pass before a pull request can be merged

## How to remove an item from the list

- Open a pull request that deletes the line of the project in question.
- Delete the submission template and substitute a description of which criteria the project is not meeting. It should be a combination of:
  - The project has not made an official release within the last year and it has open issues.
  - The project is not responding to bug reports issued within 6 months of submission.
  - The project is not meeting quality standards as indicated by the Go Report Card or Code Coverage tests.
  - The quality standards badges have been removed from the README file.
  - The project is no longer open-sourced.
  - The project is not compatible with any Go version issued within the last year (there is hopefully an open PR about this at the project)

After submitting the PR, create an issue at the project in question linking to the PR and stating the following:

>It appears that this project is not maintaining the quality standards required to continue to be listed at the Go-Awesome project.
This project is schedule to be removed within 2 weeks of this posting. To continue to be listed at Go-Awesome, please respond at:
  -- link to above PR --
  
Then, comment on your PR at Go-Awesome with a link to the removal issue at the project.

## Maintainers

To make sure every PR is checked, we have [team maintainers](MAINTAINERS). Every PR MUST be reviewed by at least one maintainer before it can get merged.

The maintainers will review your PR and notify you and tag it in case any
information is still missing. They will wait 15 days for your interaction, after
that the PR will be closed.


## Reporting issues

Please open an issue if you would like to discuss anything that could be improved or have suggestions for making the list a more valuable resource. We realize sometimes packages fall into abandonment or have breaking builds for extended periods of time, so if you see that, feel free to change its listing or let us know. We also realize that sometimes projects are just going through transitions or are more experimental in nature. These can still be cool, but we can indicate them as transitory or experimental.

Removal changes will not be applied until they have been pending for a minimum of 1 week (7 days). This grace window benefits projects that may be going through a temporary transition but are otherwise worthy of being on the list.

Thanks everyone!

## How decision are made

The official group of maintainers have the final decision on what PRs are accepted. Discussions are made openly in issues. Where there appear to be conflicts, decisions are made by a majority vote where 2/3s of the maintainers are considered a quorum.

Official maintainers must commit to responding to at least one PR per month.

Maintainers are added to or removed by the maintainer group by majority vote where 2/3s of the maintainers are considered a quorum.