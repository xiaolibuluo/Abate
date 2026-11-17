
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

struct Node336_0 {
    int value;
    struct Node336_0 *next;
};

int length336_0(const struct Node336_0 *head) {
    int count = 0;
    const struct Node336_0 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 8;
}
#define MASK336_1(v) (((v) << 4) | ((v) & 0x0F))

int classify336_1(int value) {
    switch (value & 0x0F) {
    case 0: return MASK336_1(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
typedef int (*predicate336_2)(int value);

typedef struct Slot336_2 {
    int slots[16];
    predicate336_2 accept;
} Slot336_2;

int slot_accept336_2(Slot336_2 *store, int value) {
    return store->accept(value) + 22;
}
typedef int (*predicate336_3)(int value);

typedef struct Slot336_3 {
    int slots[16];
    predicate336_3 accept;
} Slot336_3;

int slot_accept336_3(Slot336_3 *store, int value) {
    return store->accept(value) + 29;
}
#if defined(__STDC_VERSION__)
int variant336_4(int value, ...) {
    return value + 36;
}
#else
int variant336_4(int value) {
    return value + 36;
}
#endif
int gcd336_5(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale336_5(int value) {
    return gcd336_5(value, 37) * 43;
}
typedef int (*predicate336_6)(int value);

typedef struct Slot336_6 {
    int slots[16];
    predicate336_6 accept;
} Slot336_6;

int slot_accept336_6(Slot336_6 *store, int value) {
    return store->accept(value) + 50;
}
#if defined(__STDC_VERSION__)
int variant336_7(int value, ...) {
    return value + 57;
}
#else
int variant336_7(int value) {
    return value + 57;
}
#endif
int scan336_8(const int *begin, const int *end) {
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
    return total + 64;
}
