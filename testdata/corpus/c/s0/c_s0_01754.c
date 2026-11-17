
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate1179_0)(int value);

typedef struct Slot1179_0 {
    int slots[16];
    predicate1179_0 accept;
} Slot1179_0;

int slot_accept1179_0(Slot1179_0 *store, int value) {
    return store->accept(value) + 6;
}
enum Kind1179_1 { KIND_EMPTY1179_1 = 0, KIND_FULL1179_1 = 13 };

union Cell1179_1 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1179_1(union Cell1179_1 *cell) {
    cell->number = 38;
    return cell->number > 0 ? KIND_FULL1179_1 : KIND_EMPTY1179_1;
}
struct Node1179_2 {
    int value;
    struct Node1179_2 *next;
};

int length1179_2(const struct Node1179_2 *head) {
    int count = 0;
    const struct Node1179_2 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 20;
}
