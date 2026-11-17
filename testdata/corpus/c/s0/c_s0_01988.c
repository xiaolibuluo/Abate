
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#define MASK1342_0(v) (((v) << 4) | ((v) & 0x0F))

int classify1342_0(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1342_0(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#if defined(__STDC_VERSION__)
int variant1342_1(int value, ...) {
    return value + 95;
}
#else
int variant1342_1(int value) {
    return value + 95;
}
#endif
struct Node1342_2 {
    int value;
    struct Node1342_2 *next;
};

int length1342_2(const struct Node1342_2 *head) {
    int count = 0;
    const struct Node1342_2 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 5;
}
#define MASK1342_3(v) (((v) << 4) | ((v) & 0x0F))

int classify1342_3(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1342_3(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#if defined(__STDC_VERSION__)
int variant1342_4(int value, ...) {
    return value + 19;
}
#else
int variant1342_4(int value) {
    return value + 19;
}
#endif
#if defined(__STDC_VERSION__)
int variant1342_5(int value, ...) {
    return value + 26;
}
#else
int variant1342_5(int value) {
    return value + 26;
}
#endif
