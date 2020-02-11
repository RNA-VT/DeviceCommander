# Components
  - GoFire Components represent physical devices controlled through pin(s) on a Microcontroller
  - Components must inherit BaseComponent and implement the interface Component{}
  - Public methods outside of the Component{} interface should contain component workflows for execution by an api handler

## Solenoids
  - Types: Normally Closed or Normally Open
  - Modes:
    - Supply: Solenoid that controls a fuel tank, transport line or manifold inlet.
    - Outlet: exit from the pressurized system. Pilots, poofers and maybe some kind of venting. 

## Igniters
  - Types: Glowfly, Induction
  
### Future Work
- LED Driver
- Motors 
- Servos
- Inputs
  - ArtNet
  - Accels
  - Gyros
  - Other Sensors
  - Polling inputs from other micros
- Component Health Checks
  - Match in memory state per component to read state after commanded changes
- Multi-pin components