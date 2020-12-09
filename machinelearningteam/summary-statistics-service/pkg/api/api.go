package api

import (
	"context"
	"fmt"

	api "github.com/e-conomic/hiring-assignments/machinelearningteam/summary-statistics-service/proto"
)

// Server is a server implementing the proto API
type Server struct {
	api.UnimplementedDocumentSummarizerServer
}

// SummarizeDocument echoes the document provided in the request
func (s *Server) SummarizeDocument(
	ctx context.Context,
	req *api.SummarizeDocumentRequest,
) (*api.SummarizeDocumentReply, error) {
	// Echo
	fmt.Println("Received document...")
	return &api.SummarizeDocumentReply{
		Content: req.Document.GetContent(),
	}, nil
}
