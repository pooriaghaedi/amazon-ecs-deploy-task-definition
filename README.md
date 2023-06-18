# ECS Update Task definition

This project encompasses a Go script that allows for the seamless updating of an Amazon ECS task definition using a specified Docker image URI. As a versatile alternative to `aws-actions/configure-aws-credentials`, this script operates independently of any existing task definition files. It optimizes your workflow by automatically selecting the latest task definition. This feature eliminates the need for manual updating and reduces the risk of utilizing outdated task definitions, thereby improving the reliability and efficiency of your deployments.

## Usage



### Example workflow

```yaml
name: My Workflow
on: [push, pull_request]
jobs:
  build:
    runs-on: ubuntu-latest
    steps:
    - uses: actions/checkout@master

    - name: Configure AWS credentials
      uses: aws-actions/configure-aws-credentials@v1
      with:
        aws-access-key-id: ${{ secrets.AWS_ACCESS_KEY_ID }}
        aws-secret-access-key: ${{ secrets.AWS_SECRET_ACCESS_KEY }}
        aws-region: ${{ secrets.AWS_REGION }}

    - name: Update ECS Task definition
      id: task-def
      uses: pooriaghaedi/amazon-ecs-deploy-task-definition@master
      env:
        ECS_TASK_DEFINITION: "YourTaskDefinitionName"
         REGION: ${{ secrets.AWS_REGION }}
        IMAGE_URI: "nginx:latest"
        ECS_CLUSTER: "YourECSClusterName"
        ECS_SERVICE: "YourECSServiceName"        
```

### Inputs

| Input                                             | Description                                        |
|------------------------------------------------------|-----------------------------------------------|
| `ECS_TASK_DEFINITION`  | Your Task definition name    |
| `REGION`  | Your ECS task Region    |
| `IMAGE_URI`  | The Docker image URI that will be used to update your ECS task definition.   |
| `ECS_CLUSTER`  | Name of ECS cluster    |
| `ECS_SERVICE`  | AWS ECS service name   |
