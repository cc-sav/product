echo "Downloading package lists..."
apt-get update -qq

echo "Installing packages..."
apt-get install -y golang-go mono-runtime mono-devel mono-mcs

chmod +x /vagrant/scripts/*

echo "Compiling source..."
/vagrant/scripts/build.sh