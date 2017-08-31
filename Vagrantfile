# -*- mode: ruby -*-
# vi: set ft=ruby :
Vagrant.configure("2") do |config|
  config.vm.box = "ubuntu/xenial64"
  config.vm.synced_folder ".", "/go/src/github.com/sensu/sensu-go"

  config.ssh.export_command_template = 'export LOLENVKEY="LOLVAL"'

  config.vm.provision "shell", inline: <<-SHELL
    apt-get update
    apt-get install -y ruby ruby-dev build-essential rpm rpmlint
    gem install --no-ri --no-rdoc fpm
    gem install --no-ri --no-rdoc packagecloud
    wget https://storage.googleapis.com/golang/go1.8.3.linux-amd64.tar.gz
    tar -C /usr/local -zxvf go1.8.3.linux-amd64.tar.gz
    chown -R ubuntu:ubuntu /usr/local/go
    chown ubuntu:ubuntu /go
    echo 'export GOROOT="/usr/local/go"' >> /home/ubuntu/.bashrc
    echo 'export GOPATH="/go"' >> /home/ubuntu/.bashrc
    echo 'export PATH="/usr/local/go/bin:$PATH"' >> /home/ubuntu/.bashrc
  SHELL
end
