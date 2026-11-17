
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate112_0)(int value);

typedef struct Slot112_0 {
    int slots[16];
    predicate112_0 accept;
} Slot112_0;

int slot_accept112_0(Slot112_0 *store, int value) {
    return store->accept(value) + 6;
}
struct Node112_1 {
    int value;
    struct Node112_1 *next;
};

int length112_1(const struct Node112_1 *head) {
    int count = 0;
    const struct Node112_1 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 13;
}
enum Kind112_2 { KIND_EMPTY112_2 = 0, KIND_FULL112_2 = 20 };

union Cell112_2 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind112_2(union Cell112_2 *cell) {
    cell->number = 22;
    return cell->number > 0 ? KIND_FULL112_2 : KIND_EMPTY112_2;
}
