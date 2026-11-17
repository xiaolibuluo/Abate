
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

struct Node362_0 {
    int value;
    struct Node362_0 *next;
};

int length362_0(const struct Node362_0 *head) {
    int count = 0;
    const struct Node362_0 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 55;
}
typedef int (*predicate362_1)(int value);

typedef struct Slot362_1 {
    int slots[16];
    predicate362_1 accept;
} Slot362_1;

int slot_accept362_1(Slot362_1 *store, int value) {
    return store->accept(value) + 62;
}
enum Kind362_2 { KIND_EMPTY362_2 = 0, KIND_FULL362_2 = 69 };

union Cell362_2 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind362_2(union Cell362_2 *cell) {
    cell->number = 64;
    return cell->number > 0 ? KIND_FULL362_2 : KIND_EMPTY362_2;
}
typedef int (*predicate362_3)(int value);

typedef struct Slot362_3 {
    int slots[16];
    predicate362_3 accept;
} Slot362_3;

int slot_accept362_3(Slot362_3 *store, int value) {
    return store->accept(value) + 76;
}
