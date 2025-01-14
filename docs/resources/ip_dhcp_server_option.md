# routeros_ip_dhcp_server_option (Resource)
Creates a DHCP lease on the mikrotik device.

## Example Usage
```terraform
resource "routeros_ip_dhcp_server_option" "jumbo_frame_opt" {
  code    = 77
  name    = "jumbo-mtu-opt"
  value   = "0x2336"
}

resource "routeros_ip_dhcp_server_option" "tftp_option" {
  code    = 66
  name    = "tftpserver-66"
  value   = "s'10.10.10.22'"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `code` (Number) The number of the DHCP option
- `name` (String) The name of the DHCP option
- `value` (String) The value with formatting using Mikrotik settings https://wiki.mikrotik.com/wiki/Manual:IP/DHCP_Server#DHCP_Options

### Optional


### Read-Only

- `id` (String) The ID of this resource.
- `raw_value` (String) The computed value of the option as an hex value

## Import
Import is supported using the following syntax:
```shell
#The ID can be found via API or the terminal
#The command for the terminal is -> :put [/ip/dhcp-server/option/get [print show-ids]]
terraform import routeros_ip_dhcp_server_option.tftp_option "*1"
```
