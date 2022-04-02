pipeline{
    agent any
    stages{
        stage("Pre Test"){
            steps{
                echo "Installing dependencies..."
                sh 'go version'
                sh 'go get -u golang.org/x/lint/golint'
            } 
        }

        stage("build"){
            steps{
                echo "Compiling and building..."
                sh 'go build'
            }
            
        }
    }
}
