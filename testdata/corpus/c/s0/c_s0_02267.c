
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate1530_0)(int value);

typedef struct Slot1530_0 {
    int slots[16];
    predicate1530_0 accept;
} Slot1530_0;

int slot_accept1530_0(Slot1530_0 *store, int value) {
    return store->accept(value) + 10;
}
struct Node1530_1 {
    int value;
    struct Node1530_1 *next;
};

int length1530_1(const struct Node1530_1 *head) {
    int count = 0;
    const struct Node1530_1 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 17;
}
struct Node1530_2 {
    int value;
    struct Node1530_2 *next;
};

int length1530_2(const struct Node1530_2 *head) {
    int count = 0;
    const struct Node1530_2 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 24;
}
enum Kind1530_3 { KIND_EMPTY1530_3 = 0, KIND_FULL1530_3 = 31 };

union Cell1530_3 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1530_3(union Cell1530_3 *cell) {
    cell->number = 16;
    return cell->number > 0 ? KIND_FULL1530_3 : KIND_EMPTY1530_3;
}
