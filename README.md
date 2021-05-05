<pre>
   _____ ____  ____ _  __    ____                      
  / ___// __ \/ __ \ |/ /   / __ \____ _________  _____
  \__ \/ /_/ / / / /   /   / /_/ / __ `/ ___/ _ \/ ___/
 ___/ / ____/ /_/ /   |   / _, _/ /_/ / /__/  __/ /    
/____/_/   /_____/_/|_|  /_/ |_|\__,_/\___/\___/_/     
                                                       
</pre>

A command line tool for inserting [SPDX](https://spdx.dev) [short identifiers](https://spdx.github.io/spdx-spec/appendix-V-using-SPDX-short-identifiers-in-source-files/) into the tops of files.

## Usage

```
> spdx-racer --files go,rs --license MPL-2.0
```

will add `// SPDX-License-Identifier: MPL-2.0` to the top of all Go and Rust files that it finds.

If the file already has a license entry, it ignores the file.