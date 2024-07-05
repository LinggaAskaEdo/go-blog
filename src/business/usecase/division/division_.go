package division

import (
	"context"

	"github.com/linggaaskaedo/go-blog/src/business/dto"
	"github.com/linggaaskaedo/go-blog/src/business/entity"
	"github.com/linggaaskaedo/go-blog/src/common"
)

func (d *division) CreateDivision(ctx context.Context, divisionEntity entity.Division) (dto.DivisionDTO, error) {
	var divisionDTO dto.DivisionDTO

	result, err := d.division.CreateDivision(ctx, divisionEntity)
	if err != nil {
		return divisionDTO, err
	}

	divisionDTO.PublicID = common.MixerEncode(result.ID)
	divisionDTO.Name = result.Name
	divisionDTO.IsDeleted = result.IsDeleted

	return divisionDTO, nil
}

func (d *division) GetDivisioByID(ctx context.Context, divisionID int64) (dto.DivisionDTO, error) {
	var divisionDTO dto.DivisionDTO

	result, err := d.division.GetDivisioByID(ctx, divisionID)
	if err != nil {
		return divisionDTO, err
	}

	divisionDTO.PublicID = common.MixerEncode(result.ID)
	divisionDTO.Name = result.Name
	divisionDTO.IsDeleted = result.IsDeleted

	return divisionDTO, nil
}
