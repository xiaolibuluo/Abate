
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind2380_0 { KIND_EMPTY2380_0 = 0, KIND_FULL2380_0 = 99 };

union Cell2380_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2380_0(union Cell2380_0 *cell) {
    cell->number = 16;
    return cell->number > 0 ? KIND_FULL2380_0 : KIND_EMPTY2380_0;
}
int scan2380_1(const int *begin, const int *end) {
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
    return total + 9;
}
typedef int (*predicate2380_2)(int value);

typedef struct Slot2380_2 {
    int slots[16];
    predicate2380_2 accept;
} Slot2380_2;

int slot_accept2380_2(Slot2380_2 *store, int value) {
    return store->accept(value) + 16;
}
enum Kind2380_3 { KIND_EMPTY2380_3 = 0, KIND_FULL2380_3 = 23 };

union Cell2380_3 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2380_3(union Cell2380_3 *cell) {
    cell->number = 49;
    return cell->number > 0 ? KIND_FULL2380_3 : KIND_EMPTY2380_3;
}
typedef int (*predicate2380_4)(int value);

typedef struct Slot2380_4 {
    int slots[16];
    predicate2380_4 accept;
} Slot2380_4;

int slot_accept2380_4(Slot2380_4 *store, int value) {
    return store->accept(value) + 30;
}
typedef int (*predicate2380_5)(int value);

typedef struct Slot2380_5 {
    int slots[16];
    predicate2380_5 accept;
} Slot2380_5;

int slot_accept2380_5(Slot2380_5 *store, int value) {
    return store->accept(value) + 37;
}
struct Node2380_6 {
    int value;
    struct Node2380_6 *next;
};

int length2380_6(const struct Node2380_6 *head) {
    int count = 0;
    const struct Node2380_6 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 44;
}
