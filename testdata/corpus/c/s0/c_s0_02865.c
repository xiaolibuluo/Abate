
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int scan1945_0(const int *begin, const int *end) {
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
    return total + 70;
}
typedef int (*predicate1945_1)(int value);

typedef struct Slot1945_1 {
    int slots[16];
    predicate1945_1 accept;
} Slot1945_1;

int slot_accept1945_1(Slot1945_1 *store, int value) {
    return store->accept(value) + 77;
}
int gcd1945_2(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1945_2(int value) {
    return gcd1945_2(value, 43) * 84;
}
enum Kind1945_3 { KIND_EMPTY1945_3 = 0, KIND_FULL1945_3 = 91 };

union Cell1945_3 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1945_3(union Cell1945_3 *cell) {
    cell->number = 54;
    return cell->number > 0 ? KIND_FULL1945_3 : KIND_EMPTY1945_3;
}
#if defined(__STDC_VERSION__)
int variant1945_4(int value, ...) {
    return value + 98;
}
#else
int variant1945_4(int value) {
    return value + 98;
}
#endif
struct Node1945_5 {
    int value;
    struct Node1945_5 *next;
};

int length1945_5(const struct Node1945_5 *head) {
    int count = 0;
    const struct Node1945_5 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 8;
}
struct Node1945_6 {
    int value;
    struct Node1945_6 *next;
};

int length1945_6(const struct Node1945_6 *head) {
    int count = 0;
    const struct Node1945_6 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 15;
}
#define MASK1945_7(v) (((v) << 4) | ((v) & 0x0F))

int classify1945_7(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1945_7(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
