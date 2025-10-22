function Generate-PRTemplate {
    param(
        [Parameter(Mandatory=$true)]
        [ValidateNotNullOrEmpty()]
        [ValidatePattern('^[a-zA-Z0-9]([a-zA-Z0-9-]*[a-zA-Z0-9])?$')]
        [string]$RepoOwner,
        
        [Parameter(Mandatory=$true)]
        [ValidateNotNullOrEmpty()]
        [ValidatePattern('^[a-zA-Z0-9._-]+$')]
        [string]$RepoName,
        
        [Parameter(Mandatory=$true)]
        [ValidateNotNullOrEmpty()]
        [string]$PackageDescription,
        
        [ValidateSet("coveralls", "codecov")]
        [string]$CoverageService = "coveralls"
    )
    
    $template = @"
## Description
Add [$RepoName] - $PackageDescription

## Quality Check Links
forge link: https://github.com/$RepoOwner/$RepoName
pkg.go.dev: https://pkg.go.dev/github.com/$RepoOwner/$RepoName
goreportcard.com: https://goreportcard.com/report/github.com/$RepoOwner/$RepoName
"@

    if ($CoverageService -eq "codecov") {
        $template += "`ncoverage: https://codecov.io/gh/$RepoOwner/$RepoName"
    } else {
        $template += "`ncoverage: https://coveralls.io/github/$RepoOwner/$RepoName"
    }
    
    $template += @"

## Additional Information
- Follows awesome-go contribution guidelines
- Package is actively maintained
- All quality checks should pass with above links
"@

    return $template
}

Write-Host " PR Template Generator Created!" -ForegroundColor Green
Write-Host ""
Write-Host "Parameter validation:" -ForegroundColor Magenta
Write-Host "- RepoOwner: Must be valid GitHub username (alphanumeric, may contain hyphens)" -ForegroundColor Gray
Write-Host "- RepoName: Must be valid repository name (alphanumeric, dots, underscores, hyphens)" -ForegroundColor Gray
Write-Host "- CoverageService: Must be 'coveralls' or 'codecov'" -ForegroundColor Gray
