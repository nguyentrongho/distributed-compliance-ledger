data "aws_ami" "ubuntu" {
    most_recent = true
    owners      = ["099720109477"]

    filter {
        name   = "name"
        values = ["ubuntu-minimal/images/hvm-ssd/ubuntu-focal-20.04-amd64-minimal-*"]
    }

    filter {
        name   = "virtualization-type"
        values = ["hvm"]
    }
}

resource "aws_key_pair" "key_pair" {
    public_key = file(var.ssh_public_key_path)
}

resource "aws_instance" "this_nodes" {
    count = var.nodes_count

    ami           = data.aws_ami.ubuntu.id
    instance_type = "t3.medium"

    subnet_id              = element(module.this_vpc.public_subnets, 0)
    ipv6_address_count     = 1
    
    vpc_security_group_ids = [
        module.this_dev_sg.security_group_id,
        module.this_public_sg.security_group_id
    ]

    key_name   = aws_key_pair.key_pair.id
    monitoring = true

    tags = {
        Name = "Public Sentry Node ${count.index}"
    }

    root_block_device {
        encrypted   = true
        volume_size = 30
    }
}

resource "aws_instance" "this_seed_node" {
    ami           = data.aws_ami.ubuntu.id
    instance_type = "t3.medium"

    subnet_id              = element(module.this_vpc.public_subnets, 0)
    ipv6_address_count     = 1

    vpc_security_group_ids = [
        module.this_dev_sg.security_group_id,
        module.this_seed_sg.security_group_id
    ]

    key_name   = aws_key_pair.key_pair.id
    monitoring = true

    tags = {
        Name = "Public Sentries' Seed Node"
    }

    root_block_device {
        encrypted   = true
        volume_size = 30
    }
}

resource "aws_eip" "this_seed_eip" {
    instance = aws_instance.this_seed_node.id
    vpc      = true

    tags = {
        Name = "Public Sentries' Seed Node Elastic IP"
    }
}