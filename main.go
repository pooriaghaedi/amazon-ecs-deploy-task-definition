package main

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ecs"
)

func main() {

	ECS_TASK_DEFINITION := os.Getenv("ECS_TASK_DEFINITION")
	REGION := os.Getenv("REGION")
	IMAGE_URI := os.Getenv("IMAGE_URI")
	accessKeyID := os.Getenv("AWS_ACCESS_KEY_ID")
	secretAccessKey := os.Getenv("AWS_SECRET_ACCESS_KEY")
	ecsCluster := os.Getenv("ECS_CLUSTER")
	ecsService := os.Getenv("ECS_SERVICE")

	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(REGION),
		Credentials: credentials.NewStaticCredentials(accessKeyID, secretAccessKey, ""),
	})

	svc := ecs.New(sess)

	resp, err := svc.DescribeTaskDefinition(&ecs.DescribeTaskDefinitionInput{
		TaskDefinition: aws.String(ECS_TASK_DEFINITION),
	})

	if err != nil {
		fmt.Println("Error", err)
		os.Exit(1)
	}

	// check if TaskDefinition is nil
	if resp.TaskDefinition == nil {
		fmt.Println("Error: TaskDefinition is nil")
		os.Exit(1)
	}

	// // Updating the image in container definitions
	for i := range resp.TaskDefinition.ContainerDefinitions {
		resp.TaskDefinition.ContainerDefinitions[i].Image = aws.String(fmt.Sprintf(IMAGE_URI))
	}

	newTaskDefinition, err := svc.RegisterTaskDefinition(&ecs.RegisterTaskDefinitionInput{
		ContainerDefinitions:    resp.TaskDefinition.ContainerDefinitions,
		Cpu:                     resp.TaskDefinition.Cpu,
		ExecutionRoleArn:        resp.TaskDefinition.ExecutionRoleArn,
		Family:                  resp.TaskDefinition.Family,
		IpcMode:                 resp.TaskDefinition.IpcMode,
		Memory:                  resp.TaskDefinition.Memory,
		NetworkMode:             resp.TaskDefinition.NetworkMode,
		PlacementConstraints:    resp.TaskDefinition.PlacementConstraints,
		ProxyConfiguration:      resp.TaskDefinition.ProxyConfiguration,
		RequiresCompatibilities: resp.TaskDefinition.RequiresCompatibilities,
		TaskRoleArn:             resp.TaskDefinition.TaskRoleArn,
		Volumes:                 resp.TaskDefinition.Volumes,
		RuntimePlatform:         resp.TaskDefinition.RuntimePlatform,
	})

	if err != nil {
		fmt.Println("Error", err)
		os.Exit(1)
	}

	updateServiceInput := &ecs.UpdateServiceInput{
		Cluster:        aws.String(ecsCluster),
		Service:        aws.String(ecsService),
		TaskDefinition: aws.String(*newTaskDefinition.TaskDefinition.TaskDefinitionArn),
	}

	// Call UpdateService
	_, err = svc.UpdateService(updateServiceInput)

	if err != nil {
		fmt.Println("Error updating service ", err)
		os.Exit(1)
	}

	fmt.Println(*newTaskDefinition.TaskDefinition.TaskDefinitionArn)
	os.Setenv("GITHUB_OUTPUT", "taskARN="+string(*newTaskDefinition.TaskDefinition.TaskDefinitionArn))
}
