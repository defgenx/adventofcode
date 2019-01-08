mod adventofcode {
    use std::io;
    pub mod file {
        use std::io::BufRead;
        pub fn read_input() {
            let stdin = super::io::stdin();
            let mut handle = stdin.lock();
            for line in handle.lines() {
                println!("{}", line.unwrap());
            }
        }
    }
}