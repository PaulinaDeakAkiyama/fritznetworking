# Save
Get-Credential | Export-Clixml "$env:USERPROFILE\fritz.cred.xml"
# Load
$cred = Import-Clixml "$env:USERPROFILE\fritz.cred.xml"

function ConvertTo-SoapArgumentXml {
  <#
  .SYNOPSIS
    Converts a hashtable of arguments into XML elements for a SOAP request.
  .PARAMETER Arguments
    A hashtable containing the arguments to convert. Strings musnt contain XML special characters or spaces. Remember to capitalise the first letter of each word.
  .EXAMPLE
    $args = @{ NewIndex = 0 }
    $xml = ConvertTo-SoapArgumentXml -Arguments $args
    # $xml will contain:
    # <NewIndex>0</NewIndex>  
  .OUTPUTS
    A string containing the XML representation of the arguments.  
  #>
    param (
        [hashtable]$Arguments = @{}
    )

    if (-not $Arguments -or $Arguments.Count -eq 0) {
        return ""
    }

    $lines = foreach ($key in $Arguments.Keys) {
        $escapedkey = [System.Security.SecurityElement]::Escape([string]$key)
        $escapedValue = [System.Security.SecurityElement]::Escape([string]$Arguments[$key])
        "<$escapedkey>$escapedValue</$escapedkey>"
    }

    return ($lines -join "`n")
}

function Invoke-FritzRequest {
    param (
        [string]$service,
        [string]$action,
        [hashtable]$arguments = @{},
        [int]$port = 49000,
        [pscredential]$credential
    )

    $box = "192.168.178.1"
    $urlPath = $service.ToLower()
    if ($port -ne 49000) {
        $url = "https://$box`:$port/upnp/control/$urlPath"
    } else {
        $url = "http://$box`:$port/upnp/control/$urlPath"
    }

    $body = @"
    <?xml version="1.0"?>
    <s:Envelope xmlns:s="http://schemas.xmlsoap.org/soap/envelope/"
    s:encodingStyle="http://schemas.xmlsoap.org/soap/encoding/">
    <s:Body>
    <u:$action xmlns:u="urn:dslforum-org:service:$service:1">
    $(ConvertTo-SoapArgumentXml -Arguments $arguments)
    </u:$action>
    </s:Body>
    </s:Envelope>
"@

    $headers = @{
    "Content-Type" = 'text/xml; charset="utf-8"'
    "SOAPAction"   = '"urn:dslforum-org:service:{0}:1#{1}"' -f $service, $action
    }

    $response = Invoke-WebRequest -Uri $url -Method POST -Body $body -Headers $headers -Credential $credential

    return $response 
}

function Get-SecurityPort {
    $response = Invoke-FritzRequest -service "DeviceInfo" -action "GetSecurityPort"
    [xml]$xml = $response.Content
    return $xml.Envelope.Body.GetSecurityPortResponse.NewSecurityPort
}


$res = Invoke-FritzRequest -service "Hosts" -action "GetGenericHostEntry" -port (Get-SecurityPort) -credential $cred





$max = 50  # HostNumberOfEntries value

for ($i=0; $i -lt $max; $i++) {
  $bodyEntry = @"
<?xml version="1.0" encoding="utf-8"?>
<s:Envelope xmlns:s="http://schemas.xmlsoap.org/soap/envelope/" s:encodingStyle="http://schemas.xmlsoap.org/soap/encoding/">
  <s:Body>
    <u:GetGenericHostEntry xmlns:u="urn:dslforum-org:service:Hosts:1">
      <NewIndex>$i</NewIndex>
    </u:GetGenericHostEntry>
  </s:Body>
</s:Envelope>
"@

  curl.exe --digest -u "$user`:$pass" `
    -H "Content-Type: text/xml; charset=utf-8" `
    -H "SOAPACTION: ""urn:dslforum-org:service:Hosts:1#GetGenericHostEntry""" `
    --data "$bodyEntry" `
    "http://$box`:$secureport$hostsControl"
}

$securePass = ConvertTo-SecureString $pass -AsPlainText -Force
$credential = [pscredential]::new($user, $securePass)

Invoke-FritzRequest -Service "Hosts" -Action "GetHostNumberOfEntries" -Port (Get-SecurityPort) -Credential (Import-Clixml "$env:USERPROFILE\fritz.cred.xml")