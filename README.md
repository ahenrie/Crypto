# TMTO Attack on A5/1

This is a short readme since I have a report written on this.

## Requirements
- Golang

## File Structure

```bash
CryptoFinal/
├── data/
│   └── time.go
├── internal/
│   └── (empty)
├── main.go
├── README.md
├── pkg/
│   ├── a5/
│   │   ├── a5.go
│   │   ├── a5_encrypt.go
│   │   └── a5_test.go
│   └── tmto/
│       ├── precompute.go
│       └── precompute_test.go

```

## Usage
1. Run the program:
   ```bash
   go run main.go
   ```
2. Follow the prompts:
   - Enter your name when prompted.
   - Provide a 64-bit hexadecimal key (e.g., `1234567890abcdef`).

3. The program will:
   - Encrypt a plaintext message using the A5/1 cipher.
   - Generate a lookup table for decryption.
   - Attempt to decrypt the ciphertext and display the result.

## Example Output
```bash
*********************************
* Welcome to the A5/1 TMTO *
*********************************
What is your name:
John
Hello, John. We will be encrypting your name today then breaking the encryption.

*********************************
* Please enter a 64-bit hexadecimal key *
* (e.g. 1234567890abcdef) *
*********************************
1234567890abcdef

*********************************
* Encryption Complete! *
*********************************
Ciphertext: 49656c6c6f2d4a6f686e21

*********************************
* Generating Table... *
*********************************
Progress: Generated 0 entries
Planted key at 999999: 1234567890abcdef
Generated table with 1000000 unique entries.

*********************************
* Attempting to Decrypt... *
*********************************
*********************************
* Decryption Successful! *
*********************************
Found key: 1234567890abcdef
Found keystream: 49656c6c6f2d4a6f
Decrypted plaintext:  Hi! John. We inserted your key into the lookup table to save time.
```

## Notes
- The program includes a progress tracker during the table generation phase.
- Ensure the input key is a valid 64-bit hexadecimal value.
- Adjust the table size and parameters in the code as needed for performance testing.(I probably would not touch these)

## time.go
If you run this, please change the iterations down from a billion.
