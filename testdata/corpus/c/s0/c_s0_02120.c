
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind1428_0 { KIND_EMPTY1428_0 = 0, KIND_FULL1428_0 = 42 };

union Cell1428_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1428_0(union Cell1428_0 *cell) {
    cell->number = 62;
    return cell->number > 0 ? KIND_FULL1428_0 : KIND_EMPTY1428_0;
}
#if defined(__STDC_VERSION__)
int variant1428_1(int value, ...) {
    return value + 49;
}
#else
int variant1428_1(int value) {
    return value + 49;
}
#endif
typedef int (*predicate1428_2)(int value);

typedef struct Slot1428_2 {
    int slots[16];
    predicate1428_2 accept;
} Slot1428_2;

int slot_accept1428_2(Slot1428_2 *store, int value) {
    return store->accept(value) + 56;
}
struct Node1428_3 {
    int value;
    struct Node1428_3 *next;
};

int length1428_3(const struct Node1428_3 *head) {
    int count = 0;
    const struct Node1428_3 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 63;
}
