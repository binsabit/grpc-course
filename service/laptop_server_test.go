package service_test

import (
	"context"
	"testing"

	"github.com/binsabit/grpc-course/protogen/pb"
	"github.com/binsabit/grpc-course/samples"
	"github.com/binsabit/grpc-course/service"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestServerCreateLaptop(t *testing.T) {
	t.Parallel()

	laptopNoID := samples.NewLaptop()
	laptopNoID.Id = ""

	laptopInvalidID := samples.NewLaptop()
	laptopInvalidID.Id = "invalid-uuid"

	laptopDuplicateID := samples.NewLaptop()
	storeDuplicateID := service.NewInMemoryLaptopStore()
	err := storeDuplicateID.Save(laptopDuplicateID)
	assert.Nil(t, err)

	testCases := []struct {
		name   string
		laptop *pb.Laptop
		store  service.LaptopStore
		code   codes.Code
	}{
		{
			name:   "access_with_id",
			laptop: samples.NewLaptop(),
			store:  service.NewInMemoryLaptopStore(),
			code:   codes.OK,
		},
		{
			name:   "success_no_id",
			laptop: laptopNoID,
			store:  service.NewInMemoryLaptopStore(),
			code:   codes.OK,
		},
		{
			name:   "failure_invelid_id",
			laptop: laptopInvalidID,
			store:  service.NewInMemoryLaptopStore(),
			code:   codes.InvalidArgument,
		},
		{

			name:   "failure_duplicate",
			laptop: laptopDuplicateID,
			store:  storeDuplicateID,
			code:   codes.AlreadyExists,
		},
	}

	for _, cases := range testCases {
		tc := cases
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			req := &pb.CreateLaptopRequest{
				Laptop: tc.laptop,
			}
			server := service.NewLaptopServer(tc.store)
			res, err := server.CreateLaptop(context.Background(), req)
			if tc.code == codes.OK {
				assert.Nil(t, err)
				assert.NotNil(t, res)
				assert.NotEmpty(t, res.Id)
				if len(res.Id) > 0 {
					assert.Equal(t, tc.laptop.Id, res.Id)
				}
			} else {
				assert.NotNil(t, err)
				assert.Nil(t, res)
				st, ok := status.FromError(err)
				assert.True(t, ok)
				assert.Equal(t, st.Code(), tc.code)
			}
		})
	}
}
