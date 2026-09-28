import http from 'k6/http';
import { check, sleep } from 'k6';

export const options = {
  stages: [
    { duration: '30s', target: 20 }, // Ramp up to 20 virtual users over 30 seconds
    { duration: '1m', target: 20 },  // Stay at 20 users for 1 minute
    { duration: '10s', target: 0 },  // Ramp down to 0 users
  ],
};

export default function () {
  // Test root route
  const resRoot = http.get('http://localhost:8080/');
  check(resRoot, {
    'root status is 200': (r) => r.status === 200,
    'root body is correct': (r) => r.body.includes('Hello, World from Go and Docker kubernetes !'),
  });

  // Test health route
  const resHealth = http.get('http://localhost:8080/health');
  check(resHealth, {
    'health status is 200': (r) => r.status === 200,
    'health body is correct': (r) => r.body.includes('Health Route workingg!'),
  });

  sleep(1);
}
