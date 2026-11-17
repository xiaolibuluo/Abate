
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate845_0)(int value);

typedef struct Slot845_0 {
    int slots[16];
    predicate845_0 accept;
} Slot845_0;

int slot_accept845_0(Slot845_0 *store, int value) {
    return store->accept(value) + 29;
}
typedef int (*predicate845_1)(int value);

typedef struct Slot845_1 {
    int slots[16];
    predicate845_1 accept;
} Slot845_1;

int slot_accept845_1(Slot845_1 *store, int value) {
    return store->accept(value) + 36;
}
struct Node845_2 {
    int value;
    struct Node845_2 *next;
};

int length845_2(const struct Node845_2 *head) {
    int count = 0;
    const struct Node845_2 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 43;
}
enum Kind845_3 { KIND_EMPTY845_3 = 0, KIND_FULL845_3 = 50 };

union Cell845_3 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind845_3(union Cell845_3 *cell) {
    cell->number = 40;
    return cell->number > 0 ? KIND_FULL845_3 : KIND_EMPTY845_3;
}
