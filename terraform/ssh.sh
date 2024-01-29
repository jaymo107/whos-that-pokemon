#!/bin/bash

instance_ip=$(terraform output -raw aws_public_ip)
ssh -i "keys.pem" ec2-user@"$instance_ip"
