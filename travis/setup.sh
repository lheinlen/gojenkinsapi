#!/bin/bash

wget -q -O - http://pkg.jenkins-ci.org/debian/jenkins-ci.org.key | sudo apt-key add -
sudo sh -c 'echo deb http://pkg.jenkins-ci.org/debian binary/ > /etc/apt/sources.list.d/jenkins.list'
sudo apt-get update -qq
sudo apt-get install -qq jenkins

sudo service jenkins stop
sudo cp -f travis/jenkins_config.xml /var/lib/jenkins/config.xml
sudo mkdir -p /var/lib/jenkins/users/jenkinstest
sudo cp -f travis/user_config.xml /var/lib/jenkins/users/jenkinstest/config.xml
sudo service jenkins start

while ! nc -vz localhost 8080; do sleep 1; done

sleep 20 # Sleep a bit longer because even when jenkins is up, it's not up
