use std::{
    fs::File,
    io::{copy, BufRead, BufReader, Write},
    net::TcpListener,
};

fn main() {
    let server = TcpListener::bind("0.0.0.0:1919").unwrap();
    println!("LISTENING AT {}", server.local_addr().unwrap());

    let (mut stream, addr) = server.accept().unwrap();
    println!("ACCEPTED CONNECTION FROM {}", addr);

    let mut br = BufReader::new(stream.try_clone().unwrap());
    let mut filename = String::new();
    br.read_line(&mut filename).unwrap();
    filename = filename.trim().to_string();
    println!("RECEIVED THE FILENAME {}", filename);

    let mut file = File::create(filename.clone()).unwrap();
    println!("FILE {} CREATED", filename);

    stream.write(b"OK\n").unwrap();
    println!("CONFIRMATION SENT");

    let bytes = copy(&mut stream, &mut file).unwrap();

    println!("RECEIVED {} BYTES", bytes);
    println!("DONE");
}
