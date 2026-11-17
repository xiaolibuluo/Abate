
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate991_0)(int value);

typedef struct Slot991_0 {
    int slots[16];
    predicate991_0 accept;
} Slot991_0;

int slot_accept991_0(Slot991_0 *store, int value) {
    return store->accept(value) + 84;
}
enum Kind991_1 { KIND_EMPTY991_1 = 0, KIND_FULL991_1 = 91 };

union Cell991_1 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind991_1(union Cell991_1 *cell) {
    cell->number = 64;
    return cell->number > 0 ? KIND_FULL991_1 : KIND_EMPTY991_1;
}
struct Node991_2 {
    int value;
    struct Node991_2 *next;
};

int length991_2(const struct Node991_2 *head) {
    int count = 0;
    const struct Node991_2 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 98;
}
