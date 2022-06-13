// Jenkinsfile (Declarative Pipeline)
pipeline {
    agent { docker { image 'golang:1.17.5-alpine' } }
    stages {
        stage('go lang') {
            steps {
                sh 'go version'
                         }
                  }
        
        stage('Build stage') {
             steps {
                echo 'I am Building'
                             }
                   }
         stage('Deploy Stage') {
             steps {
                echo 'I am Deploying'
                               }
                   }  
        stage('Test Stage') {
             steps {
                echo 'I am Testing'
                               }
                   }  
        stage('Release Stage') {
             steps {
                echo 'I am Releasing'
                               }
                   }  
    }
}
