# Kubelumi
Pulumi demo to stand up a kube cluster.

## Demo
```
pulumi up --stack $USER --yes
pulumi destroy --stack $USER --yes && pulumi stack rm $USER
```

## Output
```
pulumi up --stack $USER --yes
The stack 'etarn' does not exist.

If you would like to create this stack now, please press <ENTER>, otherwise press ^C:
Created stack 'etarn'
Enter your passphrase to protect config/secrets:
Re-enter your passphrase to confirm:
Previewing update (etarn):
     Type                       Name            Plan
 +   pulumi:pulumi:Stack        kubelumi-etarn  create
 +   ├─ aws:ec2:Vpc             main            create
 +   ├─ aws:ec2:Subnet          us-west-2b      create
 +   ├─ aws:ec2:Subnet          us-west-2c      create
 +   ├─ aws:ec2:Subnet          us-west-2a      create
 +   ├─ aws:ec2:LaunchTemplate  cpi-lt          create
 +   ├─ aws:autoscaling:Group   cpi             create
 +   ├─ aws:ec2:LaunchTemplate  etcd-lt         create
 +   └─ aws:autoscaling:Group   etcd            create

Resources:
    + 9 to create

Updating (etarn):
     Type                       Name            Status
 +   pulumi:pulumi:Stack        kubelumi-etarn  created
 +   ├─ aws:ec2:Vpc             main            created
 +   ├─ aws:ec2:LaunchTemplate  cpi-lt          created
 +   ├─ aws:ec2:LaunchTemplate  etcd-lt         created
 +   ├─ aws:ec2:Subnet          us-west-2a      created
 +   ├─ aws:ec2:Subnet          us-west-2b      created
 +   ├─ aws:ec2:Subnet          us-west-2c      created
 +   ├─ aws:autoscaling:Group   cpi             created
 +   └─ aws:autoscaling:Group   etcd            created

Resources:
    + 9 created

Duration: 35s

Time: 0h:00m:44s
```
