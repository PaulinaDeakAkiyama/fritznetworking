$user = "fritzadmin"
$pass = "masteroffritz2002"

# Get the security port for the Hosts service

function Invoke-FritzRequest {
    param (
        [string]$service,
        [string]$action,
        [int]$port = 49000
    )

    $box = "192.168.178.1"
    $urlPath = $service.ToLower()
    $url = "http://$box`:$port/upnp/control/$urlPath"

    $body = @"
    <?xml version="1.0"?>
    <s:Envelope xmlns:s="http://schemas.xmlsoap.org/soap/envelope/"
    s:encodingStyle="http://schemas.xmlsoap.org/soap/encoding/">
    <s:Body>
    <u:$action xmlns:u="urn:dslforum-org:service:$service:1">
    </u:$action>
    </s:Body>
    </s:Envelope>
"@

    $headers = @{
    "Content-Type" = 'text/xml; charset="utf-8"'
    "SOAPAction"   = '"urn:dslforum-org:service:{0}:1#{1}"' -f $service, $action
    }

    $response = Invoke-WebRequest -Uri $url -Method POST -Body $body -Headers $headers

    return $response 
}

function Get-SecurityPort {
    $response = Invoke-FritzRequest -service "DeviceInfo" -action "GetSecurityPort"
    [xml]$xml = $response.Content
    return $xml.Envelope.Body.GetSecurityPortResponse.NewSecurityPort
}


$res = Invoke-FritzRequest -service "Hosts" -action "GetGenericHostEntry" -port (Get-SecurityPort)





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