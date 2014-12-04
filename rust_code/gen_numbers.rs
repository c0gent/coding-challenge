#![allow(experimental)]
extern crate debug;

use std::io;
use std::rand;
use std::rand::distributions::{IndependentSample, Range};
use std::collections;

fn main() {
	let mut files: collections::HashMap<&str, io::IoResult<io::File>> = collections::HashMap::new();

	files.insert("Paris", io::File::create(&Path::new("Paris.txt")));
	files.insert("London", io::File::create(&Path::new("London.txt")));
	files.insert("Berlin", io::File::create(&Path::new("Berlin.txt")));

	//let london_file = files.get(&"London");

	//let mut london_file = io::File::create(&Path::new("London.txt"));

	let rng_range = Range::new(-10000i,10000i);
	let mut rng = rand::task_rng();

	let numbers_count = 2160u; // 2160u;
	
	for (name, file) in files.mut_iter() {
		let mut i = 0u;
		for _ in range(0u, numbers_count) {

			let num = rng_range.ind_sample(&mut rng);
			let mut num_str = num.to_string();

			num_str.push_str("\n");

			let num_u8: &[u8] = num_str.as_slice().as_bytes();
	 
			match file.write(num_u8) {
				Ok(()) => (), // succeeded
			    Err(e) => println!("failed to write file: {}", e),
			}
			i += 1;
		}

		let file_name_vec: Vec<u8> = std::vec::Vec::from_slice(file.as_mut().unwrap().path().filename().unwrap());
		let file_name: std::string::String = std::string::String::from_utf8(file_name_vec).unwrap();
		println!("{} lines written to {}", i, file_name);
	}
}



/* Converting string stuff

	from_buf() -- supposed to be the way to convert from c strings according to: 
		aturon @ https://github.com/rust-lang/rust/issues/16035

	format!()	// Try it out

	u8 referred to as byte
	&[u8] referred to as bytes

	std::string::String::
		from_utf8(Vec<u8>)
		from_byte(u8)
		from_char(char)
		from_chars(&[char])
		push_bytes(&[u8])
*/	
