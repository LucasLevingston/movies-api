package usecase

import "context"

func (useCase *movieUseCase) Delete(ctx context.Context, id string) error {
	return useCase.repo.Delete(ctx, id)
}
