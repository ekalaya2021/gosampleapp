pipeline {
    // agent {
    //      label 'ec2-epus'
    // } 
    agent any
    environment {
        GIT_CREDENTIALS = credentials('GitHubCredentials')
    }
    stages {
        stage('Repo pulling') {
            steps {
                checkout scm
            }
        }
        stage('Debug GIT_CREDENTIALS') {
            steps {
                echo "username is $GIT_CREDENTIALS_USR"
                echo "password is $GIT_CREDENTIALS_PSW"
            }
        }
        stage('Build') {
            steps{
                script{
                    dockerImage = docker.build("ekalaya/gosampleapp:$BUILD_NUMBER")                
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
        stage('Update Manifest') {
              steps {
                  sh "rm -rf argo-test"
                  sh "git clone https://$GIT_CREDENTIALS_USR:$GIT_CREDENTIALS_PSW@github.com/ekalaya2021/argo-test.git"
                  sh "cd argo-test"
                  dir('argo-test') {
                    sh "sed -i 's/gosampleapp:.*/gosampleapp:${BUILD_NUMBER}/g' deployment.yaml"
                    sh "git config user.email $GIT_CREDENTIALS_USR@gmail.com"
                    sh "git config user.name $GIT_CREDENTIALS_USR"
                    sh "git add ${WORKSPACE}/argo-test/deployment.yaml"
                    sh "git commit -m 'Update image version to: ${BUILD_NUMBER}'"
                    sh "git push https://$GIT_CREDENTIALS_USR:$GIT_CREDENTIALS_PSW@github.com/$GIT_CREDENTIALS_USR/argo-test.git HEAD:master -f"
                  }
              }
            }        
        stage('Cleaning up') {
            steps {
                echo 'Cleaning up....'
                sh 'docker system prune --all --force' 
            }
        }
        stage('Send notification to mattermost'){
            steps{
                sh """
                curl -i -X POST -H 'Content-Type: application/json' -d '{"text": "Build #${BUILD_NUMBER} has been succeeded :white_check_mark: "}' https://chat.infokes.id/hooks/ojijxc5ncibijrpcfzzpanh4nc
                """
            }
        }
    }
}
