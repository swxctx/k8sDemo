# 添加 Docker 和 kubectl 到 PATH
export PATH=$PATH:/usr/local/bin

# docker 镜像打包
docker build -f build/Dockerfile -t mydemo:latest .

# k8s 部署
kubectl apply -f build/deployment.yaml
kubectl get pod