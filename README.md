# Jenkins_docker_go

Jenkins - docker - gitAction - Prometheus를 활용한 모니터링 및 CI/CD

logFile은 일부로 제거

# DB

AWS RDS 사용 중

- mysql

Endpoint : golang-sns-v2.cynkkrrli9o4.ap-northeast-2.rds.amazonaws.com

- 개인 저장 용

# mockgen

mockgen -destination=./mock/sns_mock/snsMockRunner.go -package=mocks -source=./mock/sns_mock/IMock_Sns.go
