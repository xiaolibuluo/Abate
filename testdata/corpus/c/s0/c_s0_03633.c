
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate2454_0)(int value);

typedef struct Slot2454_0 {
    int slots[16];
    predicate2454_0 accept;
} Slot2454_0;

int slot_accept2454_0(Slot2454_0 *store, int value) {
    return store->accept(value) + 91;
}
struct Node2454_1 {
    int value;
    struct Node2454_1 *next;
};

int length2454_1(const struct Node2454_1 *head) {
    int count = 0;
    const struct Node2454_1 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 98;
}
enum Kind2454_2 { KIND_EMPTY2454_2 = 0, KIND_FULL2454_2 = 8 };

union Cell2454_2 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2454_2(union Cell2454_2 *cell) {
    cell->number = 68;
    return cell->number > 0 ? KIND_FULL2454_2 : KIND_EMPTY2454_2;
}
typedef int (*predicate2454_3)(int value);

typedef struct Slot2454_3 {
    int slots[16];
    predicate2454_3 accept;
} Slot2454_3;

int slot_accept2454_3(Slot2454_3 *store, int value) {
    return store->accept(value) + 15;
}
