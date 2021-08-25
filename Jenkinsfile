pipeline{
    agent any
    
    tools {
       go "go1.17"
    }

    environment {
        PROJECT_NAME = "dummy_utils"
        GO117MODULE = 'on'
        CGO_ENABLED = 0 
        GOPATH = "${JENKINS_HOME}/jobs/${JOB_NAME}/builds/${BUILD_ID}"
    }   
    
    stages {
        
        stage ('Pre-Build') {
            steps {
                sh "echo Starting Build Process"
            }
        }

        stage ('Testing') {
            steps {
                sh """
                   # set GOPATH (already set)
                   echo \$GOPATH
                   
                   # install go dep and run dep ensure
                   go get -u golang.org/x/lint/golint

                   #go get github.com/golang/dep/cmd/dep
                   #PATH=\$PATH:\$GOPATH/bin
                   #echo \$PATH

                   # delete existing folder in GO source path and copy the latest, run dep ensure
                   #rm -rf \$GOPATH/src/\$PROJECT_NAME/
                   #mkdir \$GOPATH/src/\$PROJECT_NAME/
                   #cp -R . \$GOPATH/src/\$PROJECT_NAME/
                   #cd \$GOPATH/src/\$PROJECT_NAME/
                   #dep ensure

                   #rm -rf \$GOPATH/src/\$PROJECT_NAME/report.json
                   #rm -rf \$GOPATH/src/\$PROJECT_NAME/cover.out
                   go test ./... -json > report.json
                   go test ./... -v -coverprofile=cover.out                      

                   """
            }
        }  
                
        stage('Lint') {
            steps {
                   sh """
                   go install golang.org/x/lint/golint@latest
                   golint ./...
                   """ 
                }
        }
        
        stage('Vet') {
            steps {
                sh "go vet ./..."
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