import * as React from 'react'
import { useState } from 'react'
import styled from 'styled-components'
import { ControlConfig } from '../../containers/ControlPanelContainer'

import {
  Button,
  Grid,
  Dialog,
  DialogTitle,
  DialogContent,
  DialogContentText,
  DialogActions,
  TextField

} from '@material-ui/core'

const StyledButton = styled(Button)`
  height: 100%;
  width: 100%;
`


type ControlPanelConfigEditProps = {
  jsonConfig: string,
  setControlPanelConfig: (data: ControlConfig[]) => void
}

const ControlPanelConfigEdit = ({ jsonConfig, setControlPanelConfig }: ControlPanelConfigEditProps) => {
  const [exportJsonOpen, setExportJsonDialogOpen] = useState(false)
  const [importJsonOpen, setImportJsonDialogOpen] = useState(false)
  const [jsonInput, setJsonInput] = useState("")

  const handleLoadJson = () => {
    console.log('load the json!', jsonInput);
    console.log('load the json! (parsed)', JSON.parse(jsonInput));
    setControlPanelConfig(JSON.parse(jsonInput))
    setImportJsonDialogOpen(false)
  }
  return (
    <Grid container spacing={1}>
      <Grid item xs={2}>
        <StyledButton type="submit"
          variant="outlined"
          onClick={() => setExportJsonDialogOpen(true)}>Export Config</StyledButton>
        <Dialog
          open={exportJsonOpen}
          onClose={() => setExportJsonDialogOpen(false)}
          aria-labelledby="alert-dialog-title"
          aria-describedby="alert-dialog-description"
        >
          <DialogTitle id="alert-dialog-title">Current Layout Config</DialogTitle>
          <DialogContent>
            <DialogContentText id="alert-dialog-description">
              <pre>
                {jsonConfig}
              </pre>
            </DialogContentText>
          </DialogContent>
          <DialogActions>
            <Button onClick={() => setExportJsonDialogOpen(false)} variant="outlined">close</Button>
          </DialogActions>
        </Dialog>
      </Grid>
      <Grid item xs={2}>
        <StyledButton type="submit"
          variant="outlined"
          onClick={() => setImportJsonDialogOpen(true)}>Import Config</StyledButton>
        <Dialog
          open={importJsonOpen}
          onClose={() => setImportJsonDialogOpen(false)}
          aria-labelledby="alert-dialog-title"
          aria-describedby="alert-dialog-description"
        >
          <DialogTitle id="alert-dialog-title">Current Layout Config</DialogTitle>
          <DialogContent>
            <DialogContentText id="alert-dialog-description">

              <TextField
                id="outlined-multiline-static"
                label="Multiline"
                multiline
                fullWidth={true}
                value={jsonInput}
                onChange={(e) => setJsonInput(e.target.value)}
                rows={30}
                defaultValue="Enter JSON Here"
                variant="outlined"
              />
            </DialogContentText>
          </DialogContent>
          <DialogActions>
            <Button onClick={() => handleLoadJson()} variant="outlined">Load</Button>
          </DialogActions>
        </Dialog>

      </Grid>

    </Grid>
  )
}

export default ControlPanelConfigEdit