# -*- mode: ruby -*-
# vi: set ft=ruby :

# Vagrantfile API/syntax version. Don't touch unless you know what you're doing!
VAGRANTFILE_API_VERSION = "2"

Vagrant.configure(VAGRANTFILE_API_VERSION) do |config|
  config.vm.box = "opscode_ubuntu-13.10_chef-provisionerless.box"
  config.vm.box_url = "http://opscode-vm-bento.s3.amazonaws.com/vagrant/virtualbox/opscode_ubuntu-13.10_chef-provisionerless.box"

  config.vm.hostname = "gojenkinsapi"

  config.vm.network "private_network", ip: "192.168.33.10"

   config.vm.provision :shell, :inline => <<-EOF
     set -e
     sudo apt-get install -y curl git mercurial make binutils bison gcc build-essential

     cd /tmp
     wget -q https://godeb.s3.amazonaws.com/godeb-amd64.tar.gz
     tar zxvf godeb-amd64.tar.gz
     sudo ./godeb install
     rm godeb*

     cd /vagrant
     travis/setup.sh
   EOF
end
