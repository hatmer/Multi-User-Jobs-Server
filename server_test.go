import "testing"

func StartTest(t *testing.T) {
    s := server{}

    // set up test cases
    tests := []struct{
	op  string
        job string
    } {
        {
	    op: "start",
            job: "not a job",
        },
        {
	    op: "stop",
            jobID: "invalidID",
        },
	{   op: "status",
	    jobID: "",
        },
	{
            op: "stream",
	    jobID: "",
    },
            
	    
    }

    for _, tt := range tests {
        req := &pb.HelloRequest{Name: tt.name}
        resp, err := s.SayHello(context.Background(), req)
        if err != nil {
            t.Errorf("HelloTest(%v) got unexpected error")
        }
        if resp.Message != tt.want {
            t.Errorf("HelloText(%v)=%v, wanted %v", tt.name, resp.Message, tt.want)
        }
    }
}
