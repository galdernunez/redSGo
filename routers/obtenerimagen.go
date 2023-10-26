package routers

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"redSGo/awsgo"
	"redSGo/bd"
	"redSGo/models"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go/aws"
)

func ObtenerImagen(ctx context.Context, tipo string, request events.APIGatewayProxyRequest, claim models.Claim) models.RespAPI {

	var r models.RespAPI
	r.Status = 400
	ID := request.QueryStringParameters["id"]
	if len(ID) < 1 {
		r.Message = "el parametro ID es obligatorio"
		return r
	}

	perfil, err := bd.BuscoPerfil(ID)

	if err != nil {
		r.Status = 500
		r.Message = "usuario no encontrado"
		return r
	}

	var filename string
	switch tipo {
	case "A":
		filename = perfil.Avatar
	case "B":
		filename = "banners/" + perfil.Banner
	}
	fmt.Println("Filename: " + filename)
	svc := s3.NewFromConfig(awsgo.Cfg)

	file, err := downloadFormS3(ctx, svc, filename)
	if err != nil {
		r.Status = 500
		r.Message = "Error descargando archivo de S3: " + err.Error()
		return r
	}

	r.CustomResp = &events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       file.String(),
		Headers: map[string]string{
			"Content-Type":        "application/octet-stream",
			"Content-Disposition": fmt.Sprintf("attachment; filename=\"s \"", filename),
		},
	}
	return r
}

func downloadFormS3(ctx context.Context, svc *s3.Client, filename string) (*bytes.Buffer, error) {
	bucket := ctx.Value(models.Key("bucketName")).(string)
	obj, err := svc.GetObject(ctx, &s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(filename),
	})
	if err != nil {
		return nil, err
	}
	defer obj.Body.Close()
	fmt.Println("bucketname: " + bucket)

	file, err := ioutil.ReadAll(obj.Body)

	if err != nil {
		return nil, err
	}

	buffer := bytes.NewBuffer(file)

	return buffer, nil
}
