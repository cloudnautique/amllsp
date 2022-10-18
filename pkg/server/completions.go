package server

import (
	"strings"

	"github.com/acorn-io/amllsp/pkg/stdlib"
	"github.com/tliron/glsp"
	protocol "github.com/tliron/glsp/protocol_3_16"
	"go.lsp.dev/uri"
)

func (h *Handler) completion(ctx *glsp.Context, params *protocol.CompletionParams) (interface{}, error) {
	var completions protocol.CompletionList
	var line string

	_uri, err := uri.Parse(params.TextDocument.URI)
	if err != nil {
		return nil, err
	}
	h.log.Errorf("gotpartial result: %#v", params.PartialResultParams)
	h.log.Errorf("got text doc : %#v", params.TextDocument)
	h.log.Errorf("got workdown doc : %#v", params.WorkDoneProgressParams.WorkDoneToken)
	a := h.workspace.GetFile(_uri.Filename())

	content, err := a.StringContent()
	if err != nil {
		return nil, err
	}

	if len(content) > 0 {
		h.log.Errorf("got content: %s", content)
		//line = completionLine(string(content), params.Position)
	}

	completions.Items = h.stdLibCompletions(line)
	completions.IsIncomplete = true

	return completions, nil
}

func (h *Handler) stdLibCompletions(line string) []protocol.CompletionItem {
	var completions []protocol.CompletionItem

	fnc, err := stdlib.Functions()
	if err != nil {
		return nil
	}

	stdIndex := strings.LastIndex(line, "std.")
	if stdIndex != -1 {
		userInput := line[stdIndex+4:]
		funcStartWith := []protocol.CompletionItem{}
		funcContains := []protocol.CompletionItem{}

		for _, f := range fnc {
			if f.Name == userInput {
				break
			}
			lowerFuncName := strings.ToLower(f.Name)
			findName := strings.ToLower(userInput)

			details := f.Markdown()
			kind := protocol.CompletionItemKindFunction
			item := protocol.CompletionItem{
				Label:      f.Name,
				Kind:       &kind,
				Detail:     &details,
				InsertText: &f.Name,
			}

			if len(findName) > 0 && strings.HasPrefix(lowerFuncName, findName) {
				funcStartWith = append(funcStartWith, item)
				continue
			}

			if strings.Contains(lowerFuncName, findName) {
				funcContains = append(funcContains, item)
			}
		}

		completions = append(completions, funcStartWith...)
		completions = append(completions, funcContains...)
	}
	return completions
}

func completionLine(content string, position protocol.Position) string {
	line := strings.Split(content, "\n")[position.Line]
	charIndex := int(position.Character)
	if charIndex > len(line) {
		charIndex = len(line)
	}
	line = line[:charIndex]
	return line
}
