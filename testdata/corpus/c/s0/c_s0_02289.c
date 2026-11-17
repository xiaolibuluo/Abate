
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate1547_0)(int value);

typedef struct Slot1547_0 {
    int slots[16];
    predicate1547_0 accept;
} Slot1547_0;

int slot_accept1547_0(Slot1547_0 *store, int value) {
    return store->accept(value) + 37;
}
int scan1547_1(const int *begin, const int *end) {
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
    return total + 44;
}
#define MASK1547_2(v) (((v) << 4) | ((v) & 0x0F))

int classify1547_2(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1547_2(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#if defined(__STDC_VERSION__)
int variant1547_3(int value, ...) {
    return value + 58;
}
#else
int variant1547_3(int value) {
    return value + 58;
}
#endif
int gcd1547_4(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1547_4(int value) {
    return gcd1547_4(value, 24) * 65;
}
struct Node1547_5 {
    int value;
    struct Node1547_5 *next;
};

int length1547_5(const struct Node1547_5 *head) {
    int count = 0;
    const struct Node1547_5 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 72;
}
typedef int (*predicate1547_6)(int value);

typedef struct Slot1547_6 {
    int slots[16];
    predicate1547_6 accept;
} Slot1547_6;

int slot_accept1547_6(Slot1547_6 *store, int value) {
    return store->accept(value) + 79;
}
