import React from 'react';
import ListItem from '@material-ui/core/ListItem';
import ListItemIcon from '@material-ui/core/ListItemIcon';
import ListItemText from '@material-ui/core/ListItemText';
import ListSubheader from '@material-ui/core/ListSubheader';
import DashboardIcon from '@material-ui/icons/Dashboard';
import ShoppingCartIcon from '@material-ui/icons/ShoppingCart';
import PeopleIcon from '@material-ui/icons/People';
import BarChartIcon from '@material-ui/icons/BarChart';
import LayersIcon from '@material-ui/icons/Layers';
import AssignmentIcon from '@material-ui/icons/Assignment';
import { Link } from 'react-router-dom';
import { makeStyles } from '@material-ui/core';
import PanToolIcon from '@material-ui/icons/PanTool'
import { red } from '@material-ui/core/colors';

const useMainStyles = makeStyles({
    root: {},
    listItem: {
        color: "rgb(255,255,255)"
    },
    warning: {
        color: "#f44336"
    }
})

const CustomLink = ({ to, children, renderAs }: { to: any, children: any, renderAs: any }) => {
    return (
        <Link to={to} >
            {children}
        </Link >
    )
}

const MainListItems = () => {
    const classes = useMainStyles({})


    return (
        <div>
            <ListItem button
                className={classes.listItem}
                component={Link}
                to="/">
                <ListItemIcon>
                    <DashboardIcon />
                </ListItemIcon>
                <ListItemText primary="Dashboard" />
            </ListItem>
            <ListItem
                button
                className={classes.listItem}
                component={Link}
                to="/device-management">
                <ListItemIcon>
                    <ShoppingCartIcon />
                </ListItemIcon>
                <ListItemText primary="Device Management" />
            </ListItem>
            <ListItem
                button
                className={classes.listItem}
                component={Link}
                to="/dashboard/edit">
                <ListItemIcon>
                    <DashboardIcon />
                </ListItemIcon>
                <ListItemText primary="Dashboard Edit" />
            </ListItem>
            <ListItem button disabled>
                <ListItemIcon>
                    <PeopleIcon />
                </ListItemIcon>
                <ListItemText primary="Users" />
            </ListItem>
            <ListItem button disabled>
                <ListItemIcon>
                    <BarChartIcon />
                </ListItemIcon>
                <ListItemText primary="Reports" />
            </ListItem>
            <ListItem button disabled>
                <ListItemIcon>
                    <LayersIcon />
                </ListItemIcon>
                <ListItemText primary="Integrations" />
            </ListItem>
        </div>
    )
}

export const MainList = (
    <MainListItems />
)

const SecondaryListItems = () => {
    const classes = useMainStyles({})

    return (
        < div >
            <ListSubheader inset>Shortcuts</ListSubheader>
            <ListItem button className={classes.warning}>
                <ListItemIcon>
                    <PanToolIcon color="error" />
                </ListItemIcon>
                <ListItemText primary="Disable All" />
            </ListItem>
        </div >

    )
}

export const SecondaryList = (
    <SecondaryListItems />
)
