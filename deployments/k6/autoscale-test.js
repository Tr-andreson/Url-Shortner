import http from 'k6/http';
import { check, sleep } from 'k6';

export const options = {
  stages: [
    { duration: '30s', target: 20 },  // Warm-up: 20 virtual users
    { duration: '2m', target: 150 }, // Heavy load: ramp up to 150 users to push CPU past 70%
    { duration: '3m', target: 150 }, // Sustain load: hold 150 users so HPA scales pods up
    { duration: '30s', target: 0 },  // Cool-down: ramp down to 0 users
  ],
};

export default function () {
  // Hit your production route (or development route depending on where you test)
  // const res = http.get('http://host.docker.internal:8080/');
  const res = http.get('http://localhost:8080/');
  
  check(res, {
    'status is 200': (r) => r.status === 200,
  });

  // Short sleep to maximize request frequency per virtual user
  sleep(0.1);
}
