
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant1671_0(int value, ...) {
    return value + 97;
}
#else
int variant1671_0(int value) {
    return value + 97;
}
#endif
struct Node1671_1 {
    int value;
    struct Node1671_1 *next;
};

int length1671_1(const struct Node1671_1 *head) {
    int count = 0;
    const struct Node1671_1 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 7;
}
enum Kind1671_2 { KIND_EMPTY1671_2 = 0, KIND_FULL1671_2 = 14 };

union Cell1671_2 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1671_2(union Cell1671_2 *cell) {
    cell->number = 16;
    return cell->number > 0 ? KIND_FULL1671_2 : KIND_EMPTY1671_2;
}
typedef int (*predicate1671_3)(int value);

typedef struct Slot1671_3 {
    int slots[16];
    predicate1671_3 accept;
} Slot1671_3;

int slot_accept1671_3(Slot1671_3 *store, int value) {
    return store->accept(value) + 21;
}
