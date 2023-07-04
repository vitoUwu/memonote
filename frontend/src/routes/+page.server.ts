import { PUBLIC_API_BASE_URL } from '$env/static/public';
import type { Note } from '$lib/types';
import { fail, redirect, type Actions } from '@sveltejs/kit';
import type { PageServerLoad } from './$types';

export const load = (async ({ cookies, fetch }) => {
  const token = cookies.get('auth') as string;

  const [response, error] = await fetch(`${PUBLIC_API_BASE_URL}/api/notes/`, {
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
    token,
    notes: response as Note[]
  };
}) satisfies PageServerLoad;

export const actions = {
  default: async ({ request, fetch, cookies }) => {
    const formData = await request.formData();

    const title = formData.get('title');
    const content = formData.get('content');

    console.log('title', title, 'content', content);

    const [response, err] = await fetch(`${PUBLIC_API_BASE_URL}/api/notes/`, {
      method: 'POST',
      body: JSON.stringify({ title, content }),
      headers: {
        Authorization: cookies.get('auth') as string,
        'Content-Type': 'application/json'
      }
    })
      .then((res) => res.json())
      .then((data) => [data, null])
      .catch((err) => [null, err]);

    console.log(response);

    if (response.message || err) {
      return fail(400, { error: response.message ?? err.message ?? 'Unknown Error' });
    }

    return { success: true };
  }
} satisfies Actions;
