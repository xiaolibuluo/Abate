
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#define MASK248_0(v) (((v) << 4) | ((v) & 0x0F))

int classify248_0(int value) {
    switch (value & 0x0F) {
    case 0: return MASK248_0(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#define MASK248_1(v) (((v) << 4) | ((v) & 0x0F))

int classify248_1(int value) {
    switch (value & 0x0F) {
    case 0: return MASK248_1(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
