Write-Host -Object "Build VTApp" -BackgroundColor Blue -ForegroundColor White

$Source = "..\VTApp\src\main.go"
$Output = "..\VTApp\build\VTApp.exe"

go build -o $Output -v $Source