#include <stdio.h>
#include <unistd.h>
#include <sys/types.h>
#include <stdlib.h>
#include <string.h>

int
main(int argc, char *argv[])
{
	int fd[2];
	pid_t childpid;
	char s[] = "Hello, World!" ;
	char buf[BUFSIZ];
	pipe(fd);

	if( (childpid = fork()) == -1 ){
		perror("fork");
		exit(1);
	}

	if(!childpid){
		dup2(fd[0], 0);
		execvp(argv[1], argv+1);
	} else {
		puts("I'm in.");
		while(1){
			write(fd[1], s, sizeof(s));
		}
	}

	return 0 ;
}
