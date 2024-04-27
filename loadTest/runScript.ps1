param (
    [Parameter(Mandatory=$true)]
    [string]$scriptPath
)

$unixStylePath = (Get-Location -PSProvider FileSystem).ProviderPath
$unixStyleScriptPath = $scriptPath.Replace('\', '/')
$scriptName = Split-Path -Leaf $unixStyleScriptPath -Resolve

docker run --rm -v ${unixStylePath}:/scripts grafana/k6 run /scripts/${unixStyleScriptPath} --out json=/scripts/results/${scriptName}_results.json