package division

import (
	"context"

	"github.com/linggaaskaedo/go-blog/src/business/dto"
)

func (d *division) CreateDivision(ctx context.Context, divisionDTO dto.DivisionDTO) (dto.DivisionDTO, error) {
	_, err := d.division.CreateDivision(ctx, divisionDTO)
	if err != nil {
		return divisionDTO, err
	}

	return divisionDTO, nil
}
