PROJECT_NAME=$1
SHA1=$2 # Nab the SHA1 of the desired build from a command-line argument
BRANCH=$3
EB_BUCKET=elasticbeanstalk-us-west-2-194925301021

# Create new Elastic Beanstalk version
echo $BRANCH
DOCKERRUN_FILE=$SHA1-Dockerrun.aws.json
sed "s/<TAG>/$BRANCH/" < Dockerrun.aws.json > $DOCKERRUN_FILE
aws configure set default.region us-west-2
aws configure set region us-west-2

if [ "$BRANCH" == "production" ]; then

	aws s3 cp $DOCKERRUN_FILE s3://$EB_BUCKET/$DOCKERRUN_FILE # copy the file that beanstalk needs to spin up the containers into beanstalk's s3 bucket
	# define new application version
	aws elasticbeanstalk create-application-version --application-name $PROJECT_NAME --version-label $SHA1 --source-bundle S3Bucket=$EB_BUCKET,S3Key=$DOCKERRUN_FILE
	aws elasticbeanstalk update-environment --environment-name $PROJECT_NAME-env --version-label $SHA1 # update environment to new version

elif [ "$BRANCH" == "stage" ]; then # do the same thing, to the stage environment

	aws s3 cp $DOCKERRUN_FILE s3://$EB_BUCKET/$DOCKERRUN_FILE # copy the file that beanstalk needs to spin up the containers into beanstalk's s3 bucket
	aws elasticbeanstalk create-application-version --application-name $PROJECT_NAME-stage --version-label $SHA1 --source-bundle S3Bucket=$EB_BUCKET,S3Key=$DOCKERRUN_FILE
	aws elasticbeanstalk update-environment --environment-name $PROJECT_NAME-stg --version-label $SHA1
fi
