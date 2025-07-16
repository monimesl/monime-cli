package space

import (
	"context"
	"fmt"
	"github.com/monimesl/monime-cli/internal/text"
	"github.com/olekukonko/tablewriter"
	"os"
)

func NewService() (*Service, error) {
	return &Service{
		repository: &defaultRepository{},
	}, nil
}

type Service struct {
	repository Repository
}

func (s *Service) ShowSpaceList(ctx context.Context) error {
	space, ok, err := s.repository.GetActiveSpace(ctx)
	if err != nil {
		return err
	}
	list, err := s.repository.ListSpaces(ctx)
	if err != nil {
		return err
	}
	table := tablewriter.NewWriter(os.Stdout)
	table.Header([]string{"Name", "Id", "Alias", "URL"})
	for _, spc := range list.Items {
		row := []string{spc.Name, spc.Id, spc.Alias, spc.URL}
		if ok && spc.Id == space.Id {
			for i, d := range row {
				row[i] = text.Format(d, text.FormatOptions{Color: "green"})
			}
		}
		_ = table.Append(row)
	}
	return table.Render()
}

func (s *Service) ActivateSpace(ctx context.Context, idOrAlias string) error {
	fmt.Printf("🚀 Activating space: %s\n", text.FormatToGreen(idOrAlias))
	list, err := s.repository.ListSpaces(ctx)
	if err != nil {
		return err
	}
	spc, ok := list.GetSpace(idOrAlias)
	if !ok {
		return fmt.Errorf("no space with alias '%s' exist", idOrAlias)
	}
	if err = s.repository.ActivateSpace(spc); err != nil {
		return err
	}
	name := text.FormatToGreen(spc.Name)
	alias := text.FormatToGreen(idOrAlias)
	fmt.Printf("✅  Space %s with alias %s is now active.\n", name, alias)
	return nil
}
