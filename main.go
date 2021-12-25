package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"

	persian "github.com/mavihq/persian"
	ptime "github.com/yaa110/go-persian-calendar"
)

func HandleByEvent(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	const layout = "2006-01-02"
	p := strings.Split(req.Path, "/")
	paramter := strings.Trim(p[len(p)-1], "/")
	if len(paramter) == 8 {
		paramter = "20" + paramter // Change YY-mm-dd -> 20YY-mm-dd
	}
	paramter = strings.ReplaceAll(paramter, ".", "-") // change YYYY.mm.dd -> YYYY-mm-dd
	gdate, err := time.Parse(layout, paramter)

	if err != nil {
		return events.APIGatewayProxyResponse{
			Headers: map[string]string{
				"content-type": "text/plain;charset=UTF8",
			},
			StatusCode: 400,
			Body:       fmt.Sprintf(`usagee : https://API_THE_URL/YYYY-mm-dd  Err:%v`, err),
		}, err
	}

	jdate := ptime.New(gdate)

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Headers: map[string]string{
			"Content-Type": "text/plain;charset=UTF8",
		},
		Body: persian.ToPersianDigits(jdate.Format("E dd MMM yyyy")),
	}, nil
}

func main() {
	lambda.Start(HandleByEvent)
}
