
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate2448_0)(int value);

typedef struct Slot2448_0 {
    int slots[16];
    predicate2448_0 accept;
} Slot2448_0;

int slot_accept2448_0(Slot2448_0 *store, int value) {
    return store->accept(value) + 13;
}
enum Kind2448_1 { KIND_EMPTY2448_1 = 0, KIND_FULL2448_1 = 20 };

union Cell2448_1 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2448_1(union Cell2448_1 *cell) {
    cell->number = 15;
    return cell->number > 0 ? KIND_FULL2448_1 : KIND_EMPTY2448_1;
}
struct Node2448_2 {
    int value;
    struct Node2448_2 *next;
};

int length2448_2(const struct Node2448_2 *head) {
    int count = 0;
    const struct Node2448_2 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 27;
}
