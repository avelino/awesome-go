function Generate-PRTemplate {
    param(
        [Parameter(Mandatory=$true)]
        [string]$RepoOwner,
        
        [Parameter(Mandatory=$true)]
        [string]$RepoName,
        
        [Parameter(Mandatory=$true)]
        [string]$PackageDescription,
        
        [string]$CoverageService = "coveralls" # or "codecov"
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

# Example usage:
Write-Host " PR Template Generator Created!" -ForegroundColor Green
Write-Host ""
Write-Host "Usage example:" -ForegroundColor Yellow
Write-Host 'Generate-PRTemplate -RepoOwner "username" -RepoName "package-name" -PackageDescription "Brief description of what the package does"' -ForegroundColor Cyan
Write-Host ""
Write-Host "This will generate a properly formatted PR body with all required links." -ForegroundColor White
