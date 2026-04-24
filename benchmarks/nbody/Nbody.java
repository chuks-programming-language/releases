public class Nbody {
    public static void main(String[] args) {
        int N = 5;
        int steps = 500000;

        double[] x = {0.0, 4.841431, 8.343367, 12.894369, 15.379697};
        double[] y = {0.0, -1.160320, 4.124799, -15.201544, -25.632172};
        double[] z = {0.0, -0.103860, -0.403093, -0.223581, 0.179258};

        double[] vx = {0.0, 0.001660, 0.007699, 0.002965, 0.002681};
        double[] vy = {0.0, 0.007699, 0.006286, 0.002378, 0.001628};
        double[] vz = {0.0, -0.000069, -0.000169, -0.000030, -0.000015};

        double[] mass = {39.478418, 0.037694, 0.011287, 0.001724, 0.000203};

        double dt = 0.01;

        for (int step = 0; step < steps; step++) {
            for (int i = 0; i < N; i++) {
                for (int j = i + 1; j < N; j++) {
                    double dx = x[i] - x[j];
                    double dy = y[i] - y[j];
                    double dz = z[i] - z[j];
                    double dist2 = dx * dx + dy * dy + dz * dz;
                    double mag = dt / (dist2 * 100.0 + 0.001);

                    double fx = dx * mag;
                    double fy = dy * mag;
                    double fz = dz * mag;

                    vx[i] -= fx * mass[j];
                    vy[i] -= fy * mass[j];
                    vz[i] -= fz * mass[j];
                    vx[j] += fx * mass[i];
                    vy[j] += fy * mass[i];
                    vz[j] += fz * mass[i];
                }
            }
            for (int i = 0; i < N; i++) {
                x[i] += dt * vx[i];
                y[i] += dt * vy[i];
                z[i] += dt * vz[i];
            }
        }

        double energy = 0.0;
        for (int i = 0; i < N; i++) {
            energy += 0.5 * mass[i] * (vx[i]*vx[i] + vy[i]*vy[i] + vz[i]*vz[i]);
        }
        System.out.println(energy);
    }
}
