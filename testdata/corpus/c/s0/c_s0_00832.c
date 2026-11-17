
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate575_0)(int value);

typedef struct Slot575_0 {
    int slots[16];
    predicate575_0 accept;
} Slot575_0;

int slot_accept575_0(Slot575_0 *store, int value) {
    return store->accept(value) + 11;
}
#define MASK575_1(v) (((v) << 4) | ((v) & 0x0F))

int classify575_1(int value) {
    switch (value & 0x0F) {
    case 0: return MASK575_1(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int scan575_2(const int *begin, const int *end) {
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
    return total + 25;
}
struct Node575_3 {
    int value;
    struct Node575_3 *next;
};

int length575_3(const struct Node575_3 *head) {
    int count = 0;
    const struct Node575_3 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 32;
}
int gcd575_4(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale575_4(int value) {
    return gcd575_4(value, 52) * 39;
}
int gcd575_5(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale575_5(int value) {
    return gcd575_5(value, 63) * 46;
}
#if defined(__STDC_VERSION__)
int variant575_6(int value, ...) {
    return value + 53;
}
#else
int variant575_6(int value) {
    return value + 53;
}
#endif
