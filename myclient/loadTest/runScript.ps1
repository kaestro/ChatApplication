$containerName = "k6_container"
$projectRoot = Resolve-Path "..\.."
$scriptNameWithoutExtension = [System.IO.Path]::GetFileNameWithoutExtension($args[0])

# Check if the container already exists
$containerExists = (docker ps -a -f "name=$containerName" --format "{{.Names}}") -eq $containerName
Write-Host "Container exists: $containerExists"

if ($containerExists) {
    # Remove the existing container
    docker rm $containerName
}

# Run the script inside the container
docker run `
    --name $containerName `
    -v "${projectRoot}:/chatApplication" `
    grafana/k6 run `
    "/chatApplication/tests/loadTest/${scriptNameWithoutExtension}.js" `
    --out json="/chatApplication/tests/loadTest/results/${scriptNameWithoutExtension}_results.json"
