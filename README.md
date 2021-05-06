<pre>
   _____ ____  ____ _  __    ____                      
  / ___// __ \/ __ \ |/ /   / __ \____ _________  _____
  \__ \/ /_/ / / / /   /   / /_/ / __ `/ ___/ _ \/ ___/
 ___/ / ____/ /_/ /   |   / _, _/ /_/ / /__/  __/ /    
/____/_/   /_____/_/|_|  /_/ |_|\__,_/\___/\___/_/     
                                                       
</pre>

A command line tool for inserting [SPDX](https://spdx.dev) [short identifiers](https://spdx.github.io/spdx-spec/appendix-V-using-SPDX-short-identifiers-in-source-files/) into the tops of files.

* [Usage](#usage)
   * [Adding a License](#adding-a-license)
   * [Removing a License](#removing-a-license)

## Usage

### Adding a License

```
> spdx-racer --files go,rs --license MPL-2.0
```

will insert `// SPDX-License-Identifier: MPL-2.0` at the top of all Go and Rust files in the current directory.

If the file already has a license entry, it ignores the file.

### Removing a License

To delete a license, add the `--delete` or `-d` flag to the command:

```
> spdx-racer --files go,rs --license MPL-2.0 -d
```

will delete `// SPDX-License-Identifier: MPL-2.0` from the top of all Go and Rust files in the current directory.
