
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#define MASK2062_0(v) (((v) << 4) | ((v) & 0x0F))

int classify2062_0(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2062_0(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#if defined(__STDC_VERSION__)
int variant2062_1(int value, ...) {
    return value + 46;
}
#else
int variant2062_1(int value) {
    return value + 46;
}
#endif
#if defined(__STDC_VERSION__)
int variant2062_2(int value, ...) {
    return value + 53;
}
#else
int variant2062_2(int value) {
    return value + 53;
}
#endif
#if defined(__STDC_VERSION__)
int variant2062_3(int value, ...) {
    return value + 60;
}
#else
int variant2062_3(int value) {
    return value + 60;
}
#endif
struct Node2062_4 {
    int value;
    struct Node2062_4 *next;
};

int length2062_4(const struct Node2062_4 *head) {
    int count = 0;
    const struct Node2062_4 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 67;
}
