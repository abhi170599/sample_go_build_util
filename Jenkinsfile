pipeline{
    agent any
    stages{
        stage("Check"){
            steps{
                sh "go version"
                sh "echo \$GOPATH"
                sh "$WORKSPACE"
            }
        }

        stage("Test"){
            steps{
                sh "go test ./utils/" 
                sh "go test ./utils/"
            }
        }

    }
    post{
        always{
            echo "========always========"
        }
        success{
            echo "========pipeline executed successfully ========"
        }
        failure{
            echo "========pipeline execution failed========"
        }
    }
}