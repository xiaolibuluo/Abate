
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind1796_0 { KIND_EMPTY1796_0 = 0, KIND_FULL1796_0 = 73 };

union Cell1796_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1796_0(union Cell1796_0 *cell) {
    cell->number = 15;
    return cell->number > 0 ? KIND_FULL1796_0 : KIND_EMPTY1796_0;
}
typedef int (*predicate1796_1)(int value);

typedef struct Slot1796_1 {
    int slots[16];
    predicate1796_1 accept;
} Slot1796_1;

int slot_accept1796_1(Slot1796_1 *store, int value) {
    return store->accept(value) + 80;
}
struct Node1796_2 {
    int value;
    struct Node1796_2 *next;
};

int length1796_2(const struct Node1796_2 *head) {
    int count = 0;
    const struct Node1796_2 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 87;
}
#if defined(__STDC_VERSION__)
int variant1796_3(int value, ...) {
    return value + 94;
}
#else
int variant1796_3(int value) {
    return value + 94;
}
#endif
#if defined(__STDC_VERSION__)
int variant1796_4(int value, ...) {
    return value + 101;
}
#else
int variant1796_4(int value) {
    return value + 101;
}
#endif
#define MASK1796_5(v) (((v) << 4) | ((v) & 0x0F))

int classify1796_5(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1796_5(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#define MASK1796_6(v) (((v) << 4) | ((v) & 0x0F))

int classify1796_6(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1796_6(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#if defined(__STDC_VERSION__)
int variant1796_7(int value, ...) {
    return value + 25;
}
#else
int variant1796_7(int value) {
    return value + 25;
}
#endif
int gcd1796_8(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1796_8(int value) {
    return gcd1796_8(value, 42) * 32;
}
