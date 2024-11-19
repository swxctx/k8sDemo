pipeline {
    agent any
    stages {
        stage('Checkout') {
            steps {
                // 从 Git 仓库中拉取代码
                checkout scm
            }
        }
        stage('Build and Deploy') {
            steps {
                // 执行部署脚本
                sh './build/deploy.sh'
            }
        }
    }
    post {
        always {
            // 构建后的操作，例如清理工作区或者发送通知
            echo 'Build finished'
        }
    }
}
