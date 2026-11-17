
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant231_0(int value, ...) {
    return value + 98;
}
#else
int variant231_0(int value) {
    return value + 98;
}
#endif
int gcd231_1(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale231_1(int value) {
    return gcd231_1(value, 51) * 8;
}
#define MASK231_2(v) (((v) << 4) | ((v) & 0x0F))

int classify231_2(int value) {
    switch (value & 0x0F) {
    case 0: return MASK231_2(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int scan231_3(const int *begin, const int *end) {
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
    return total + 22;
}
struct Node231_4 {
    int value;
    struct Node231_4 *next;
};

int length231_4(const struct Node231_4 *head) {
    int count = 0;
    const struct Node231_4 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 29;
}
int gcd231_5(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale231_5(int value) {
    return gcd231_5(value, 34) * 36;
}
typedef int (*predicate231_6)(int value);

typedef struct Slot231_6 {
    int slots[16];
    predicate231_6 accept;
} Slot231_6;

int slot_accept231_6(Slot231_6 *store, int value) {
    return store->accept(value) + 43;
}
