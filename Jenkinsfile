// Jenkinsfile (Declarative Pipeline)
pipeline {
    agent any
    stages {
        stage('compile') {
            steps {
                sh 'cd server'
                sh 'echo server completed successfully...'
                sh 'go build ./...'
                sh 'echo project build completed sucessfully...'
            }
        }
    }
}
