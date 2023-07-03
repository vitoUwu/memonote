import type { Note } from '$lib/types';
import { redirect } from '@sveltejs/kit';
import type { PageServerLoad } from './$types';

export const load = (async ({ cookies, fetch }) => {
  const token = cookies.get('auth') as string;

  const [response, error] = await fetch('http://localhost:3000/api/notes/', {
    method: 'GET',
    headers: { Authorization: token }
  })
    .then((res) => res.json())
    .then((data) => [data, null])
    .catch((err) => [null, err]);

  if (error) {
    console.error(error);
    return {
      notes: []
    };
  }

  if (response.message === 'Unauthorized') {
    throw redirect(303, '/logout');
  }

  return {
    notes: response as Note[]
  };
}) satisfies PageServerLoad;
