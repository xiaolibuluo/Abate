
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#define MASK1146_0(v) (((v) << 4) | ((v) & 0x0F))

int classify1146_0(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1146_0(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#define MASK1146_1(v) (((v) << 4) | ((v) & 0x0F))

int classify1146_1(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1146_1(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#if defined(__STDC_VERSION__)
int variant1146_2(int value, ...) {
    return value + 76;
}
#else
int variant1146_2(int value) {
    return value + 76;
}
#endif
#define MASK1146_3(v) (((v) << 4) | ((v) & 0x0F))

int classify1146_3(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1146_3(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
struct Node1146_4 {
    int value;
    struct Node1146_4 *next;
};

int length1146_4(const struct Node1146_4 *head) {
    int count = 0;
    const struct Node1146_4 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 90;
}
