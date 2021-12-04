import {
  AccessTime,
  AddAlarm,
  AccountBalance,
  AddLocation,
  Apps,
  AssignmentInd,
} from '@mui/icons-material';

/**
 * Создание мока задачи
 */
export const mockTask = (overrides) => ({
  title: 'Title',
  icon: 0,
  description: 'Description',
  importance: 5,
  urgency: 5,
  deadline: new Date().toString(),
  estimated: 86400,
  ...overrides,
});

export const icons = {
  ACCESS_TIME: {
    id: 0,
    name: 'AccessTime',
    component: AccessTime,
  },
  ADD_ALARM: {
    id: 1,
    name: 'AddAlarm',
    component: AddAlarm,
  },
  ACCOUNT_BALANCE: {
    id: 2,
    name: 'AccountBalance',
    component: AccountBalance,
  },
  ADD_ALARM: {
    id: 3,
    name: 'AddLocation',
    component: AddLocation,
  },
  APPS: {
    id: 4,
    name: 'Apps',
    component: Apps,
  },
  ASSIGMENT_IND: {
    id: 5,
    name: 'AssignmentInd',
    component: AssignmentInd,
  },
};

export function getIconById(id) {
  return Object.values(icons).find((icon) => icon.id === id);
}
