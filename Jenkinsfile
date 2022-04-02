pipeline{
    agent any
    environment {
        containerName = 'sample-go'
        scmUrl = 'https://github.com/jutionck/sample-gorm-gin.git'
        branch = 'master'
        imageName = 'sample-go:v1'
    }
    stages{

        stage("Cleaning up"){
            steps{
                echo "Cleaning up..."
                sh 'docker rm -f "${containerName}" || true'
            } 
        }

        stage("Clone source") {
            steps {
                git branch: "${branch}", url: "${scmUrl}"
            }
        }

        stage("Docker build"){
            steps{
                echo "Compiling and building..."
                sh 'docker build -t "${imageName}" .'
            }
        }
    }
}
