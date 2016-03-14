require 'spec_helper'
# This checks if nginx is installed
describe package('nginx') do 
  it {should be_installed}
end

# This checks if service is running
describe service('nginx') do
  it {should be_enabled}
  it {should be_running}
end

# This checks if port 80 is running
describe port(80) do
  it {should be_listening}
end

# This chekcs if nginx confi is there
describe file('/etc/nginx/nginx.conf') do
it {should be_file}
end

