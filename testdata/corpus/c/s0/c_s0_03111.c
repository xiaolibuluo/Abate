
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#define MASK2105_0(v) (((v) << 4) | ((v) & 0x0F))

int classify2105_0(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2105_0(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#if defined(__STDC_VERSION__)
int variant2105_1(int value, ...) {
    return value + 23;
}
#else
int variant2105_1(int value) {
    return value + 23;
}
#endif
struct Node2105_2 {
    int value;
    struct Node2105_2 *next;
};

int length2105_2(const struct Node2105_2 *head) {
    int count = 0;
    const struct Node2105_2 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 30;
}
