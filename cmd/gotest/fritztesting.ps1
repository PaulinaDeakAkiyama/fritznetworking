# Load saved credentials
$cred = Import-Clixml "$env:USERPROFILE\fritz.cred.xml"

# Ignore self-signed FRITZ!Box certificate
Add-Type @"
using System.Net;
using System.Security.Cryptography.X509Certificates;

public class TrustAllCertsPolicy : ICertificatePolicy {
    public bool CheckValidationResult(
        ServicePoint srvPoint,
        X509Certificate certificate,
        WebRequest request,
        int certificateProblem) {
        return true;
    }
}
"@

[System.Net.ServicePointManager]::CertificatePolicy = New-Object TrustAllCertsPolicy

function ConvertTo-SoapArgumentXml {
    param(
        [hashtable]$Arguments = @{}
    )

    foreach ($key in $Arguments.Keys) {
        $value = [System.Security.SecurityElement]::Escape(
            [string]$Arguments[$key]
        )

        "<$key>$value</$key>"
    }
}

function Invoke-FritzRequest {
    param(
        [string]$Service,
        [string]$Action,
        [hashtable]$Arguments = @{},
        [int]$Port,
        [pscredential]$Credential
    )

    $url = "https://fritz.box:$Port/upnp/control/$($Service.ToLower())"

    $body = @"
<?xml version="1.0" encoding="utf-8"?>
<s:Envelope
    xmlns:s="http://schemas.xmlsoap.org/soap/envelope/">
  <s:Body>
    <u:$Action xmlns:u="urn:dslforum-org:service:$Service`:1">
$(ConvertTo-SoapArgumentXml $Arguments)
    </u:$Action>
  </s:Body>
</s:Envelope>
"@

    $headers = @{
        SOAPAction = "urn:dslforum-org:service:$Service`:1#$Action"
    }

    Invoke-WebRequest `
        -Uri $url `
        -Method POST `
        -Credential $Credential `
        -Headers $headers `
        -Body $body `
        -ContentType "text/xml; charset=utf-8"
}

function Get-SecurityPort {
    param(
        [pscredential]$Credential
    )

    $body = @'
<?xml version="1.0" encoding="utf-8"?>
<s:Envelope xmlns:s="http://schemas.xmlsoap.org/soap/envelope/">
<s:Body>
<u:GetSecurityPort xmlns:u="urn:dslforum-org:service:DeviceInfo:1"/>
</s:Body>
</s:Envelope>
'@

    $headers = @{
        SOAPAction = 'urn:dslforum-org:service:DeviceInfo:1#GetSecurityPort'
    }

    $response = Invoke-WebRequest `
        -Uri "http://fritz.box:49000/upnp/control/deviceinfo" `
        -Method POST `
        -Credential $Credential `
        -Headers $headers `
        -Body $body `
        -ContentType "text/xml; charset=utf-8"

    [xml]$xml = $response.Content
    $xml.Envelope.Body.GetSecurityPortResponse.NewSecurityPort
}

# Discover HTTPS port once
$securePort = Get-SecurityPort -Credential $cred

Write-Host "Security Port: $securePort"

# Get number of hosts
$response = Invoke-FritzRequest `
    -Service Hosts `
    -Action GetHostNumberOfEntries `
    -Port $securePort `
    -Credential $cred

[xml]$xml = $response.Content

$count = [int]$xml.Envelope.Body.GetHostNumberOfEntriesResponse.NewHostNumberOfEntries

Write-Host "Hosts: $count"

# Enumerate hosts (FRITZ host indexes are 1-based)
for ($i = 1; $i -le 10; $i++) {
    try {
        $response = Invoke-FritzRequest `
            -Service Hosts `
            -Action GetGenericHostEntry `
            -Arguments @{
                NewIndex = $i
            } `
            -Port $securePort `
            -Credential $cred

        [xml]$xml = $response.Content

        $hostentry = $xml.Envelope.Body.GetGenericHostEntryResponse

        [pscustomobject]@{
            Index     = $i
            HostName  = $hostentry.NewHostName
            IPAddress = $hostentry.NewIPAddress
            MAC       = $hostentry.NewMACAddress
            Active    = $hostentry.NewActive
        }
    }
    catch {
        Write-Warning "Failed to retrieve host index $i, Error: $($_.Exception.Message)"
    }
}