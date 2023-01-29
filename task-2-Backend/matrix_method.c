#include <stdint.h>
#include <stdio.h>

typedef struct poly_s {
    int64_t x0, x1, x2;
} poly_s;

poly_s mul(poly_s a, poly_s b) {
    return (poly_s){
        a.x0*b.x0 + a.x2*b.x1 + a.x1*b.x2,
        a.x1*b.x0 + a.x0*b.x1 + a.x2*b.x2 + a.x2*b.x1 + a.x1*b.x2,
        a.x2*b.x0 + a.x1*b.x1 + a.x0*b.x2 + a.x2*b.x2
    };
}

poly_s polypow(poly_s x, int n) {
    poly_s r = {1, 0, 0};
    while(n) {
        if (n & 1) r = mul(r, x);
        x = mul(x, x);
        n >>= 1;
    }
    return r;
}

int main() {
    for (int i = 0; i < 30; i++) {
        printf("%d: %ld\n", i, mul((poly_s){0, 2, 1}, polypow((poly_s){0, 1, 0}, i)).x0);
    }
}
