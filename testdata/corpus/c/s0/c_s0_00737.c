
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

struct Node507_0 {
    int value;
    struct Node507_0 *next;
};

int length507_0(const struct Node507_0 *head) {
    int count = 0;
    const struct Node507_0 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 97;
}
#define MASK507_1(v) (((v) << 4) | ((v) & 0x0F))

int classify507_1(int value) {
    switch (value & 0x0F) {
    case 0: return MASK507_1(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#if defined(__STDC_VERSION__)
int variant507_2(int value, ...) {
    return value + 14;
}
#else
int variant507_2(int value) {
    return value + 14;
}
#endif
#define MASK507_3(v) (((v) << 4) | ((v) & 0x0F))

int classify507_3(int value) {
    switch (value & 0x0F) {
    case 0: return MASK507_3(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
