import * as React from 'react'
import { FunctionComponent, useState } from 'react'
import { Button } from '@material-ui/core'
import DraggableCore, { DraggableData, DraggableEvent } from "react-draggable"
import styled from 'styled-components'

const StyledDiv = styled.div`
  display: inline-block;
`

const SolenoidButton = ({ solenoid, xPos, yPos }: { solenoid: any, xPos: number, yPos: number }) => {
  const [isOpen, setIsOpen] = useState(false)

  console.log(solenoid.name, xPos, yPos);

  console.log(solenoid)

  const triggerSolenoid = () => {
    if (isOpen) {
      solenoid.close()
      setIsOpen(false)
    } else {
      solenoid.open()
      setIsOpen(true)
    }
  }

  const onStop = (e: DraggableEvent, data: DraggableData) => {
    console.log('draggableData onStop', data);
  }

  return (

    <DraggableCore
      handle={".handle"}
      defaultPosition={{ x: xPos, y: yPos }}
      grid={[25, 25]}
      scale={1}
      onStop={onStop}
      key={solenoid.name + solenoid.uid}>
      <StyledDiv>
        <div className="handle">X</div>
        <Button
          variant="outlined"
          onMouseDown={triggerSolenoid}
          onMouseUp={triggerSolenoid}>
          <h4>{solenoid.name}</h4>
        </Button>
      </StyledDiv>

    </DraggableCore>
  )
}

export default SolenoidButton