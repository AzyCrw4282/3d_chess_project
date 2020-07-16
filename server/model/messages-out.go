package model

const (
	//ViewMenu for displaying the game menu
	ViewMenu = "view-menu"
	//ViewBoard for displaying the game board
	ViewBoard = "view-board"
)

//Creates msgs to be sent to the front-end. A neat way to do so.

//CreateMessageSecret for easy access
func CreateMessageSecret(secret, id string) *MessageSecret {
	return &MessageSecret{Msg: "secret", Secret: secret, ID: id}
}

//MessageSecret for easy access
type MessageSecret struct {
	Msg    string `json:"msg"`
	Secret string `json:"secret"`
	ID     string `json:"id"`
}

//CreateMessageView for easy access
func CreateMessageView(view string) *MessageView {
	return &MessageView{Msg: "view", View: view}
}

//MessageView for easy access
type MessageView struct {
	Msg  string `json:"msg"`
	View string `json:"view"`
}

//CreateMessageError is a general error message constructor
func CreateMessageError(t, m string) *MessageError {
	return &MessageError{Msg: "error", Title: t, Body: m}
}

//MessageError is a general error message
type MessageError struct {
	Msg   string `json:"msg"`
	Title string `json:"title"`
	Body  string `json:"body"`
}

func CreateMessageListOfGames(games *ListOfGames) *MessageListOfGames {
	return &MessageListOfGames{Msg: "list-games", Games: games}
}

type MessageListOfGames struct {
	Msg   string       `json:"msg"`
	Games *ListOfGames `json:"games"`
}

func CreateMessageShareState(game *Game) *MessageShareState {
	return &MessageShareState{Msg: "share-state", Game: game}
}

type MessageShareState struct {
	Msg  string `json:"msg"`
	Game *Game  `json:"game"`
}
