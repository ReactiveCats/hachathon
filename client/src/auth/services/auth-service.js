import { wrappedFetch } from '../../shared/wrapped-fetch';

class AuthService {
  static ACCESS_TOKEN_STORAGE_KEY = 'accessToken';

  login(user) {
    return wrappedFetch('/api/auth/login', {
      method: 'POST',
      body: user,
    });
  }

  signup(user) {
    return wrappedFetch('/api/auth/signup', {
      method: 'POST',
      body: user,
    });
  }

  logout() {
    this.removeAccessToken();
  }

  saveAccessToken(accessToken) {
    window.localStorage.setItem(
      AuthService.ACCESS_TOKEN_STORAGE_KEY,
      accessToken,
    );
  }

  getAccessToken() {
    return window.localStorage.getItem(AuthService.ACCESS_TOKEN_STORAGE_KEY);
  }

  removeAccessToken() {
    window.localStorage.removeItem(AuthService.ACCESS_TOKEN_STORAGE_KEY);
  }
}

export const authService = new AuthService();
