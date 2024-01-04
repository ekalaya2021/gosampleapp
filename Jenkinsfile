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
                dockerImage = docker.build("ekalaya/gosampleapp:latest")                
            }
        }
        stage('Publish') {
            steps {
                echo 'Publishing..'
                withDockerRegistry([ credentialsId: "dockerhubkred", url: "https://hub.docker.com" ]) {
                dockerImage.push()
                }                
            }
        }

        stage('Test') {
            steps {
                echo 'Cleanup image..'
                sh 'docker system prune --all --force'
            }
        }
        stage('Deploy') {
            steps {
                echo 'Deploying....'
            }
        }
    }
}
