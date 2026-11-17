
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

struct Node940_0 {
    int value;
    struct Node940_0 *next;
};

int length940_0(const struct Node940_0 *head) {
    int count = 0;
    const struct Node940_0 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 100;
}
typedef int (*predicate940_1)(int value);

typedef struct Slot940_1 {
    int slots[16];
    predicate940_1 accept;
} Slot940_1;

int slot_accept940_1(Slot940_1 *store, int value) {
    return store->accept(value) + 10;
}
enum Kind940_2 { KIND_EMPTY940_2 = 0, KIND_FULL940_2 = 17 };

union Cell940_2 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind940_2(union Cell940_2 *cell) {
    cell->number = 23;
    return cell->number > 0 ? KIND_FULL940_2 : KIND_EMPTY940_2;
}
