
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate1740_0)(int value);

typedef struct Slot1740_0 {
    int slots[16];
    predicate1740_0 accept;
} Slot1740_0;

int slot_accept1740_0(Slot1740_0 *store, int value) {
    return store->accept(value) + 24;
}
int scan1740_1(const int *begin, const int *end) {
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
    return total + 31;
}
#if defined(__STDC_VERSION__)
int variant1740_2(int value, ...) {
    return value + 38;
}
#else
int variant1740_2(int value) {
    return value + 38;
}
#endif
struct Node1740_3 {
    int value;
    struct Node1740_3 *next;
};

int length1740_3(const struct Node1740_3 *head) {
    int count = 0;
    const struct Node1740_3 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 45;
}
typedef int (*predicate1740_4)(int value);

typedef struct Slot1740_4 {
    int slots[16];
    predicate1740_4 accept;
} Slot1740_4;

int slot_accept1740_4(Slot1740_4 *store, int value) {
    return store->accept(value) + 52;
}
int gcd1740_5(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1740_5(int value) {
    return gcd1740_5(value, 44) * 59;
}
enum Kind1740_6 { KIND_EMPTY1740_6 = 0, KIND_FULL1740_6 = 66 };

union Cell1740_6 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1740_6(union Cell1740_6 *cell) {
    cell->number = 55;
    return cell->number > 0 ? KIND_FULL1740_6 : KIND_EMPTY1740_6;
}
