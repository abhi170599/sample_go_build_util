pipeline{
    agent any

    tools {
        go 'go1.17'
    }
    environment {
        GO117MODULE = 'on'
        CGO_ENABLED = 0 
        GOPATH = "${JENKINS_HOME}/jobs/${JOB_NAME}/builds/${BUILD_ID}"
    }   
    
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