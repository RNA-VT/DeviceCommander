# Device Commander

Device Commander is a utility for the management and control of RNA-VT devices, including:

- Solenoid Controllers
- ArtNet LED Drivers
- Sensors

## Functional Breakdown

### Device Registration

- DeviceCommander will periodically probe every ip address in the configured subnet for a `/device-registration` endpoint.
- Devices that respond appropriately will be registered & available to configure and add to a control panel.
- Known devices will have their last configuration restored on re-registration.

### Device Configuration

- Registered devices may be passed static configuration based on their device type
- Applied configurations are saved to a database.

### Health Check

- DeviceCommander performs a periodic heartbeat check. 
- Compliant devices must respond to a GET `/health` with a 200
- 

### Editable Control Panel

## Notes

DeviceCommander was originally forked from  [GoFire](https://github.com/RNA-VT/GoFire)