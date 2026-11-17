
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#define MASK696_0(v) (((v) << 4) | ((v) & 0x0F))

int classify696_0(int value) {
    switch (value & 0x0F) {
    case 0: return MASK696_0(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
struct Node696_1 {
    int value;
    struct Node696_1 *next;
};

int length696_1(const struct Node696_1 *head) {
    int count = 0;
    const struct Node696_1 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 39;
}
#if defined(__STDC_VERSION__)
int variant696_2(int value, ...) {
    return value + 46;
}
#else
int variant696_2(int value) {
    return value + 46;
}
#endif
#if defined(__STDC_VERSION__)
int variant696_3(int value, ...) {
    return value + 53;
}
#else
int variant696_3(int value) {
    return value + 53;
}
#endif
