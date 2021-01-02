# Device Commander

Device Commander is a utility for the management and control of RNA[ ]VT devices, including:

[ ] Solenoid Controllers
[ ] ArtNet LED Drivers
[ ] Sensors

## Functional Breakdown

### Device Registration

[ ] DeviceCommander will periodically probe every ip address in the configured subnet for a `/registration` endpoint.
[ ] Devices that respond appropriately will be registered & available to configure and add to a control panel.
[ ] Known devices will have their last configuration restored on re[ ]registration.

### Device Configuration

[ ] Registered devices may be passed static configuration based on their device type
[ ] Applied configurations are saved to a database.

### Health Check

[ ] DeviceCommander performs a periodic heartbeat check to all registered devices.
[ ] Compliant devices will respond to a GET `/health` with a 200.

### Device Directory

[ ] Lists all currently registered devices
[ ] For each device, it displays:
  [ ] Device Name
  [ ] UUID / Serial No.
  [ ] (Stretch) Board
  [ ] Health Status
    [ ] (stretch) Drill down to see paginated health check response, ordered by time
  [ ] Configuration status
    [ ] (stretch) Drill down to see config history, ordered by time
  [ ] Current session start date & time
    [ ] (stretch) First session start date & time

### Individual Device Edit

[ ] View device info
[ ] Update configuration values
[ ] Save Button
  [ ] Config blob is stored to the db
[ ] Push Button
  [ ] Config is sent to the device.
  [ ] Compliant devices will respond to a PUT at `/configuration`
  [ ] Success/Fail are stored to the db with the config blob

### Editable Control Panel

[ ] Registered & healthy devices are available to add to an open layout
[ ] Layout config export
[ ] Save layout config to database
[ ] Enable/Disable edit mode

## Lineage

DeviceCommander was originally forked from  [GoFire](https://github.com/RNA[ ]VT/GoFire)