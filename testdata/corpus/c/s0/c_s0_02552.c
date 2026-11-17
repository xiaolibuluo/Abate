
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind1733_0 { KIND_EMPTY1733_0 = 0, KIND_FULL1733_0 = 30 };

union Cell1733_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1733_0(union Cell1733_0 *cell) {
    cell->number = 62;
    return cell->number > 0 ? KIND_FULL1733_0 : KIND_EMPTY1733_0;
}
#if defined(__STDC_VERSION__)
int variant1733_1(int value, ...) {
    return value + 37;
}
#else
int variant1733_1(int value) {
    return value + 37;
}
#endif
typedef int (*predicate1733_2)(int value);

typedef struct Slot1733_2 {
    int slots[16];
    predicate1733_2 accept;
} Slot1733_2;

int slot_accept1733_2(Slot1733_2 *store, int value) {
    return store->accept(value) + 44;
}
enum Kind1733_3 { KIND_EMPTY1733_3 = 0, KIND_FULL1733_3 = 51 };

union Cell1733_3 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1733_3(union Cell1733_3 *cell) {
    cell->number = 34;
    return cell->number > 0 ? KIND_FULL1733_3 : KIND_EMPTY1733_3;
}
struct Node1733_4 {
    int value;
    struct Node1733_4 *next;
};

int length1733_4(const struct Node1733_4 *head) {
    int count = 0;
    const struct Node1733_4 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 58;
}
