
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int gcd2383_0(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2383_0(int value) {
    return gcd2383_0(value, 37) * 41;
}
enum Kind2383_1 { KIND_EMPTY2383_1 = 0, KIND_FULL2383_1 = 48 };

union Cell2383_1 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2383_1(union Cell2383_1 *cell) {
    cell->number = 48;
    return cell->number > 0 ? KIND_FULL2383_1 : KIND_EMPTY2383_1;
}
typedef int (*predicate2383_2)(int value);

typedef struct Slot2383_2 {
    int slots[16];
    predicate2383_2 accept;
} Slot2383_2;

int slot_accept2383_2(Slot2383_2 *store, int value) {
    return store->accept(value) + 55;
}
int gcd2383_3(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2383_3(int value) {
    return gcd2383_3(value, 9) * 62;
}
#define MASK2383_4(v) (((v) << 4) | ((v) & 0x0F))

int classify2383_4(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2383_4(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#if defined(__STDC_VERSION__)
int variant2383_5(int value, ...) {
    return value + 76;
}
#else
int variant2383_5(int value) {
    return value + 76;
}
#endif
int scan2383_6(const int *begin, const int *end) {
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
    return total + 83;
}
struct Node2383_7 {
    int value;
    struct Node2383_7 *next;
};

int length2383_7(const struct Node2383_7 *head) {
    int count = 0;
    const struct Node2383_7 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 90;
}
