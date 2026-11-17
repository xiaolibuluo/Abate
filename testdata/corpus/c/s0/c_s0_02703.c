
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind1837_0 { KIND_EMPTY1837_0 = 0, KIND_FULL1837_0 = 24 };

union Cell1837_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1837_0(union Cell1837_0 *cell) {
    cell->number = 58;
    return cell->number > 0 ? KIND_FULL1837_0 : KIND_EMPTY1837_0;
}
int gcd1837_1(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1837_1(int value) {
    return gcd1837_1(value, 69) * 31;
}
typedef int (*predicate1837_2)(int value);

typedef struct Slot1837_2 {
    int slots[16];
    predicate1837_2 accept;
} Slot1837_2;

int slot_accept1837_2(Slot1837_2 *store, int value) {
    return store->accept(value) + 38;
}
int gcd1837_3(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1837_3(int value) {
    return gcd1837_3(value, 30) * 45;
}
#if defined(__STDC_VERSION__)
int variant1837_4(int value, ...) {
    return value + 52;
}
#else
int variant1837_4(int value) {
    return value + 52;
}
#endif
struct Node1837_5 {
    int value;
    struct Node1837_5 *next;
};

int length1837_5(const struct Node1837_5 *head) {
    int count = 0;
    const struct Node1837_5 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 59;
}
int scan1837_6(const int *begin, const int *end) {
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
    return total + 66;
}
