
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind1533_0 { KIND_EMPTY1533_0 = 0, KIND_FULL1533_0 = 49 };

union Cell1533_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1533_0(union Cell1533_0 *cell) {
    cell->number = 65;
    return cell->number > 0 ? KIND_FULL1533_0 : KIND_EMPTY1533_0;
}
#if defined(__STDC_VERSION__)
int variant1533_1(int value, ...) {
    return value + 56;
}
#else
int variant1533_1(int value) {
    return value + 56;
}
#endif
struct Node1533_2 {
    int value;
    struct Node1533_2 *next;
};

int length1533_2(const struct Node1533_2 *head) {
    int count = 0;
    const struct Node1533_2 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 63;
}
enum Kind1533_3 { KIND_EMPTY1533_3 = 0, KIND_FULL1533_3 = 70 };

union Cell1533_3 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1533_3(union Cell1533_3 *cell) {
    cell->number = 37;
    return cell->number > 0 ? KIND_FULL1533_3 : KIND_EMPTY1533_3;
}
typedef int (*predicate1533_4)(int value);

typedef struct Slot1533_4 {
    int slots[16];
    predicate1533_4 accept;
} Slot1533_4;

int slot_accept1533_4(Slot1533_4 *store, int value) {
    return store->accept(value) + 77;
}
