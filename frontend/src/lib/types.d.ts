export type User = {
	id: string;
	username: string;
	createdAt: string;
	updatedAt: string;
};

export type Note = {
	id: string;
	content: string;
	user_id: string;
	created_at: string;
	updated_at: string;
};
