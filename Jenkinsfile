pipeline {
    agent any

    stages {
        stage('Build') {
            steps {
                echo 'Building..'
                sh   'ls'
                sh   'pwd'
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
		sh    'docker build -f Dockerfile.go .'
            }
        }
    }
}
