package routers

import (
	"bytes"
	"context"
	"encoding/base64"
	"io"
	"mime"
	"mime/multipart"
	"redSGo/bd"
	"redSGo/models"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

type readSeeker struct {
	io.Reader
}

func (rs *readSeeker) Seek(offset int64, whence int) (int64, error) {
	return 0, nil
}

func UploadImage(ctx context.Context, tipo string, request events.APIGatewayProxyRequest, claim models.Claim) models.RespAPI {
	var r models.RespAPI
	r.Status = 400
	IDUser := claim.ID.Hex()

	var filename string
	var usuario models.User

	bucket := aws.String(ctx.Value(models.Key("bucketName")).(string))

	switch tipo {
	case "A":
		filename = "avatars/" + IDUser + ".jpg"
		usuario.Avatar = filename
	case "B":
		filename = "banners/" + IDUser + ".jpg"
		usuario.Banner = filename
	}

	mediaType, params, err := mime.ParseMediaType(request.Headers["Content-Type"])
	if err != nil {
		r.Status = 500
		r.Message = err.Error()
		return r
	}

	if strings.HasPrefix(mediaType, "mulltipart/") {
		body, err := base64.StdEncoding.DecodeString(request.Body)
		if err != nil {
			r.Status = 500
			r.Message = err.Error()
			return r
		}
		mr := multipart.NewReader(bytes.NewReader(body), params["boundaty"])
		p, err := mr.NextPart()
		if err != nil && err != io.EOF {
			r.Status = 500
			r.Message = err.Error()
			return r
		}

		if err != io.EOF {
			if p.FileName() != "" {
				buff := bytes.NewBuffer(nil)
				if _, err := io.Copy(buff, p); err != nil {
					r.Status = 500
					r.Message = err.Error()
					return r
				}
				sess, err := session.NewSession(&aws.Config{
					Region: aws.String("us-east-1"),
				})
				if err != nil && err != io.EOF {
					r.Status = 500
					r.Message = err.Error()
					return r
				}

				uploader := s3manager.NewUploader(sess)
				_, err = uploader.Upload(&s3manager.UploadInput{
					Bucket: bucket,
					Key:    aws.String(filename),
					Body:   &readSeeker{buff},
				})
				if err != nil && err != io.EOF {
					r.Status = 500
					r.Message = err.Error()
					return r
				}
			}
		}

		status, err := bd.ModificarPerfil(usuario, IDUser)
		if err != nil || !status {
			r.Status = 400
			r.Message = "Error al modificar el avatar/banener del usurio" + err.Error()
			return r
		}

	} else {
		r.Message = "Debe enviar imagen con el 'Content-Type' de tipo 'multipart/"
		return r
	}

	r.Status = 200
	r.Message = "image upload succesfull"
	return r
}
