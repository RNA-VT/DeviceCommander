import * as React from 'react'
import { useState } from 'react'
import { Button } from '@material-ui/core'
import DraggableCore, { DraggableData, DraggableEvent } from "react-draggable"
import styled from 'styled-components'

const StyledDiv = styled.div`
  display: inline-block;
`

type ControlButtonProps = {
  xPos: number,
  yPos: number,
  label: string
}

const ControlButton = ({ xPos, yPos, label }: ControlButtonProps) => {
  const [isOpen, setIsOpen] = useState(false)


  const triggerSolenoid = () => {
    if (isOpen) {
      setIsOpen(false)
    } else {
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
      onStop={onStop}>
      <StyledDiv>
        <div className="handle">X</div>
        <Button
          variant="outlined"
          onMouseDown={triggerSolenoid}
          onMouseUp={triggerSolenoid}>
          <h4>{label}</h4>
        </Button>
      </StyledDiv>

    </DraggableCore>
  )
}

export default ControlButton