import { JWT_SECRET } from '$env/static/private';
import type { Handle } from '@sveltejs/kit';
import jwt from 'jsonwebtoken';
import type { User } from './lib/types';

const routesBypass = ["/login", "/register"];

const redirectTo = (path: string) => {
  return new Response('Redirect', { status: 303, headers: { Location: path } });
};

export const handle = (async ({ event, resolve }) => {
  console.debug(event.request.method, event.url.pathname, Boolean(event.cookies.get('auth')));

  if (event.url.pathname === '/logout') {
    event.cookies.delete('auth');
    event.locals.user = null;
    return await resolve(event);
  }


  const token = event.cookies.get('auth');

  if (!token) {
    if (!routesBypass.includes(event.url.pathname)) {
      return redirectTo('/logout');
    }

    return await resolve(event);
  }

  try {
    jwt.verify(token, JWT_SECRET);
  } catch (_) {
    event.cookies.delete('auth');
    event.locals.user = null;
    return redirectTo('/logout');
  }

  const user = await event
    .fetch('http://localhost:3000/api/users/me', { headers: { Authorization: token } })
    .then((res) => res.json() as Promise<User>)
    .catch(console.error);

  if (!user) {
    return redirectTo('/logout');
  }

  event.locals.user = user;

  return await resolve(event);
}) satisfies Handle;
