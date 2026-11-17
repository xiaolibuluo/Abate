
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant1943_0(int value, ...) {
    return value + 44;
}
#else
int variant1943_0(int value) {
    return value + 44;
}
#endif
#define MASK1943_1(v) (((v) << 4) | ((v) & 0x0F))

int classify1943_1(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1943_1(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
struct Node1943_2 {
    int value;
    struct Node1943_2 *next;
};

int length1943_2(const struct Node1943_2 *head) {
    int count = 0;
    const struct Node1943_2 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 58;
}
int gcd1943_3(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1943_3(int value) {
    return gcd1943_3(value, 40) * 65;
}
int scan1943_4(const int *begin, const int *end) {
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
    return total + 72;
}
typedef int (*predicate1943_5)(int value);

typedef struct Slot1943_5 {
    int slots[16];
    predicate1943_5 accept;
} Slot1943_5;

int slot_accept1943_5(Slot1943_5 *store, int value) {
    return store->accept(value) + 79;
}
int gcd1943_6(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1943_6(int value) {
    return gcd1943_6(value, 12) * 86;
}
enum Kind1943_7 { KIND_EMPTY1943_7 = 0, KIND_FULL1943_7 = 93 };

union Cell1943_7 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1943_7(union Cell1943_7 *cell) {
    cell->number = 23;
    return cell->number > 0 ? KIND_FULL1943_7 : KIND_EMPTY1943_7;
}
