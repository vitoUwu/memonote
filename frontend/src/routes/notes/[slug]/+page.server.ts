import { PUBLIC_API_BASE_URL } from "$env/static/public";
import type { Note } from "$lib/types";
import { fail, type Actions } from "@sveltejs/kit";
import type { PageServerLoad } from "./$types";

export const load = (async ({ params, fetch, cookies }) => {
  const [response, err] = await fetch(`${PUBLIC_API_BASE_URL}/api/notes/${params.slug}`, {
    method: "GET",
    headers: {
      Authorization: cookies.get("auth") as string
    }
  })
    .then((res) => res.json())
    .then((data) => [data, null])
    .catch((err) => [null, err]);

  if (response.message || err) {
    console.log(response, err)
    return { error: response.message ?? err.message ?? 'Unknown Error' }
  }

  return {
    note: response as Note,
  }
}) satisfies PageServerLoad;

export const actions = {
  default: async ({ cookies, fetch, request, params }) => {
    const formData = await request.formData();

    const title = formData.get('title');
    const content = formData.get('content');

    const [response, err] = await fetch(`${PUBLIC_API_BASE_URL}/api/notes/${params.slug}`, {
      method: 'PATCH',
      headers: {
        Authorization: cookies.get('auth') as string,
        'Content-Type': 'application/json'
      },
      body: JSON.stringify({ title, content })
    })
      .then((res) => res.json())
      .then((data) => [data, null])
      .catch((err) => [null, err]);

    if (response?.message || err) {
      return fail(400, { error: response?.message ?? err.message ?? "Unknown Error" });
    }

    return {
      success: true
    }
  }
} satisfies Actions