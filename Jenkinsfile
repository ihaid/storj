node('node') {
  properties([disableConcurrentBuilds()])
  try {
    currentBuild.result = "SUCCESS"

    stage('Checkout') {
      checkout scm

      echo "Current build result: ${currentBuild.result}"
    }

    stage('Run Versions Test') {
      echo "Running migration test"

      env.STORJ_SIM_POSTGRES = 'postgres://postgres@localhost:58723/teststorj?sslmode=disable'

      echo "$STORJ_SIM_POSTGRES"
      sh 'docker run --rm -p 58723:5432 -d --name postgres postgres:9.6'
      sh '''until $(docker logs postgres | grep "database system is ready to accept connections" > /dev/null)
            do printf '.'
            sleep 5
            done
        '''
      sh 'docker exec postgres createdb -U postgres teststorj'
      sh 'git fetch --no-tags --progress -- https://github.com/storj/storj.git +refs/heads/master:refs/remotes/origin/master'
      sh 'git branch -a -v --no-abbrev'
      sh 'docker run -u $UID:$GID --rm -i -v $PWD:$PWD -w $PWD --entrypoint $PWD/scripts/test-sim-versions.sh -e STORJ_SIM_POSTGRES --link postgres:postgres -e CC=gcc storjlabs/golang:1.13.5'
      sh 'docker rm -f postgres'
    }

    stage('Build Binaries') {
      sh 'make binaries'

      stash name: "storagenode-binaries", includes: "release/**/storagenode*.exe"

      echo "Current build result: ${currentBuild.result}"
    }

    stage('Build Windows Installer') {
      node('windows') {
        checkout scm

        unstash "storagenode-binaries"

        bat 'installer\\windows\\build.bat'

        stash name: "storagenode-installer", includes: "release/**/storagenode*.msi"

        echo "Current build result: ${currentBuild.result}"
      }
    }

    stage('Sign Windows Installer') {
      unstash "storagenode-installer"

      sh 'make sign-windows-installer'

      echo "Current build result: ${currentBuild.result}"
    }

    stage('Build Images') {
      sh 'make images'

      echo "Current build result: ${currentBuild.result}"
    }

    stage('Push Images') {
      echo 'Push to Repo'
      sh 'make push-images'
      echo "Current build result: ${currentBuild.result}"
    }

    stage('Upload') {
      sh 'make binaries-upload'
      echo "Current build result: ${currentBuild.result}"
    }

  }
  catch (err) {
    echo "Caught errors! ${err}"
    echo "Setting build result to FAILURE"
    currentBuild.result = "FAILURE"

    /*
    slackSend color: 'danger', message: "@channel ${env.BRANCH_NAME} build failed during stage ${env.STAGE_NAME} ${env.BUILD_URL}"

    mail from: 'builds@storj.io',
      replyTo: 'builds@storj.io',
      to: 'builds@storj.io',
      subject: "storj/storj branch ${env.BRANCH_NAME} build failed",
      body: "Project build log: ${env.BUILD_URL}"
    */

      throw err

  }
  finally {

    stage('Cleanup') {
      sh 'docker rm -f postgres || true'
      sh 'make clean-images'
      deleteDir()
    }

  }
}
