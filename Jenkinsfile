// Jenkinsfile (Declarative Pipeline)
pipeline {
    agent any
    stages {
        stage('build') {
            steps {
                sh 'go mod tidy'
                sh 'cd server'
                sh 'echo server completed successfully...'
                sh 'go build ./...'
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
