package utils

import "strings"

func SelectNote(knowledgeBasePath string) (noteName string, err error) {

	notes, err := ReadAllNotes(knowledgeBasePath)

	if err != nil {
		panic(err)
	}

	var noteNames []string
	var notesMap = make(map[string]Note)

	for _, note := range notes {
		noteNames = append(noteNames, note.Name)
		notesMap[note.Name] = note
	}

	reader := strings.NewReader(strings.Join(noteNames, "\n"))

	noteName, err = Fzf(reader)

	if err != nil {
		return "", err
	}

	notePath := notesMap[noteName].Path

	return notePath, nil
}
