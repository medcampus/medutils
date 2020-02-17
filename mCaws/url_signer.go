package mCaws

import (
	"github.com/aws/aws-sdk-go/service/cloudfront/sign"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"time"
)

func CloudfrontSigner(url string, signer *sign.URLSigner, timeToValid time.Time) (string, error) {
	policy := &sign.Policy{
		Statements: []sign.Statement{
			{
				Resource: url,
				Condition: sign.Condition{
					DateLessThan: &sign.AWSEpochTime{
						Time: timeToValid,
					},
				},
			},
		},
	}

	signedUrl, err := signer.SignWithPolicy(url, policy)
	if err != nil {
		return "", status.Errorf(codes.Internal, "error signing url %s, error %v", url, err)
	}

	return signedUrl, nil
}
