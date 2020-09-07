#!/usr/bin/env node
import 'source-map-support/register';
import * as cdk from '@aws-cdk/core';
import { ResourcesStack } from '../lib/resources-stack';

const app = new cdk.App();
new ResourcesStack(app, 'ResourcesStack');
