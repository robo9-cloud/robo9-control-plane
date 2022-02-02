
Robo 9 - Serverless Cloud Infrastructure Pipeline
=================================================


Introduction
------------

The Robo9 project is dedicated to making deployments of infrastructure workflows on the cloud simple, portable and scalable. Our goal is not to recreate other services, but to provide a straightforward way to deploy best-of-breed open-source systems using infrastructure as code to the cloud(AWS, GCP, Azure and more). Anywhere you are running docker containers, you should be able to run Robo9.

Getting started with Robo9
--------------------------
Read the architecture overview for an introduction to the architecture of Robo9 and to see how you can use Robo9 to manage your Infrastructure as Code Workflows.

Follow *Installing Robo9* to set up your environment and install Robo9.

You can sign up for Robo9-Cloud for a managed solution which will scale across the cloud. However, if you looking to experiment with this tool you should be able to deploy this to your internal cloud environment if you prefer self hosted solutions.


What Is Robo9?
-----------------
Robo9 is a tool that will run your infrastructure as code on AWS, GCP and (coming soon) Azure.

We all know there are tools like terraform, ansible, chef, packer which solve the problem of configuring infrastructure. Where it becomes challenging
is when you scale this to multiple team member then to multiple teams. It can be complex to to set up and configure.
Also managing statefile and having fine grained control over approving the infrastructure as code to be deployed is currently not available in open source.

To use Robo9, the basic workflow for selfhosted robo9 is:

* Download and run the Robo9 CLI binary.
* Customize the resulting configuration files.
* Run the specified script to deploy your containers to your specific environment.

The workflow for Managed Robo9:

* You can sign up for robo9-developer account by visiting <TBD_URL>/sign-up
* You can then create a workflow by choosing your preferred tool (Terraform, Packer, Ansible, Bash) and your preferred cloud.
* Secrets are encypted in Google Secret Manager and only you will have access to these secrets.
* You can then trigger your workflow to run on any of the major clouds (AWS, GCP, Azure). See detailed docs on managed Robo9-Cloud.

The Robo9 Mission
-----------------
Our goal is to make scaling infrastructure pipelines and deploying them to production as simple as possible, by letting serverless functions do what it’s great at:

* Easy, repeatable, portable deployments on a diverse infrastructure (for example, experimenting on a laptop, then moving to an on-premises cluster or to the cloud)
* Deploying and managing  Terraform
* Building Cloud Images using Packer
* Scaling based On Demand
* Running Infrastructure pipeline on demand
* Auditing and Approval of Pipeline by Administrators

Ultimately, we want to provide a way to run your infrastructure on demand or even schedule it to run at a particular time. Be able to see who ran what and when in case of outages or misconfigurations.


Roadmaps
----------
To see what’s coming up in future versions of Robo9, refer to the Robo9 roadmap.

The following components also have roadmaps:

- Robo9-CLI
- Robo9-UI
- Robo9-Serverless

Getting involved
----------------
There are many ways to contribute to Robo9, and we welcome contributions!

Read the contributor’s guide to get started on the code, and learn about the community on the community page.


.. toctree::
    :caption: Table of Contents
    :name: mastertoc

    gettingstarted
