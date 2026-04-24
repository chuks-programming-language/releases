const N = 5;
const steps = 500000;

const x = [0.0, 4.841431, 8.343367, 12.894369, 15.379697];
const y = [0.0, -1.16032, 4.124799, -15.201544, -25.632172];
const z = [0.0, -0.10386, -0.403093, -0.223581, 0.179258];

const vx = [0.0, 0.00166, 0.007699, 0.002965, 0.002681];
const vy = [0.0, 0.007699, 0.006286, 0.002378, 0.001628];
const vz = [0.0, -0.000069, -0.000169, -0.00003, -0.000015];

const mass = [39.478418, 0.037694, 0.011287, 0.001724, 0.000203];

const dt = 0.01;

for (let step = 0; step < steps; step++) {
  for (let i = 0; i < N; i++) {
    for (let j = i + 1; j < N; j++) {
      const dx = x[i] - x[j];
      const dy = y[i] - y[j];
      const dz = z[i] - z[j];
      const dist2 = dx * dx + dy * dy + dz * dz;
      const mag = dt / (dist2 * 100.0 + 0.001);

      const fx = dx * mag;
      const fy = dy * mag;
      const fz = dz * mag;

      vx[i] -= fx * mass[j];
      vy[i] -= fy * mass[j];
      vz[i] -= fz * mass[j];
      vx[j] += fx * mass[i];
      vy[j] += fy * mass[i];
      vz[j] += fz * mass[i];
    }
  }
  for (let i = 0; i < N; i++) {
    x[i] += dt * vx[i];
    y[i] += dt * vy[i];
    z[i] += dt * vz[i];
  }
}

let energy = 0.0;
for (let i = 0; i < N; i++) {
  energy += 0.5 * mass[i] * (vx[i] * vx[i] + vy[i] * vy[i] + vz[i] * vz[i]);
}
console.log(energy);
