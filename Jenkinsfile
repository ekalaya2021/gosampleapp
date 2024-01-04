pipeline {
    agent any

    stages {
        stage('Repo pulling') {
            steps {
                checkout scm
            }
        }

        stage('Build') {
            steps {
                echo 'Building..'
                sh 'docker build -t gosampleapp .'
                
            }
        }
        stage('Test') {
            steps {
                echo 'Testing..'
            }
        }
        stage('Deploy') {
            steps {
                echo 'Deploying....'
            }
        }
    }
}
