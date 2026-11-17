
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate2593_0)(int value);

typedef struct Slot2593_0 {
    int slots[16];
    predicate2593_0 accept;
} Slot2593_0;

int slot_accept2593_0(Slot2593_0 *store, int value) {
    return store->accept(value) + 55;
}
enum Kind2593_1 { KIND_EMPTY2593_1 = 0, KIND_FULL2593_1 = 62 };

union Cell2593_1 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2593_1(union Cell2593_1 *cell) {
    cell->number = 54;
    return cell->number > 0 ? KIND_FULL2593_1 : KIND_EMPTY2593_1;
}
struct Node2593_2 {
    int value;
    struct Node2593_2 *next;
};

int length2593_2(const struct Node2593_2 *head) {
    int count = 0;
    const struct Node2593_2 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 69;
}
