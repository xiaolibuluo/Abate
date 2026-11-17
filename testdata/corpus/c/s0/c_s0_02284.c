
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant1544_0(int value, ...) {
    return value + 95;
}
#else
int variant1544_0(int value) {
    return value + 95;
}
#endif
struct Node1544_1 {
    int value;
    struct Node1544_1 *next;
};

int length1544_1(const struct Node1544_1 *head) {
    int count = 0;
    const struct Node1544_1 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 5;
}
enum Kind1544_2 { KIND_EMPTY1544_2 = 0, KIND_FULL1544_2 = 12 };

union Cell1544_2 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1544_2(union Cell1544_2 *cell) {
    cell->number = 42;
    return cell->number > 0 ? KIND_FULL1544_2 : KIND_EMPTY1544_2;
}
enum Kind1544_3 { KIND_EMPTY1544_3 = 0, KIND_FULL1544_3 = 19 };

union Cell1544_3 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1544_3(union Cell1544_3 *cell) {
    cell->number = 53;
    return cell->number > 0 ? KIND_FULL1544_3 : KIND_EMPTY1544_3;
}
typedef int (*predicate1544_4)(int value);

typedef struct Slot1544_4 {
    int slots[16];
    predicate1544_4 accept;
} Slot1544_4;

int slot_accept1544_4(Slot1544_4 *store, int value) {
    return store->accept(value) + 26;
}
enum Kind1544_5 { KIND_EMPTY1544_5 = 0, KIND_FULL1544_5 = 33 };

union Cell1544_5 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1544_5(union Cell1544_5 *cell) {
    cell->number = 14;
    return cell->number > 0 ? KIND_FULL1544_5 : KIND_EMPTY1544_5;
}
