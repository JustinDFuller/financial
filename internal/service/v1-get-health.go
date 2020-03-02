package service

import context "context"

func (s *service) getHealth(ctx context.Context, req interface{}) (interface{}, error) {
	return "ok", nil
}
