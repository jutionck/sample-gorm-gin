pipeline{
    agent any
    stages{
        stage("Pre Test"){
            steps{
                echo "Pre Test Setup..."
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
