
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#define MASK830_0(v) (((v) << 4) | ((v) & 0x0F))

int classify830_0(int value) {
    switch (value & 0x0F) {
    case 0: return MASK830_0(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#define MASK830_1(v) (((v) << 4) | ((v) & 0x0F))

int classify830_1(int value) {
    switch (value & 0x0F) {
    case 0: return MASK830_1(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
