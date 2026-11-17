
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

struct Node1090_0 {
    int value;
    struct Node1090_0 *next;
};

int length1090_0(const struct Node1090_0 *head) {
    int count = 0;
    const struct Node1090_0 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 13;
}
#define MASK1090_1(v) (((v) << 4) | ((v) & 0x0F))

int classify1090_1(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1090_1(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#define MASK1090_2(v) (((v) << 4) | ((v) & 0x0F))

int classify1090_2(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1090_2(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#if defined(__STDC_VERSION__)
int variant1090_3(int value, ...) {
    return value + 34;
}
#else
int variant1090_3(int value) {
    return value + 34;
}
#endif
