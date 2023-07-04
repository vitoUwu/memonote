import { PUBLIC_API_BASE_URL } from '$env/static/public';
import { fail, redirect, type Actions } from '@sveltejs/kit';
import type { PageServerLoad } from '../$types';

export const load = (async ({ cookies }) => {
	if (cookies.get('auth')) {
		throw redirect(303, '/');
	}
}) satisfies PageServerLoad;

export const actions = {
	default: async (event) => {
		const data = await event.request.formData();
		const username = data.get('username');
		const password = data.get('password');
		const confirmPassword = data.get('confirm-password');

		if (!username) {
			return fail(400, { field: 'username', error: 'Username is required' });
		}

		if (!password) {
			return fail(400, { field: 'password', error: 'Password is required' });
		}

		if (!confirmPassword) {
			return fail(400, { field: 'confirmPassword', error: 'Confirm Password is required' });
		}

		if (confirmPassword !== password) {
			return fail(400, { error: 'Password and Confirm Passord must be equal each other' });
		}

		const [response, error] = await event
			.fetch(`${PUBLIC_API_BASE_URL}/api/users/register`, {
				method: 'POST',
				headers: {
					'Content-Type': 'application/json'
				},
				body: JSON.stringify({ username, password })
			})
			.then((res) => res.json())
			.then((data) => [data, null])
			.catch((err) => [null, err]);

		if (error || !response['access_token']) {
			if (response.message) {
				return fail(400, { error: response.message });
			}
			return fail(400, {
				error: 'An error has occurred while creating your account. Try again later'
			});
		}

		event.cookies.set('auth', response.access_token, {
			maxAge: 60 * 60 * 24, // 1 day
			httpOnly: true,
			sameSite: 'strict',
			secure: true
		});

		throw redirect(303, '/');
	}
} satisfies Actions;
