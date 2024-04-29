param (
    [Parameter(Mandatory=$true)]
    [string]$scriptPath
)

$unixStylePath = (Get-Location -PSProvider FileSystem).ProviderPath
$unixStyleScriptPath = $scriptPath.Replace('\', '/').Replace('./', '')
$scriptName = Split-Path -Leaf $unixStyleScriptPath -Resolve
$scriptNameWithoutExtension = [IO.Path]::GetFileNameWithoutExtension($scriptName)

$projectRoot = Resolve-Path "${unixStylePath}/../.."

# Ensure the results directory exists
$resultsDirectory = "${projectRoot}/tests/loadTest/results"
if (!(Test-Path -Path $resultsDirectory)) {
    New-Item -ItemType Directory -Force -Path $resultsDirectory
}

docker run --rm `
    -v "${projectRoot}:/chatApplication" `
    grafana/k6 run `
    "/chatApplication/tests/loadTest/${scriptNameWithoutExtension}.js" `
    --out json="/chatApplication/tests/loadTest/results/${scriptNameWithoutExtension}_results.json"