#include <stdio.h>
#include <ctype.h>
#include <string.h>
#include "common.h"

/*
   simple reader for /etc/hosts
   it only supports comments, blank lines and lines consisting of an ipv4 hostname pair.
   this is required so we can return entries from the host db without messing up the
   non-thread-safe state of libc's gethostent().

*/

struct hostsreader {
	FILE *f;
	char* ip, *name;
};

int hostsreader_open(struct hostsreader *ctx) {
	if(!(ctx->f = fopen("/etc/hosts", "r"))) return 0;
	return 1;
}

void hostsreader_close(struct hostsreader *ctx) {
	fclose(ctx->f);
}

int hostsreader_get(struct hostsreader *ctx, char* buf, size_t bufsize) {
	while(1) {
		if(!fgets(buf, bufsize, ctx->f)) return 0;
		if(*buf == '#') continue;
		char *p = buf;
		size_t l = bufsize;
		ctx->ip = p;
		while(*p && !isspace(*p) && l) {
			p++;
			l--;
		}
		if(!l || !*p || p == ctx->ip) continue;
		*p = 0;
		p++;
		while(*p && isspace(*p) && l) {
			p++;
			l--;
		}
		if(!l || !*p) continue;
		ctx->name = p;
		while(*p && !isspace(*p) && l) {
			p++;
			l--;
		}
		if(!l || !*p) continue;
		*p = 0;
		if(pc_isnumericipv4(ctx->ip)) return 1;
	}
}

char* hostsreader_get_ip_for_name(const char* name, char* buf, size_t bufsize) {
	struct hostsreader ctx;
	char *res = 0;
	if(!hostsreader_open(&ctx)) return 0;
	while(hostsreader_get(&ctx, buf, bufsize)) {
		if(!strcmp(ctx.name, name)) {
			res = ctx.ip;
			break;
		}
	}
	hostsreader_close(&ctx);
	return res;
}

#include "ip_type.h"
#include <netinet/in.h>
#include <sys/socket.h>
#include <arpa/inet.h>
ip_type4 hostsreader_get_numeric_ip_for_name(const char* name) {
	char *hres;
	char buf[320];
	if((hres = hostsreader_get_ip_for_name(name, buf, sizeof buf))) {
		struct in_addr c;
		inet_aton(hres, &c);
		ip_type4 res;
		memcpy(res.octet, &c.s_addr, 4);
		return res;
	} else return IPT4_INVALID;
}

#ifdef HOSTSREADER_TEST
#include "ip_type.c"
int main(int a, char**b) {
	char buf[256];
	if(a != 2) return 1;
	char * ret = hostsreader_get_ip_for_name(b[1], buf, sizeof buf);
	printf("%s\n", ret ? ret : "null");
}
#endif
