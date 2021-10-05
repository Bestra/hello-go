package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
}

onst extract = ({ filePath, binary, os }) => {
	return new Promise(async (resolve, reject) => {
		const tmpPath = path.dirname(filePath);
		await unzip({
			from: filePath,
			to: tmpPath,
		});
		const installer = new Installer({
			engine: binary,
			path: tmpPath,
		});
		switch (os) {
			case 'win32':
			case 'win64': {
				installer.installBinary(
					{ 'qjs.exe': `${binary}.exe` },
					{ symlink: false }
				);
				installer.unstallLibrary('libwinpthread-1.dll');
				installer.unstallScript({
					name: `${binary}.cmd`,
					generateScript: (targetPath) => {
						return `
							@echo off
							"${targetPath}\\${binary}.exe" %*
						`;
					}
				});
				break;
			}
			case 'linux32':
			case 'linux64': {
				installer.installBinary({ 'qjs': binary });
				installer.installScript({
					name: binary,
					generateScript: (targetPath) => {
						return `
							#!/usr/bin/env bash
							"${targetPath}/${binary}" "$@"
						`;
					}
				});
				break;
			}
		}
		resolve();
	});
};

