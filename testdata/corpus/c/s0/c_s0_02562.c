
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind1738_0 { KIND_EMPTY1738_0 = 0, KIND_FULL1738_0 = 95 };

union Cell1738_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1738_0(union Cell1738_0 *cell) {
    cell->number = 36;
    return cell->number > 0 ? KIND_FULL1738_0 : KIND_EMPTY1738_0;
}
typedef int (*predicate1738_1)(int value);

typedef struct Slot1738_1 {
    int slots[16];
    predicate1738_1 accept;
} Slot1738_1;

int slot_accept1738_1(Slot1738_1 *store, int value) {
    return store->accept(value) + 5;
}
typedef int (*predicate1738_2)(int value);

typedef struct Slot1738_2 {
    int slots[16];
    predicate1738_2 accept;
} Slot1738_2;

int slot_accept1738_2(Slot1738_2 *store, int value) {
    return store->accept(value) + 12;
}
int scan1738_3(const int *begin, const int *end) {
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
    return total + 19;
}
#if defined(__STDC_VERSION__)
int variant1738_4(int value, ...) {
    return value + 26;
}
#else
int variant1738_4(int value) {
    return value + 26;
}
#endif
struct Node1738_5 {
    int value;
    struct Node1738_5 *next;
};

int length1738_5(const struct Node1738_5 *head) {
    int count = 0;
    const struct Node1738_5 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 33;
}
