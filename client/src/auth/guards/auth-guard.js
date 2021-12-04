import { useRouter } from 'next/router';
import { useEffect } from 'react';
import { useAuthContext } from '../context/auth-context';

export function AuthGuard({ children }) {
  const router = useRouter();

  const { state, dispatch } = useAuthContext();

  useEffect(() => {
    if (!state.authorized) {
      router.push({
        pathname: '/auth',
      });
    }
  }, []);

  return children;
}
