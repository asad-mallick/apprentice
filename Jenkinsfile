// Jenkinsfile (Declarative Pipeline)
pipeline {
    agent any
    tools {
        go 'go-lang'
    }
    stages {
        stage('build') {
            steps {
                sh 'go mod tidy'
                sh "pwd"
                    dir('server') {
                      sh "pwd"
                    }
            //  sh "pwd"
                sh 'echo server completed successfully...'
                sh 'go build -o server/server.go'
            //  sh 'go build ./...'
                sh 'echo project build completed sucessfully...'
            }
        }
        stage ('execute') {
            steps {
                sh './server'
            }
        }
    }
}
