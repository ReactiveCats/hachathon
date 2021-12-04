import { wrappedFetch } from '../../shared/wrapped-fetch';

class AuthService {
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
}

export const authService = new AuthService();
