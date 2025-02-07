use core::f64;

fn main() {
    const TRAIN: &[[f64; 2]] = &[[0.0, 0.0], [1.0, 2.0], [2.0, 4.0], [3.0, 6.0], [4.0, 8.0]];
    const LEARNING_RATE: f64 = 0.1;

    let mut weight: f64 = 10.0;

    for epoch in 1..10 {
        let mut loss: f64 = 0.0;
        for input in TRAIN {
            let x = input[0];
            let y = input[1];
            let yh = x * weight;
            loss = (y - yh).powi(2);
            let gradient = 2.0 * x * (yh - y);
            weight -= LEARNING_RATE * gradient;
        }
        println!("Epoch {epoch} Weight {weight} Loss {loss}");
    }
}
