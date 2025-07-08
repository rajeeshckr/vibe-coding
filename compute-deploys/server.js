import express from 'express';
import { exec } from 'child_process';

const app = express();
const port = 3001;

app.get('/api/clusters', (req, res) => {
  exec("kubectl get ClusterInfo -A --context pod998 -ojson | jq '.items[] | select (.metadata.name!=\"current\") | .metadata.name'", (error, stdout, stderr) => {
    if (error) {
      console.error(`exec error: ${error}`);
      return res.status(500).send('Error fetching clusters');
    }
    const clusters = stdout.trim().split('\n').map(cluster => ({ value: cluster.replace(/"/g, ''), label: cluster.replace(/"/g, '') }));
    res.json(clusters);
  });
});

app.listen(port, () => {
  console.log(`Server listening at http://localhost:${port}`);
});
