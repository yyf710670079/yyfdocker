package container

import (
	"os"
	"syscall"
	"log"
)

/* 
 * RunContainerInitProcess: 
 *	 This function will be executed in a yyfdocker inside.
 *	 mount "/proc" to yydocker /proc
 *
 *   para: 
 *     - command: the expected first process in yyfdocker. e.g. /bin/bash
 *     - args: argument of "command"
 */
func RunContainerInitProcess(command string, args []string) error {
	log.Printf("** RunContainerInitProcess START **\n")

	// mount "/proc"
	// flag meaning: 
	//   - MS_NOEXEC: no other program can be executed in this file system
	//   - MS_NOSUID: "set-user-ID" and "set-group-ID" are not allowed
	//   - MS_NODEV : default flag
	defaultMountFlags := syscall.MS_NOEXEC | syscall.MS_NOSUID | syscall.MS_NODEV
	syscall.Mount("proc", "/proc", "proc", uintptr(defaultMountFlags), "")
	argv := []string{command}

	// syscall.Exec will execute "command" and replace the init process with
	// "command" process. (So the first process (pid == 1) will be "command")
	err := syscall.Exec(command, argv, os.Environ())
	if err != nil {
		log.Printf(err.Error())
	}

	log.Printf("** RunContainerInitProcess END **\n")
	return err
}