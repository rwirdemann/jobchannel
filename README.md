# Job channel Mock

## Deployment

```
make deploy
```

## Endpoints

### POST /jobs
```
curl -v -X POST \
  https://jkyszaoly5.execute-api.eu-central-1.amazonaws.com/dev/jobs \
  -H 'Content-Type: application/xml' \
  -d '<job>
    <id>1234</id>
    <title>Kranfahrer</title>
    <desciption>Im Hafen Hamburg</desciption>
</job>'

# HTTP Status Codes
200: body contains valid xml
400: body contains invalid xml

```

## AWS CLI Hints
```
aws lambda list-function
aws lambda get-function-configuration --function-name recruiting-dev-jobs
```

```
saw watch /aws/lambda/recruiting-dev-jobs
```