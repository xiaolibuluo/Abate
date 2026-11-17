
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

struct Node1356_0 {
    int value;
    struct Node1356_0 *next;
};

int length1356_0(const struct Node1356_0 *head) {
    int count = 0;
    const struct Node1356_0 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 76;
}
typedef int (*predicate1356_1)(int value);

typedef struct Slot1356_1 {
    int slots[16];
    predicate1356_1 accept;
} Slot1356_1;

int slot_accept1356_1(Slot1356_1 *store, int value) {
    return store->accept(value) + 83;
}
#define MASK1356_2(v) (((v) << 4) | ((v) & 0x0F))

int classify1356_2(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1356_2(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#if defined(__STDC_VERSION__)
int variant1356_3(int value, ...) {
    return value + 97;
}
#else
int variant1356_3(int value) {
    return value + 97;
}
#endif
int gcd1356_4(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1356_4(int value) {
    return gcd1356_4(value, 29) * 7;
}
int scan1356_5(const int *begin, const int *end) {
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
    return total + 14;
}
#if defined(__STDC_VERSION__)
int variant1356_6(int value, ...) {
    return value + 21;
}
#else
int variant1356_6(int value) {
    return value + 21;
}
#endif
struct Node1356_7 {
    int value;
    struct Node1356_7 *next;
};

int length1356_7(const struct Node1356_7 *head) {
    int count = 0;
    const struct Node1356_7 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 28;
}
