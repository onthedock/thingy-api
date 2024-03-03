#!/usr/bin/env bash
sudo apt-get install make --yes
grep --quiet --fixed-strings --line-regexp 'source .devcontainer/git-completion.bash' ~/.bashrc || echo 'source .devcontainer/git-completion.bash' >> ~/.bashrc

# Install Terraform
wget -O- https://apt.releases.hashicorp.com/gpg | sudo gpg --dearmor -o /usr/share/keyrings/hashicorp-archive-keyring.gpg
echo "deb [signed-by=/usr/share/keyrings/hashicorp-archive-keyring.gpg] https://apt.releases.hashicorp.com $(lsb_release -cs) main" | sudo tee /etc/apt/sources.list.d/hashicorp.list
sudo apt update && sudo apt install terraform
