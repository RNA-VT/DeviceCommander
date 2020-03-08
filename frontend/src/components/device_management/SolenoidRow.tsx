import * as React from 'react'
import { useState } from 'react'
import Solenoid from '../../utils/Solenoid'

type SolenoidRowProps = {
    solenoid: Solenoid,
    cellClasses: any
}

const SolenoidRow = ({ solenoid, cellClasses }: SolenoidRowProps) => {
    return (
        <>
            <tr key={solenoid.uid}>
                <td className={cellClasses.cells}>{solenoid.uid}</td>
                <td className={cellClasses.cells}>{solenoid.name}</td>
                <td className={cellClasses.cells}>{solenoid.headerPin}</td>
                <td className={cellClasses.cells}>{solenoid.type}</td>
                <td className={cellClasses.cells}>{solenoid.enabled ? 'Enabled' : 'Disabled'}</td>
            </tr>
        </>
    )
}

export default SolenoidRow