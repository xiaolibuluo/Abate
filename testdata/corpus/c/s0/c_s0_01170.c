
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant802_0(int value, ...) {
    return value + 52;
}
#else
int variant802_0(int value) {
    return value + 52;
}
#endif
#if defined(__STDC_VERSION__)
int variant802_1(int value, ...) {
    return value + 59;
}
#else
int variant802_1(int value) {
    return value + 59;
}
#endif
enum Kind802_2 { KIND_EMPTY802_2 = 0, KIND_FULL802_2 = 66 };

union Cell802_2 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind802_2(union Cell802_2 *cell) {
    cell->number = 33;
    return cell->number > 0 ? KIND_FULL802_2 : KIND_EMPTY802_2;
}
#define MASK802_3(v) (((v) << 4) | ((v) & 0x0F))

int classify802_3(int value) {
    switch (value & 0x0F) {
    case 0: return MASK802_3(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#define MASK802_4(v) (((v) << 4) | ((v) & 0x0F))

int classify802_4(int value) {
    switch (value & 0x0F) {
    case 0: return MASK802_4(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
