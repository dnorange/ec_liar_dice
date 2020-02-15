pipeline {
  agent any
  stages {
    stage('Build') {
      steps {
        sh './deploy/devops/scripts/build-images.sh'
      }
    }
    stage('Test') {
      steps {
        sh './deploy/devops/scripts/test.sh'
      }
    }
    stage('Push Images') {
      steps {
        sh './deploy/devops/scripts/push-images.sh'
      }
    }
    stage ('Verify') {
      steps {
        sh './deploy/devops/scripts/verify.sh'
      }
    }
  }
}
