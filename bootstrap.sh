adduser -G sudo kacchan
mkdir -p /home/kacchan/.ssh
cp /root/.ssh/authorized_keys /home/kacchan/.ssh/authorized_keys
su -l kacchan passwd
