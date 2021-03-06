import * as React from 'react'
import { useState } from 'react'
import { Button, FormControlLabel, Switch } from '@material-ui/core'
import DraggableCore, { DraggableData, DraggableEvent } from "react-draggable"
import styled from 'styled-components'

const StyledDiv = styled.div`
  display: inline-block;
`

type ControlSwitchProps = {
  componentUID: string,
  xPos: number,
  yPos: number,
  label: string,
  setPosition: (uid: string, xPos: number, yPos: number) => void,
  onAction: () => void,
  offAction: () => void
}

const ControlSwitch = ({ componentUID, xPos, yPos, label, setPosition, onAction, offAction }: ControlSwitchProps) => {
  const [isOpen, setIsOpen] = useState(false)

  if (isOpen == false) {
    offAction()
  } else {
    onAction()
  }

  const handleOnChange = () => {
    setIsOpen(!isOpen)
    if (isOpen == false) {
      onAction()
    } else {
      offAction()
    }
  }

  const onStop = (e: DraggableEvent, data: DraggableData) => {
    console.log('draggableData onStop', data);
    setPosition(componentUID, data.x, data.y)
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
        <FormControlLabel
          control={
            <Switch
              checked={isOpen}
              onChange={() => handleOnChange()}
              name="checkedB"
            />
          }
          label={label}
        />
      </StyledDiv>

    </DraggableCore>
  )
}

export default ControlSwitch