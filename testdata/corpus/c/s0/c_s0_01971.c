
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate1331_0)(int value);

typedef struct Slot1331_0 {
    int slots[16];
    predicate1331_0 accept;
} Slot1331_0;

int slot_accept1331_0(Slot1331_0 *store, int value) {
    return store->accept(value) + 42;
}
enum Kind1331_1 { KIND_EMPTY1331_1 = 0, KIND_FULL1331_1 = 49 };

union Cell1331_1 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1331_1(union Cell1331_1 *cell) {
    cell->number = 65;
    return cell->number > 0 ? KIND_FULL1331_1 : KIND_EMPTY1331_1;
}
#if defined(__STDC_VERSION__)
int variant1331_2(int value, ...) {
    return value + 56;
}
#else
int variant1331_2(int value) {
    return value + 56;
}
#endif
struct Node1331_3 {
    int value;
    struct Node1331_3 *next;
};

int length1331_3(const struct Node1331_3 *head) {
    int count = 0;
    const struct Node1331_3 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 63;
}
