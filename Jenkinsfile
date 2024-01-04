pipeline {
    agent any

    stages {
        stage('Repo pulling') {
            steps {
                checkout scm
            }
        }

        stage('Build') {
            dockerImage = docker.build("ekalaya/gosampleapp:latest")                
        }
        stage('Publish') {
                withDockerRegistry([ credentialsId: "dockerhubkred", url: "https://hub.docker.com" ]) {
                dockerImage.push()
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
