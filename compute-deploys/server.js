import express from 'express';
import { exec } from 'child_process';
import axios from 'axios';
import dotenv from 'dotenv';

dotenv.config({ override: true }); // Will override existing shell variables

const app = express();
const port = 3001;

app.use(express.json());

const DD_API_KEY = process.env.DD_API_KEY;
const DD_APP_KEY = process.env.DD_APP_KEY;

app.get('/api/clusters', (req, res) => {
  console.log('Fetching clusters...');
  exec("kubectl get ClusterInfo -A --context pod998 -ojson | jq '.items[] | select (.metadata.name!=\"current\") | .metadata.name'", (error, stdout, stderr) => {
    if (error) {
      console.error(`exec error: ${error}`);
      return res.status(500).send('Error fetching clusters');
    }
    const clusters = stdout.trim().split('\n').map(cluster => ({ value: cluster.replace(/"/g, ''), label: cluster.replace(/"/g, '') }));
    res.json(clusters);
  });
});

app.post('/api/search', async (req, res) => {
  console.log('Received /api/search request with body:', req.body);
  const { clusters, services, startDate, endDate } = req.body;

  if (!clusters || clusters.length === 0 || !services || services.length === 0 || !startDate || !endDate) {
    return res.status(400).json({ error: 'Missing required filter parameters. Please select at least one cluster and one service.' });
  }

  // Use the first selected cluster and service
  const clusterName = clusters[0].value;
  const serviceName = services[0].name;

  // Construct the specific query
  const query = `@requestObject.metadata.annotations.moniker.spinnaker.io/cluster:${clusterName} service:k8s_audit_logs @objectRef.namespace:kube-system @objectRef.resource:daemonsets @objectRef.name:${serviceName} @user.username:"system:serviceaccount:spinnaker:spinnaker-remote" @verb:update`;

  // Function to format date to ISO string with +10:00 offset
  const formatToAEST = (dateString) => {
    const date = new Date(dateString);
    // Manually construct the date string with the desired offset
    // This is a simplified approach. For robust timezone handling, a library like date-fns-tz would be better.
    const year = date.getFullYear();
    const month = String(date.getMonth() + 1).padStart(2, '0');
    const day = String(date.getDate()).padStart(2, '0');
    const hours = String(date.getHours()).padStart(2, '0');
    const minutes = String(date.getMinutes()).padStart(2, '0');
    const seconds = String(date.getSeconds()).padStart(2, '0');
    return `${year}-${month}-${day}T${hours}:${minutes}:${seconds}+10:00`;
  };

  const fromDate = formatToAEST(startDate);
  const toDate = formatToAEST(endDate);

  const payload = {
    filter: {
      query: query,
      from: fromDate,
      to: toDate,
    },
    page: {
      limit: 1000,
    },
  };

  try {
    console.log('Sending request to Datadog v2 API with payload:', JSON.stringify(payload, null, 2));
    const response = await axios.post('https://api.datadoghq.com/api/v2/logs/events/search', payload, {
      headers: {
        'DD-API-KEY': DD_API_KEY,
        'DD-APPLICATION-KEY': DD_APP_KEY,
        'Content-Type': 'application/json'
      }
    });
    console.log('Datadog API response status:', response.status);
    res.json(response.data);
  } catch (error) {
    if (error.response) {
      // The request was made and the server responded with a status code
      // that falls out of the range of 2xx
      console.error('Datadog API Error Status:', error.response.status);
      console.error('Datadog API Error Data:', error.response.data);
      res.status(error.response.status).send(error.response.data);
    } else if (error.request) {
      // The request was made but no response was received
      console.error('Datadog API No Response:', error.request);
      res.status(500).send('No response from Datadog API');
    } else {
      // Something happened in setting up the request that triggered an Error
      console.error('Error setting up Datadog request:', error.message);
      res.status(500).send('Error fetching data from Datadog');
    }
  }
});

app.listen(port, () => {
  console.log(`Server listening at http://localhost:${port}`);
});


process.on('exit', (code) => {
  console.log('Process exiting with code:', code);
});


console.log('Server process is alive')

// Keep the process alive to prevent it from exiting prematurely
setInterval(() => {}, 1 << 30);
