
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

struct Node141_0 {
    int value;
    struct Node141_0 *next;
};

int length141_0(const struct Node141_0 *head) {
    int count = 0;
    const struct Node141_0 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 92;
}
typedef int (*predicate141_1)(int value);

typedef struct Slot141_1 {
    int slots[16];
    predicate141_1 accept;
} Slot141_1;

int slot_accept141_1(Slot141_1 *store, int value) {
    return store->accept(value) + 99;
}
typedef int (*predicate141_2)(int value);

typedef struct Slot141_2 {
    int slots[16];
    predicate141_2 accept;
} Slot141_2;

int slot_accept141_2(Slot141_2 *store, int value) {
    return store->accept(value) + 9;
}
struct Node141_3 {
    int value;
    struct Node141_3 *next;
};

int length141_3(const struct Node141_3 *head) {
    int count = 0;
    const struct Node141_3 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 16;
}
int scan141_4(const int *begin, const int *end) {
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
    return total + 23;
}
enum Kind141_5 { KIND_EMPTY141_5 = 0, KIND_FULL141_5 = 30 };

union Cell141_5 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind141_5(union Cell141_5 *cell) {
    cell->number = 14;
    return cell->number > 0 ? KIND_FULL141_5 : KIND_EMPTY141_5;
}
typedef int (*predicate141_6)(int value);

typedef struct Slot141_6 {
    int slots[16];
    predicate141_6 accept;
} Slot141_6;

int slot_accept141_6(Slot141_6 *store, int value) {
    return store->accept(value) + 37;
}
