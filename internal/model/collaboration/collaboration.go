package collaboration

type PlaylistCollaboration struct {
	UserId     int
	PlaylistId int
}

type AddCollaboratorRequest struct {
	UserId     int `json:"user_id"`
}

type RemoveCollaboratorRequest struct {
	UserId     int `json:"user_id"`
}
