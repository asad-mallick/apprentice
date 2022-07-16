// Jenkinsfile (Declarative Pipeline)
pipeline {
    agent any
    stages {
        stage('compile') {
            steps {
                sh 'cd server'
                echo 'cd server completed successfully...'
                sh 'go build ./...'
                echo 'project build completed sucessfully...'
            }
        }
    }
}
