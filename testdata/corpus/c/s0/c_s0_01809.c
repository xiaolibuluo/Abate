
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int gcd1214_0(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1214_0(int value) {
    return gcd1214_0(value, 28) * 73;
}
#if defined(__STDC_VERSION__)
int variant1214_1(int value, ...) {
    return value + 80;
}
#else
int variant1214_1(int value) {
    return value + 80;
}
#endif
int scan1214_2(const int *begin, const int *end) {
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
    return total + 87;
}
typedef int (*predicate1214_3)(int value);

typedef struct Slot1214_3 {
    int slots[16];
    predicate1214_3 accept;
} Slot1214_3;

int slot_accept1214_3(Slot1214_3 *store, int value) {
    return store->accept(value) + 94;
}
typedef int (*predicate1214_4)(int value);

typedef struct Slot1214_4 {
    int slots[16];
    predicate1214_4 accept;
} Slot1214_4;

int slot_accept1214_4(Slot1214_4 *store, int value) {
    return store->accept(value) + 101;
}
struct Node1214_5 {
    int value;
    struct Node1214_5 *next;
};

int length1214_5(const struct Node1214_5 *head) {
    int count = 0;
    const struct Node1214_5 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 11;
}
int scan1214_6(const int *begin, const int *end) {
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
    return total + 18;
}
enum Kind1214_7 { KIND_EMPTY1214_7 = 0, KIND_FULL1214_7 = 25 };

union Cell1214_7 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1214_7(union Cell1214_7 *cell) {
    cell->number = 44;
    return cell->number > 0 ? KIND_FULL1214_7 : KIND_EMPTY1214_7;
}
