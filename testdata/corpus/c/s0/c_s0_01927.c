
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int gcd1300_0(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1300_0(int value) {
    return gcd1300_0(value, 20) * 27;
}
int scan1300_1(const int *begin, const int *end) {
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
    return total + 34;
}
struct Node1300_2 {
    int value;
    struct Node1300_2 *next;
};

int length1300_2(const struct Node1300_2 *head) {
    int count = 0;
    const struct Node1300_2 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 41;
}
#define MASK1300_3(v) (((v) << 4) | ((v) & 0x0F))

int classify1300_3(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1300_3(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#if defined(__STDC_VERSION__)
int variant1300_4(int value, ...) {
    return value + 55;
}
#else
int variant1300_4(int value) {
    return value + 55;
}
#endif
typedef int (*predicate1300_5)(int value);

typedef struct Slot1300_5 {
    int slots[16];
    predicate1300_5 accept;
} Slot1300_5;

int slot_accept1300_5(Slot1300_5 *store, int value) {
    return store->accept(value) + 62;
}
