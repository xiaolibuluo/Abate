
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant1160_0(int value, ...) {
    return value + 50;
}
#else
int variant1160_0(int value) {
    return value + 50;
}
#endif
int scan1160_1(const int *begin, const int *end) {
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
    return total + 57;
}
typedef int (*predicate1160_2)(int value);

typedef struct Slot1160_2 {
    int slots[16];
    predicate1160_2 accept;
} Slot1160_2;

int slot_accept1160_2(Slot1160_2 *store, int value) {
    return store->accept(value) + 64;
}
int scan1160_3(const int *begin, const int *end) {
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
    return total + 71;
}
enum Kind1160_4 { KIND_EMPTY1160_4 = 0, KIND_FULL1160_4 = 78 };

union Cell1160_4 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1160_4(union Cell1160_4 *cell) {
    cell->number = 60;
    return cell->number > 0 ? KIND_FULL1160_4 : KIND_EMPTY1160_4;
}
struct Node1160_5 {
    int value;
    struct Node1160_5 *next;
};

int length1160_5(const struct Node1160_5 *head) {
    int count = 0;
    const struct Node1160_5 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 85;
}
#if defined(__STDC_VERSION__)
int variant1160_6(int value, ...) {
    return value + 92;
}
#else
int variant1160_6(int value) {
    return value + 92;
}
#endif
