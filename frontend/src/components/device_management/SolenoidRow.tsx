import * as React from 'react'
import { useState } from 'react'
import { TableCell, TableRow } from '@material-ui/core'
import Solenoid from '../../utils/Solenoid'

type SolenoidRowProps = {
    solenoid: Solenoid,
    cellClasses: any
}

const SolenoidRow = ({ solenoid, cellClasses }: SolenoidRowProps) => {
    return (
        <>
            <TableRow key={solenoid.uid}>
                <TableCell className={cellClasses.cells}>{solenoid.uid}</TableCell>
                <TableCell className={cellClasses.cells}>{solenoid.name}</TableCell>
                <TableCell className={cellClasses.cells}>{solenoid.headerPin}</TableCell>
                <TableCell className={cellClasses.cells}>{solenoid.type}</TableCell>
                <TableCell className={cellClasses.cells}>{solenoid.enabled ? 'Enabled' : 'Disabled'}</TableCell>
            </TableRow>
        </>
    )
}

export default SolenoidRow