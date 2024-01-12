pipeline {
    // agent {
    //      label 'ec2-fleet'
    // } 
    agent any 
    environment {
        GIT_CREDENTIALS = credentials('GitHubCredentials')
        AWS_ACCOUNT_ID="421567267553"
        AWS_DEFAULT_REGION="ap-southeast-3" 
        IMAGE_REPO_NAME=" infokes-ecr"        
        REPOSITORY_URI = "${AWS_ACCOUNT_ID}.dkr.ecr.${AWS_DEFAULT_REGION}.amazonaws.com/${IMAGE_REPO_NAME}"
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
                    dockerImage = docker.build("infokes-ecr/gosampleapp:$BUILD_NUMBER")                
                }
            }
        }
        stage('Publish') {
            steps{
                script{
                    withDockerRegistry([credentialsId:"infokes-admin",url:"https://421567267553.dkr.ecr.ap-southeast-3.amazonaws.com"]){
                    // withDockerRegistry([ credentialsId: "dockerhubcred", url: "" ]) {
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
    post{
        always{
            script{
                cleanWs()
            }
        }
    }
}
