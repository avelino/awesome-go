# PR Quality Check Resolution Helper Script
# Usage: Follow the steps below to resolve quality check failures

Write-Host " PR Quality Check Resolution Guide" -ForegroundColor Cyan
Write-Host "=====================================" -ForegroundColor Cyan
Write-Host ""
Write-Host "1. Open PR in browser:" -ForegroundColor Yellow
Write-Host "   https://github.com/avelino/awesome-go/pull/5890" -ForegroundColor White
Write-Host ""
Write-Host "2. Look for the 'pr-quality-check' comment and note which checks failed" -ForegroundColor Yellow
Write-Host ""
Write-Host "3. Common required format for PR body:" -ForegroundColor Yellow
Write-Host ""
Write-Host "forge link: https://github.com/username/repo-name" -ForegroundColor Green
Write-Host "pkg.go.dev: https://pkg.go.dev/github.com/username/repo-name" -ForegroundColor Green  
Write-Host "goreportcard.com: https://goreportcard.com/report/github.com/username/repo-name" -ForegroundColor Green
Write-Host "coverage: https://coveralls.io/github/username/repo-name" -ForegroundColor Green
Write-Host ""
Write-Host "4. Replace 'username/repo-name' with actual repository details" -ForegroundColor Yellow
Write-Host ""
Write-Host "5. Edit the PR description to include the missing links" -ForegroundColor Yellow
Write-Host ""
Write-Host "6. The workflow will automatically re-run and validate" -ForegroundColor Yellow
Write-Host ""
Write-Host " Tip: Our enhanced workflow now shows detailed failure reasons!" -ForegroundColor Cyan
