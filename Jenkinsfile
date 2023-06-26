pipeline {
  agent any
  stages {
    stage('Stage1') {
      steps {
        sh 'echo "Stage 1 - Step 1"'
        sh 'echo "Stage 1 - step 2"'
      }
    }

    stage('Stage2') {
      parallel {
        stage('Stage2') {
          steps {
            sh 'echo "Stage 2 - Step 1"'
          }
        }

        stage('Stage3') {
          steps {
            sh 'echo "Stage 3 - Step 1"'
          }
        }

      }
    }

  }
}