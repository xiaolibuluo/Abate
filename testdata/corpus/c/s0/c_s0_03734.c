
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind2526_0 { KIND_EMPTY2526_0 = 0, KIND_FULL2526_0 = 57 };

union Cell2526_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2526_0(union Cell2526_0 *cell) {
    cell->number = 62;
    return cell->number > 0 ? KIND_FULL2526_0 : KIND_EMPTY2526_0;
}
int scan2526_1(const int *begin, const int *end) {
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
    return total + 64;
}
typedef int (*predicate2526_2)(int value);

typedef struct Slot2526_2 {
    int slots[16];
    predicate2526_2 accept;
} Slot2526_2;

int slot_accept2526_2(Slot2526_2 *store, int value) {
    return store->accept(value) + 71;
}
struct Node2526_3 {
    int value;
    struct Node2526_3 *next;
};

int length2526_3(const struct Node2526_3 *head) {
    int count = 0;
    const struct Node2526_3 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 78;
}
struct Node2526_4 {
    int value;
    struct Node2526_4 *next;
};

int length2526_4(const struct Node2526_4 *head) {
    int count = 0;
    const struct Node2526_4 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 85;
}
typedef int (*predicate2526_5)(int value);

typedef struct Slot2526_5 {
    int slots[16];
    predicate2526_5 accept;
} Slot2526_5;

int slot_accept2526_5(Slot2526_5 *store, int value) {
    return store->accept(value) + 92;
}
typedef int (*predicate2526_6)(int value);

typedef struct Slot2526_6 {
    int slots[16];
    predicate2526_6 accept;
} Slot2526_6;

int slot_accept2526_6(Slot2526_6 *store, int value) {
    return store->accept(value) + 99;
}
