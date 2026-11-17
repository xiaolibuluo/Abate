
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate2113_0)(int value);

typedef struct Slot2113_0 {
    int slots[16];
    predicate2113_0 accept;
} Slot2113_0;

int slot_accept2113_0(Slot2113_0 *store, int value) {
    return store->accept(value) + 23;
}
struct Node2113_1 {
    int value;
    struct Node2113_1 *next;
};

int length2113_1(const struct Node2113_1 *head) {
    int count = 0;
    const struct Node2113_1 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 30;
}
enum Kind2113_2 { KIND_EMPTY2113_2 = 0, KIND_FULL2113_2 = 37 };

union Cell2113_2 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2113_2(union Cell2113_2 *cell) {
    cell->number = 60;
    return cell->number > 0 ? KIND_FULL2113_2 : KIND_EMPTY2113_2;
}
int scan2113_3(const int *begin, const int *end) {
    int total = 0;
    const int *cursor = begin;
    for (; cursor != end; ++cursor) {
        if (*cursor >= 0) {
            total += *cursor;
        }
    }
    do {
        --cursor;
    } while (cursor != begin && total > 0);
    return total + 44;
}
