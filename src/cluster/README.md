# Cluster Module
  - Cluster represents a set of Microcontrollers running instances of GoFire

  - Clusters are responsible for:
    - Determining and setting Master/Slave state of this microcontroller
    - If this microcontroller is in master mode, managing de/registration of other microcontrollers, ID generation and handling commands for local components. 
    - If this microcontroller is in slave mode, reporting to master and handling commands for local components. 
    - Microcontroller Deregistration (Graceful exit)
    - Heartbeat/Health Check - Master polls slaves at an interval and deregisters non responsive nodes

  ## Future Work:
  - Broadcast all stop (Safety event)
  - Retries for all registration events
