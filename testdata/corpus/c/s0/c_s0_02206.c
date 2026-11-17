
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate1484_0)(int value);

typedef struct Slot1484_0 {
    int slots[16];
    predicate1484_0 accept;
} Slot1484_0;

int slot_accept1484_0(Slot1484_0 *store, int value) {
    return store->accept(value) + 91;
}
enum Kind1484_1 { KIND_EMPTY1484_1 = 0, KIND_FULL1484_1 = 98 };

union Cell1484_1 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1484_1(union Cell1484_1 *cell) {
    cell->number = 38;
    return cell->number > 0 ? KIND_FULL1484_1 : KIND_EMPTY1484_1;
}
#if defined(__STDC_VERSION__)
int variant1484_2(int value, ...) {
    return value + 8;
}
#else
int variant1484_2(int value) {
    return value + 8;
}
#endif
struct Node1484_3 {
    int value;
    struct Node1484_3 *next;
};

int length1484_3(const struct Node1484_3 *head) {
    int count = 0;
    const struct Node1484_3 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 15;
}
typedef int (*predicate1484_4)(int value);

typedef struct Slot1484_4 {
    int slots[16];
    predicate1484_4 accept;
} Slot1484_4;

int slot_accept1484_4(Slot1484_4 *store, int value) {
    return store->accept(value) + 22;
}
