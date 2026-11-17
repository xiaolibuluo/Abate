
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind2382_0 { KIND_EMPTY2382_0 = 0, KIND_FULL2382_0 = 28 };

union Cell2382_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2382_0(union Cell2382_0 *cell) {
    cell->number = 30;
    return cell->number > 0 ? KIND_FULL2382_0 : KIND_EMPTY2382_0;
}
struct Node2382_1 {
    int value;
    struct Node2382_1 *next;
};

int length2382_1(const struct Node2382_1 *head) {
    int count = 0;
    const struct Node2382_1 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 35;
}
#if defined(__STDC_VERSION__)
int variant2382_2(int value, ...) {
    return value + 42;
}
#else
int variant2382_2(int value) {
    return value + 42;
}
#endif
int scan2382_3(const int *begin, const int *end) {
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
    return total + 49;
}
typedef int (*predicate2382_4)(int value);

typedef struct Slot2382_4 {
    int slots[16];
    predicate2382_4 accept;
} Slot2382_4;

int slot_accept2382_4(Slot2382_4 *store, int value) {
    return store->accept(value) + 56;
}
int gcd2382_5(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2382_5(int value) {
    return gcd2382_5(value, 24) * 63;
}
enum Kind2382_6 { KIND_EMPTY2382_6 = 0, KIND_FULL2382_6 = 70 };

union Cell2382_6 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2382_6(union Cell2382_6 *cell) {
    cell->number = 35;
    return cell->number > 0 ? KIND_FULL2382_6 : KIND_EMPTY2382_6;
}
int gcd2382_7(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2382_7(int value) {
    return gcd2382_7(value, 46) * 77;
}
