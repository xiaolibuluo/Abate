
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate1193_0)(int value);

typedef struct Slot1193_0 {
    int slots[16];
    predicate1193_0 accept;
} Slot1193_0;

int slot_accept1193_0(Slot1193_0 *store, int value) {
    return store->accept(value) + 91;
}
struct Node1193_1 {
    int value;
    struct Node1193_1 *next;
};

int length1193_1(const struct Node1193_1 *head) {
    int count = 0;
    const struct Node1193_1 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 98;
}
#if defined(__STDC_VERSION__)
int variant1193_2(int value, ...) {
    return value + 8;
}
#else
int variant1193_2(int value) {
    return value + 8;
}
#endif
enum Kind1193_3 { KIND_EMPTY1193_3 = 0, KIND_FULL1193_3 = 15 };

union Cell1193_3 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1193_3(union Cell1193_3 *cell) {
    cell->number = 36;
    return cell->number > 0 ? KIND_FULL1193_3 : KIND_EMPTY1193_3;
}
struct Node1193_4 {
    int value;
    struct Node1193_4 *next;
};

int length1193_4(const struct Node1193_4 *head) {
    int count = 0;
    const struct Node1193_4 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 22;
}
