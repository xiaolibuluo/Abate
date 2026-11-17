
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant762_0(int value, ...) {
    return value + 17;
}
#else
int variant762_0(int value) {
    return value + 17;
}
#endif
#define MASK762_1(v) (((v) << 4) | ((v) & 0x0F))

int classify762_1(int value) {
    switch (value & 0x0F) {
    case 0: return MASK762_1(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#define MASK762_2(v) (((v) << 4) | ((v) & 0x0F))

int classify762_2(int value) {
    switch (value & 0x0F) {
    case 0: return MASK762_2(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#define MASK762_3(v) (((v) << 4) | ((v) & 0x0F))

int classify762_3(int value) {
    switch (value & 0x0F) {
    case 0: return MASK762_3(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#if defined(__STDC_VERSION__)
int variant762_4(int value, ...) {
    return value + 45;
}
#else
int variant762_4(int value) {
    return value + 45;
}
#endif
struct Node762_5 {
    int value;
    struct Node762_5 *next;
};

int length762_5(const struct Node762_5 *head) {
    int count = 0;
    const struct Node762_5 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 52;
}
struct Node762_6 {
    int value;
    struct Node762_6 *next;
};

int length762_6(const struct Node762_6 *head) {
    int count = 0;
    const struct Node762_6 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 59;
}
