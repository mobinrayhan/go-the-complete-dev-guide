package dtosV1

import "mobin.dev/internals/model"

func ToNoteResponse(n model.Note) NoteResponse {
	return NoteResponse{
		Id:        n.Id,
		UserId:    n.UserId,
		Title:     n.Title,
		Body:      n.Body,
		Tags:      n.Tags,
		CreatedAt: n.CreatedAt,
		UpdatedAt: n.UpdatedAt,
	}
}

func ToNoteResponses(notes []model.Note) []NoteResponse {
	res := make([]NoteResponse, len(notes))

	for i, n := range notes {
		res[i] = ToNoteResponse(n)
	}

	return res
}
