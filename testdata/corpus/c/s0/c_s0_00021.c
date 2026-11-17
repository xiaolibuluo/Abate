
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate12_0)(int value);

typedef struct Slot12_0 {
    int slots[16];
    predicate12_0 accept;
} Slot12_0;

int slot_accept12_0(Slot12_0 *store, int value) {
    return store->accept(value) + 64;
}
enum Kind12_1 { KIND_EMPTY12_1 = 0, KIND_FULL12_1 = 71 };

union Cell12_1 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind12_1(union Cell12_1 *cell) {
    cell->number = 43;
    return cell->number > 0 ? KIND_FULL12_1 : KIND_EMPTY12_1;
}
struct Node12_2 {
    int value;
    struct Node12_2 *next;
};

int length12_2(const struct Node12_2 *head) {
    int count = 0;
    const struct Node12_2 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 78;
}
