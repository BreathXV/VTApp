Write-Host -Object "Build VTApp" -BackgroundColor Blue -ForegroundColor White

$Source = "/src/main.go"
$Output = "/build/VTApp.exe"

go build -o $Source -v $Output