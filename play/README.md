# Ansible Playbook For Contently on AWS

The Ansible playbook is used for design and deploying servers for Contently applications.

### Installation
Setup on OSX:

1. ```brew install ansible```
2. ``brew install python``
3. ``pip install -U pip setuptools``
4. ``pip install boto``
5. ``Setup boto Env``
    - ex: vim ~/.boto
    - [credentials]
    - aws_access_key = <your access_key >
    - aws_secret_access_key = <your secret_key >
2. ```git clone git@github.com:contently/playbook.git```

Run a playbook to create or update some servers, fix state, run tasks etc.
 - ex: Nat playbook
ansible-playbook site.yml -i local.yml -vvvv 
```
To run playbook with vault 
ansible-playbook xxxxx.yml -i local.yml --ask-vault -vvv
provide the vault password
```

Edit vars encrypted in vault.
change xxxx.yml to common.yml for instance.
Vault will prompt for a password to decrypt.  Ask someone from ops for it.
```
ansible-vault edit vars/xxxx.yml
```

Create vars encrypted in vault.
change xxxx.yml to common.yml for instance.
Vault will prompt for a password to decrypt.  Ask someone from ops for it.
```
ansible-vault create vars/xxxx.yml
```
To decrypt a vault file 
ansible-vault decrypt vars/xxxx.yml
See Ansible docs for more info: http://docs.ansible.com/ansible/

### DevOps checklist
When building new servers, they have support to check off the following items to ensure reliability and scalability. Copy and paste this evolving checklist in to any new server issues.

- [ ] Highly available
    - [ ] Auto-scaling/fault-tolerant/self-healing
    - [ ] Cross zone —> Cross region
- [ ] Deployable
    - [ ] Instructions, etc.
- [ ] Testing
    - [ ] Build local
    - [ ] QA —> Staging —> Production
- [ ] Security
    - [ ] Security groups (ports, etc)
    - [ ] User access/VPN
- [ ] Monitoring
    - [ ] Alerts
- [ ] Tagging
    - [ ] App, Component, Env
