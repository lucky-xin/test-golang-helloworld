#!/bin/bash

# 默认参数
SONAR_HOST_URL=""
SONAR_LOGIN=""

# 解析命令行参数
while [[ $# -gt 0 ]]; do
  case $1 in
    --host-url)
      SONAR_HOST_URL="$2"
      shift 2
      ;;
    --login)
      SONAR_LOGIN="$2"
      shift 2
      ;;
    --help)
      echo "用法: $0 [选项]"
      echo "选项:"
      echo "  --host-url URL    SonarQube 服务器地址 (默认: http://8.145.35.103:9000)"
      echo "  --login TOKEN     SonarQube 登录令牌"
      echo "  --help            显示此帮助信息"
      exit 0
      ;;
    *)
      echo "未知参数: $1"
      echo "使用 --help 查看帮助信息"
      exit 1
      ;;
  esac
done

echo "=== SonarQube 扫描开始 ==="
echo "服务器地址: $SONAR_HOST_URL"
echo "登录令牌: ${SONAR_LOGIN:0:10}..."

# 运行SonarQube扫描
docker run --rm -u root:root \
  -v ./:/usr/src \
  -v ./.sonar:/root/.sonar \
  -w /usr/src \
  --name sonar-scanner-cli \
  xin8/devops/sonar-scanner-cli:latest \
  sonar-scanner \
  -Dsonar.host.url="$SONAR_HOST_URL" \
  -Dsonar.login="$SONAR_LOGIN" \
  -Dsonar.projectKey=test-golang-helloworld \
  -Dsonar.projectName=test-golang-helloworld \
  -Dsonar.projectVersion=1.5.2 \
  -Dsonar.sourceEncoding=UTF-8 \
  -Dsonar.projectBaseDir=. \
  -Dsonar.sources=encryption,config \
  -Dsonar.tests=test \
  -Dsonar.language=golang \
  -Dsonar.coverage.exclusions=**/t/**,**/tests/**,**/test/** \
  -Dsonar.exclusions=**/*.pb.go,**/vendor/**,**/node_modules/**,**/*.pb.go,**/testdata/** \
  -Dsonar.test.inclusions=**/*_test.go \
  -Dsonar.coverageReportPaths=reports/sonar-coverage.xml \
  -Dsonar.go.tests.reportPaths=reports/test-suite-report-json.txt \
  -Dsonar.verbose=true
