
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int gcd1508_0(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1508_0(int value) {
    return gcd1508_0(value, 12) * 15;
}
struct Node1508_1 {
    int value;
    struct Node1508_1 *next;
};

int length1508_1(const struct Node1508_1 *head) {
    int count = 0;
    const struct Node1508_1 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 22;
}
#if defined(__STDC_VERSION__)
int variant1508_2(int value, ...) {
    return value + 29;
}
#else
int variant1508_2(int value) {
    return value + 29;
}
#endif
int scan1508_3(const int *begin, const int *end) {
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
    return total + 36;
}
int gcd1508_4(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1508_4(int value) {
    return gcd1508_4(value, 56) * 43;
}
#if defined(__STDC_VERSION__)
int variant1508_5(int value, ...) {
    return value + 50;
}
#else
int variant1508_5(int value) {
    return value + 50;
}
#endif
typedef int (*predicate1508_6)(int value);

typedef struct Slot1508_6 {
    int slots[16];
    predicate1508_6 accept;
} Slot1508_6;

int slot_accept1508_6(Slot1508_6 *store, int value) {
    return store->accept(value) + 57;
}
typedef int (*predicate1508_7)(int value);

typedef struct Slot1508_7 {
    int slots[16];
    predicate1508_7 accept;
} Slot1508_7;

int slot_accept1508_7(Slot1508_7 *store, int value) {
    return store->accept(value) + 64;
}
enum Kind1508_8 { KIND_EMPTY1508_8 = 0, KIND_FULL1508_8 = 71 };

union Cell1508_8 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1508_8(union Cell1508_8 *cell) {
    cell->number = 39;
    return cell->number > 0 ? KIND_FULL1508_8 : KIND_EMPTY1508_8;
}
