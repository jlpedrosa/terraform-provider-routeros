# routeros_interface_ethernet (Resource)




<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String) Name of the ethernet interface.

### Optional

- `advertise` (String) Advertised speed and duplex modes for Ethernet interfaces over twisted pair, 
				only applies when auto-negotiation is enabled. Advertising higher speeds than 
				the actual interface supported speed will have no effect, multiple options are allowed.
- `arp` (String) Address Resolution Protocol mode:
				disabled - the interface will not use ARP
				enabled - the interface will use ARP
				local-proxy-arp - the router performs proxy ARP on the interface and sends replies to the same interface
				proxy-arp - the router performs proxy ARP on the interface and sends replies to other interfaces
				reply-only - the interface will only reply to requests originated from matching IP address/MAC address combinations which are entered as static entries in the ARP table. No dynamic entries will be automatically stored in the ARP table. Therefore for communications to be successful, a valid static entry must already exist.
- `auto_negotiation` (Boolean) When enabled, the interface "advertises" its maximum capabilities to achieve the best connection possible.
					Note1: Auto-negotiation should not be disabled on one end only, otherwise Ethernet Interfaces may not work properly.
					Note2: Gigabit Ethernet and NBASE-T Ethernet links cannot work with auto-negotiation disabled.
- `bandwidth` (Number) Sets max rx/tx bandwidth in kbps that will be handled by an interface. TX limit is supported on all Atheros switch-chip ports. 
				RX limit is supported only on Atheros8327/QCA8337 switch-chip ports.
- `cable_setting` (String) Changes the cable length setting (only applicable to NS DP83815/6 cards)
- `combo_mode` (String) When auto mode is selected, the port that was first connected will establish the link. In case this link fails, the other port will try to establish a new link. If both ports are connected at the same time (e.g. after reboot), 
				the priority will be the SFP/SFP+ port. When sfp mode is selected, the interface will only work through SFP/SFP+ cage.
				When copper mode is selected, the interface will only work through RJ45 Ethernet port.
- `comment` (String)
- `disable_running_check` (Boolean) Disable running check. If this value is set to 'no', the router automatically detects whether the NIC is connected with a device in the network or not.
						Default value is 'yes' because older NICs do not support it. (only applicable to x86)
- `full_duplex` (Boolean) Defines whether the transmission of data appears in two directions simultaneously, only applies when auto-negotiation is disabled.
- `l2mtu` (Number) Layer2 Maximum transmission unit. see (https://wiki.mikrotik.com/wiki/Maximum_Transmission_Unit_on_RouterBoards)
- `mac_address ` (String) Media Access Control number of an interface.
- `mdix_enable` (Boolean) Whether the MDI/X auto cross over cable correction feature is enabled for the port (Hardware specific, e.g. ether1 on RB500 can be set to yes/no. Fixed to 'yes' on other hardware.)
- `mtu` (Number) Layer3 Maximum transmission unit
- `poe-out` (String) PoE settings: (https://wiki.mikrotik.com/wiki/Manual:PoE-Out)
- `poe-priority` (Number) PoE settings: (https://wiki.mikrotik.com/wiki/Manual:PoE-Out)
- `rx_flow_control` (String) When set to on, the port will process received pause frames and suspend transmission if required.
					auto is the same as on except when auto-negotiation=yes flow control status is resolved by taking into account what other end advertises.
- `sfp-shutdown-temperature` (Number) The temperature in Celsius at which the interface will be temporarily turned off due to too high detected SFP module temperature (introduced v6.48).The default value for SFP/SFP+/SFP28 interfaces is 95, and for QSFP+/QSFP28 interfaces 80 (introduced v7.6).
- `speed` (String) Sets interface data transmission speed which takes effect only when auto-negotiation is disabled.
- `tx_flow_control` (String) When set to on, the port will generate pause frames to the upstream device to temporarily stop the packet transmission. 
					Pause frames are only generated when some routers output interface is congested and packets cannot be transmitted anymore. 
					Auto is the same as on except when auto-negotiation=yes flow control status is resolved by taking into account what other end advertises.

### Read-Only

- `id` (String) The ID of this resource.
- `orig-mac-address` (String) Original Media Access Control number of an interface. (read only)
- `running` (Boolean) Whether interface is running. Note that some interface does not have running check and they are always reported as "running"
- `slave` (Boolean) Whether interface is configured as a slave of another interface (for example Bonding)
- `switch` (Number) ID to which switch chip interface belongs to.

