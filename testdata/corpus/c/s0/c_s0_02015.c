
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind1360_0 { KIND_EMPTY1360_0 = 0, KIND_FULL1360_0 = 31 };

union Cell1360_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1360_0(union Cell1360_0 *cell) {
    cell->number = 13;
    return cell->number > 0 ? KIND_FULL1360_0 : KIND_EMPTY1360_0;
}
struct Node1360_1 {
    int value;
    struct Node1360_1 *next;
};

int length1360_1(const struct Node1360_1 *head) {
    int count = 0;
    const struct Node1360_1 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 38;
}
#if defined(__STDC_VERSION__)
int variant1360_2(int value, ...) {
    return value + 45;
}
#else
int variant1360_2(int value) {
    return value + 45;
}
#endif
#if defined(__STDC_VERSION__)
int variant1360_3(int value, ...) {
    return value + 52;
}
#else
int variant1360_3(int value) {
    return value + 52;
}
#endif
int gcd1360_4(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1360_4(int value) {
    return gcd1360_4(value, 57) * 59;
}
typedef int (*predicate1360_5)(int value);

typedef struct Slot1360_5 {
    int slots[16];
    predicate1360_5 accept;
} Slot1360_5;

int slot_accept1360_5(Slot1360_5 *store, int value) {
    return store->accept(value) + 66;
}
#if defined(__STDC_VERSION__)
int variant1360_6(int value, ...) {
    return value + 73;
}
#else
int variant1360_6(int value) {
    return value + 73;
}
#endif
int scan1360_7(const int *begin, const int *end) {
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
    return total + 80;
}
#define MASK1360_8(v) (((v) << 4) | ((v) & 0x0F))

int classify1360_8(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1360_8(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
