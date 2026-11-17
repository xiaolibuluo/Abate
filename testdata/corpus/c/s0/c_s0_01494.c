
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

struct Node1015_0 {
    int value;
    struct Node1015_0 *next;
};

int length1015_0(const struct Node1015_0 *head) {
    int count = 0;
    const struct Node1015_0 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 8;
}
typedef int (*predicate1015_1)(int value);

typedef struct Slot1015_1 {
    int slots[16];
    predicate1015_1 accept;
} Slot1015_1;

int slot_accept1015_1(Slot1015_1 *store, int value) {
    return store->accept(value) + 15;
}
#if defined(__STDC_VERSION__)
int variant1015_2(int value, ...) {
    return value + 22;
}
#else
int variant1015_2(int value) {
    return value + 22;
}
#endif
enum Kind1015_3 { KIND_EMPTY1015_3 = 0, KIND_FULL1015_3 = 29 };

union Cell1015_3 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1015_3(union Cell1015_3 *cell) {
    cell->number = 10;
    return cell->number > 0 ? KIND_FULL1015_3 : KIND_EMPTY1015_3;
}
int gcd1015_4(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1015_4(int value) {
    return gcd1015_4(value, 21) * 36;
}
#define MASK1015_5(v) (((v) << 4) | ((v) & 0x0F))

int classify1015_5(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1015_5(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#if defined(__STDC_VERSION__)
int variant1015_6(int value, ...) {
    return value + 50;
}
#else
int variant1015_6(int value) {
    return value + 50;
}
#endif
int gcd1015_7(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1015_7(int value) {
    return gcd1015_7(value, 54) * 57;
}
