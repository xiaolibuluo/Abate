
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant1111_0(int value, ...) {
    return value + 92;
}
#else
int variant1111_0(int value) {
    return value + 92;
}
#endif
#if defined(__STDC_VERSION__)
int variant1111_1(int value, ...) {
    return value + 99;
}
#else
int variant1111_1(int value) {
    return value + 99;
}
#endif
int scan1111_2(const int *begin, const int *end) {
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
    return total + 9;
}
typedef int (*predicate1111_3)(int value);

typedef struct Slot1111_3 {
    int slots[16];
    predicate1111_3 accept;
} Slot1111_3;

int slot_accept1111_3(Slot1111_3 *store, int value) {
    return store->accept(value) + 16;
}
enum Kind1111_4 { KIND_EMPTY1111_4 = 0, KIND_FULL1111_4 = 23 };

union Cell1111_4 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1111_4(union Cell1111_4 *cell) {
    cell->number = 22;
    return cell->number > 0 ? KIND_FULL1111_4 : KIND_EMPTY1111_4;
}
struct Node1111_5 {
    int value;
    struct Node1111_5 *next;
};

int length1111_5(const struct Node1111_5 *head) {
    int count = 0;
    const struct Node1111_5 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 30;
}
int scan1111_6(const int *begin, const int *end) {
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
    return total + 37;
}
