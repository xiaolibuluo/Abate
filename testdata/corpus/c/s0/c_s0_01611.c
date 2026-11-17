
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

struct Node1086_0 {
    int value;
    struct Node1086_0 *next;
};

int length1086_0(const struct Node1086_0 *head) {
    int count = 0;
    const struct Node1086_0 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 58;
}
typedef int (*predicate1086_1)(int value);

typedef struct Slot1086_1 {
    int slots[16];
    predicate1086_1 accept;
} Slot1086_1;

int slot_accept1086_1(Slot1086_1 *store, int value) {
    return store->accept(value) + 65;
}
enum Kind1086_2 { KIND_EMPTY1086_2 = 0, KIND_FULL1086_2 = 72 };

union Cell1086_2 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1086_2(union Cell1086_2 *cell) {
    cell->number = 69;
    return cell->number > 0 ? KIND_FULL1086_2 : KIND_EMPTY1086_2;
}
