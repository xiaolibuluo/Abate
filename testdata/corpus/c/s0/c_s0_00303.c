
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind209_0 { KIND_EMPTY209_0 = 0, KIND_FULL209_0 = 6 };

union Cell209_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind209_0(union Cell209_0 *cell) {
    cell->number = 69;
    return cell->number > 0 ? KIND_FULL209_0 : KIND_EMPTY209_0;
}
struct Node209_1 {
    int value;
    struct Node209_1 *next;
};

int length209_1(const struct Node209_1 *head) {
    int count = 0;
    const struct Node209_1 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 13;
}
typedef int (*predicate209_2)(int value);

typedef struct Slot209_2 {
    int slots[16];
    predicate209_2 accept;
} Slot209_2;

int slot_accept209_2(Slot209_2 *store, int value) {
    return store->accept(value) + 20;
}
