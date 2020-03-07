import * as React from 'react'
import { useState } from 'react'
import { Button, Dialog, makeStyles } from '@material-ui/core'
import EditIcon from '@material-ui/icons/Edit'

import AppBar from '@material-ui/core/AppBar';
import Toolbar from '@material-ui/core/Toolbar';
import IconButton from '@material-ui/core/IconButton';
import Typography from '@material-ui/core/Typography';
import CloseIcon from '@material-ui/icons/Close';
import Slide from '@material-ui/core/Slide';
import { TransitionProps } from '@material-ui/core/transitions';

const useStyles = makeStyles(theme => ({
    appBar: {
        position: 'relative',
    },
    title: {
        marginLeft: theme.spacing(2),
        flex: 1,
    },
}))

const Transition = React.forwardRef<unknown, TransitionProps>(function Transition(props, ref) {
    return <Slide direction="up" ref={ref} {...props} />;
});

const SolenoidRow = ({ solenoid, cellClasses }: { solenoid: any, cellClasses: any }) => {
    const classes = useStyles()
    const [open, setOpen] = useState(false)

    const handleClose = () => {
        setOpen(false)
    }

    const handleEditOpen = () => {
        setOpen(true)
    }

    return (
        <>
            <tr key={solenoid.UID}>
                <td className={cellClasses.cells}>{solenoid.UID}</td>
                <td className={cellClasses.cells}>{solenoid.Name}</td>
                <td className={cellClasses.cells}>{solenoid.HeaderPin}</td>
                <td className={cellClasses.cells}>{solenoid.Type}</td>
                <td className={cellClasses.cells}>{solenoid.Enabled ? 'Enabled' : 'Disabled'}</td>
                <td className={cellClasses.cells}>
                    <Button onClick={handleEditOpen}><EditIcon /></Button>
                </td>
            </tr>
            <Dialog fullScreen open={open} onClose={handleClose} TransitionComponent={Transition}>
                <AppBar className={classes.appBar}>
                    <Toolbar>
                        <IconButton edge="start" color="inherit" onClick={handleClose} aria-label="close">
                            <CloseIcon />
                        </IconButton>
                        <Typography variant="h6" className={classes.title}>
                            Sound
            </Typography>
                        <Button autoFocus color="inherit" onClick={handleClose}>
                            save
            </Button>
                    </Toolbar>
                </AppBar>
            </Dialog>
        </>
    )
}

export default SolenoidRow