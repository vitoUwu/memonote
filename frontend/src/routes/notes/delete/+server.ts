import { PUBLIC_API_BASE_URL } from "$env/static/public";
import { json } from "@sveltejs/kit";
import type { RequestHandler } from "./$types";

export const POST = (async ({ request, fetch, cookies }) => {
  const body = await request.json();
  const [response, err] = await fetch(`${PUBLIC_API_BASE_URL}/api/notes/${body.id as string}`, {
    method: "DELETE",
    headers: {
      Authorization: cookies.get('auth') as string
    },
  })
    .then(async (res) => {
      if (res.status !== 200) {
        const data = await res.json();
        return [data, null];
      }
      return [res, null]
    })
    .catch((err) => [null, err]);

  if (response?.status !== 200 || err) {
    return json({ error: response.message ?? err.message ?? "Unknown Error" });
  }

  return json({ success: true });
}) satisfies RequestHandler