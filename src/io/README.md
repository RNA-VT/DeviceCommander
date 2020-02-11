# IO Module
- IO represents the controllable concepts available on a microcontroller

## Mocks
- IOs must declare an interface to support mocking
- Mocking should be handled in Init()
- Mock or Real should be invisible outside of the IO module

## GPIO 
- Digital pin control object
- rpi_pin.go contains the mapping of the various naming conventions for pins on a RasPi

### Future Work
- Digital Input
- Analog Input/Output
- Serial In/Out
- 