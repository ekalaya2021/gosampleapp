pipeline {
    agent {
         label 'ec2-epus'
    }

    stages {
        stage('Repo pulling') {
            steps {
                checkout scm
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
                withCredentials([usernamePassword(credentialsId: 'GitHubCredentials', passwordVariable: 'GIT_PASSWORD', usernameVariable: 'GIT_USERNAME')]) {
                  sh "rm -rf argocd-test"
                  sh "git clone https://github.com/ekalaya2021/argocd-test.git"
                  sh "cd argocd-test"
                  dir('argocd-test') {
                    sh "sed -i 's/gosampleapp:.*/gosampleapp:${BUILD_NUMBER}/g' deployment.yaml"
                    sh "git config user.email ekalaya2021@gmail.com"
                    sh "git config user.name ekalaya2021"
                    sh "git add ${WORKSPACE}/argocd-test/deployment.yaml"
                    sh "git commit -m 'Update image version to: ${BUILD_NUMBER}'"
                    sh"git push https://${GIT_USERNAME}:${GIT_PASSWORD}@github.com/ekalaya2021/argocd-test.git HEAD:master -f"
                  }
                }
              }
            }        
        stage('Cleaning up') {
            steps {
                echo 'Cleaning up....'
                sh 'docker system prune --all --force' 
            }
        }
    }
}
