
pipeline {
    agent { docker { image 'golang:1.17.5-alpine' } }
    stages {
        stage('build') {
            steps {
                sh 'go version'
                sh 'echo "Hello World!"'
           //     catchError(buildResult: 'NOT_BUILT', message: 'this is the thing', stageResult: 'FAILURE') {
    // some block
// }
            }
        }
    }
}
