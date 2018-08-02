package errors

/*
Erreurs applicatives utilisées dans le store.
Elles permettent d'apporter des informations sur le code HTTP à retourner.

Au début, un service retourne ses propres types d'erreurs, mais cela devient vite chiant à gérer car il faut faire la correspondance avec les codes retour HTTP, si on veut un peu de maitrise.

L'idée est donc de faire comme dans mattermost-server, à savoir définir le code HTTP au niveau de l'erreur levée.

*/
type AppError struct {
	Message  string
	HTTPCode int
}

func (e *AppError) Error() string {
	return e.Message
}

func InvalidParameter(message string) error {
	return &AppError{
		Message:  message,
		HTTPCode: 400,
	}
}

func MissingResource(message string) error {
	return &AppError{
		Message:  message,
		HTTPCode: 404,
	}
}

func ServerError(message string) error {
	return &AppError{
		Message:  message,
		HTTPCode: 500,
	}
}
