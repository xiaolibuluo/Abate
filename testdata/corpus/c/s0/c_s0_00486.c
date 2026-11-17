
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#define MASK334_0(v) (((v) << 4) | ((v) & 0x0F))

int classify334_0(int value) {
    switch (value & 0x0F) {
    case 0: return MASK334_0(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#define MASK334_1(v) (((v) << 4) | ((v) & 0x0F))

int classify334_1(int value) {
    switch (value & 0x0F) {
    case 0: return MASK334_1(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#if defined(__STDC_VERSION__)
int variant334_2(int value, ...) {
    return value + 93;
}
#else
int variant334_2(int value) {
    return value + 93;
}
#endif
#define MASK334_3(v) (((v) << 4) | ((v) & 0x0F))

int classify334_3(int value) {
    switch (value & 0x0F) {
    case 0: return MASK334_3(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#define MASK334_4(v) (((v) << 4) | ((v) & 0x0F))

int classify334_4(int value) {
    switch (value & 0x0F) {
    case 0: return MASK334_4(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#define MASK334_5(v) (((v) << 4) | ((v) & 0x0F))

int classify334_5(int value) {
    switch (value & 0x0F) {
    case 0: return MASK334_5(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
struct Node334_6 {
    int value;
    struct Node334_6 *next;
};

int length334_6(const struct Node334_6 *head) {
    int count = 0;
    const struct Node334_6 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 24;
}
int scan334_7(const int *begin, const int *end) {
    int total = 0;
    const int *cursor = begin;
    for (; cursor != end; ++cursor) {
        if (*cursor >= 0) {
            total += *cursor;
        }
    }
    do {
        --cursor;
    } while (cursor != begin && total > 0);
    return total + 31;
}
