import { AccessTime, AddAlarm, AccountBalance, AddLocation, Apps, AssignmentInd } from '@mui/icons-material';

export const ICONS = {
    ACCESS_TIME: {
        id: '0',
        name: 'AccessTime',
        component: AccessTime
    },
    ADD_ALARM: {
        id: '1',
        name: 'AddALarm',
        component: AddAlarm,
    },
    ACCOUNT_BALANCE: {
        id: '2',
        name: 'AccountBalance',
        component: AccountBalance,
    },
    ADD_ALARM: {
        id: '3',
        name: 'AddLocation',
        component: AddLocation,
    },
    APPS: {
        id: '4',
        name: 'Apps',
        component: Apps,
    },
    ASSIGMENT_IND: {
        id: '5',
        name: 'AssignmentInd',
        component: AssignmentInd,
    },
};

export function getIcons() {
    return Object.values(ICONS);
}

export function getIconById(id) {
    const icons = getIcons();
    return icons.find(icon => icon.id === id);
}