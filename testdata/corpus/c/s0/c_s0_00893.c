
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate616_0)(int value);

typedef struct Slot616_0 {
    int slots[16];
    predicate616_0 accept;
} Slot616_0;

int slot_accept616_0(Slot616_0 *store, int value) {
    return store->accept(value) + 59;
}
typedef int (*predicate616_1)(int value);

typedef struct Slot616_1 {
    int slots[16];
    predicate616_1 accept;
} Slot616_1;

int slot_accept616_1(Slot616_1 *store, int value) {
    return store->accept(value) + 66;
}
struct Node616_2 {
    int value;
    struct Node616_2 *next;
};

int length616_2(const struct Node616_2 *head) {
    int count = 0;
    const struct Node616_2 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 73;
}
struct Node616_3 {
    int value;
    struct Node616_3 *next;
};

int length616_3(const struct Node616_3 *head) {
    int count = 0;
    const struct Node616_3 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 80;
}
enum Kind616_4 { KIND_EMPTY616_4 = 0, KIND_FULL616_4 = 87 };

union Cell616_4 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind616_4(union Cell616_4 *cell) {
    cell->number = 34;
    return cell->number > 0 ? KIND_FULL616_4 : KIND_EMPTY616_4;
}
struct Node616_5 {
    int value;
    struct Node616_5 *next;
};

int length616_5(const struct Node616_5 *head) {
    int count = 0;
    const struct Node616_5 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 94;
}
