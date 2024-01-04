pipeline {
    agent any

    stages {
        stage('Repo pulling') {
            steps {
                checkout scm
            }
        }

        stage('Build') {
            steps{
                script{
                    dockerImage = docker.build("ekalaya/gosampleapp:latest")                
                }
            }
        }
        stage('Publish') {
            steps{
                script{
                    withDockerRegistry([ credentialsId: "dockerhubcred", url: "registry-1.docker.io" ]) {
                        dockerImage.push()
                    }
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
