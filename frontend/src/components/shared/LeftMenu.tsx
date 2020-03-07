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

const useMainStyles = makeStyles({
    root: {},
    listItem: {
        color: "rgb(255,255,255)"
    }
})

const CustomLink = ({ to, children, renderAs }: { to: any, children: any, renderAs: any }) => {
    return (
        <Link to={to} >
            {children}
        </Link >
    )
}

const MainList = () => {
    const classes = useMainStyles({})

    return (
        <div>
            <ListItem button>
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
            <ListItem button>
                <ListItemIcon>
                    <PeopleIcon />
                </ListItemIcon>
                <ListItemText primary="Customers" />
            </ListItem>
            <ListItem button>
                <ListItemIcon>
                    <BarChartIcon />
                </ListItemIcon>
                <ListItemText primary="Reports" />
            </ListItem>
            <ListItem button>
                <ListItemIcon>
                    <LayersIcon />
                </ListItemIcon>
                <ListItemText primary="Integrations" />
            </ListItem>
        </div>
    )
}

export const MainListItems = (
    <MainList />
)

export const SecondaryListItems = (
    <div>
        <ListSubheader inset>Saved reports</ListSubheader>
        <ListItem button>
            <ListItemIcon>
                <AssignmentIcon />
            </ListItemIcon>
            <ListItemText primary="Current month" />
        </ListItem>
        <ListItem button>
            <ListItemIcon>
                <AssignmentIcon />
            </ListItemIcon>
            <ListItemText primary="Last quarter" />
        </ListItem>
        <ListItem button>
            <ListItemIcon>
                <AssignmentIcon />
            </ListItemIcon>
            <ListItemText primary="Year-end sale" />
        </ListItem>
    </div>
);
