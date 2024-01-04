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
                    withDockerRegistry([ credentialsId: "dockerhubcred", url: "" ]) {
                        dockerImage.push()
                    }
                }
            }
        }
        stage('Deploy') {
            steps {
                echo 'Deploying....'
                sh 'docker system prune --all --force' 
            }
        }
    }
}
