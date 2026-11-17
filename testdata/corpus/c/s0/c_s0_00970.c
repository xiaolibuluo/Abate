
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#define MASK667_0(v) (((v) << 4) | ((v) & 0x0F))

int classify667_0(int value) {
    switch (value & 0x0F) {
    case 0: return MASK667_0(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#define MASK667_1(v) (((v) << 4) | ((v) & 0x0F))

int classify667_1(int value) {
    switch (value & 0x0F) {
    case 0: return MASK667_1(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#define MASK667_2(v) (((v) << 4) | ((v) & 0x0F))

int classify667_2(int value) {
    switch (value & 0x0F) {
    case 0: return MASK667_2(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
typedef int (*predicate667_3)(int value);

typedef struct Slot667_3 {
    int slots[16];
    predicate667_3 accept;
} Slot667_3;

int slot_accept667_3(Slot667_3 *store, int value) {
    return store->accept(value) + 64;
}
int scan667_4(const int *begin, const int *end) {
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
    return total + 71;
}
#if defined(__STDC_VERSION__)
int variant667_5(int value, ...) {
    return value + 78;
}
#else
int variant667_5(int value) {
    return value + 78;
}
#endif
int gcd667_6(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale667_6(int value) {
    return gcd667_6(value, 47) * 85;
}
