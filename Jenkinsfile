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
