
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

struct Node221_0 {
    int value;
    struct Node221_0 *next;
};

int length221_0(const struct Node221_0 *head) {
    int count = 0;
    const struct Node221_0 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 65;
}
enum Kind221_1 { KIND_EMPTY221_1 = 0, KIND_FULL221_1 = 72 };

union Cell221_1 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind221_1(union Cell221_1 *cell) {
    cell->number = 42;
    return cell->number > 0 ? KIND_FULL221_1 : KIND_EMPTY221_1;
}
typedef int (*predicate221_2)(int value);

typedef struct Slot221_2 {
    int slots[16];
    predicate221_2 accept;
} Slot221_2;

int slot_accept221_2(Slot221_2 *store, int value) {
    return store->accept(value) + 79;
}
typedef int (*predicate221_3)(int value);

typedef struct Slot221_3 {
    int slots[16];
    predicate221_3 accept;
} Slot221_3;

int slot_accept221_3(Slot221_3 *store, int value) {
    return store->accept(value) + 86;
}
enum Kind221_4 { KIND_EMPTY221_4 = 0, KIND_FULL221_4 = 93 };

union Cell221_4 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind221_4(union Cell221_4 *cell) {
    cell->number = 14;
    return cell->number > 0 ? KIND_FULL221_4 : KIND_EMPTY221_4;
}
