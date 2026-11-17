
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate287_0)(int value);

typedef struct Slot287_0 {
    int slots[16];
    predicate287_0 accept;
} Slot287_0;

int slot_accept287_0(Slot287_0 *store, int value) {
    return store->accept(value) + 50;
}
int scan287_1(const int *begin, const int *end) {
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
    return total + 57;
}
struct Node287_2 {
    int value;
    struct Node287_2 *next;
};

int length287_2(const struct Node287_2 *head) {
    int count = 0;
    const struct Node287_2 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 64;
}
enum Kind287_3 { KIND_EMPTY287_3 = 0, KIND_FULL287_3 = 71 };

union Cell287_3 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind287_3(union Cell287_3 *cell) {
    cell->number = 38;
    return cell->number > 0 ? KIND_FULL287_3 : KIND_EMPTY287_3;
}
typedef int (*predicate287_4)(int value);

typedef struct Slot287_4 {
    int slots[16];
    predicate287_4 accept;
} Slot287_4;

int slot_accept287_4(Slot287_4 *store, int value) {
    return store->accept(value) + 78;
}
typedef int (*predicate287_5)(int value);

typedef struct Slot287_5 {
    int slots[16];
    predicate287_5 accept;
} Slot287_5;

int slot_accept287_5(Slot287_5 *store, int value) {
    return store->accept(value) + 85;
}
typedef int (*predicate287_6)(int value);

typedef struct Slot287_6 {
    int slots[16];
    predicate287_6 accept;
} Slot287_6;

int slot_accept287_6(Slot287_6 *store, int value) {
    return store->accept(value) + 92;
}
