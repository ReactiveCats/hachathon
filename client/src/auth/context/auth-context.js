import { createContext, useReducer, useContext } from 'react';
import { authService } from '../services/auth-service';
import { makeDispatchWithEffects } from '../../shared/make-dispatch-with-effects';

const initialState = {
  authorized: false,
  error: null,
};

const AuthContext = createContext();
AuthContext.displayName = 'AuthContext';

export const AUTH_LOGIN = 'auth login';
export const AUTH_SIGNUP = 'auth signup';
export const AUTH_LOGOUT = 'auth logout';
const AUTH_SUCCESS = 'auth success';
const AUTH_ERROR = 'auth error';

export function useAuthContext() {
  return useContext(AuthContext);
}

function saveAccessToken(accessToken) {
  window.localStorage.setItem('accessToken', accessToken);
}

function removeAccessToken() {
  window.localStorage.removeItem('accessToken');
}

function login(user, dispatch) {
  authService
    .login(user)
    .then(({ accessToken }) => {
      saveAccessToken(accessToken);
    })
    .catch((err) => {
      dispatch({ type: AUTH_ERROR, data: err });
    });
}

function signup(user, dispatch) {
  authService
    .signup(user)
    .then(({ accessToken }) => {
      saveAccessToken(accessToken);
    })
    .catch((err) => {
      dispatch({ type: AUTH_ERROR, data: err });
    });
}

function logout(data, dispatch) {
  removeAccessToken();
}

const authEffects = {
  [AUTH_LOGIN]: login,
  [AUTH_SIGNUP]: signup,
  [AUTH_LOGOUT]: logout,
};

function authReducer(state = initialState, action) {
  const { type, data = {} } = action;

  console.log('action', action);

  switch (type) {
    case AUTH_LOGIN:
    case AUTH_SIGNUP:
    case AUTH_LOGOUT:
      return initialState;
    case AUTH_SUCCESS:
      return {
        authorized: true,
        error: null,
      };
    case AUTH_ERROR:
      return {
        authorized: false,
        error: data,
      };
    default:
      throw new Error('Unsupported action type');
  }
}

export function AuthProvider({ children }) {
  const [state, dispatch] = useReducer(authReducer, initialState);

  const dispatchWithEffects = makeDispatchWithEffects(dispatch, authEffects);

  return (
    <AuthContext.Provider value={{ state, dispatch: dispatchWithEffects }}>
      {children}
    </AuthContext.Provider>
  );
}
