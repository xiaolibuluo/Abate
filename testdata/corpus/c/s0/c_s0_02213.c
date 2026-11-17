
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate1489_0)(int value);

typedef struct Slot1489_0 {
    int slots[16];
    predicate1489_0 accept;
} Slot1489_0;

int slot_accept1489_0(Slot1489_0 *store, int value) {
    return store->accept(value) + 59;
}
enum Kind1489_1 { KIND_EMPTY1489_1 = 0, KIND_FULL1489_1 = 66 };

union Cell1489_1 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1489_1(union Cell1489_1 *cell) {
    cell->number = 12;
    return cell->number > 0 ? KIND_FULL1489_1 : KIND_EMPTY1489_1;
}
#if defined(__STDC_VERSION__)
int variant1489_2(int value, ...) {
    return value + 73;
}
#else
int variant1489_2(int value) {
    return value + 73;
}
#endif
struct Node1489_3 {
    int value;
    struct Node1489_3 *next;
};

int length1489_3(const struct Node1489_3 *head) {
    int count = 0;
    const struct Node1489_3 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 80;
}
int scan1489_4(const int *begin, const int *end) {
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
