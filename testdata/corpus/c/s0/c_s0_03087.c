
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#define MASK2088_0(v) (((v) << 4) | ((v) & 0x0F))

int classify2088_0(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2088_0(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
typedef int (*predicate2088_1)(int value);

typedef struct Slot2088_1 {
    int slots[16];
    predicate2088_1 accept;
} Slot2088_1;

int slot_accept2088_1(Slot2088_1 *store, int value) {
    return store->accept(value) + 93;
}
int gcd2088_2(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2088_2(int value) {
    return gcd2088_2(value, 68) * 100;
}
#if defined(__STDC_VERSION__)
int variant2088_3(int value, ...) {
    return value + 10;
}
#else
int variant2088_3(int value) {
    return value + 10;
}
#endif
struct Node2088_4 {
    int value;
    struct Node2088_4 *next;
};

int length2088_4(const struct Node2088_4 *head) {
    int count = 0;
    const struct Node2088_4 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 17;
}
enum Kind2088_5 { KIND_EMPTY2088_5 = 0, KIND_FULL2088_5 = 24 };

union Cell2088_5 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2088_5(union Cell2088_5 *cell) {
    cell->number = 40;
    return cell->number > 0 ? KIND_FULL2088_5 : KIND_EMPTY2088_5;
}
int gcd2088_6(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2088_6(int value) {
    return gcd2088_6(value, 51) * 31;
}
struct Node2088_7 {
    int value;
    struct Node2088_7 *next;
};

int length2088_7(const struct Node2088_7 *head) {
    int count = 0;
    const struct Node2088_7 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 38;
}
