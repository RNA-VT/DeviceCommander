import * as React from 'react'
import { useState } from 'react'
import SolenoidEdit from './SolenoidEdit'

type SolenoidRowProps = {
    solenoid: any,
    cellClasses: any
}

const SolenoidRow = ({ solenoid, cellClasses }: SolenoidRowProps) => {
    const [open, setOpen] = useState(false)

    return (
        <>
            <tr key={solenoid.UID}>
                <td className={cellClasses.cells}>{solenoid.UID}</td>
                <td className={cellClasses.cells}>{solenoid.Name}</td>
                <td className={cellClasses.cells}>{solenoid.HeaderPin}</td>
                <td className={cellClasses.cells}>{solenoid.Type}</td>
                <td className={cellClasses.cells}>{solenoid.Enabled ? 'Enabled' : 'Disabled'}</td>
            </tr>
        </>
    )
}

export default SolenoidRow