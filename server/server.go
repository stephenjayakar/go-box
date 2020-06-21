import (
	pb "github.com/stephenjayakar/go-box/gobox"
)

type goBoxServer struct {
	pb.UnimplementedGoBoxServer
}

func (s *goBoxServer) Meow(ctx context.Context) (string, error) {
	return "meow"
}