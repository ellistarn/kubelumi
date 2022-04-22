package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/autoscaling"
	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/ec2"
	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/ssm"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(cluter)
}

func cluter(ctx *pulumi.Context) error {

	// Networking
	vpc, err := ec2.NewVpc(ctx, "main", &ec2.VpcArgs{
		CidrBlock: pulumi.String("10.0.0.0/16"),
	})
	if err != nil {
		return err
	}

	var subnets []pulumi.StringInput

	for zone, cidr := range map[string]string{
		"us-west-2a": "10.0.1.0/24",
		"us-west-2b": "10.0.2.0/24",
		"us-west-2c": "10.0.3.0/24",
	} {
		subnet, err := ec2.NewSubnet(ctx, zone, &ec2.SubnetArgs{
			CidrBlock:        pulumi.String(cidr),
			VpcId:            vpc.ID(),
			AvailabilityZone: pulumi.String(zone),
		})

		subnets = append(subnets, subnet.ID().ToStringOutput())
		if err != nil {
			return err
		}
	}

	if err := cpi(ctx, pulumi.StringArray(subnets)); err != nil {
		return err
	}

	if err := etcd(ctx, pulumi.StringArray(subnets)); err != nil {
		return err
	}
	return nil
}

func cpi(ctx *pulumi.Context, subnets pulumi.StringArray) error {
	parameter, err := ssm.LookupParameter(ctx, &ssm.LookupParameterArgs{
		Name: "/aws/service/eks/optimized-ami/1.21/amazon-linux-2/recommended/image_id",
	}, nil)
	if err != nil {
		return err
	}
	launchTemplate, err := ec2.NewLaunchTemplate(ctx, "cpi-lt", &ec2.LaunchTemplateArgs{
		ImageId:      pulumi.String(parameter.Value),
		InstanceType: pulumi.String("r5.large"),
	})
	if err != nil {
		return err
	}
	_, err = autoscaling.NewGroup(ctx, "cpi", &autoscaling.GroupArgs{
		VpcZoneIdentifiers: subnets,
		DesiredCapacity:    pulumi.Int(3),
		MaxSize:            pulumi.Int(3),
		MinSize:            pulumi.Int(1),
		LaunchTemplate: &autoscaling.GroupLaunchTemplateArgs{
			Id:      launchTemplate.ID(),
			Version: pulumi.String("$Latest"),
		},
	})
	if err != nil {
		return err
	}
	return nil
}

func etcd(ctx *pulumi.Context, subnets pulumi.StringArray) error {
	parameter, err := ssm.LookupParameter(ctx, &ssm.LookupParameterArgs{
		Name: "/aws/service/eks/optimized-ami/1.21/amazon-linux-2/recommended/image_id",
	}, nil)
	if err != nil {
		return err
	}
	launchTemplate, err := ec2.NewLaunchTemplate(ctx, "etcd-lt", &ec2.LaunchTemplateArgs{
		ImageId:      pulumi.String(parameter.Value),
		InstanceType: pulumi.String("r5.large"),
	})
	if err != nil {
		return err
	}
	_, err = autoscaling.NewGroup(ctx, "etcd", &autoscaling.GroupArgs{
		VpcZoneIdentifiers: subnets,
		DesiredCapacity:    pulumi.Int(3),
		MaxSize:            pulumi.Int(3),
		MinSize:            pulumi.Int(1),
		LaunchTemplate: &autoscaling.GroupLaunchTemplateArgs{
			Id:      launchTemplate.ID(),
			Version: pulumi.String("$Latest"),
		},
	})
	if err != nil {
		return err
	}
	return nil
}
