pub mod file {
    use std::io::prelude::*;
    use std::io::{Lines, BufReader, Bytes};
    use std::fs::File;
    use std::error::Error;
    use std::env;
    use std::path::Path;

    fn get_args() -> String  {
        let args: Vec<String> = env::args().collect();
        args[1].to_string()
    }

    fn load_file(fileinput: String) -> File {
        // Open file in ro mode
        let file = match File::open(Path::new(&fileinput)) {
            Err(why) => panic!("couldn't open exercise file cause: {}", why.description()),
            Ok(file) => file,
        };
        file
    }

    pub fn read() -> String {
        let path_arg = get_args();
        let mut file_handler = load_file(path_arg);
        // Init string var
        let mut s = String::new();
        match file_handler.read_to_string(&mut s) {
            Err(why) => panic!("couldn't read cause: {}", why.description()),
            Ok(_) => {},
        };
        s
    }

    pub fn read_data() -> Bytes<File> {
        let path_arg = get_args();
        let mut file_handler = load_file(path_arg);
        // Init string var
        file_handler.bytes()
    }

    pub fn read_stream() -> Lines<BufReader<File>> {
        let path_arg = get_args();
        BufReader::new(load_file(path_arg)).lines()
    }
}
