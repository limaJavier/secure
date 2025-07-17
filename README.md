# 🔐 Secure

A fast, reliable, and secure CLI tool to manage passwords, built with Go.

> Because storing your passwords shouldn't depend on a browser.

## ✨ Features

- 🔒 Strong symmetric encryption with AES-GCM
- 🧠 Key derivation using Argon2id (memory-hard, resistant to GPU attacks)
- 🧰 Extensible architecture for future algorithms and storage backends
- 📁 Local, offline vault storage — no cloud, no tracking
- 🧼 Full CRUD support via simple CLI commands

## ⚙️ Installation

```bash
git clone https://github.com/limaJavier/secure.git
cd secure
go mod tidy
go build -o bin/secure cmd/main.go
```

## 🚀 Usage

### 🆕 Create a user:

**secure init**

```console
$ secure init
Enter username: johndoe
Enter password:
Confirm password:
```

### ➕ Create a password-entry:

**secure create**

```console
$ secure create
Enter username: johndoe
Enter password:
Enter name: GitHub Account Password
Enter description: Password for johndoe's GitHub Account
Enter password:
Confirm password: 
```

### 📋 Retrieve all password-entries:

**secure retrieve**

```console
$ secure retrieve
Enter username: johndoe
Enter password:
--------------------
ID: 1
Name: GitHub Account Password
Description: Password for johndoe's GitHub Account
Password: doepasswordgithub
--------------------
```

### ✏️ Update a password-entry

**secure update <entry-id>**

```console
$ secure update 1
Enter username: johndoe
Enter password:
Enter name: GitHub Account Password Updated
Enter description:
Enter password:
```

**Note**: Leave a field blank to skip updating it.

### ❌ Delete a password-entry

**secure delete <entry-id>**

```console
$ secure delete 1
Enter username: johndoe
Enter password:
```

### 🆘 Help

**secure --help**

For a full list of commands and flags:

```bash
secure --help
```

And optionally per command:

```bash
secure create --help
```

## 🛡️ Security Design

* Secrets are encrypted with AES-256-GCM using unique nonces for each entry.
* Each user gets a unique, randomly generated master key. This master key is encrypted using a key derived from the user's password (via Argon2id), and stored securely.
* All password entries are encrypted using the master key, ensuring forward secrecy.
* Everything is stored locally on a sqlite database — no third-party servers or network calls.

## 🗂️ Storage

All encrypted entries are saved under:

```plaintext
$HOME/.local/share/secure/secure.db
```

You can safely back up this file — but **you must remember your master password** to decrypt it. If lost, recovery is impossible by design.

## 🧪 Running Tests

```bash
go test ./...
```

## 🧭 Coming Soon

* User management improvements (update & delete users)
* Encrypted database backups
* Session-based authentication (like sudo) — unlock once, use for a limited time

## 🤝 Contributing

Pull requests are welcome! For major changes, please open an issue first to discuss your ideas.

Please:

* Write clear commit messages
* Include appropriate test coverage
* Follow the existing code style and structure
* Squash all your commits


## 📄 License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## 🙌 Acknowledgements

Thanks to:

* The [Go cryptography community](https://golang.org/pkg/crypto/)
* The authors of [Argon2](https://github.com/P-H-C/phc-winner-argon2)
* CLI inspiration from tools like `pass`, `gh`, and `git`
