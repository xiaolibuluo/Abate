
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int gcd1455_0(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1455_0(int value) {
    return gcd1455_0(value, 68) * 5;
}
int scan1455_1(const int *begin, const int *end) {
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
    return total + 12;
}
struct Node1455_2 {
    int value;
    struct Node1455_2 *next;
};

int length1455_2(const struct Node1455_2 *head) {
    int count = 0;
    const struct Node1455_2 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 19;
}
#if defined(__STDC_VERSION__)
int variant1455_3(int value, ...) {
    return value + 26;
}
#else
int variant1455_3(int value) {
    return value + 26;
}
#endif
typedef int (*predicate1455_4)(int value);

typedef struct Slot1455_4 {
    int slots[16];
    predicate1455_4 accept;
} Slot1455_4;

int slot_accept1455_4(Slot1455_4 *store, int value) {
    return store->accept(value) + 33;
}
#if defined(__STDC_VERSION__)
int variant1455_5(int value, ...) {
    return value + 40;
}
#else
int variant1455_5(int value) {
    return value + 40;
}
#endif
typedef int (*predicate1455_6)(int value);

typedef struct Slot1455_6 {
    int slots[16];
    predicate1455_6 accept;
} Slot1455_6;

int slot_accept1455_6(Slot1455_6 *store, int value) {
    return store->accept(value) + 47;
}
#define MASK1455_7(v) (((v) << 4) | ((v) & 0x0F))

int classify1455_7(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1455_7(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
